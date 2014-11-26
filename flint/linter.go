package flint

import (
	"errors"
)

type Flags struct {
	RunReadme       bool
	RunContributing bool
	RunLicense      bool
	RunBootstrap    bool
	RunTestScript   bool
}

type Project interface {
	CheckReadme() bool
	CheckContributing() bool
	CheckLicense() bool
	CheckBootstrap() bool
	CheckTestScript() bool
}

type Linter struct{}

func (l *Linter) Run(p Project, flags *Flags) (summary *Summary, err error) {
	if p == nil {
		return nil, errors.New("Must supply a project")
	}

	summary = &Summary{}

	if flags.RunReadme {
		if !p.CheckReadme() {
			summary.AppendError(ReadmeNotFound)
		}
	}
	if flags.RunContributing {
		if !p.CheckContributing() {
			summary.AppendError(ContributingNotFound)
		}
	}
	if flags.RunLicense {
		if !p.CheckLicense() {
			summary.AppendError(LicenseNotFound)
		}
	}
	if flags.RunBootstrap {
		if !p.CheckBootstrap() {
			summary.AppendError(BootstrapNotFound)
		}
	}
	if flags.RunTestScript {
		if !p.CheckTestScript() {
			summary.AppendError(TestScriptNotFound)
		}
	}

	return summary, nil
}
