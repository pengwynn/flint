package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

type lint struct {
	Path   string
	Errors []string
	status int
}

func (l *lint) findFile(pattern string) bool {
	matches, _ := filepath.Glob(l.Path + "/" + pattern)
	return len(matches) > 0
}

func (l *lint) CheckReadme() {
	if !l.findFile("README*") {
		l.Errors = append(l.Errors, "[ERROR] README not found")
	}
}

func (l *lint) CheckContributing() {
	if !l.findFile("CONTRIBUTING*") {
		l.Errors = append(l.Errors, "[ERROR] CONTRIBUTING not found")
	}
}

func (l *lint) CheckLicense() {
	if !l.findFile("LICENSE*") {
		l.Errors = append(l.Errors, "[ERROR] LICENSE not found")
	}
}

func (l *lint) CheckBootstrap() {
	if !l.findFile("script/bootstrap") {
		l.Errors = append(l.Errors, "[ERROR] Bootstrap script not found")
	}
}

func (l *lint) CheckTest() {
	if !l.findFile("script/test") {
		l.Errors = append(l.Errors, "[ERROR] Test script not found")
	}
}

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
		linter := &lint{Path: path}

		if !c.Bool("skip-readme") {
			linter.CheckReadme()
		}
		if !c.Bool("skip-contributing") {
			linter.CheckContributing()
		}
		if !c.Bool("skip-license") {
			linter.CheckLicense()
		}
		if !c.Bool("skip-bootstrap") {
			linter.CheckBootstrap()
		}
		if !c.Bool("skip-test") {
			linter.CheckTest()
		}

		if len(linter.Errors) > 0 {
			for _, element := range linter.Errors {
				fmt.Println(element)
			}
			os.Exit(1)
		}
	}

	app.Run(os.Args)
}
