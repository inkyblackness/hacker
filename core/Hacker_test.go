package core

import (
	"fmt"
	"os"

	"github.com/inkyblackness/hacker/styling"

	check "gopkg.in/check.v1"
)

type HackerSuite struct {
	hacker *Hacker

	testDirectories map[string][]os.FileInfo
}

var _ = check.Suite(&HackerSuite{})

func (suite *HackerSuite) SetUpTest(c *check.C) {
	suite.testDirectories = make(map[string][]os.FileInfo)

	suite.hacker = NewHacker(styling.NullStyle())
	suite.hacker.fileAccess = &fileAccess{
		readDir: func(path string) (info []os.FileInfo, err error) {
			var ok bool
			info, ok = suite.testDirectories[path]
			if !ok {
				err = fmt.Errorf("Not existing")
			}
			return
		}}

}

func (suite *HackerSuite) TestLoadOfUnknownLocationResultsInErrorMessage(c *check.C) {
	result := suite.hacker.Load("nonExisting1", "nonExisting2")

	c.Check(result, check.Equals, "Can't access directories")
}

func (suite *HackerSuite) TestLoadOfWrongLocationResultsInErrorMessage(c *check.C) {
	suite.testDirectories["dir1"] = []os.FileInfo{testFile("file1.res"), testFile("file2.res")}
	suite.testDirectories["dir2"] = []os.FileInfo{testFile("file3.res"), testFile("file4.res")}

	result := suite.hacker.Load("dir1", "dir2")

	c.Check(result, check.Equals, "Could not resolve release")
}

func (suite *HackerSuite) TestLoadOfKnownLocationResultsInConfirmation(c *check.C) {
	hdFiles, cdFiles := DataFiles(&dosCdRelease)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.testDirectories["dir2"] = testFiles(cdFiles...)

	result := suite.hacker.Load("dir1", "dir2")

	c.Check(result, check.Equals, "Loaded release [DOS CD Release]")
}

func (suite *HackerSuite) TestLoadAllowsOptionalSecondPath(c *check.C) {
	hdFiles, _ := DataFiles(&dosHdDemo)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)

	result := suite.hacker.Load("dir1", "")

	c.Check(result, check.Equals, "Loaded release [DOS HD Demo]")
}

func (suite *HackerSuite) TestLoadOfKnownSwitchedLocationResultsInConfirmation(c *check.C) {
	hdFiles, cdFiles := DataFiles(&dosCdDemo)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.testDirectories["dir2"] = testFiles(cdFiles...)

	result := suite.hacker.Load("dir2", "dir1")

	c.Check(result, check.Equals, "Loaded release [DOS CD Demo]")
}
