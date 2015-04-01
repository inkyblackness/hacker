package core

import (
	"fmt"

	"github.com/inkyblackness/res"
)

type chunkDataNode struct {
	parentNode *resourceDataNode
	chunkID    res.ResourceID
}

func newChunkDataNode(parentNode *resourceDataNode, chunkID res.ResourceID) *chunkDataNode {
	node := &chunkDataNode{
		parentNode: parentNode,
		chunkID:    chunkID}

	return node
}

func (node *chunkDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *chunkDataNode) Info() string {
	info := "Chunk: " + node.Id()

	return info
}

func (node *chunkDataNode) Id() string {
	return fmt.Sprintf("%04X", node.chunkID)
}

func (node *chunkDataNode) Resolve(path string) DataNode {
	return nil
}
