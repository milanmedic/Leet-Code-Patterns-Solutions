package main

import (
	"fmt"

	Graph "bfs.com/m/v2/structures.com/Graph"
)

func main() {
	gn0 := Graph.NewGraphNode(0)
	gn1 := Graph.NewGraphNode(1)
	gn2 := Graph.NewGraphNode(2)
	gn3 := Graph.NewGraphNode(3)
	gn4 := Graph.NewGraphNode(4)

	gn0.AddNeighbour(gn1)
	gn0.AddNeighbour(gn2)
	gn0.AddNeighbour(gn3)
	gn1.AddNeighbour(gn0)
	gn1.AddNeighbour(gn3)
	gn2.AddNeighbour(gn0)
	gn2.AddNeighbour(gn3)
	gn3.AddNeighbour(gn1)
	gn3.AddNeighbour(gn2)
	gn3.AddNeighbour(gn4)
	gn3.AddNeighbour(gn0)
	gn4.AddNeighbour(gn3)

	//foundNode := gn0.BreadthFirstSearch(4)
	//fmt.Println(foundNode)
	path := gn0.ShortestPath(4)
	fmt.Println(path)

	fmt.Println("Endee")
}
