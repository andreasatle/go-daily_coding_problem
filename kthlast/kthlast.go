// Package kthlast contains Daily Coding Problem #26
//
// This problem was asked by Google.
//
// Given a singly linked list and an integer k, remove the kth last element
// from the list. k is guaranteed to be smaller than the length of the list.
//
// The list is very long, so making more than one pass is prohibitively expensive.
//
// Do this in constant space and in one pass.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package kthlast

// Node contains a node in a linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList contains a pointer to the first node
type LinkedList struct {
	List *Node
}

// RemoveKthLast removes the kth last node from a linked list
func (ll *LinkedList) RemoveKthLast(k int) *LinkedList {
	if k <= 0 {
		return ll
	}
	ll.removeKthLast(k, 0, ll.List, ll.List, ll.List)
	return ll
}

// Slice returns a slice with the values of the linked list.
func (ll *LinkedList) Slice() []int {
	return ll.List.slice([]int{})
}

func (ll *LinkedList) removeKthLast(k, cnt int, currN, prevK, currK *Node) {
	if currN == nil && cnt < k {
		return
	}
	if currN == nil && cnt == k {
		ll.List = currK.Next
		return
	}
	if currN == nil {
		prevK.Next = currK.Next
		return
	}

	if cnt >= k {
		ll.removeKthLast(k, cnt+1, currN.Next, currK, currK.Next)
	} else {
		ll.removeKthLast(k, cnt+1, currN.Next, prevK, currK)
	}

}

func (n *Node) slice(vals []int) []int {
	if n == nil {
		return vals
	}
	return n.Next.slice(append(vals, n.Value))
}
