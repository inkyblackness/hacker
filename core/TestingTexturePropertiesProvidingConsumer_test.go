package core

import (
	"github.com/inkyblackness/res"
)

type TestingTexturePropertiesProvidingConsumer struct {
	textureData [][]byte
}

// TextureCount returns the amount of textures available
func (provider *TestingTexturePropertiesProvidingConsumer) TextureCount() uint32 {
	return uint32(len(provider.textureData))
}

// Provide returns the data for the requested TextureID.
func (provider *TestingTexturePropertiesProvidingConsumer) Provide(id res.TextureID) []byte {
	return provider.textureData[int(id)]
}

// Consume adds to the list
func (provider *TestingTexturePropertiesProvidingConsumer) Consume(data []byte) {
	provider.textureData = append(provider.textureData, data)
}

func (provider *TestingTexturePropertiesProvidingConsumer) Finish() {

}
