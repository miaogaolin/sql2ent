// https://entgo.io/zh/docs/schema-fields
package main

import (
	"errors"

	"github.com/miaogaolin/sql2ent/parser"

	ddlParser "github.com/miaogaolin/ddlparser/parser"
)

func Parse(sql string) (string, error) {
	p := ddlParser.NewParser()
	tables, err := p.From([]byte(sql))
	if err != nil {
		return "", err
	}

	if len(tables) > 1 || len(tables) == 0 {
		return "", errors.New("only supports one table")
	}

	sch, err := parser.ParseSchema(tables[0])
	if err != nil {
		return "", err
	}

	return parser.ParseTpl(sch)
}
