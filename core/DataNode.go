package core

type dataNode interface {
	parent() dataNode

	info() string

	id() string

	resolve(string) dataNode
}
