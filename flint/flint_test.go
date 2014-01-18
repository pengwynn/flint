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
	{"", 1},
	{"README", 0},
	{"README.md", 0},
	{"README.rst", 0},
	{"docs/README.rst", 1},
	{"docs/README.md", 1},
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
			assert.Equal(t, "[ERROR] README not found", lint.Errors[0])
		}
	}
}

var contributingTests = []scenarios{
	{"", 1},
	{"CONTRIBUTING", 0},
	{"CONTRIBUTING.md", 0},
	{"CONTRIBUTING.rst", 0},
	{"docs/CONTRIBUTING.rst", 1},
	{"docs/CONTRIBUTING.md", 1},
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
			assert.Equal(t, "[ERROR] CONTRIBUTING guide not found", lint.Errors[0])
		}
	}
}

var licenseTests = []scenarios{
	{"", 1},
	{"LICENSE", 0},
	{"LICENSE.md", 0},
	{"LICENSE.rst", 0},
	{"docs/LICENSE.rst", 1},
	{"docs/LICENSE.md", 1},
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
			assert.Equal(t, "[ERROR] LICENSE not found", lint.Errors[0])
		}
	}
}

var bootstrapScriptTests = []scenarios{
	{"", 1},
	{"script/bootstrap", 0},
	{"util/script/bootstrap", 1},
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
			assert.Equal(t, "[ERROR] Bootstrap script not found", lint.Errors[0])
		}
	}
}

var testScriptTests = []scenarios{
	{"", 1},
	{"script/test", 0},
	{"util/script/test", 1},
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
			assert.Equal(t, "[ERROR] Test script not found", lint.Errors[0])
		}
	}
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
