package flint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteProjectFetchRequiresFullName(t *testing.T) {
	project := &RemoteProject{}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	if assert.NotNil(t, err) {
		assert.Equal(t, "Must supply FullName as owner/repository", err.Error())
	}
}

func TestRemoteProjectPopulatesProjectInfo(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.Equal(t, "Ruby toolkit for the GitHub API", project.Description)
	assert.Equal(t, "http://octokit.github.io/octokit.rb/", project.Homepage)
}

func TestRemoteProjectPopulatesTree(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckReadme())
}

func TestRemoteProjectCheckReadme(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckReadme())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckReadme())
}

func TestRemoteProjectCheckContributing(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckContributing())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckContributing())
}

func TestRemoteProjectCheckLicense(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckLicense())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckLicense())
}

func TestRemoteProjectCheckCopying(t *testing.T) {
	project := &RemoteProject{FullName: "projects/copying"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckLicense())

}

func TestRemoteProjectCheckChangelog(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckChangelog())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckChangelog())

	project = &RemoteProject{FullName: "projects/has-releases"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckChangelog())

	project = &RemoteProject{FullName: "projects/no-changelog"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckChangelog())
}

func TestRemoteProjectCheckBootstrap(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckBootstrap())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckBootstrap())

	project = &RemoteProject{FullName: "projects/no-files"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckBootstrap())
}

func TestRemoteProjectCheckTestScript(t *testing.T) {
	project := &RemoteProject{FullName: "octokit/octokit.rb"}
	fetcher := &FakeProjectFetcher{}
	err := project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckTestScript())

	project = &RemoteProject{FullName: "projects/lowercase-names"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.True(t, project.CheckTestScript())

	project = &RemoteProject{FullName: "projects/no-files"}
	err = project.Fetch(fetcher)
	assert.Nil(t, err)
	assert.False(t, project.CheckTestScript())
}

type FakeProjectFetcher struct {
}

func (f *FakeProjectFetcher) FetchRepository(nwo string) (repository *Repository, err error) {
	repository = &Repository{}
	switch nwo {
	case "octokit/octokit.rb":
		repository = &Repository{
			"Ruby toolkit for the GitHub API",
			"http://octokit.github.io/octokit.rb/",
		}
	}

	return repository, nil
}

func (f *FakeProjectFetcher) FetchTree(nwo string) (paths []string, err error) {
	switch nwo {
	case "octokit/octokit.rb":
		paths = []string{
			"CONTRIBUTING.md",
			"CHANGELOG.md",
			"LICENSE.md",
			"README.md",
			"lib",
			"lib/octokit.rb",
			"script/bootstrap",
			"script/test",
		}
	case "projects/lowercase-names":
		paths = []string{
			"contributing",
			"license",
			"changelog",
			"readme",
			"script/bootstrap",
			"script/test",
		}
	case "projects/no-files":
		paths = []string{}
	case "projects/copying":
		paths = []string{
			"COPYING",
			"contributing",
			"readme",
			"script/bootstrap",
			"script/test",
		}
	case "projects/has-releases":
		paths = []string{}
	}

	return paths, nil
}

func (f *FakeProjectFetcher) FetchReleases(nwo string) (releases []string, err error) {
	switch nwo {
	case "projects/has-releases":
		releases = []string{
			"v0.0.1",
			"v1.0.0",
		}
	case "projects/no-changelog":
		releases = []string{}
	}

	return releases, nil
}
