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

func (ge *graphEdge) ShortestPath(searchedValue int) []int {
	previous := solveShortestPath(ge, searchedValue)
	path := reconstructShortestPath(previous, searchedValue, ge.value)
	return path
}

func solveShortestPath(ge *graphEdge, searchedValue int) map[int]int {
	if ge.value == searchedValue {
		return nil
	}

	queue := Queue.NewQueue()
	queue.Enqueue(ge)

	previous := map[int]int{}
	visited := map[int]bool{}

	previous[ge.value] = 0x7FFFFFFFFFFFFFFF //max int
	for queue.Count > 0 {
		edge := queue.Dequeue().(*graphEdge)
		visited[edge.value] = true
		if edge.value == searchedValue {
			return previous
		}
		for _, neighbour := range edge.neighbours {
			_, found := visited[neighbour.value]
			if !found && !queue.Contains(neighbour.value) {
				if _, found := previous[neighbour.value]; !found {
					previous[neighbour.value] = edge.value
				}
				queue.Enqueue(neighbour)
			}
		}
	}

	return previous
}

func reconstructShortestPath(previous map[int]int, searchedValue int, startNode int) []int {
	path := []int{}
	at := searchedValue

	for {
		_, found := previous[at]
		if !found {
			break
		}
		path = append(path, at)
		at = previous[at]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if path[0] == startNode {
		return path
	}

	return path
}
