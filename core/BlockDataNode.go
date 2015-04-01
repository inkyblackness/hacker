package core

import (
	"fmt"
)

type blockDataNode struct {
	parentNode DataNode
	blockIndex uint16

	data []byte
}

func newBlockDataNode(parentNode DataNode, blockIndex uint16, data []byte) *blockDataNode {
	node := &blockDataNode{
		parentNode: parentNode,
		blockIndex: blockIndex,
		data:       data}

	return node
}

func (node *blockDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *blockDataNode) Info() string {
	info := ""

	for index, value := range node.data {
		if index == 0 {
		} else if (index % 16) == 0 {
			info += "\n"
		} else if (index % 8) == 0 {
			info += "  "
		} else {
			info += " "
		}
		info += fmt.Sprintf("%02X", value)
	}

	return info
}

func (node *blockDataNode) ID() string {
	return fmt.Sprintf("%d", node.blockIndex)
}

func (node *blockDataNode) Resolve(path string) DataNode {
	return nil
}

func (node *blockDataNode) Data() []byte {
	return node.data
}
