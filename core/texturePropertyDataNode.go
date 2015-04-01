package core

import (
	"fmt"

	"github.com/inkyblackness/res"
)

type texturePropertyDataNode struct {
	parentNode DataNode
	id         res.TextureID

	data []byte
}

func newTexturePropertyDataNode(parentNode DataNode, id res.TextureID, data []byte) *texturePropertyDataNode {
	node := &texturePropertyDataNode{
		parentNode: parentNode,
		id:         id,
		data:       data}

	return node
}

func (node *texturePropertyDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *texturePropertyDataNode) Info() string {
	info := ""

	return info
}

func (node *texturePropertyDataNode) ID() string {
	return fmt.Sprintf("%d", node.id)
}

func (node *texturePropertyDataNode) Resolve(path string) DataNode {
	return nil
}

func (node *texturePropertyDataNode) Data() []byte {
	return node.data
}
