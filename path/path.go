// Package path contains Daily Coding Problem #72
//
// This problem was asked by Google.
//
// In a directed graph, each node is assigned an uppercase letter.
// We define a path's value as the number of most frequently-occurring
// letter along that path. For example, if a path in the graph goes
// through "ABACA", the value of the path is 3, since there are
// 3 occurrences of 'A' on the path.
//
// Given a graph with n nodes and m directed edges, return the largest value
// path of the graph. If the largest value is infinite, then return null.
//
// The graph is represented with a string and an edge list. The i-th character
// represents the uppercase letter of the i-th node. Each tuple in the
// edge list (i, j) means there is a directed edge from the i-th node to the j-th node.
// Self-edges are possible, as well as multi-edges.
//
// For example, the following input graph:
//
//  ABACA
//  [(0, 1),
//   (0, 2),
//   (2, 3),
//   (3, 4)]
// Would have maximum value 3 using the path of vertices
//  [0, 2, 3, 4], (A, A, C, A).
//
// The following input graph:
//
//  A
//  [(0, 0)]
// Should return null, since we have an infinite loop.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package path

import (
	"errors" // Empty implements a set
	"math"
)

// Empty implements a set
type Empty struct{}

// VertexSet is a set of vertices
type VertexSet map[int]Empty

// AltGraph is a different representation of the graph
type AltGraph map[int]VertexSet

// ValueCounter counts the occurancies of each rune-value in the graph-traversal
type ValueCounter map[rune]int

// Infinity is the largest possible integer (null doesn't exist in Go)
const Infinity = math.MaxInt64

// Edge contains an edge of a directed graph
type Edge struct {
	u, v int
}

// Vertices contains a string with all the values at the vertices
type Vertices string

// Graph contains the definition of a directed graph
type Graph struct {
	V Vertices
	E []Edge
}

// NewGraph creates a new instance of Graph
func NewGraph(V Vertices, E []Edge) *Graph {
	E2 := make([]Edge, len(E))
	copy(E2, E)
	g := &Graph{V: V, E: E2}
	return g
}

// NewVertex creates a new graph without any edges
func NewVertex(V Vertices) *Graph {
	g := &Graph{V: V}
	return g
}

// AddEdge adds an edge to the Graph
func (g *Graph) AddEdge(u, v int) *Graph {
	g.E = append(g.E, Edge{u, v})
	return g
}

// maxVertexSet returns a set of vertices with edges in the graph
func (g *Graph) maxVertexSet() VertexSet {
	V := VertexSet{}
	for i := 0; i < len(g.E); i++ {
		V[g.E[i].u] = Empty{}
		V[g.E[i].v] = Empty{}
	}
	return V
}

// startVertices returns the set of vertices that has no incoming edges
func (g *Graph) startVertices() VertexSet {
	startV := g.maxVertexSet()
	for _, e := range g.E {
		delete(startV, e.v)
	}
	return startV
}

// CreateAltGraph creates an alternative represntation of the graph using sets
func (g *Graph) CreateAltGraph() AltGraph {
	altGraph := AltGraph{}
	for _, e := range g.E {
		if _, ok := altGraph[e.u]; !ok {
			altGraph[e.u] = VertexSet{}
		}
		altGraph[e.u][e.v] = Empty{}
	}
	return altGraph
}

// LargestValueForStartValue returns the values of the graph starting at vertex v
func (g *Graph) LargestValueForStartValue(v int, vertSet VertexSet, altGraph AltGraph, valCounter ValueCounter) int {

	// Check for cyclic graph
	if _, ok := vertSet[v]; ok {
		return Infinity
	}

	// Tag vertex v as used
	vertSet[v] = Empty{}
	valCounter[rune(g.V[v])]++

	maxVal := 0

	// Count values when last vertex
	if len(altGraph[v]) == 0 {
		for _, val := range valCounter {
			maxVal = max(maxVal, val)
		}
		return maxVal
	}

	// Recursion for next vertex in the graph
	for nextVert := range altGraph[v] {
		maxVal = max(maxVal, g.LargestValueForStartValue(nextVert, vertSet, altGraph, valCounter))
		valCounter[rune(g.V[nextVert])]--
	}
	return maxVal
}

// LargestValue computes the largest value for any starting vertex
func (g *Graph) LargestValue() (int, error) {
	maxVal := 0
	altGraph := g.CreateAltGraph()
	startV := g.startVertices()
	if len(altGraph) > 0 && len(startV) == 0 {
		return Infinity, errors.New("cyclic graph")
	}
	for v := range startV {
		maxVal = max(maxVal, g.LargestValueForStartValue(v, VertexSet{}, altGraph, ValueCounter{}))
	}

	if maxVal == Infinity {
		return Infinity, errors.New("cyclic graph")
	}
	return maxVal, nil
}

// max is a help function returning max(a,b)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
