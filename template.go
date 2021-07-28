package sql2ent

const TemplateSchema = `package schema

import (
    "entgo.io/ent"
	"entgo.io/ent/schema/field"
	
	{{ range .Imports }} "{.}" {{ end }}
	
)

type {{ .TableName}} struct {
    ent.Schema
}

func ({{ .TableName}}) Fields() []ent.Field {
	{{ if .IsHaveFields }}
    return []ent.Field{
        {{ range .Fields }}
			{{ . }},
		{{ end }}
    }
	{{ end }}
}

func ({{ .TableName}}) Edges() []ent.Edge {
	return nil
}`
