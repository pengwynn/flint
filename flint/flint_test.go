package flint

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReportsMissingReadme(t *testing.T) {
	setup := Setup()
	defer setup.Teardown()

	lint := &Lint{Path: setup.Path}
	lint.CheckReadme()
	assert.Equal(t, len(lint.Errors), 1)
	msg := "[ERROR] README not found"
	assert.Equal(t, msg, lint.Errors[0])
}

func TestFindsReadme(t *testing.T) {
	setup := Setup()
	setup.WriteFile("README.md", "The README")
	defer setup.Teardown()

	lint := &Lint{Path: setup.Path}
	lint.CheckReadme()
	assert.Equal(t, len(lint.Errors), 0)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type SetupResult struct{ Path string }

func (s *SetupResult) Cleanup() {
	os.RemoveAll(s.Path)
}

func (s *SetupResult) WriteFile(path string, content string) {
	bytes := []byte(content)
	err := ioutil.WriteFile(s.Path+"/"+path, bytes, 0777)
	check(err)
}

func (s *SetupResult) Teardown() {
	s.Cleanup()
}

func Setup() *SetupResult {
	result := &SetupResult{Path: "../tmp/test-project"}
	result.Cleanup() // Cleanup after previous failures
	os.MkdirAll(result.Path, 0777)
	return result
}
