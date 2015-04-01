package core

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
)

type nullChunkProvider struct{}

func NewEmptyChunkProvider() chunk.Provider {
	return &nullChunkProvider{}
}

// IDs returns an empty list
func (provider *nullChunkProvider) IDs() []res.ResourceID {
	return nil
}

// Provide returns nil
func (provider *nullChunkProvider) Provide(id res.ResourceID) chunk.BlockHolder {
	return nil
}
