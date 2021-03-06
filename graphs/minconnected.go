// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

type state byte

const (
	undiscovered state = iota
	discovered
	processed
)

// Vertex represents a Graph vertex with its connection edges to another vertices.
type Vertex struct {
	label string
	state state // Default is undiscovered.
	edges []*Vertex
}

func hasCycle(curr, prev *Vertex) bool {
	if curr.state == discovered { // Base case.
		return true
	}

	curr.state = discovered // In process.
	for _, next := range curr.edges {
		if next != prev && next.state != processed && hasCycle(next, curr) {
			return true
		}
	}
	curr.state = processed // Done.
	return false
}

// IsMinimallyConnected returns true if graph is minimally connected.
// The time complexity is O(v+e) where v is the number of vertices and e is the
// number of edges. However, if given graph is an undirected graph with no cycles
// then the time complexity is O(v). The O(v) additional space is needed.
func IsMinimallyConnected(graph []*Vertex) bool {
	if len(graph) == 0 {
		return true
	}
	return !hasCycle(graph[0], nil)
}
