package core

import (
	"fmt"
)

type blockDataNode struct {
	rawDataNode
}

func newBlockDataNode(parentNode DataNode, blockIndex uint16, data []byte) *blockDataNode {
	node := &blockDataNode{rawDataNode{
		parentNode: parentNode,
		id:         fmt.Sprintf("%d", blockIndex),
		data:       data}}

	return node
}

func (node *blockDataNode) Info() string {
	info := ""

	return info
}
