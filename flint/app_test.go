package flint

import (
	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _run runFunc

func TestNewApp(t *testing.T) {
	app := NewApp()
	assert.Equal(t, "flint", app.Name)
}

func Test_newProject_BuildsLocalProjectByDefault(t *testing.T) {
	setupAppTest()
	defer tearDownAppTest()

	var project Project
	run = func(c *cli.Context) {
		project = newProject(c)
	}
	app := NewApp()
	app.Run([]string{""})
	assert.IsType(t, &LocalProject{}, project)
}

func Test_newProject_BuildsRemoteProjectWithGitHubFlag(t *testing.T) {
	setupAppTest()
	defer tearDownAppTest()

	var project Project
	run = func(c *cli.Context) {
		project = newProject(c)
	}
	app := NewApp()
	app.Run([]string{"", "--github", "pengwynn/flint"})
	assert.IsType(t, &RemoteProject{}, project)
}

func Test_newFlagsFromContext_TranslatesBoolFlags(t *testing.T) {
	setupAppTest()
	defer tearDownAppTest()

	var flags *Flags
	run = func(c *cli.Context) {
		flags = newFlagsFromContext(c)
	}
	app := NewApp()
	app.Run([]string{""})

	assert.True(t, flags.RunReadme)
	assert.True(t, flags.RunContributing)
	assert.True(t, flags.RunLicense)
	assert.True(t, flags.RunBootstrap)
	assert.True(t, flags.RunTestScript)

	app.Run([]string{"", "--skip-readme", "--skip-license", "--skip-test-script"})

	assert.False(t, flags.RunReadme)
	assert.True(t, flags.RunContributing)
	assert.False(t, flags.RunLicense)
	assert.True(t, flags.RunBootstrap)
	assert.False(t, flags.RunTestScript)

	app.Run([]string{"", "--skip-scripts"})

	assert.True(t, flags.RunReadme)
	assert.True(t, flags.RunContributing)
	assert.True(t, flags.RunLicense)
	assert.False(t, flags.RunBootstrap)
	assert.False(t, flags.RunTestScript)
}

type fakeGitHubFetcher struct{}

func (f *fakeGitHubFetcher) FetchRepository(nwo string) (repo *Repository, err error) {
	return &Repository{}, nil
}

func (f *fakeGitHubFetcher) FetchTree(nwo string) ([]string, error) {
	return []string{}, nil
}

func setupAppTest() {
	_run = run
	newGitHubFetcher = func(c *cli.Context) RemoteRepositoryFetcher {
		return &fakeGitHubFetcher{}
	}
}

func tearDownAppTest() {
	run = _run
}
