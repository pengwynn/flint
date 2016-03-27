package flint

import (
	"errors"
	"regexp"
)

type Repository struct {
	Description string
	Homepage    string
}

type RemoteRepositoryFetcher interface {
	FetchRepository(string) (*Repository, error)
	FetchTree(string) ([]string, error)
	FetchReleases(string) ([]string, error)
}

type RemoteProject struct {
	FullName string
	paths    []string
	releases []string
	Repository
}

func (r *RemoteProject) Fetch(fetcher RemoteRepositoryFetcher) error {
	if len(r.FullName) == 0 {
		return errors.New("Must supply FullName as owner/repository")
	}

	info, err := fetcher.FetchRepository(r.FullName)

	if err != nil {
		return err
	}
	r.Repository.Description = info.Description
	r.Homepage = info.Homepage

	paths, err := fetcher.FetchTree(r.FullName)

	if err != nil {
		return err
	}
	r.paths = paths

	releases, err := fetcher.FetchReleases(r.FullName)

	if err != nil {
		return err
	}
	r.releases = releases

	return nil
}

func (l *RemoteProject) searchPath(re *regexp.Regexp) bool {
	for _, path := range l.paths {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

func (l *RemoteProject) CheckReadme() bool {
	return l.searchPath(regexp.MustCompile(`README`))
}

func (l *RemoteProject) CheckContributing() bool {
	return l.searchPath(regexp.MustCompile(`CONTRIBUTING`))
}

func (l *RemoteProject) CheckLicense() bool {
	return l.searchPath(regexp.MustCompile(`LICENSE|COPYING`))
}

func (l *RemoteProject) CheckChangelog() bool {
	return l.searchPath(regexp.MustCompile(`CHANGELOG`)) || len(l.releases) > 1
}

func (l *RemoteProject) CheckBootstrap() bool {
	return l.searchPath(regexp.MustCompile(`script/bootstrap`))
}

func (l *RemoteProject) CheckTestScript() bool {
	return l.searchPath(regexp.MustCompile(`script/test`))
}

func (l *RemoteProject) CheckCodeOfConduct() bool {
	return l.searchPath(regexp.MustCompile(`CODE_OF_CONDUCT`))
}
