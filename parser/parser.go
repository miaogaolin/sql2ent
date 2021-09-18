// https://entgo.io/zh/docs/schema-fields
package parser

import (
	"errors"

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

	sch, err := ParseSchema(tables[0])
	if err != nil {
		return "", err
	}

	return ParseTpl(sch)
}
