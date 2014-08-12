package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/pengwynn/flint/flint"
)

func main() {
	app := cli.NewApp()
	app.Name = "flint"
	app.Usage = "Check a project for common sources of contributor friction"
	app.Version = "0.0.3"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"skip-readme", "skip check for README", ""},
		cli.BoolFlag{"skip-contributing", "skip check for contributing guide", ""},
		cli.BoolFlag{"skip-license", "skip check for license", ""},
		cli.BoolFlag{"skip-bootstrap", "skip check for bootstrap script", ""},
		cli.BoolFlag{"skip-test", "skip check for test script", ""},
		cli.BoolFlag{"skip-scripts", "skip check for all scripts", ""},
		cli.BoolFlag{"no-color", "skip coloring the terminal output", ""},
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
				if !c.Bool("no-color") { // if not skipping output color
					if element.Level == 0 { // if [FIXME]
						color.White(element.Message)
					} else { // if [ERROR]
						color.Yellow(element.Message)
					}
				} else {
					fmt.Println(element.Message)
				}
			}
			level := linter.Severity()
			if level > 0 {
				if !c.Bool("no-color") {
					color.Red("[CRITICAL] Some critical problems found. Please fix right away!")
				} else {
					fmt.Println("[CRITICAL] Some critical problems found. Please fix right away!")
				}
			}
			os.Exit(level)
		} else {
			if !c.Bool("no-color") {
				color.Green("[OK] All is well!")
			} else {
				fmt.Println("[OK] All is well!")
			}
		}
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
