package core

import (
	"bytes"
	"fmt"

	"github.com/inkyblackness/res/serial"
)

type Table interface {
	Size() int
	Entry(index int) interface{}
}

type tableDataNode struct {
	parentDataNode

	table Table
}

func newTableDataNode(parentNode DataNode, id string, data []byte, table Table) *tableDataNode {
	entryCount := table.Size()
	node := &tableDataNode{
		parentDataNode: makeParentDataNode(parentNode, id, entryCount),
		table:          table}
	decoder := serial.NewPositioningDecoder(bytes.NewReader(data))
	startOffset := 0

	for i := 0; i < entryCount; i++ {
		entry := table.Entry(i)

		serial.MapData(entry, decoder)
		endOffset := int(decoder.CurPos())
		node.addChild(newBlockDataNode(node, uint16(i), data[startOffset:endOffset], entry))
		startOffset = endOffset
	}

	return node
}

func (node *tableDataNode) Info() string {
	return fmt.Sprintf("Entries: %d\n", node.table.Size())
}

func (node *tableDataNode) Data() []byte {
	entryCount := node.table.Size()
	buffer := serial.NewByteStore()
	encoder := serial.NewPositioningEncoder(buffer)

	for i := 0; i < entryCount; i++ {
		entry := node.table.Entry(i)

		serial.MapData(entry, encoder)
	}

	return buffer.Data()
}

func (node *tableDataNode) UnknownData() []byte {
	return nil
}
