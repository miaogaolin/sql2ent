package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

var (
	buildVersion = "1.0.0"
	commands     = []cli.Command{
		{
			Name:  "mysql",
			Usage: `generate ent schema`,
			Subcommands: []cli.Command{
				{
					Name:  "ddl",
					Usage: `generate mysql ent schema from ddl`,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "src, s",
							Usage: "the path or path globbing patterns of the ddl",
						},
						cli.StringFlag{
							Name:  "dir, d",
							Value: "./ent/schema",
							Usage: "the target dir",
						},
					},
					Action: MysqlDDL,
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Version = fmt.Sprintf("%s %s/%s", buildVersion, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println(aurora.Red("error: " + err.Error()))
	}
}
