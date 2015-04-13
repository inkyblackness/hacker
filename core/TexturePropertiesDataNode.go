package core

import (
	"fmt"
	"strings"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/textprop"
)

type textpropConsumerFactory func() textprop.Consumer

type texturePropertiesDataNode struct {
	parentDataNode
	consumerFactory textpropConsumerFactory
}

func NewTexturePropertiesDataNode(parentNode DataNode, name string,
	provider textprop.Provider, consumerFactory textpropConsumerFactory) DataNode {
	node := &texturePropertiesDataNode{
		parentDataNode:  makeParentDataNode(parentNode, strings.ToLower(name), int(provider.TextureCount())),
		consumerFactory: consumerFactory}

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

func (node *texturePropertiesDataNode) save() string {
	consumer := node.consumerFactory()
	defer consumer.Finish()

	for _, child := range node.Children() {
		consumer.Consume(child.Data())
	}

	return node.ID() + "\n"
}
