package graph

type graphEdge struct {
	value      int
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

func (ge *graphEdge) DepthFirstSearch() map[int]bool {
	var visited map[int]bool = make(map[int]bool)

	return solveDFS(ge, visited)
}

func solveDFS(ge *graphEdge, visited map[int]bool) map[int]bool {
	if _, found := visited[ge.value]; !found {
		visited[ge.value] = true

		for _, neighbour := range ge.neighbours {
			if _, found := visited[neighbour.value]; !found {
				solveDFS(neighbour, visited)
			}
		}
	}
	return visited
}
