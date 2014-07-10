package flint

import (
	"path/filepath"
)

type lintError struct {
	Level   int
	Message string
}

type Lint struct {
	Path   string
	Errors []*lintError
	status int
}

func (l *Lint) findFile(pattern string) bool {
	search := filepath.Join(l.Path, pattern)
	matches, _ := filepath.Glob(search)
	return len(matches) > 0
}

func (l *Lint) Severity() int {
	severity := 0
	for _, e := range l.Errors {
		if e.Level > severity {
			severity = e.Level
		}
	}

	return severity
}

func (l *Lint) Check(file, message, fixme string) {
	if !l.findFile(file) {
		l.Errors = append(l.Errors, &lintError{2, message})
		l.Errors = append(l.Errors, &lintError{0, fixme})
	}
}

func (l *Lint) CheckReadme() {
	l.Check(
		"README*",
		"[ERROR] README not found",
		"[FIXME] Every project begins with a README. http://bit.ly/1dqUYQF")
}

func (l *Lint) CheckContributing() {
	l.Check(
		"CONTRIBUTING*",
		"[ERROR] CONTRIBUTING guide not found",
		"[FIXME] Add a CONTRIBUTING guide for potential contributors. http://git.io/z-TiGg")
}

func (l *Lint) CheckLicense() {
	l.Check(
		"LICENSE*",
		"[ERROR] LICENSE not found",
		"[FIXME] Add a license to protect yourself and your users. http://choosealicense.com/")
}

func (l *Lint) CheckBootstrap() {
	l.Check(
		"script/bootstrap",
		"[ERROR] Bootstrap script not found",
		"[FIXME] A bootstrap script makes setup a snap. http://bit.ly/JZjVL6")
}

func (l *Lint) CheckTest() {
	l.Check(
		"script/test",
		"[ERROR] Test script not found",
		"[FIXME] Make it easy to run the test suite regardless of project type. http://bit.ly/JZjVL6")
}
