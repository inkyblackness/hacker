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

	for i := uint16(0); i < holder.BlockCount(); i++ {
		blockData := holder.BlockData(i)
		table := getTableForBlock(chunkID, blockData)

		if table != nil {
			node.addChild(newTableDataNode(node, fmt.Sprintf("%d", i), blockData, table))
		} else {
			dataStruct := getDataStructForBlock(chunkID)
			node.addChild(newBlockDataNode(node, i, blockData, dataStruct))
		}

	}

	return node
}

func getTableForBlock(chunkID res.ResourceID, blockData []byte) (table Table) {
	if isLevelChunk(chunkID, 5) {
		entryCount := 64 * 64
		table = data.NewTable(entryCount, func() interface{} { return data.DefaultTileMapEntry() })
	} else if isLevelChunk(chunkID, 8) {
		entryCount := len(blockData) / data.LevelObjectEntrySize
		table = data.NewTable(entryCount, func() interface{} { return data.DefaultLevelObjectEntry() })
	} else if isLevelChunk(chunkID, 9) {
		entryCount := len(blockData) / data.LevelObjectCrossReferenceSize
		table = data.NewTable(entryCount, func() interface{} { return data.DefaultLevelObjectCrossReference() })
	} else if isLevelChunk(chunkID, 10) {
		entryCount := len(blockData) / data.LevelWeaponEntrySize
		table = data.NewTable(entryCount, func() interface{} { return data.NewLevelWeaponEntry() })
	}

	return
}

func getDataStructForBlock(chunkID res.ResourceID) (dataStruct interface{}) {
	if chunkID == res.ResourceID(0x0FA0) {
		dataStruct = data.NewString("")
	} else if chunkID == res.ResourceID(0x0FA1) {
		dataStruct = data.DefaultGameState()
	} else if isLevelChunk(chunkID, 4) {
		dataStruct = data.DefaultLevelInformation()
	}

	return
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
