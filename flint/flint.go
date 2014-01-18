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

func (l *Lint) CheckReadme() {
	if !l.findFile("README*") {
		l.Errors = append(l.Errors, &lintError{2, "[ERROR] README not found"})
		l.Errors = append(l.Errors, &lintError{0, "[FIXME] Every project begins with a README. http://bit.ly/1dqUYQF"})
	}
}

func (l *Lint) CheckContributing() {
	if !l.findFile("CONTRIBUTING*") {
		l.Errors = append(l.Errors, &lintError{2, "[ERROR] CONTRIBUTING guide not found"})
		l.Errors = append(l.Errors, &lintError{0, "[FIXME] Add a CONTRIBUTING guide for potential contributors. http://git.io/z-TiGg"})
	}
}

func (l *Lint) CheckLicense() {
	if !l.findFile("LICENSE*") {
		l.Errors = append(l.Errors, &lintError{2, "[ERROR] LICENSE not found"})
		l.Errors = append(l.Errors, &lintError{0, "[FIXME] Add a license to protect yourself and your users. http://choosealicense.com/"})
	}
}

func (l *Lint) CheckBootstrap() {
	if !l.findFile("script/bootstrap") {
		l.Errors = append(l.Errors, &lintError{2, "[ERROR] Bootstrap script not found"})
		l.Errors = append(l.Errors, &lintError{0, "[FIXME] A bootstrap script makes setup a snap. http://bit.ly/JZjVL6"})
	}
}

func (l *Lint) CheckTest() {
	if !l.findFile("script/test") {
		l.Errors = append(l.Errors, &lintError{2, "[ERROR] Test script not found"})
		l.Errors = append(l.Errors, &lintError{0, "[FIXME] Make it easy to run the test suite regardless of project type. http://bit.ly/JZjVL6"})
	}
}
