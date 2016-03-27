package flint

import (
	"fmt"
)

type LintError struct {
	Level   int
	Message string
}

var labels = map[int]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
}

func (e *LintError) Error() (out string) {
	label := labels[e.Level]
	out = fmt.Sprintf("[%s] %s", label, e.Message)

	return
}

var ReadmeNotFoundError = &LintError{
	2,
	"README not found",
}
var ReadmeNotFoundInfo = &LintError{
	0,
	"Every project begins with a README. http://bit.ly/1dqUYQF",
}

var ContributingNotFoundError = &LintError{
	2,
	"CONTRIBUTING guide not found",
}

var ContributingNotFoundInfo = &LintError{
	0,
	"Add a guide for potential contributors. http://git.io/z-TiGg",
}

var LicenseNotFoundError = &LintError{
	2,
	"LICENSE not found",
}

var LicenseNotFoundInfo = &LintError{
	0,
	"Add a license to protect yourself and your users. http://choosealicense.com/",
}

var ChangelogNotFoundError = &LintError{
	1,
	"CHANGELOG not found",
}

var ChangelogNotFoundInfo = &LintError{
	0,
	"Add a changelog to show what's new or improved with each release. http://keepachangelog.com/",
}

var BootstrapNotFoundError = &LintError{
	1,
	"Bootstrap script not found",
}

var BootstrapNotFoundInfo = &LintError{
	0,
	"A bootstrap script makes setup a snap. http://bit.ly/JZjVL6",
}

var TestScriptNotFoundError = &LintError{
	1,
	"Test script not found",
}

var TestScriptNotFoundInfo = &LintError{
	0,
	"Make it easy to run the test suite regardless of project type. http://bit.ly/JZjVL6",
}

var CodeOfConductNotFoundError = &LintError{
	1,
	"CODE_OF_CONDUCT not found",
}
var CodeOfConductNotFoundInfo = &LintError{
	0,
	"Let people know what to expect when they participate in the project",
}
