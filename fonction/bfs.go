package anouar

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1) // pour un graphe non orienté
}

var visited = make(map[int]bool)

func BFS(g *Graph, start int) {
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(int)
		fmt.Println(node)

		for _, neighbor := range g.vertices[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
				fmt.Println(visited)

			}
		}
	}
}

// func main() {
// 	g := NewGraph()
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 4)

// 	fmt.Println("BFS starting from node 0:")
// 	BFS(g, 0)
// }
