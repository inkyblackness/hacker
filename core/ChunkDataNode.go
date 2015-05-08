package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/data"
)

type chunkDataNode struct {
	parentDataNode

	chunkID res.ResourceID
	holder  chunk.BlockHolder
}

func isLevelChunk(chunkID res.ResourceID, relativeID int) bool {
	result := false

	if (chunkID >= res.ResourceID(4000)) && (int(chunkID)%100) == relativeID {
		result = true
	}

	return result
}

func newChunkDataNode(parentNode DataNode, chunkID res.ResourceID, holder chunk.BlockHolder) *chunkDataNode {
	node := &chunkDataNode{
		parentDataNode: makeParentDataNode(parentNode, fmt.Sprintf("%04X", chunkID), int(holder.BlockCount())),
		chunkID:        chunkID,
		holder:         holder}
	var dataStruct interface{} = nil

	if chunkID == res.ResourceID(0x0FA1) {
		dataStruct = data.DefaultGameState()
	} else if isLevelChunk(chunkID, 5) {
		dataStruct = data.DefaultTileMap(64, 64)
	}

	for i := uint16(0); i < holder.BlockCount(); i++ {
		node.addChild(newBlockDataNode(node, i, holder.BlockData(i), dataStruct))
	}

	return node
}

func (node *chunkDataNode) Info() string {
	info := fmt.Sprintf("Available blocks: %d\nContent type: 0x%02X", node.holder.BlockCount(), node.holder.ContentType())

	return info
}

func (node *chunkDataNode) saveTo(consumer chunk.Consumer) {
	blockNodes := node.Children()
	blocks := make([][]byte, len(blockNodes))
	for index, blockNode := range blockNodes {
		blocks[index] = blockNode.Data()
	}
	newHolder := chunk.NewBlockHolder(node.holder.ChunkType(), node.holder.ContentType(), blocks)
	consumer.Consume(node.chunkID, newHolder)
}
