package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/objprop"
)

type objectPropertyDataNode struct {
	parentNode DataNode
	id         res.ObjectID

	properties map[string]DataNode
}

func newObjectPropertyDataNode(parentNode DataNode, id res.ObjectID, provider objprop.Provider) *objectPropertyDataNode {
	node := &objectPropertyDataNode{
		parentNode: parentNode,
		id:         id,
		properties: make(map[string]DataNode)}

	objData := provider.Provide(id)
	node.properties["generic"] = newGenericPropertyDataNode(node, objData.Generic)
	node.properties["specific"] = newSpecificPropertyDataNode(node, objData.Specific)
	node.properties["common"] = newCommonPropertyDataNode(node, objData.Common)

	return node
}

func (node *objectPropertyDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *objectPropertyDataNode) Info() string {
	info := ""

	return info
}

func (node *objectPropertyDataNode) ID() string {
	return fmt.Sprintf("%d-%d-%d", node.id.Class, node.id.Subclass, node.id.Type)
}

func (node *objectPropertyDataNode) Resolve(path string) (resolved DataNode) {
	temp, existing := node.properties[path]

	if existing {
		resolved = temp
	}

	return
}

func (node *objectPropertyDataNode) Data() []byte {
	return nil
}
