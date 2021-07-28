package sql2ent

const TemplateSchema = `package schema

import (
    "entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	
	{{ range .Imports }} "{{ . }}" {{ end }}
	
)

// {{ .TableName }} holds the schema definition for the {{ .TableName }} entity.
type {{ .TableName }} struct {
    ent.Schema
}

// Fields of the {{ .TableName }}.
func ({{ .TableName}}) Fields() []ent.Field {
	{{ if .IsHaveFields }}
    return []ent.Field{
        {{ range .Fields }}
			{{ . }},
		{{ end }}
    }
	{{ end }}
}

// Edges of the {{ .TableName }}.
func ({{ .TableName}}) Edges() []ent.Edge {
	return nil
}`
