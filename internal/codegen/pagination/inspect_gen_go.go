package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"unicode"
)

func inspectGeneratedCode(rootDir string, ops []PaginatedOperation) ([]PaginatedOperation, error) {
	var result []PaginatedOperation

	for _, op := range ops {
		// Assume generated code is in gen/<service>
		genDir := filepath.Join(rootDir, "gen", op.Service)

		// Parse all files in the directory
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, genDir, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse dir %s: %w", genDir, err)
		}

		var pkg *ast.Package
		for _, p := range pkgs {
			if strings.HasSuffix(p.Name, "api") { // e.g. integrationsapi
				pkg = p
				break
			}
		}
		if pkg == nil {
			// Fallback to any package found
			for _, p := range pkgs {
				pkg = p
				break
			}
		}
		if pkg == nil {
			return nil, fmt.Errorf("no package found in %s", genDir)
		}

		// Determine MethodName from OperationID (PascalCase)
		methodName := toPascalCase(op.OperationID)
		op.MethodName = methodName
		op.ItemPackage = pkg.Name

		// Helper to find type definition
		findTypeSpec := func(name string) *ast.TypeSpec {
			for _, file := range pkg.Files {
				for _, decl := range file.Decls {
					if genDecl, ok := decl.(*ast.GenDecl); ok {
						for _, spec := range genDecl.Specs {
							if typeSpec, ok := spec.(*ast.TypeSpec); ok {
								if typeSpec.Name.Name == name {
									return typeSpec
								}
							}
						}
					}
				}
			}
			return nil
		}

		// Find the request builder struct
		// Convention: Api<MethodName>Request
		requestStructName := "Api" + methodName + "Request"
		op.RequestType = requestStructName

		// Find the struct definition and its methods
		var requestStruct *ast.TypeSpec
		var methods []*ast.FuncDecl

		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if typeSpec.Name.Name == requestStructName {
								requestStruct = typeSpec
							}
						}
					}
				}
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
						recvType := funcDecl.Recv.List[0].Type
						// Handle pointer receiver
						if starExpr, ok := recvType.(*ast.StarExpr); ok {
							recvType = starExpr.X
						}
						if ident, ok := recvType.(*ast.Ident); ok {
							if ident.Name == requestStructName {
								methods = append(methods, funcDecl)
							}
						}
					}
				}
			}
		}

		if requestStruct != nil {
			op.CallStyle = "builder"

			// Check for Execute method
			hasExecute := false
			for _, m := range methods {
				if m.Name.Name == "Execute" {
					hasExecute = true
					// Determine ResponseType from Execute return values
					// func (r Api...) Execute() (*ResponseType, *http.Response, error)
					if m.Type.Results != nil && len(m.Type.Results.List) == 3 {
						if starExpr, ok := m.Type.Results.List[0].Type.(*ast.StarExpr); ok {
							if ident, ok := starExpr.X.(*ast.Ident); ok {
								op.ResponseType = ident.Name
							}
						}
					}
				}
			}

			if !hasExecute {
				return nil, fmt.Errorf("request builder %s has no Execute method", requestStructName)
			}

			// Check for Cursor and Limit methods
			// They might be named differently, e.g. Cursor, Page, Size, Limit
			// We need to match them with op.CursorParam and op.LimitParam

			// Helper to find method for param
			findMethodForParam := func(paramName string) string {
				// Try exact match (PascalCase)
				target := toPascalCase(paramName)
				for _, m := range methods {
					if m.Name.Name == target {
						return target
					}
				}
				return ""
			}

			if op.HasCursor {
				method := findMethodForParam(op.CursorParam)
				if method == "" {
					return nil, fmt.Errorf("no method found for cursor param %s in %s", op.CursorParam, requestStructName)
				}
				op.CursorGoName = method
			}

			if op.HasLimit {
				method := findMethodForParam(op.LimitParam)
				if method == "" {
					return nil, fmt.Errorf("no method found for limit param %s in %s", op.LimitParam, requestStructName)
				}
				op.LimitGoName = method
			}

			// Fix Tag to match APIClient field
			// Assume Tag + "API"
			op.Tag = toPascalCase(op.Tag) + "API"

			// Find factory method to determine required params order and names
			// func (a *ApiService) MethodName(ctx context.Context, arg1 type1, arg2 type2) ApiRequest
			var factoryMethod *ast.FuncDecl
			for _, file := range pkg.Files {
				for _, decl := range file.Decls {
					if funcDecl, ok := decl.(*ast.FuncDecl); ok {
						if funcDecl.Name.Name == op.MethodName {
							// Check receiver is ApiService
							if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
								// We assume it's the right service if the method name matches
								factoryMethod = funcDecl
								break
							}
						}
					}
				}
				if factoryMethod != nil {
					break
				}
			}

			if factoryMethod == nil {
				return nil, fmt.Errorf("factory method %s not found", op.MethodName)
			}

			// Map factory arguments to RequiredParams
			// Args: ctx, reqParam1, reqParam2...
			// We need to reorder op.RequiredParams to match the factory method arguments.

			var orderedRequiredParams []Param
			// Skip the first argument (ctx)
			if len(factoryMethod.Type.Params.List) > 0 {
				for _, field := range factoryMethod.Type.Params.List {
					// Each field might have multiple names (e.g. func(a, b int))
					for _, name := range field.Names {
						if name.Name == "ctx" {
							continue
						}

						// Find matching param in op.RequiredParams
						found := false
						for _, p := range op.RequiredParams {
							// Match by name (ignoring case/format slightly?)
							// OpenAPI: connection_id, Go: connectionId
							if toCamelCase(p.Name) == name.Name {
								p.GoName = toPascalCase(p.Name) // Use PascalCase for struct field
								orderedRequiredParams = append(orderedRequiredParams, p)
								found = true
								break
							}
						}
						if !found {
							// If not found in RequiredParams, maybe it's a param that OpenAPI thinks is optional but Go thinks is required?
							// Or maybe name mismatch.
							// For now, fail if we can't match.
							return nil, fmt.Errorf("could not match Go argument %s to any required param for %s", name.Name, op.OperationID)
						}
					}
				}
			}
			op.RequiredParams = orderedRequiredParams

			// Map OptionalParams to setter methods
			for i := range op.OptionalParams {
				p := &op.OptionalParams[i]
				methodName := findMethodForParam(p.Name)
				if methodName != "" {
					p.GoName = methodName // This is the setter name (PascalCase), which is fine for struct field too?
					// Wait, for OptionalParams, we use them as:
					// if opts.GoName != nil { req = req.GoName(*opts.GoName) }
					// So GoName must be the struct field name AND the setter name.
					// Setter name is PascalCase. Struct field name should be PascalCase.
					// So this works.
				} else {
					// If no setter found, maybe it's not supported in the generated code?
					// Or maybe the name mapping is different.
					// Warn and skip? Or error?
					// For now, error.
					return nil, fmt.Errorf("no setter found for optional param %s", p.Name)
				}
			}

			// Determine ItemType from ResponseType
			// We need to find the ResponseType struct and look for "Items" field
			// ResponseType is e.g. CursorPageTypeVarCustomizedConnectionRead
			// It should have a field "Items" []ConnectionRead

			var responseStruct *ast.TypeSpec
			for _, file := range pkg.Files {
				for _, decl := range file.Decls {
					if genDecl, ok := decl.(*ast.GenDecl); ok {
						for _, spec := range genDecl.Specs {
							if typeSpec, ok := spec.(*ast.TypeSpec); ok {
								if typeSpec.Name.Name == op.ResponseType {
									responseStruct = typeSpec
								}
							}
						}
					}
				}
			}

			if responseStruct == nil {
				return nil, fmt.Errorf("response struct %s not found", op.ResponseType)
			}

			if structType, ok := responseStruct.Type.(*ast.StructType); ok {
				for _, field := range structType.Fields.List {
					if len(field.Names) > 0 && field.Names[0].Name == "Items" {
						// Found Items field
						if arrayType, ok := field.Type.(*ast.ArrayType); ok {
							if ident, ok := arrayType.Elt.(*ast.Ident); ok {
								op.ItemType = ident.Name
							} else if starExpr, ok := arrayType.Elt.(*ast.StarExpr); ok {
								if ident, ok := starExpr.X.(*ast.Ident); ok {
									op.ItemType = ident.Name // It might be a pointer, but we want the type name
								}
							}
						}
					}
					if len(field.Names) > 0 && field.Names[0].Name == "NextPage" {
						if ident, ok := field.Type.(*ast.Ident); ok {
							op.NextPageType = ident.Name
						}
					}
				}
			}

			if op.ItemType == "" {
				return nil, fmt.Errorf("could not determine ItemType for %s", op.ResponseType)
			}

			// Refine OptionalParams types
			for i, p := range op.OptionalParams {
				methodName := toPascalCase(p.Name)
				for _, m := range methods {
					if m.Name.Name == methodName {
						if len(m.Type.Params.List) > 0 {
							argType := m.Type.Params.List[0].Type
							// Handle pointer
							if starExpr, ok := argType.(*ast.StarExpr); ok {
								argType = starExpr.X
							}
							if ident, ok := argType.(*ast.Ident); ok {
								if !isBasicType(ident.Name) {
									typeSpec := findTypeSpec(ident.Name)
									if typeSpec != nil {
										if _, ok := typeSpec.Type.(*ast.StructType); ok {
											// It is a struct
											op.OptionalParams[i].Type = pkg.Name + "." + ident.Name
											op.OptionalParams[i].IsEnum = false
										} else {
											// It is likely an alias (enum)
											op.OptionalParams[i].IsEnum = true
											op.OptionalParams[i].EnumType = pkg.Name + "." + ident.Name
										}
									} else {
										// Fallback
										op.OptionalParams[i].IsEnum = true
										op.OptionalParams[i].EnumType = pkg.Name + "." + ident.Name
									}
								}
							}
						}
						break
					}
				}
			}

		} else {
			// Fallback to flat style (not implemented yet as per requirement "If request builder... If flat...")
			// But the generated code I saw IS request builder.
			// So I'll stick to builder for now.
			return nil, fmt.Errorf("only request builder style supported, could not find %s", requestStructName)
		}

		result = append(result, op)
	}

	return result, nil
}

func isBasicType(t string) bool {
	switch t {
	case "string", "bool", "int", "int32", "int64", "float32", "float64", "byte", "rune":
		return true
	}
	return false
}

func toPascalCase(s string) string {
	var sb strings.Builder
	nextUpper := true
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			nextUpper = true
			continue
		}
		if nextUpper {
			sb.WriteRune(unicode.ToUpper(r))
			nextUpper = false
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func toCamelCase(s string) string {
	s = toPascalCase(s)
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
