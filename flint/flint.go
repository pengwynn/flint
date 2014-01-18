package flint

import (
	"path/filepath"
)

type Lint struct {
	Path   string
	Errors []string
	status int
}

func (l *Lint) findFile(pattern string) bool {
	search := filepath.Join(l.Path, pattern)
	matches, _ := filepath.Glob(search)
	return len(matches) > 0
}

func (l *Lint) CheckReadme() {
	if !l.findFile("README*") {
		l.Errors = append(l.Errors, "[ERROR] README not found")
	}
}

func (l *Lint) CheckContributing() {
	if !l.findFile("CONTRIBUTING*") {
		l.Errors = append(l.Errors, "[ERROR] CONTRIBUTING guide not found")
	}
}

func (l *Lint) CheckLicense() {
	if !l.findFile("LICENSE*") {
		l.Errors = append(l.Errors, "[ERROR] LICENSE not found")
	}
}

func (l *Lint) CheckBootstrap() {
	if !l.findFile("script/bootstrap") {
		l.Errors = append(l.Errors, "[ERROR] Bootstrap script not found")
	}
}

func (l *Lint) CheckTest() {
	if !l.findFile("script/test") {
		l.Errors = append(l.Errors, "[ERROR] Test script not found")
	}
}
