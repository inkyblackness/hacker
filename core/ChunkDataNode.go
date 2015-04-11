package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
)

type chunkDataNode struct {
	parentDataNode

	holder chunk.BlockHolder
}

func newChunkDataNode(parentNode DataNode, chunkID res.ResourceID, holder chunk.BlockHolder) *chunkDataNode {
	node := &chunkDataNode{
		parentDataNode: makeParentDataNode(parentNode, fmt.Sprintf("%04X", chunkID), int(holder.BlockCount())),
		holder:         holder}

	for i := uint16(0); i < holder.BlockCount(); i++ {
		node.addChild(newBlockDataNode(node, i, holder.BlockData(i)))
	}

	return node
}

func (node *chunkDataNode) Info() string {
	info := fmt.Sprintf("Available blocks: %d\nContent type: 0x%02X", node.holder.BlockCount(), node.holder.ContentType())

	return info
}
