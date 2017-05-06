package flint

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubFetcher struct {
	*github.Client
}

func (g *GitHubFetcher) FetchRepository(nwo string) (repo *Repository, err error) {
	if g.Client == nil {
		return nil, errors.New("GitHub client required")
	}

	owner, name, err := g.ParseFullName(nwo)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	info, _, err := g.Client.Repositories.Get(ctx, owner, name)

	if err != nil {
		return nil, err
	}
	repo = &Repository{*info.Description, *info.Homepage}

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

	ctx := context.Background()
	tree, _, err := g.Client.Git.GetTree(ctx, owner, name, "master", true)
	if err != nil {
		return nil, err
	}

	for _, entry := range tree.Entries {
		paths = append(paths, *entry.Path)
	}

	return
}

func (g *GitHubFetcher) FetchReleases(nwo string) (releases []string, err error) {
	if g.Client == nil {
		return nil, errors.New("GitHub client required")
	}

	owner, name, err := g.ParseFullName(nwo)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	items, _, err := g.Client.Repositories.ListReleases(ctx, owner, name, nil)

	if err != nil {
		return nil, err
	}

	for _, release := range items {
		if body := *release.Body; len(body) > 0 {
			releases = append(releases, *release.TagName)
		}
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
	client := github.NewClient(nil)

	return &GitHubFetcher{client}
}

func NewGitHubFetcherWithToken(token string) *GitHubFetcher {
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	client := github.NewClient(tc)

	return &GitHubFetcher{client}
}
