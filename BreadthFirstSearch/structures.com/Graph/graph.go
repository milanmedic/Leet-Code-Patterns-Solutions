package graph

import (
	Queue "bfs.com/m/v2/structures.com/Queue"
)

type graphEdge struct {
	value      int
	visited    bool
	neighbours map[int]*graphEdge
}

func NewGraphNode(value int) *graphEdge {
	return &graphEdge{value: value, neighbours: map[int]*graphEdge{}}
}

func (ge *graphEdge) AddNeighbour(neighbour *graphEdge) {
	if !ge.Contains(neighbour) {
		ge.neighbours[neighbour.value] = neighbour
	}
}

func (ge *graphEdge) Contains(neighbour *graphEdge) bool {
	if _, found := ge.neighbours[neighbour.value]; found {
		return true
	}
	return false
}

func (ge *graphEdge) BreadthFirstSearch(searchedNode int) *graphEdge {
	if ge.value == searchedNode {
		return ge
	}

	queue := Queue.NewQueue()
	queue.Enqueue(ge)

	for queue.Count > 0 {
		edge := queue.Dequeue().(*graphEdge)
		edge.visited = true
		if edge.value == searchedNode {
			return edge
		}
		for _, neighbour := range edge.neighbours {
			if !neighbour.visited && !queue.Contains(neighbour.value) {
				queue.Enqueue(neighbour)
			}
		}
	}

	return nil
}
