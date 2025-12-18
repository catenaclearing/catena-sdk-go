package main

import (
	_ "embed"
	"text/template"
)

const wrapperTemplate = `
// {{.MethodName}}PaginationOptions holds the options for the {{.MethodName}} paginated operation.
type {{.MethodName}}PaginationOptions struct {
	{{- range .RequiredParams}}
	// {{.GoName}} is a required parameter.
	{{.GoName}} {{.Type}}
	{{- end}}
	{{- range .OptionalParams}}
	{{.GoName}} *{{.Type}}
	{{- end}}
	{{- if .HasCursor}}
	// {{.CursorGoName}} is the cursor for the next page.
	{{.CursorGoName}} string
	{{- end}}
	{{- if .HasLimit}}
	// {{.LimitGoName}} is the maximum number of items to return per page.
	{{.LimitGoName}} *int
	{{- end}}
}

// {{.MethodName}}Each iterates over all items in the {{.MethodName}} operation.
func {{.MethodName}}Each(c *catena.Client, ctx context.Context, opts {{.MethodName}}PaginationOptions, yield func({{.ItemPackage}}.{{.ItemType}}) error) error {
	fetch := func(ctx context.Context, cursor string) ([]{{.ItemPackage}}.{{.ItemType}}, string, error) {
		req := c.{{.ServiceAccessor}}().{{.Tag}}.{{.MethodName}}(ctx{{range .RequiredParams}}, opts.{{.GoName}}{{end}})
		
		{{- range .OptionalParams}}
		if opts.{{.GoName}} != nil {
			{{- if .IsEnum}}
			req = req.{{.GoName}}({{.EnumType}}(*opts.{{.GoName}}))
			{{- else}}
			req = req.{{.GoName}}(*opts.{{.GoName}})
			{{- end}}
		}
		{{- end}}

		{{- if .HasCursor}}
		if cursor != "" {
			req = req.{{.CursorGoName}}(cursor)
		} else if opts.{{.CursorGoName}} != "" {
			req = req.{{.CursorGoName}}(opts.{{.CursorGoName}})
		}
		{{- end}}

		{{- if .HasLimit}}
		if opts.{{.LimitGoName}} != nil {
			req = req.{{.LimitGoName}}(int32(*opts.{{.LimitGoName}}))
		}
		{{- end}}

		resp, _, err := req.Execute()
		if err != nil {
			return nil, "", err
		}

		if resp == nil {
			return nil, "", nil
		}

		var next string
		{{- if eq .NextPageType "NullableString"}}
		if resp.NextPage.IsSet() && resp.NextPage.Get() != nil {
			next = *resp.NextPage.Get()
		}
		{{- else}}
		if resp.NextPage != nil {
			next = *resp.NextPage
		}
		{{- end}}

		return resp.Items, next, nil
	}

	return pagination.Each(ctx, opts.{{.CursorGoName}}, fetch, yield)
}
`

var tmpl = template.Must(template.New("wrapper").Parse(wrapperTemplate))
