package flint

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeverity(t *testing.T) {
	list := &Summary{}
	list.Errors = append(list.Errors, &LintError{2, "README not found", "Add a README, yo"})
	assert.Equal(t, 2, list.Severity())
}

func TestPrintNoColor(t *testing.T) {
	list := &Summary{}
	list.Errors = append(list.Errors, &LintError{2, "README not found", "Add a README, yo"})
	list.Errors = append(list.Errors, &LintError{1, "Bootstrap not found", "Add a bootstrap, yo"})
	buf := &bytes.Buffer{}

	list.Print(buf, false)
	expected := `[ERROR] README not found
[FIXME] Add a README, yo
[WARNING] Bootstrap not found
[FIXME] Add a bootstrap, yo
[CRITICAL] Some critical problems found.
`

	assert.Equal(t, expected, buf.String())
}

func TestPrintColor(t *testing.T) {
	list := &Summary{}
	list.Errors = append(list.Errors, &LintError{2, "README not found", "Add a README, yo"})
	list.Errors = append(list.Errors, &LintError{1, "Bootstrap not found", "Add a bootstrap, yo"})
	buf := &bytes.Buffer{}

	list.Print(buf, true)
	expected := "\x1b[31m[ERROR] README not found\n\x1b[0m"
	expected += "[FIXME] Add a README, yo\n"
	expected += "\x1b[33m[WARNING] Bootstrap not found\n\x1b[0m"
	expected += "[FIXME] Add a bootstrap, yo\n"
	expected += "\x1b[31m[CRITICAL] Some critical problems found.\n\x1b[0m"

	assert.Equal(t, expected, buf.String())
}

func TestPrintSuccessNoColor(t *testing.T) {
	list := &Summary{}
	buf := &bytes.Buffer{}

	list.Print(buf, false)
	expected := "[OK] All is well!\n"

	assert.Equal(t, expected, buf.String())
}

func TestPrintSuccessColor(t *testing.T) {
	list := &Summary{}
	buf := &bytes.Buffer{}

	list.Print(buf, true)
	expected := "\x1b[32m[OK] All is well!\n\x1b[0m"

	assert.Equal(t, expected, buf.String())
}
