package flint

import (
	"errors"
	"fmt"
	"github.com/octokit/go-octokit/octokit"
	"strings"
)

type GitHubFetcher struct {
	*octokit.Client
}

func (g *GitHubFetcher) FetchRepository(nwo string) (repo *Repository, err error) {
	if g.Client == nil {
		return nil, errors.New("GitHub client required")
	}

	owner, name, err := g.ParseFullName(nwo)
	if err != nil {
		return nil, err
	}
	info, result := g.Repositories().One(nil, octokit.M{"owner": owner, "repo": name})
	if result.HasError() {
		return nil, result.Err
	}

	repo = &Repository{info.Description, info.Homepage}

	return
}

func (g *GitHubFetcher) FetchTree(nwo string) (paths []string, err error) {
	if g.Client == nil {
		return nil, errors.New("GitHub client required")
	}

	owner, name, err := g.ParseFullName(nwo)
	if err != nil {
		return nil, err
	}
	url, err := octokit.GitTreesURL.Expand(octokit.M{
		"owner":     owner,
		"repo":      name,
		"sha":       "master",
		"recursive": 1,
	})
	if err != nil {
		return nil, err
	}
	tree, result := g.GitTrees(url).One()
	if result.HasError() {
		return nil, result.Err
	}

	for _, entry := range tree.Tree {
		paths = append(paths, entry.Path)
	}

	return
}

func (g *GitHubFetcher) ParseFullName(nwo string) (owner, name string, err error) {
	parts := strings.Split(nwo, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("Invalid GitHub repository: %s", nwo)
		return
	}
	owner = parts[0]
	name = parts[1]

	return
}

func NewGitHubFetcher() *GitHubFetcher {
	client := octokit.NewClient(nil)

	return &GitHubFetcher{client}
}

func NewGitHubFetcherWithToken(token string) *GitHubFetcher {
	client := octokit.NewClient(&octokit.TokenAuth{AccessToken: token})

	return &GitHubFetcher{client}
}
