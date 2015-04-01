package core

import (
	check "gopkg.in/check.v1"
)

type RootDataNodeSuite struct {
	node *rootDataNode
}

var _ = check.Suite(&RootDataNodeSuite{})

func (suite *RootDataNodeSuite) SetUpTest(c *check.C) {
	suite.node = newRootDataNode(&dosCdRelease)
	suite.node.addLocation(newLocationDataNode(suite.node, HD, "/hdPath/", []string{"file1", "file2"}))
	suite.node.addLocation(newLocationDataNode(suite.node, CD, "/cdPath/", []string{"fileA", "fileB"}))

}

func (suite *RootDataNodeSuite) TestResolveOfDotDotReturnsNil(c *check.C) {
	result := suite.node.resolve("..")

	c.Check(result, check.IsNil)
}

func (suite *RootDataNodeSuite) TestResolveOfHdReturnsHdLocation(c *check.C) {
	result := suite.node.resolve("hd")

	c.Check(result, check.Equals, suite.node.locations[HD])
}
