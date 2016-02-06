package flint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinterRequiresProject(t *testing.T) {
	linter := &Linter{}
	summary, err := linter.Run(nil, nil)
	assert.Nil(t, summary)
	assert.NotNil(t, err)
	assert.Equal(t, "Must supply a project", err.Error())
}

func TestLinterReturnsSummary(t *testing.T) {
	linter := &Linter{}
	flags := &Flags{true, true, true, true, true, true} // Run all
	project := &FakeProject{Flags: flags}
	summary, err := linter.Run(project, flags)
	assert.Nil(t, err)
	if assert.NotNil(t, summary) {
		assert.Equal(t, 0, len(summary.Errors))
	}
}

func TestLinterReportsErrors(t *testing.T) {
	linter := &Linter{}
	flags := &Flags{true, true, true, true, true, true} // Run all
	results := &Flags{}                                 // Return false for each check
	project := &FakeProject{Flags: results}
	assert.Equal(t, results.RunReadme, false)
	summary, err := linter.Run(project, flags)
	assert.Nil(t, err)
	if assert.NotNil(t, summary) {
		assert.Equal(t, 12, len(summary.Errors))
	}
}

type FakeProject struct {
	*Flags
}

func (p *FakeProject) CheckReadme() bool {
	return p.RunReadme
}

func (p *FakeProject) CheckContributing() bool {
	return p.RunContributing
}

func (p *FakeProject) CheckLicense() bool {
	return p.RunLicense
}

func (p *FakeProject) CheckChangelog() bool {
	return p.RunChangelog
}

func (p *FakeProject) CheckBootstrap() bool {
	return p.RunBootstrap
}

func (p *FakeProject) CheckTestScript() bool {
	return p.RunTestScript
}
