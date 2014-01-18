package flint

import (
	"fmt"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type scenarios struct {
	path string // fixture path to write
	n    int    // Lint.Errors count
}

var readmeTests = []scenarios{
	{"", 2},
	{"README", 0},
	{"README.md", 0},
	{"README.rst", 0},
	{"docs/README.rst", 2},
	{"docs/README.md", 2},
}

func TestCheckReadme(t *testing.T) {
	for _, tt := range readmeTests {
		setup := Setup()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "The README")
		}

		lint := &Lint{Path: setup.Path}
		lint.CheckReadme()

		msg := fmt.Sprintf("Fixture: %s, Errors: %d", tt.path, tt.n)
		assert.Equal(t, len(lint.Errors), tt.n, msg)
		if tt.n > 0 {
			assert.Equal(t, "[ERROR] README not found", lint.Errors[0].Message)
			assert.Equal(t, "[FIXME] Every project begins with a README. http://bit.ly/1dqUYQF", lint.Errors[1].Message)
		}
	}
}

var contributingTests = []scenarios{
	{"", 2},
	{"CONTRIBUTING", 0},
	{"CONTRIBUTING.md", 0},
	{"CONTRIBUTING.rst", 0},
	{"docs/CONTRIBUTING.rst", 2},
	{"docs/CONTRIBUTING.md", 2},
}

func TestCheckContributing(t *testing.T) {
	for _, tt := range contributingTests {
		setup := Setup()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "Patches welcome")
		}

		lint := &Lint{Path: setup.Path}
		lint.CheckContributing()

		msg := fmt.Sprintf("Fixture: %s, Errors: %d", tt.path, tt.n)
		assert.Equal(t, len(lint.Errors), tt.n, msg)
		if tt.n > 0 {
			assert.Equal(t, "[ERROR] CONTRIBUTING guide not found", lint.Errors[0].Message)
			assert.Equal(t, "[FIXME] Add a CONTRIBUTING guide for potential contributors. http://git.io/z-TiGg", lint.Errors[1].Message)
		}
	}
}

var licenseTests = []scenarios{
	{"", 2},
	{"LICENSE", 0},
	{"LICENSE.md", 0},
	{"LICENSE.rst", 0},
	{"docs/LICENSE.rst", 2},
	{"docs/LICENSE.md", 2},
}

func TestCheckLicense(t *testing.T) {
	for _, tt := range licenseTests {
		setup := Setup()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "MIT")
		}

		lint := &Lint{Path: setup.Path}
		lint.CheckLicense()

		msg := fmt.Sprintf("Fixture: %s, Errors: %d", tt.path, tt.n)
		assert.Equal(t, len(lint.Errors), tt.n, msg)
		if tt.n > 0 {
			assert.Equal(t, "[ERROR] LICENSE not found", lint.Errors[0].Message)
			assert.Equal(t, "[FIXME] Add a license to protect yourself and your users. http://choosealicense.com/", lint.Errors[1].Message)
		}
	}
}

var bootstrapScriptTests = []scenarios{
	{"", 2},
	{"script/bootstrap", 0},
	{"util/script/bootstrap", 2},
}

func TestCheckBootstrapScript(t *testing.T) {
	for _, tt := range bootstrapScriptTests {
		setup := Setup()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "MIT")
		}

		lint := &Lint{Path: setup.Path}
		lint.CheckBootstrap()

		msg := fmt.Sprintf("Fixture: %s, Errors: %d", tt.path, tt.n)
		assert.Equal(t, len(lint.Errors), tt.n, msg)
		if tt.n > 0 {
			assert.Equal(t, "[ERROR] Bootstrap script not found", lint.Errors[0].Message)
			assert.Equal(t, "[FIXME] A bootstrap script makes setup a snap. http://bit.ly/JZjVL6", lint.Errors[1].Message)
		}
	}
}

var testScriptTests = []scenarios{
	{"", 2},
	{"script/test", 0},
	{"util/script/test", 2},
}

func TestCheckTestScript(t *testing.T) {
	for _, tt := range testScriptTests {
		setup := Setup()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "MIT")
		}

		lint := &Lint{Path: setup.Path}
		lint.CheckTest()

		msg := fmt.Sprintf("Fixture: %s, Errors: %d", tt.path, tt.n)
		assert.Equal(t, len(lint.Errors), tt.n, msg)
		if tt.n > 0 {
			assert.Equal(t, "[ERROR] Test script not found", lint.Errors[0].Message)
			assert.Equal(t, "[FIXME] Make it easy to run the test suite regardless of project type. http://bit.ly/JZjVL6", lint.Errors[1].Message)
		}
	}
}

func TestSeverity(t *testing.T) {
	lint := &Lint{}
	lint.Errors = append(lint.Errors, &lintError{2, "[ERROR] README not found"})
	assert.Equal(t, 2, lint.Severity())
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
	dest := filepath.Join(s.Path, path)
	os.MkdirAll(filepath.Dir(dest), 0777)
	err := ioutil.WriteFile(dest, bytes, 0777)
	check(err)
}

func (s *SetupResult) Teardown() {
	s.Cleanup()
}

func Setup() *SetupResult {
	fixturePath := filepath.Join("..", "tmp", "test-project")
	result := &SetupResult{Path: fixturePath}
	result.Cleanup() // Cleanup after previous failures
	os.MkdirAll(result.Path, 0777)
	return result
}
