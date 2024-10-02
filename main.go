package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

func Rougroupe(allPaths [][]string) map[int][][]string {
	res := make(map[int][][]string)
	indix := 0
	for _, path := range allPaths {
		passed := false
		if len(res) == 0 {
			res[indix] = append(res[indix], path)
		} else {
			for i, way := range res {
				if !HandulWay(way, path) { // comparer entre le tableau actuel est le tableau de l'indice i dans la cart
					res[i] = append(res[i], path)
					passed = true
				}
			}
			if !passed {
				indix++
				res[indix] = append(res[indix], path) // crier autre indice de la cart pour stocker la nouvel parcour
			}
		}
	}

	for _, Paths := range allPaths {
		for i, r := range res {
			if !HandulWay(r, Paths) {
				res[i] = append(res[i], Paths)
			}
		}
	}
	return res
}

func HandulWay(Paths [][]string, way []string) bool { // comparer entre un tableau et un tableau bidimentionel si il n'y a pas un element commine entre ce tableau et les tableau de [][]
	for _, t := range Paths {
		if !Com2Tab(t, way) { // il ya un element commun
			return true
		}
	}
	return false
}

func Com2Tab(path1, path2 []string) bool { // comparer entre deux tableau et verifie si il partage un element si le cas la fonction retourn  false
	rooms1 := make(map[string]bool)
	if len(path2) == 2 && len(path1) == 2 {
		return false
	}
	for _, room := range path1[1 : len(path1)-1] {
		rooms1[room] = true
	}
	for _, room := range path2[1 : len(path2)-1] {
		if rooms1[room] {
			return false
		}
	}
	return true
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
	for i := 1; i < len(str1); i++ {

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
	test := &Graph{}
	for i := 0; i < len(vertexe); i++ {
		test.AddVertex(vertexe[i])
	}
	for i := 0; i < len(Edges); i++ {
		Tab := strings.Split(Edges[i], "-")
		test.AddEdge(Tab[0], Tab[1])
	}

	visited := make(map[*Vertix]bool)
	start := test.GetVertex(vertexe[0])
	end := test.GetVertex(vertexe[len(vertexe)-1])
	var paths [][]string

	test.dfs(start, end, visited, []string{}, &paths)
	d := Rougroupe(paths)
	r := []string{"1","2"}
	b:= FindPaths(d)
	b = append(b, r)

	fmt.Println(Chois(b,1))
}

func FindPaths(m map[int][][]string) [][]string {
	if len(m) == 0 {
		return nil
	}

	maxlin := 0
	for i := range m {
		if len(m[i]) > len(m[maxlin]) {
			maxlin = i
		}
	}

	return m[maxlin]
}



func Chois(s [][]string, n int) [][]string {
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})
	




	
}

// func Parcour(slice [][]string, n int) {
// 	pp := make([][]string, len(slice))
// 	for i := 0; i < len(slice); i++ {
// 		pp[i] = make([]string, len(slice[i]))
// 	}
// 	j := 1
// 	for {
// 		move(pp)
// 		for i := 0; i < len(pp); i++ {
// 			if j <= n {
// 				pp[i][0] = "L" + strconv.Itoa(j)
// 				j++
// 			}
// 		}
// 		if checknil(pp) {
// 			break
// 		}
// 		for i := 0; i < len(pp); i++ {
// 			for j := 0; j < len(pp[i]); j++ {
// 				if pp[i][j] != "" {
// 					fmt.Print(pp[i][j], "-", slice[i][j], " ")
// 				}
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func checknil(pp [][]string) bool {
// 	for i := 0; i < len(pp); i++ {
// 		for j := 0; j < len(pp[i]); j++ {
// 			if pp[i][j] != "" {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func move(arr [][]string) [][]string {
// 	for i := 0; i < len(arr); i++ {
// 		for j := len(arr[i]) - 1; j > 0; j-- {
// 			arr[i][j] = arr[i][j-1]
// 		}
// 		arr[i][0] = ""
// 	}
// 	return arr
// }
