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
}

type RemoteProject struct {
	FullName string
	paths    []string
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
	return l.searchPath(regexp.MustCompile(`(?i)README`))
}

func (l *RemoteProject) CheckLowercaseReadme() bool {
	return l.searchPath(regexp.MustCompile(`[Rr]eadme`))
}

func (l *RemoteProject) CheckContributing() bool {
	return l.searchPath(regexp.MustCompile(`(?i)CONTRIBUTING`))
}

func (l *RemoteProject) CheckLowercaseContributing() bool {
	return l.searchPath(regexp.MustCompile(`[Cc]ontributing`))
}

func (l *RemoteProject) CheckLicense() bool {
	return l.searchPath(regexp.MustCompile(`(?i)LICENSE`))
}

func (l *RemoteProject) CheckLowercaseLicense() bool {
	return l.searchPath(regexp.MustCompile(`[Ll]icense`))
}

func (l *RemoteProject) CheckBootstrap() bool {
	return l.searchPath(regexp.MustCompile(`script/bootstrap`))
}

func (l *RemoteProject) CheckTestScript() bool {
	return l.searchPath(regexp.MustCompile(`script/test`))
}
