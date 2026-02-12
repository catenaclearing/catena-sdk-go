package main

type PaginatedOperation struct {
	Service         string // e.g., "integrations", "orgs"
	ServiceAccessor string // e.g., "Integrations", "Orgs"
	Tag             string // e.g., "Connections"
	OperationID     string // e.g., "list_connections"
	MethodName      string // e.g., "ListConnections"
	CursorParam     string // e.g., "cursor"
	LimitParam      string // e.g., "size"
	ResponseSchema  string // e.g., "_CursorPage_TypeVar_Customized_ConnectionRead_"
	ItemType        string // e.g., "ConnectionRead" (Go type)
	ItemPackage     string // e.g., "integrationsapi"
	RequestType     string // e.g., "ApiListConnectionsRequest"
	ResponseType    string // e.g., "CursorPageTypeVarCustomizedConnectionRead"
	RequiredParams  []Param
	OptionalParams  []Param
	CallStyle       string // "builder" or "flat"
	HasLimit        bool
	HasCursor       bool
	CursorGoName    string // e.g. "Cursor"
	LimitGoName     string // e.g. "Size"
	NextPageType    string // e.g. "NullableString"
}

type Param struct {
	Name        string // OpenAPI name
	GoName      string // Go struct field name / arg name
	Type        string // Go type
	IsPointer   bool
	In          string // "path", "query"
	IsEnum      bool
	IsEnumSlice bool
	EnumType    string
}
