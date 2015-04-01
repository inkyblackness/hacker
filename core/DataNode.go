package core

// DataNode represents a container with data.
type DataNode interface {
	// Parent returns the parent node or nil if none known.
	Parent() DataNode
	// Info returns human readable information about this node.
	Info() string
	// ID returns the identification for this node. The returned value must be
	// the same by which the parent resolves this node.
	ID() string
	// Resolve returns a DataNode this node knows for the given ID.
	Resolve(string) DataNode
}
