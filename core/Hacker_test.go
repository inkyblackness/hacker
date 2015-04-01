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
	suite.hacker.fileAccess = fileAccess{
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

func (suite *HackerSuite) TestLoadSetsUpRootNodeForHdOnly(c *check.C) {
	hdFiles, _ := DataFiles(&dosHdDemo)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.hacker.Load("dir1", "")

	c.Assert(suite.hacker.root, check.Not(check.IsNil))
	c.Check(suite.hacker.root.locations[HD].filePath, check.Equals, "dir1")
	_, exists := suite.hacker.root.locations[CD]
	c.Check(exists, check.Equals, false)
}

func (suite *HackerSuite) TestLoadSetsUpRootNodeForRelease(c *check.C) {
	hdFiles, cdFiles := DataFiles(&dosCdRelease)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.testDirectories["dir2"] = testFiles(cdFiles...)
	suite.hacker.Load("dir1", "dir2")

	c.Assert(suite.hacker.root, check.Not(check.IsNil))
	c.Check(suite.hacker.root.locations[HD].filePath, check.Equals, "dir1")
	c.Check(suite.hacker.root.locations[CD].filePath, check.Equals, "dir2")
}

func (suite *HackerSuite) TestLoadSetsUpRootNodeForSwappedPaths(c *check.C) {
	hdFiles, cdFiles := DataFiles(&dosCdRelease)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.testDirectories["dir2"] = testFiles(cdFiles...)
	suite.hacker.Load("dir2", "dir1")

	c.Assert(suite.hacker.root, check.Not(check.IsNil))
	c.Check(suite.hacker.root.locations[HD].filePath, check.Equals, "dir1")
	c.Check(suite.hacker.root.locations[CD].filePath, check.Equals, "dir2")
}

func (suite *HackerSuite) TestInfoWithoutDataReturnsHintToLoad(c *check.C) {
	result := suite.hacker.Info()

	c.Check(result, check.Equals, `No data loaded. Use the [load "path1" "path2"] command.`)
}

func (suite *HackerSuite) givenAStandardSetup() {
	hdFiles, cdFiles := DataFiles(&dosCdRelease)
	suite.testDirectories["dir1"] = testFiles(hdFiles...)
	suite.testDirectories["dir2"] = testFiles(cdFiles...)
	suite.hacker.Load("dir1", "dir2")

}

func (suite *HackerSuite) TestInfoAfterLoadReturnsReleaseInfo(c *check.C) {
	suite.givenAStandardSetup()

	result := suite.hacker.Info()

	c.Check(result, check.Equals, suite.hacker.root.Info())
}

func (suite *HackerSuite) TestChangeDirectoryChangesCurrentNode(c *check.C) {
	suite.givenAStandardSetup()

	suite.hacker.ChangeDirectory("hd")

	c.Check(suite.hacker.Info(), check.Equals, suite.hacker.root.locations[HD].Info())
}

func (suite *HackerSuite) TestChangeDirectoryHandlesStartingSlash(c *check.C) {
	suite.givenAStandardSetup()
	suite.hacker.ChangeDirectory("hd")

	suite.hacker.ChangeDirectory("/cd")

	c.Check(suite.hacker.Info(), check.Equals, suite.hacker.root.locations[CD].Info())
}

func (suite *HackerSuite) TestChangeDirectoryHandlesDotDot(c *check.C) {
	suite.givenAStandardSetup()
	suite.hacker.ChangeDirectory("hd")

	suite.hacker.ChangeDirectory("../cd")

	c.Check(suite.hacker.Info(), check.Equals, suite.hacker.root.locations[CD].Info())
}

func (suite *HackerSuite) TestChangeDirectoryIgnoresTrailingSlash(c *check.C) {
	suite.givenAStandardSetup()

	suite.hacker.ChangeDirectory("hd/")

	c.Check(suite.hacker.Info(), check.Equals, suite.hacker.root.locations[HD].Info())
}

func (suite *HackerSuite) TestCurrentDirctoryReturnsCurrentPath(c *check.C) {
	suite.givenAStandardSetup()

	suite.hacker.ChangeDirectory("hd")

	c.Check(suite.hacker.CurrentDirectory(), check.Equals, "/hd")
}
