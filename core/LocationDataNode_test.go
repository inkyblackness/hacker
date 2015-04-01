package core

import (
	check "gopkg.in/check.v1"
)

type LocationDataNodeSuite struct {
	parentNode           DataNode
	fileDataNodeProvider *TestingFileDataNodeProvider
	locationDataNode     DataNode
}

var _ = check.Suite(&LocationDataNodeSuite{})

func (suite *LocationDataNodeSuite) SetUpTest(c *check.C) {
	suite.fileDataNodeProvider = NewTestingFileDataNodeProvider()
	suite.locationDataNode = newLocationDataNode(suite.parentNode, HD,
		"/filePath", []string{"file1.res", "file2.res"}, suite.fileDataNodeProvider)
}

func (suite *LocationDataNodeSuite) TestResolveOfUnknownFileReturnsNil(c *check.C) {
	var dataNode DataNode = NewTestingDataNode("invalid")
	suite.fileDataNodeProvider.nodesByFileName["unknown.res"] = dataNode

	result := suite.locationDataNode.Resolve("unknown.res")

	c.Check(result, check.IsNil)
}

func (suite *LocationDataNodeSuite) TestResolveOfKnownFileReturnsDataNode(c *check.C) {
	var dataNode DataNode = NewTestingDataNode("id")
	suite.fileDataNodeProvider.nodesByFileName["file1.res"] = dataNode

	result := suite.locationDataNode.Resolve("file1.res")

	c.Check(result, check.Equals, dataNode)
}

func (suite *LocationDataNodeSuite) TestResolveOfKnownFileReturnsSameDataNodeSecondTime(c *check.C) {
	var dataNode1 DataNode = NewTestingDataNode("firstFile")
	var dataNode2 DataNode = NewTestingDataNode("secondFile")
	suite.fileDataNodeProvider.nodesByFileName["file1.res"] = dataNode1
	suite.locationDataNode.Resolve("file1.res")
	suite.fileDataNodeProvider.nodesByFileName["file1.res"] = dataNode2

	result := suite.locationDataNode.Resolve("file1.res")

	c.Check(result.ID(), check.Equals, "firstFile")
}

func (suite *LocationDataNodeSuite) TestResolveOfKnownFailingFileReturnsNil(c *check.C) {
	result := suite.locationDataNode.Resolve("file2.res")

	c.Check(result, check.IsNil)
}

func (suite *LocationDataNodeSuite) TestResolveOfKnownFailingFileReturnsDataNodeSecondTimeWhenOk(c *check.C) {
	var dataNode DataNode = NewTestingDataNode("file2.res")
	suite.locationDataNode.Resolve("file2.res")
	suite.fileDataNodeProvider.nodesByFileName["file2.res"] = dataNode

	result := suite.locationDataNode.Resolve("file2.res")

	c.Check(result.ID(), check.Equals, "file2.res")
}
