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
	color.Output = out
	if len(l.Errors) > 0 {
		for _, e := range l.Errors {
			if colored {
				if e.Level == 0 { // [INFO]
					fmt.Fprintln(out, e.Error())
				}
				if e.Level == 1 { // [WARNING]
					color.Yellow(e.Error())
				}
				if e.Level == 2 { // [ERROR]
					color.Red(e.Error())
				}
			} else {
				fmt.Fprintln(out, e.Error())
			}
		}
		level := l.Severity()
		if level > 1 {
			message := "[CRITICAL] Some critical problems found."
			if colored {
				color.Red(message)
			} else {
				fmt.Fprintln(out, message)
			}
		}
	} else {
		message := "[OK] All is well!"
		if colored {
			color.Green(message)
		} else {
			fmt.Fprintln(out, message)
		}
	}
}
