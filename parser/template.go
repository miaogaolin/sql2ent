package parser

import (
	"bytes"
	"go/format"
	"html/template"
)

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

func ParseTpl(sch *Schema) (string, error) {
	tpl, err := template.New("tpl").Parse(TemplateSchema)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, sch)
	if err != nil {
		return "", err
	}
	res, err := format.Source(b.Bytes())
	return string(res), err
}
