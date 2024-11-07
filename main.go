package main

import (
	"fmt"
	"strconv"
	"strings"

	Lemin "Lemin/fonction"
)

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
	nam         string
	adjacenlist []*Vertix
}

// AddVertex  ajouter une noude au graph
func (g *Graph) AddVertex(key string) {
	if !Contain(g.Vertices, key) {
		g.Vertices = append(g.Vertices, &Vertix{nam: key})
	} else {
		err := fmt.Errorf("la noude %v est deja exist! ", key)
		fmt.Println(err)
	}
}

// AddEdge add an edge to the graphe

func (g *Graph) AddEdge(from, to string) {
	// Get Vertex
	FromVertex := g.GetVertex(from)
	ToVertex := g.GetVertex(to)

	// check err
	if FromVertex == nil || ToVertex == nil {
		err := fmt.Errorf("invalid Edge %v==>%v", from, to)
		fmt.Println(err)
	} else if Contain(FromVertex.adjacenlist, to) {
		err := fmt.Errorf("%v==>%v  est deja exist", from, to)
		fmt.Println(err)

	} else {
		FromVertex.adjacenlist = append(FromVertex.adjacenlist, ToVertex)
		ToVertex.adjacenlist = append(ToVertex.adjacenlist, FromVertex)
	}
}

func (G *Graph) GetVertex(k string) *Vertix {
	for i, v := range G.Vertices {
		if k == v.nam {
			return G.Vertices[i]
		}
	}
	return nil
}

func Contain(s []*Vertix, key string) bool {
	for _, v := range s {
		if key == v.nam {
			return true
		}
	}
	return false
}

func (G *Graph) Print() {
	for _, v := range G.Vertices {
		fmt.Printf("vertice : %s \n", v.nam)
		for _, v := range v.adjacenlist {
			fmt.Printf(" %s \n", v.nam)
		}
	}
}

func (g *Graph) dfs(start, end *Vertix, visited map[*Vertix]bool, path []string, paths *[][]string) {
	visited[start] = true
	path = append(path, start.nam)

	if start == end {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*paths = append(*paths, tmp)
	}
	for _, neigh := range start.adjacenlist {
		if !visited[neigh] {
			g.dfs(neigh, end, visited, path, paths)
		}
	}
	delete(visited, start)
}

func main() {
	var paths [][]string
	vertexe, Edges, insects, str, start, end := Lemin.Handlfile()
	test := &Graph{}
	for i := 0; i < len(vertexe); i++ {
		test.AddVertex(vertexe[i])
	}
	for i := 0; i < len(Edges); i++ {
		Tab := strings.Split(Edges[i], "-")
		test.AddEdge(Tab[0], Tab[1])
	}

	visited := make(map[*Vertix]bool)
	startV := test.GetVertex(start)
	endV := test.GetVertex(end)

	test.dfs(startV, endV, visited, []string{}, &paths)

	MapOfPaths := Lemin.Rougroupe(paths)
	BestPaths := Lemin.FindPaths(MapOfPaths)
	fmt.Println(BestPaths)

	NumberOfAnts, _ := strconv.Atoi(insects)
	AntsInPaths := Lemin.MoveAnts(NumberOfAnts, BestPaths)
	resfinal := Lemin.SortTable(AntsInPaths)
	fmt.Println(str)
	fmt.Println()
	for _, v := range resfinal {
		fmt.Println(strings.Join(v, " "))
	}
}
