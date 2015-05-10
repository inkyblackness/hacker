package core

import (
	"bytes"
	"fmt"

	"github.com/inkyblackness/res/data"
	"github.com/inkyblackness/res/serial"
)

type NodeDataSource struct {
	archiveNode  DataNode
	hacker       *Hacker
	currentLevel int
}

func NewNodeDataSource(archiveNode DataNode, hacker *Hacker) *NodeDataSource {
	source := &NodeDataSource{
		archiveNode: archiveNode,
		hacker:      hacker}

	source.currentLevel = int(source.GameState().CurrentLevel)

	return source
}

func (source *NodeDataSource) levelEntryPath(table int, index int) string {
	return fmt.Sprintf("%04X/0/%d", 4000+100*source.currentLevel+table, index)
}

func (source *NodeDataSource) ObjectEntryPath(class int, index int) string {
	return source.levelEntryPath(10+class, index)
}

func (source *NodeDataSource) mapData(data interface{}, path string) {
	blockNode := source.hacker.resolveFrom(source.archiveNode, path)
	coder := serial.NewDecoder(bytes.NewReader(blockNode.Data()))
	serial.MapData(data, coder)
}

func (source *NodeDataSource) GameState() *data.GameState {
	gameState := data.DefaultGameState()

	source.mapData(gameState, "0FA1/0")

	return gameState
}

func (source *NodeDataSource) Tile(x int, y int) *data.TileMapEntry {
	entry := data.DefaultTileMapEntry()
	path := source.levelEntryPath(5, y*64+x)

	source.mapData(entry, path)

	return entry
}

func (source *NodeDataSource) LevelObjectCrossReference(index uint16) *data.LevelObjectCrossReference {
	ref := data.DefaultLevelObjectCrossReference()

	source.mapData(ref, source.levelEntryPath(9, int(index)))

	return ref
}

func (source *NodeDataSource) LevelObject(index uint16) *data.LevelObjectEntry {
	entry := data.DefaultLevelObjectEntry()

	source.mapData(entry, source.levelEntryPath(8, int(index)))

	return entry
}
