package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/objprop"
)

type objectPropertiesDataNode struct {
	parentDataNode
}

func NewObjectPropertiesDataNode(parentNode DataNode, name string,
	provider objprop.Provider, classes []objprop.ClassDescriptor) DataNode {
	node := &objectPropertiesDataNode{parentDataNode: makeParentDataNode(parentNode, name, 0)}

	for classIndex, classDesc := range classes {
		for subclassIndex, subclassDesc := range classDesc.Subclasses {
			for typeIndex := uint32(0); typeIndex < subclassDesc.TypeCount; typeIndex++ {
				id := res.MakeObjectID(res.ObjectClass(classIndex), res.ObjectSubclass(subclassIndex), res.ObjectType(typeIndex))
				subnode := newObjectPropertyDataNode(node, id, provider)
				node.addChild(subnode)
			}
		}
	}

	return node
}

func (node *objectPropertiesDataNode) Info() string {
	info := fmt.Sprintf("Objects available: %d", len(node.Children()))

	return info
}
