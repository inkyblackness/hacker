package core

import (
	check "gopkg.in/check.v1"
)

type BlockDataNodeSuite struct {
	parentNode DataNode

	blockDataNode DataNode
}

var _ = check.Suite(&BlockDataNodeSuite{})

func (suite *BlockDataNodeSuite) TestInfoReturnsDataAsHexDump(c *check.C) {
	data := make([]byte, 17)
	for index := range data {
		data[index] = byte(index)
	}
	suite.blockDataNode = newBlockDataNode(suite.parentNode, uint16(0), data)

	result := suite.blockDataNode.Info()

	c.Check(result, check.Equals, "00 01 02 03 04 05 06 07  08 09 0A 0B 0C 0D 0E 0F\n10")
}
