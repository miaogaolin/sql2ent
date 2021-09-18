package parser

import (
	"errors"
	"fmt"
	"html/template"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/model/sql/util"

	"github.com/iancoleman/strcase"
	"github.com/tal-tech/go-zero/core/collection"

	"github.com/miaogaolin/sql2ent/converter"

	"github.com/miaogaolin/ddlparser/parser"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
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

func ParseSchema(e *parser.Table) (*Schema, error) {

	imports := collection.NewSet()
	sch := &Schema{
		TableName: strcase.ToCamel(e.Name),
	}
	fields, err := parserFields(e)
	if err != nil {
		return nil, err
	}

	var schFields []template.HTML
	for _, v := range fields {
		field := fmt.Sprintf(`field.%s("%s")`, v.FieldFuncName, v.Name.Source())

		field += fmt.Sprintf(`.SchemaType(map[string]string{
                dialect.MySQL:    "%s",   // Override MySQL.
            })`, v.DataTypeSource)

		if strings.Contains(v.DataTypeSource, "UNSIGNED") {
			field += ".NonNegative()"
		}

		if !v.IsNotNull {
			field += ".Optional()"
		}

		pkgs, defaultVal := converter.ConvertDefaultValue(v.DataType, v.DefaultValue.Value, v.DefaultValue.IsHas)
		addImports(imports, pkgs...)
		field += defaultVal

		if v.Comment != "" {
			field += fmt.Sprintf(`.Comment("%s")`, v.Comment)
		}

		if v.IsUnique || v.IsPrimary {
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
func parserFields(e *parser.Table) ([]Field, error) {
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

func convertColumns(columns []*parser.Column, primaryColumn string) ([]Field, error) {
	var fields []Field

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
		field.Comment = util.TrimNewLine(comment)
		field.IsPrimary = isPrimary
		field.IsAutoIncrement = isAutoIncrement
		field.IsNotNull = isNotNull
		field.DefaultValue = column.Constraint.DefaultValue
		field.DataTypeSource = column.DataTypeSource
		field.DataType = column.DataType
		field.IsUnique = isUnique

		fields = append(fields, field)
	}
	return fields, nil
}
