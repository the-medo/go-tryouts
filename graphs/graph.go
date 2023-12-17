package graphs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Graph(t *testing.T) {

	var adjacencyList = map[int][]int{
		1: {2, 4},
		2: {3, 5, 1},
		3: {6, 2},
		4: {1, 5, 7},
		5: {2, 6, 8, 4},
		6: {3, 0, 9, 5},
		7: {4, 8},
		8: {5, 9, 7},
		9: {6, 0, 8},
	}

	g := NewGraph(WithAdjacencyList(adjacencyList))

	for key, list := range adjacencyList {
		require.ElementsMatch(t, list, g.Neighbors(key))
	}
}

// NewGraph returns a new graph.
func NewGraph(opts ...GraphOption) *Graph {
	g := &Graph{Vertices: map[int]*Vertex{}}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// GraphOption is a functional option for the graph constructor.
type GraphOption func(this *Graph)

// WithAdjacencyList is a graph option to initialize the graph with an adjacency list.
func WithAdjacencyList(list map[int][]int) GraphOption {
	return func(this *Graph) {
		for vertex, edges := range list {
			// add vertex
			if _, ok := this.Vertices[vertex]; !ok {
				this.AddVertex(vertex, vertex)
			}

			// add edges to vertex
			for _, edge := range edges {
				// add edge as vertex, if not added
				if _, ok := this.Vertices[edge]; !ok {
					this.AddVertex(edge, edge)
				}

				this.AddEdge(vertex, edge, 0) // no weights in this adjacency list
			}
		}
	}
}

// Graph represents a set of vertices connected by edges.
type Graph struct {
	Vertices map[int]*Vertex
}

// Vertex is a node in the graph that stores the int value at that node
// along with a map to the vertices it is connected to via edges.
type Vertex struct {
	Val   int
	Edges map[int]*Edge
}

// Edge represents an edge in the graph and the destination vertex.
type Edge struct {
	Weight int
	Vertex *Vertex
}

// AddVertex adds a vertex to the graph with no edges.
func (this *Graph) AddVertex(key, val int) {
	this.Vertices[key] = &Vertex{Val: val, Edges: map[int]*Edge{}}
}

// AddEdge adds an edge between existing source and existing destination vertex.
func (this *Graph) AddEdge(srcKey, destKey int, weight int) {
	// check if src & dest exist
	if _, ok := this.Vertices[srcKey]; !ok {
		return
	}
	if _, ok := this.Vertices[destKey]; !ok {
		return
	}

	// add edge src --> dest
	this.Vertices[srcKey].Edges[destKey] = &Edge{Weight: weight, Vertex: this.Vertices[destKey]}
}

// Neighbors returns all vertex values that have an edge from
// the provided src vertex.
func (this *Graph) Neighbors(srcKey int) []int {
	result := []int{}

	if _, ok := this.Vertices[srcKey]; !ok {
		return result
	}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.Val)
	}

	return result
}
