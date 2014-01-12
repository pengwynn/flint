package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "flint"
	app.Usage = "Check a project for common sources of contributor friction"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"skip-readme", "skip check for README"},
		cli.BoolFlag{"skip-contributing", "skip check for contributing guide"},
		cli.BoolFlag{"skip-license", "skip check for license"},
		cli.BoolFlag{"skip-bootstrap", "skip check for bootstrap script"},
		cli.BoolFlag{"skip-test", "skip check for test script"},
	}
	app.Action = func(c *cli.Context) {
		path, _ := os.Getwd()
		if len(c.Args()) > 0 {
			path = c.Args()[0]
		}
		println(path)

		if !c.Bool("skip-readme") {
			checkReadme(path)
		}
		if !c.Bool("skip-contributing") {
			checkContributing(path)
		}
	}

	app.Run(os.Args)
}

func checkReadme(path string) {
	println("Checking", path)
}
