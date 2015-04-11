package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/textprop"
)

type texturePropertiesDataNode struct {
	parentDataNode
}

func NewTexturePropertiesDataNode(parentNode DataNode, name string, provider textprop.Provider) DataNode {
	node := &texturePropertiesDataNode{
		parentDataNode: makeParentDataNode(parentNode, name, int(provider.TextureCount()))}

	for i := uint32(0); i < provider.TextureCount(); i++ {
		id := res.TextureID(i)
		node.addChild(newTexturePropertyDataNode(node, id, provider.Provide(id)))
	}

	return node
}

func (node *texturePropertiesDataNode) Info() string {
	info := fmt.Sprintf("Textures available: %d", len(node.Children()))

	return info
}
