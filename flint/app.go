package flint

import (
	"github.com/codegangsta/cli"
	"github.com/octokit/go-octokit/octokit"

	"fmt"
	"os"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "flint"
	app.Usage = "Check a project for common sources of contributor friction"
	app.Version = "0.0.4"
	app.Flags = []cli.Flag{
		cli.BoolFlag{"skip-readme", "skip check for README", ""},
		cli.BoolFlag{"skip-contributing", "skip check for contributing guide", ""},
		cli.BoolFlag{"skip-license", "skip check for license", ""},
		cli.BoolFlag{"skip-bootstrap", "skip check for bootstrap script", ""},
		cli.BoolFlag{"skip-test-script", "skip check for test script", ""},
		cli.BoolFlag{"skip-scripts", "skip check for all scripts", ""},
		cli.BoolFlag{"no-color", "skip coloring the terminal output", ""},
		cli.StringFlag{
			Name:  "github, g",
			Value: "",
			Usage: "GitHub repository as owner/repo",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			EnvVar: "FLINT_TOKEN",
			Usage:  "GitHub API access token",
		},
	}
	app.Action = func(c *cli.Context) {
		run(c)
	}

	return app
}

type runFunc func(*cli.Context)

var run = func(c *cli.Context) {
	project := newProject(c)
	flags := newFlagsFromContext(c)
	linter := &Linter{}
	summary, err := linter.Run(project, flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	if summary != nil {
		summary.Print(os.Stdout, !c.Bool("no-color"))
		os.Exit(summary.Severity())
	}
}

func newProject(c *cli.Context) Project {
	github := c.String("github")
	if len(github) > 0 {
		project := &RemoteProject{FullName: github}
		fetcher := newGitHubFetcher(c)
		err := project.Fetch(fetcher)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return project
	} else {
		path, _ := os.Getwd()
		if len(c.Args()) > 0 {
			path = c.Args()[0]
		}
		return &LocalProject{Path: path}
	}
}

func newFlagsFromContext(c *cli.Context) *Flags {
	runBootstrap := !c.Bool("skip-bootstrap")
	runTestScript := !c.Bool("skip-test-script")
	if c.Bool("skip-scripts") {
		runBootstrap = false
		runTestScript = false
	}

	flags := &Flags{
		RunReadme:       !c.Bool("skip-readme"),
		RunContributing: !c.Bool("skip-contributing"),
		RunLicense:      !c.Bool("skip-license"),
		RunBootstrap:    runBootstrap,
		RunTestScript:   runTestScript,
	}
	return flags
}

var newGitHubFetcher = func(c *cli.Context) RemoteRepositoryFetcher {
	var client *octokit.Client

	token := c.String("token")
	if len(token) > 0 {
		client = octokit.NewClient(octokit.TokenAuth{AccessToken: token})
	} else {
		client = octokit.NewClient(nil)
	}

	return &GitHubFetcher{client}
}
