// https://entgo.io/zh/docs/schema-fields
package sql2ent

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"html/template"

	"github.com/miaogaolin/sql2ent/utils"

	"github.com/miaogaolin/ddlparser/parser"
	"github.com/miaogaolin/sql2ent/converter"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"

	"github.com/iancoleman/strcase"

	"github.com/tal-tech/go-zero/core/collection"
)

type (
	Schema struct {
		TableName    string
		Fields       []template.HTML
		IsHaveFields bool
		Imports      []string
	}

	// Field describes a table field
	Field struct {
		Name            stringx.String
		IsPrimary       bool
		IsUnique        bool
		IsAutoIncrement bool
		IsNotNull       bool
		DefaultValue    parser.DefaultValue
		FieldFuncName   string
		DataTypeSource  string
		DataType        parser.DataType
		Comment         string
		SeqInIndex      int
		OrdinalPosition int
	}
)

func Parse(sql string) (string, error) {
	p := parser.NewParser()
	tables, err := p.From([]byte(sql))
	if err != nil {
		return "", err
	}

	if len(tables) > 1 || len(tables) == 0 {
		return "", errors.New("only supports one table")
	}

	sch, err := parserSchema(tables[0])
	if err != nil {
		return "", err
	}

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

func parserSchema(e *parser.Table) (*Schema, error) {

	imports := collection.NewSet()
	sch := &Schema{
		TableName: strcase.ToCamel(e.Name),
	}
	fields, err := parserFields(e)
	if err != nil {
		return nil, err
	}

	var schFields []template.HTML
	for k, v := range fields {
		if v.IsPrimary {
			// ent containers primary key.
			continue
		}
		field := fmt.Sprintf(`field.%s("%s")`, v.FieldFuncName, k)

		field += fmt.Sprintf(`.SchemaType(map[string]string{
                dialect.MySQL:    "%s",   // Override MySQL.
            })`, v.DataTypeSource)

		if !v.IsNotNull {
			field += ".Optional()"
		}

		pkgs, defaultVal := converter.ConvertDefaultValue(v.DataType, v.DefaultValue.Value, v.DefaultValue.IsHas)
		addImports(imports, pkgs...)
		field += defaultVal

		if v.Comment != "" {
			field += fmt.Sprintf(`.Comment("%s")`, v.Comment)
		}

		if v.IsUnique {
			field += ".Unique()"
		}

		schFields = append(schFields, template.HTML(field))
	}
	sch.Fields = schFields
	sch.IsHaveFields = len(schFields) > 0
	sch.Imports = imports.KeysStr()
	return sch, nil

}

func addImports(imports *collection.Set, news ...string) {
	for _, v := range news {
		if !imports.Contains(v) {
			imports.AddStr(v)
		}
	}
}
func parserFields(e *parser.Table) (map[string]*Field, error) {
	primaryKey := ""
	primaryKeys := collection.NewSet()
	for _, e := range e.Constraints {
		if len(e.ColumnPrimaryKey) == 1 {
			primaryKeys.AddStr(e.ColumnPrimaryKey[0])
			primaryKey = e.ColumnPrimaryKey[0]
		}
	}

	columns := e.Columns

	for _, c := range columns {
		if c.Constraint != nil {
			// ent contains the primary key.
			if c.Constraint.Primary || primaryKeys.Contains(c.Name) {
				primaryKeys.AddStr(c.Name)
				primaryKey = c.Name
				continue
			}
		}
	}

	if primaryKeys.Count() > 1 {
		return nil, errors.New("unexpected join primary key")
	}

	fields, err := convertColumns(columns, primaryKey)
	return fields, err
}

func convertColumns(columns []*parser.Column, primaryColumn string) (map[string]*Field, error) {
	var fieldM = make(map[string]*Field)

	for _, column := range columns {
		if column == nil {
			continue
		}

		var (
			comment         string
			isPrimary       bool
			isAutoIncrement bool
			isNotNull       bool
			isUnique        bool
		)
		if column.Constraint != nil {
			comment = column.Constraint.Comment

			if column.Name == primaryColumn {
				isPrimary = true
			}
			if column.Constraint.AutoIncrement {
				isAutoIncrement = true

			}
			if column.Constraint.NotNull {
				isNotNull = true
			}

			isUnique = column.Constraint.Unique
		}

		fieldFuncName, err := converter.ConvertField(column.DataType.Type())
		if err != nil {
			return nil, err
		}

		var field Field
		field.Name = stringx.From(column.Name)
		field.FieldFuncName = fieldFuncName
		field.Comment = utils.TrimNewLine(comment)
		field.IsPrimary = isPrimary
		field.IsAutoIncrement = isAutoIncrement
		field.IsNotNull = isNotNull
		field.DefaultValue = column.Constraint.DefaultValue
		field.DataTypeSource = column.DataTypeSource
		field.DataType = column.DataType
		field.IsUnique = isUnique

		fieldM[field.Name.Source()] = &field
	}
	return fieldM, nil
}
