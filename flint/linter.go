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

	if flags.RunReadme && !p.CheckReadme() {
		summary.AppendError(ReadmeNotFoundError)
		summary.AppendError(ReadmeNotFoundInfo)
	}
	if flags.RunContributing && !p.CheckContributing() {
		summary.AppendError(ContributingNotFoundError)
		summary.AppendError(ContributingNotFoundInfo)
	}
	if flags.RunLicense && !p.CheckLicense() {
		summary.AppendError(LicenseNotFoundError)
		summary.AppendError(LicenseNotFoundInfo)
	}
	if flags.RunBootstrap && !p.CheckBootstrap() {
		summary.AppendError(BootstrapNotFoundError)
		summary.AppendError(BootstrapNotFoundInfo)
	}
	if flags.RunTestScript && !p.CheckTestScript() {
		summary.AppendError(TestScriptNotFoundError)
		summary.AppendError(TestScriptNotFoundInfo)
	}

	return summary, nil
}
