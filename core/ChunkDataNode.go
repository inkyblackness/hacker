package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
)

type chunkDataNode struct {
	parentNode DataNode
	chunkID    res.ResourceID

	holder         chunk.BlockHolder
	blockDataNodes []*blockDataNode
}

func newChunkDataNode(parentNode DataNode, chunkID res.ResourceID, holder chunk.BlockHolder) *chunkDataNode {
	node := &chunkDataNode{
		parentNode:     parentNode,
		chunkID:        chunkID,
		holder:         holder,
		blockDataNodes: make([]*blockDataNode, holder.BlockCount())}

	for i := uint16(0); i < holder.BlockCount(); i++ {
		node.blockDataNodes[i] = newBlockDataNode(node, i, holder.BlockData(i))
	}

	return node
}

func (node *chunkDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *chunkDataNode) Info() string {
	info := fmt.Sprintf("Available blocks: %d\nContent type: 0x%02X", node.holder.BlockCount(), node.holder.ContentType())

	return info
}

func (node *chunkDataNode) ID() string {
	return fmt.Sprintf("%04X", node.chunkID)
}

func (node *chunkDataNode) Resolve(path string) (resolved DataNode) {
	for _, temp := range node.blockDataNodes {
		if temp.ID() == path {
			resolved = temp
		}
	}

	return
}
