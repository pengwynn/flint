package flint

import (
	"fmt"
	"github.com/fatih/color"
	"io"
)

type Summary struct {
	Errors []*LintError
}

func (list *Summary) AppendError(err *LintError) {
	list.Errors = append(list.Errors, err)
}

func (list *Summary) Severity() int {
	severity := 0
	for _, e := range list.Errors {
		if e.Level > severity {
			severity = e.Level
		}
	}

	return severity
}

func (l *Summary) Print(out io.Writer, colored bool) {
	info := func(a ...interface{}) { fmt.Fprintln(out, a...) }
	warn := info
	fail := info
	success := info

	color.Output = out
	if colored {
		warn = color.New(color.FgYellow).PrintlnFunc()
		fail = color.New(color.FgRed).PrintlnFunc()
		success = color.New(color.FgGreen).PrintlnFunc()
	}
	if len(l.Errors) == 0 {
		message := "[OK] All is well!"
		success(message)
		return
	}

	for _, e := range l.Errors {
		switch e.Level {
		case 0:
			info(e.Error())
		case 1:
			warn(e.Error())
		case 2:
			fail(e.Error())
		}
	}
	if l.Severity() > 1 {
		message := "[CRITICAL] Some critical problems found."
		fail(message)
	}
}
