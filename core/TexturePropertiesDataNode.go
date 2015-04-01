package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/textprop"
)

type texturePropertiesDataNode struct {
	parentNode DataNode
	name       string

	propertyDataNodes []*texturePropertyDataNode
}

func NewTexturePropertiesDataNode(parentNode DataNode, name string, provider textprop.Provider) DataNode {
	node := &texturePropertiesDataNode{parentNode: parentNode,
		name:              name,
		propertyDataNodes: make([]*texturePropertyDataNode, provider.TextureCount())}

	for i := uint32(0); i < provider.TextureCount(); i++ {
		id := res.TextureID(i)
		node.propertyDataNodes[i] = newTexturePropertyDataNode(node, id, provider.Provide(id))
	}

	return node
}

func (node *texturePropertiesDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *texturePropertiesDataNode) Info() string {
	info := fmt.Sprintf("Textures available: %d", len(node.propertyDataNodes))

	return info
}

func (node *texturePropertiesDataNode) ID() string {
	return node.name
}

func (node *texturePropertiesDataNode) Resolve(path string) (resolved DataNode) {
	for _, temp := range node.propertyDataNodes {
		if temp.ID() == path {
			resolved = temp
		}
	}

	return
}

func (node *texturePropertiesDataNode) Data() []byte {
	return nil
}
