package core

import (
	"fmt"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/objprop"
)

type objectPropertiesDataNode struct {
	parentNode DataNode
	name       string

	propertyDataNodes map[string]*objectPropertyDataNode
}

func NewObjectPropertiesDataNode(parentNode DataNode, name string,
	provider objprop.Provider, classes []objprop.ClassDescriptor) DataNode {
	node := &objectPropertiesDataNode{parentNode: parentNode,
		name:              name,
		propertyDataNodes: make(map[string]*objectPropertyDataNode)}

	for classIndex, classDesc := range classes {
		for subclassIndex, subclassDesc := range classDesc.Subclasses {
			for typeIndex := uint32(0); typeIndex < subclassDesc.TypeCount; typeIndex++ {
				id := res.MakeObjectID(res.ObjectClass(classIndex), res.ObjectSubclass(subclassIndex), res.ObjectType(typeIndex))
				subnode := newObjectPropertyDataNode(node, id, provider)
				node.propertyDataNodes[subnode.ID()] = subnode
			}
		}
	}

	return node
}

func (node *objectPropertiesDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *objectPropertiesDataNode) Info() string {
	info := fmt.Sprintf("Objects available: %d", len(node.propertyDataNodes))

	return info
}

func (node *objectPropertiesDataNode) ID() string {
	return node.name
}

func (node *objectPropertiesDataNode) Resolve(path string) (resolved DataNode) {
	temp, existing := node.propertyDataNodes[path]

	if existing {
		resolved = temp
	}

	return
}

func (node *objectPropertiesDataNode) Data() []byte {
	return nil
}
