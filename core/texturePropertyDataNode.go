package core

import (
	"fmt"

	"github.com/inkyblackness/res"
)

type texturePropertyDataNode struct {
	rawDataNode
}

func newTexturePropertyDataNode(parentNode DataNode, id res.TextureID, data []byte) *texturePropertyDataNode {
	node := &texturePropertyDataNode{rawDataNode{
		parentNode: parentNode,
		id:         fmt.Sprintf("%d", id),
		data:       data}}

	return node
}

func (node *texturePropertyDataNode) Info() string {
	info := ""

	return info
}
