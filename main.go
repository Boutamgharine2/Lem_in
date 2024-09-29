package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	anouar "anouar/fonction"
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
		err := fmt.Errorf("%v==>%v  est deja exist!", from, to)
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

func Supartion(s string) []string {
	var T []string
	c := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '-' {
			c += string(s[i])
		} else {
			T = append(T, c)
			c = ""
		}
	}
	T = append(T, c)
	return T
}

// func BFS(g *Vertix, start string) {
// 	visited := make(map[string]bool)
// 	queue := list.New()

// 	visited[start] = true
// 	queue.PushBack(start)

// 	for queue.Len() > 0 {
// 		element := queue.Front()
// 		queue.Remove(element)
// 		node := element.Value.(int)
// 		fmt.Println(node)

// 		for _, neighbor := range g.adjacenlist {
// 			if !visited[neighbor.nam] {
// 				visited[neighbor.nam] = true
// 				queue.PushBack(neighbor)
// 				fmt.Println(visited)

// 			}
// 		}
// 	}
// }

// func (g *Graph) BFS(start, end string) []string {
// var visited = make(map[string]bool)
// 	order := []string{}
// 	queue := list.New()

// 	// Trouver le sommet de départ
// 	var startVertex *Vertix
// 	for _, v := range g.Vertices {
// 		if v.nam == start {
// 			startVertex = v
// 			break
// 		}
// 	}

// 	if startVertex == nil {
// 		return order // Retourne vide si le sommet de départ n'est pas trouvé
// 	}

// 	// Ajouter le sommet de départ à la queue
// 	queue.PushBack(startVertex)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		element := queue.Front()
// 		queue.Remove(element)
// 		current := element.Value.(*Vertix)

// 		// Ajouter le sommet courant à l'ordre
// 		order = append(order, current.nam)

// 		// Vérifier si nous avons atteint le sommet de fin
// 		if current.nam == end {
// 			return order // Retourner l'ordre si fin atteinte
// 		}

// 		// Vérifier les voisins
// 		for _, neighbor := range current.adjacenlist {
// 			if !visited[neighbor.nam] {
// 				visited[neighbor.nam] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}

//		return nil // Retourner l'ordre même si le sommet de fin n'est pas atteint
//	}

func (g *Graph) dfs(start string, end string, visited map[string]bool, path []string, paths *[][]string) {
	visited[start] = true
	path = append(path, start)

	if start == end {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*paths = append(*paths, tmp)
	}
	s := g.GetVertex(start)
	for _, neigh := range s.adjacenlist {
		if !visited[neigh.nam] {
			g.dfs(neigh.nam, end, visited, path, paths)
		}
	}
	visited[start] = false
}

func main() {
	var Edges []string
	var vertexe []string
	var Romm []string
	v := os.Args
	if len(v) != 2 {
		log.Fatal("invalid Arguments!")
	}
	file, err := os.ReadFile(v[1])
	if err != nil {
		log.Fatal(err)
	}
	str := string(file)
	str1 := strings.Split(str, "\n")
	// insect := str1[0]
	for i := 1; i < len(str1)-1; i++ {

		if strings.Contains(str1[i], "-") {
			Edges = append(Edges, str1[i])
		}
		Romm = append(Romm, anouar.Roms(str1[i]))

	}
	for i := 0; i < len(Romm); i++ {
		if Romm[i] != "" {
			vertexe = append(vertexe, Romm[i])
		}
	}
	// fmt.Println(insect)
	// fmt.Println(vertexe)
	test := &Graph{}
	for i := 0; i < len(vertexe); i++ {
		test.AddVertex(vertexe[i])
	}
	for i := 0; i < len(Edges); i++ {
		Tab := Supartion(Edges[i])
		test.AddEdge(Tab[0], Tab[1])
	}

	// fmt.Println(Romm[1])
	// BFS(test,Romm[1])
	// test.Print()
	// fmt.Println(Romm[1])
	fmt.Println(vertexe[len(vertexe)-1])
	fmt.Println(vertexe[0])
	visited := make(map[string]bool)

	var (
		path  []string
		paths [][]string
	)
	test.dfs(vertexe[0], vertexe[len(vertexe)-1], visited, path, &paths)
	fmt.Println(paths)
}
