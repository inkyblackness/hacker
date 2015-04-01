package core

import (
	"fmt"
)

type blockDataNode struct {
	parentNode *chunkDataNode
	blockIndex uint16
}

func (node *blockDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *blockDataNode) Info() string {
	info := "Chunk: " + node.ID()

	return info
}

func (node *blockDataNode) ID() string {
	return fmt.Sprintf("%d", node.blockIndex)
}

func (node *blockDataNode) Resolve(path string) DataNode {
	return nil
}
