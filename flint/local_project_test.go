package flint

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type scenarios struct {
	path   string
	result bool
}

var readmeTests = []scenarios{
	{"", false},
	{"README", true},
	{"README.md", true},
	{"README.rst", true},
	{"docs/README.rst", false},
	{"docs/README.md", false},
}

func TestLocalProjectFindsReadme(t *testing.T) {
	for _, tt := range readmeTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "The README")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckReadme()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var contributingTests = []scenarios{
	{"", false},
	{"CONTRIBUTING", true},
	{"CONTRIBUTING.md", true},
	{"CONTRIBUTING.rst", true},
	{"docs/CONTRIBUTING.rst", false},
	{"docs/CONTRIBUTING.md", false},
}

func TestLocalProjectFindsContributing(t *testing.T) {
	for _, tt := range contributingTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "The CONTRIBUTING")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckContributing()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var licenseTests = []scenarios{
	{"", false},
	{"LICENSE", true},
	{"COPYING", true},
	{"LICENSE.md", true},
	{"LICENSE.rst", true},
	{"docs/LICENSE.rst", false},
	{"docs/LICENSE.md", false},
	{"docs/COPYING", false},
}

func TestLocalProjectFindsLicense(t *testing.T) {
	for _, tt := range licenseTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "The LICENSE")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckLicense()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var changelogTests = []scenarios{
	{"", false},
	{"CHANGELOG", true},
	{"CHANGELOG.md", true},
	{"CHANGELOG.rst", true},
	{"docs/CHANGELOG.rst", false},
	{"docs/CHANGELOG.md", false},
}

func TestLocalProjectFindsChangelog(t *testing.T) {
	for _, tt := range changelogTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "The CHANGELOG")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckChangelog()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var bootstrapTests = []scenarios{
	{"", false},
	{"script/bootstrap", true},
	{"util/script/bootstrap", false},
}

func TestLocalProjectFindsBootstrap(t *testing.T) {
	for _, tt := range bootstrapTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "#!/bin/sh")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckBootstrap()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var testScriptTests = []scenarios{
	{"", false},
	{"script/test", true},
	{"util/script/test", false},
}

func TestLocalProjectFindsTestScript(t *testing.T) {
	for _, tt := range testScriptTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "#!/bin/sh")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckTestScript()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

var CodeOfConductTests = []scenarios{
	{"", false},
	{"CODE_OF_CONDUCT", true},
	{"CODE_OF_CONDUCT.md", true},
}

func TestLocalProjectFindsCodeOfConduct(t *testing.T) {
	for _, tt := range CodeOfConductTests {
		setup := setupLocalProjectTest()
		defer setup.Teardown()

		if len(tt.path) > 0 {
			setup.WriteFile(tt.path, "#!/bin/sh")
		}

		project := &LocalProject{Path: setup.Path}
		actual := project.CheckCodeOfConduct()

		msg := fmt.Sprintf("Path: '%s', Errors: %d", tt.path, tt.result)
		assert.Equal(t, tt.result, actual, msg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type LocalTest struct{ Path string }

func (s *LocalTest) Cleanup() {
	os.RemoveAll(s.Path)
}

func (s *LocalTest) WriteFile(path string, content string) {
	bytes := []byte(content)
	dest := filepath.Join(s.Path, path)
	os.MkdirAll(filepath.Dir(dest), 0777)
	err := ioutil.WriteFile(dest, bytes, 0777)
	check(err)
}

func (s *LocalTest) Teardown() {
	s.Cleanup()
}

func setupLocalProjectTest() *LocalTest {
	fixturePath := filepath.Join("..", "tmp", "test-project")
	result := &LocalTest{Path: fixturePath}
	result.Cleanup() // Cleanup after previous failures
	os.MkdirAll(result.Path, 0777)
	return result
}
