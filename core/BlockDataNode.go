package core

import (
	"fmt"
)

type blockDataNode struct {
	parentNode *chunkDataNode
	blockIndex uint16
}

func newBlockDataNode(parentNode *chunkDataNode, blockIndex uint16) *blockDataNode {
	node := &blockDataNode{
		parentNode: parentNode,
		blockIndex: blockIndex}

	return node
}

func (node *blockDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *blockDataNode) Info() string {
	info := ""

	return info
}

func (node *blockDataNode) ID() string {
	return fmt.Sprintf("%d", node.blockIndex)
}

func (node *blockDataNode) Resolve(path string) DataNode {
	return nil
}
