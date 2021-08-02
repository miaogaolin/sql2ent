package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/miaogaolin/sql2ent/util"

	"github.com/miaogaolin/sql2ent/gen"
	"github.com/urfave/cli"
)

const (
	flagSrc = "src"
	flagDir = "dir"
)

func MysqlDDL(cli *cli.Context) error {
	src := cli.String(flagSrc)
	dir := cli.String(flagDir)
	src = strings.TrimSpace(src)
	if len(src) == 0 {
		return errors.New("expected path or path globbing patterns, but nothing found")
	}

	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}
	fmt.Println(files)

	if len(files) == 0 {
		return errors.New("sql not matched")
	}

	g := gen.NewMysqlGenerator(dir)

	for _, f := range files {
		fmt.Println(f)
		err := g.FromFile(f)
		if err != nil {
			return err
		}
	}
	return nil
}
