// Package palin contains Daily Coding Problem #104
//
// This problem was asked by Google.
//
// Determine whether a doubly linked list is a palindrome.
// What if it’s singly linked?
//
// For example,
//  1 -> 4 -> 3 -> 4 -> 1
// returns True while
//  1 -> 4
// returns False.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package palin

import "fmt"

// Value is the type of the entry in the double linked list
type Value int

// DoubleNode is an entry in a double linked list
type DoubleNode struct {
	Val  Value
	Prev *DoubleNode
	Next *DoubleNode
}

// DoubleLinkedList is a double linked list (duh!)
type DoubleLinkedList struct {
	First *DoubleNode
	Last  *DoubleNode
}

// NewDoubleNode returns a pointer to a new DoubleNode
func NewDoubleNode(val Value) *DoubleNode {
	return &DoubleNode{val, nil, nil}
}

// NewDoubleLinkedList returns a pointer to a new double linked list
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// InsertFirst inserts a value first in the double linked list
func (list *DoubleLinkedList) InsertFirst(val Value) {
	node := NewDoubleNode(val)
	node.Next = list.First
	list.First = node

	if node.Next != nil {
		node.Next.Prev = node
	}

	if list.Last == nil {
		list.Last = node
	}
}

// InsertLast inserts a value last in the double linked list
func (list *DoubleLinkedList) InsertLast(val Value) {
	node := NewDoubleNode(val)
	node.Prev = list.Last
	list.Last = node

	if node.Prev != nil {
		node.Prev.Next = node
	}

	if list.First == nil {
		list.First = node
	}
}

// Print output the double linked list to screen
func (list *DoubleLinkedList) Print() {
	for curr := list.First; curr != nil; curr = curr.Next {
		fmt.Println("Value:", curr.Val)
	}
}

// PrintReverse output the double linked list to screen in reverse order
func (list *DoubleLinkedList) PrintReverse() {
	for curr := list.Last; curr != nil; curr = curr.Prev {
		fmt.Println("Reverse Value:", curr.Val)
	}
}

// Palindrome returns true if the values make a palindrome
func (list *DoubleLinkedList) Palindrome() bool {
	for first, last := list.First, list.Last; ; first, last = first.Next, last.Prev {
		if first == last {
			return true
		}
		if first.Val != last.Val {
			return false
		}
	}
}

// PalindromeSingle returns true if the values make a palindrome, without
// using list.Last or node.Prev for nodes in the list.
func (list *DoubleLinkedList) PalindromeSingle() bool {
	var last *DoubleNode = nil
	for first := list.First; ; first = first.Next {
		for curr := first; ; curr = curr.Next {
			if curr.Next == last {
				last = curr
				break
			}
		}
		if first == last {
			return true
		}
		if first.Val != last.Val {
			return false
		}
	}
}
