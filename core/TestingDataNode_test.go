package core

type TestingDataNode struct {
	id   string
	data []byte
}

func NewTestingDataNode(id string) *TestingDataNode {
	node := &TestingDataNode{id: id}

	return node
}

// Parent returns the parent node or nil if none known.
func (node *TestingDataNode) Parent() DataNode {
	return nil
}

// Children returns no children.
func (node *TestingDataNode) Children() []DataNode {
	return nil
}

// Info returns human readable information about this node.
func (node *TestingDataNode) Info() string {
	return "testing"
}

// ID returns the identification for this node. The returned value must be
// the same by which the parent resolves this node.
func (node *TestingDataNode) ID() string {
	return node.id
}

// Resolve returns a DataNode this node knows for the given ID.
func (node *TestingDataNode) Resolve(string) DataNode {
	return nil
}

func (node *TestingDataNode) Data() []byte {
	return node.data
}
