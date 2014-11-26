package flint

type LintError struct {
	Level   int
	Message string
	Fixme   string
}

func (e *LintError) FormattedMessage() string {
	label := ""
	if e.Level == 1 {
		label = "[WARNING]"
	}
	if e.Level == 2 {
		label = "[ERROR]"
	}

	return label + " " + e.Message
}

func (e *LintError) FormattedFixme() string {
	return "[FIXME] " + e.Fixme
}

var ReadmeNotFound = &LintError{
	2,
	"README not found",
	"Every project begins with a README. http://bit.ly/1dqUYQF",
}

var ContributingNotFound = &LintError{
	2,
	"CONTRIBUTING guide not found",
	"Add a guide for potential contributors. http://git.io/z-TiGg",
}

var LicenseNotFound = &LintError{
	2,
	"LICENSE not found",
	"Add a license to protect yourself and your users. http://choosealicense.com/",
}

var BootstrapNotFound = &LintError{
	1,
	"Bootstrap script not found",
	"A bootstrap script makes setup a snap. http://bit.ly/JZjVL6",
}

var TestScriptNotFound = &LintError{
	1,
	"Test script not found",
	"Make it easy to run the test suite regardless of project type. http://bit.ly/JZjVL6",
}
