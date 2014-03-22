package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/pengwynn/flint/flint"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "flint"
	app.Usage = "Check a project for common sources of contributor friction"
	app.Version = "0.0.3"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"skip-readme", "skip check for README"},
		cli.BoolFlag{"skip-contributing", "skip check for contributing guide"},
		cli.BoolFlag{"skip-license", "skip check for license"},
		cli.BoolFlag{"skip-bootstrap", "skip check for bootstrap script"},
		cli.BoolFlag{"skip-test", "skip check for test script"},
		cli.BoolFlag{"skip-scripts", "skip check for all scripts"},
	}
	app.Action = func(c *cli.Context) {
		path, _ := os.Getwd()
		if len(c.Args()) > 0 {
			path = c.Args()[0]
		}
		linter := &flint.Lint{Path: path}

		if !c.Bool("skip-readme") {
			linter.CheckReadme()
		}
		if !c.Bool("skip-contributing") {
			linter.CheckContributing()
		}
		if !c.Bool("skip-license") {
			linter.CheckLicense()
		}
		if !c.Bool("skip-scripts") {
			if !c.Bool("skip-bootstrap") {
				linter.CheckBootstrap()
			}
			if !c.Bool("skip-test") {
				linter.CheckTest()
			}
		}

		if len(linter.Errors) > 0 {
			for _, element := range linter.Errors {
				fmt.Println(element.Message)
			}
			level := linter.Severity()
			if level > 0 {
				fmt.Println("[CRITICAL] Some critical problems found. Please fix right away!")
			}
			os.Exit(level)
		} else {
			fmt.Println("[OK] All is well!")
		}
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
