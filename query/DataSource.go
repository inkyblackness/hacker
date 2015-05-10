package query

import (
	"github.com/inkyblackness/res/data"
)

type DataSource interface {
	GameState() *data.GameState
	Tile(x int, y int) *data.TileMapEntry
	LevelObjectCrossReference(index uint16) *data.LevelObjectCrossReference
	LevelObject(index uint16) *data.LevelObjectEntry

	ObjectEntryPath(class int, index int) string
}
