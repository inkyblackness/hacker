package core

type dataNode interface {
	parent() dataNode

	info() string

	resolve(string) dataNode
}
