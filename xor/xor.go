// Package xor contains Daily Coding Problem #6
// This problem was asked by Google.
// An XOR linked list is a more memory efficient doubly linked list.
// Instead of each node holding next and prev fields, it holds a field
// named both, which is an XOR of the next node and the previous node.
// Implement an XOR linked list; it has an add(element) which adds the
// element to the end, and a get(index) which returns the node at index.
// If using a language that has no pointers (such as Python), you can
// assume you have access to get_pointer and dereference_pointer functions
// that converts between nodes and memory addresses.
// Note(GoLang):
// It might be very dangerous with the xor version of a linked list,
// since we are losing the reference to each element.
// It is not clear if the garbage collection will remove unlinked nodes.
// Maybe a better solution is to implement the linked list in C,
// and use cgo to make is go-ish...
// Then we would have full control!
package xor

import (
	"errors"
	"unsafe"
)

//package xor

// Value is the type of the value (int in this example)
type Value int

// UPtr is short for an unsigned int pointer
type UPtr uintptr

// Node is one node of the doubly linked list
type Node struct {
	Val  Value
	both UPtr
}

// DLinkedList is an instance of a doubly linked list
type DLinkedList struct {
	first *Node
	last  *Node
}

// uPtr returns a pointer converted to unsigned int
func (n *Node) uPtr() UPtr {
	return UPtr(unsafe.Pointer(n))
}

// nodePtr returns a pointer converted to *Node
func nodePtr(u UPtr) *Node {
	return (*Node)(unsafe.Pointer(u))
}

// xor returns n xor n2
func (n *Node) xor(n2 *Node) UPtr {
	return n.uPtr() ^ n2.uPtr()
}

// xor returns u xor n
func (u UPtr) xor(n *Node) UPtr {
	return n.uPtr() ^ u
}

// Insert a value in the doubly linked list
// Remark: The node is returned to avoid garbage collection.
//         This is VERY dangerous!
func (lst *DLinkedList) Insert(v Value) *Node {
	if lst.last == nil {
		n := &Node{v, UPtr(0)}
		lst.last = n
		lst.first = n
		return n
	}
	if lst.first == lst.last {
		n := &Node{v, lst.first.uPtr()}
		lst.first.both = n.uPtr()
		lst.last = n
		return n
	}

	n := &Node{v, lst.last.uPtr()}
	lst.last.both = lst.last.both.xor(n)
	lst.last = n
	return n
}

// NewDLinkedList returns a new instance of a doubly linked list
func NewDLinkedList() *DLinkedList {
	lst := &DLinkedList{nil, nil}
	return lst
}

// Slice makes a slice of values out of a doubly linked list
func (lst *DLinkedList) Slice() []Value {
	value := []Value{}
	var n, p *Node
	for n, p = lst.first, nil; ; n, p = nodePtr(n.both.xor(p)), n {

		value = append(value, n.Val)
		if n == lst.last {
			break
		}
	}
	return value
}

// Get returns the ith node of the linked list
func (lst *DLinkedList) Get(index int) (*Node, error) {
	i := 0
	var n, p *Node
	for n, p = lst.first, nil; ; i, n, p = i+1, nodePtr(n.both.xor(p)), n {
		if i == index {
			return n, nil
		}
		if n == lst.last {
			return nil, errors.New("index out of bounds")
		}
	}

}
