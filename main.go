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
	vertexe, Edges, insects := Lemin.Handlfile()
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

	test.dfs(start, end, visited, []string{}, &paths)

	MapOfPaths := Lemin.Rougroupe(paths)
	BestPaths := Lemin.FindPaths(MapOfPaths)
	n, _ := strconv.Atoi(insects)
	moveAnts(n, BestPaths)
}

func moveAnts(numAnts int, paths [][]string) {
	var (
		resfinal []string
		matrix   [][]string
	)

	for i := 0; i < len(paths); i++ {
		for k := 0; k < numAnts; k++ {
			for j := 1; j < len(paths[i]); j++ { // Commencer à 0

				// Créer la chaîne pour chaque fourmi et chemin
				restem := "L" + TAbloOfAnts(numAnts)[k] + "-" + paths[i][j]
				resfinal = append(resfinal, restem)
			}
			matrix = append(matrix, resfinal)
			resfinal = nil
		}
	}
	tableau := (HandlTab(matrix))
	fmt.Println(tableau)

	// Afficher le résultat
	// for _, res := range tableau {
	// 	for _,rese := range res {
	// 		fmt.Print(rese)
	// 	}
	// }
}

func HandlTab(tab [][]string) [][]string {
	fmt.Println(tab)
	var checkpathee []string
	var checkformis []string
	// str := ""
	var res [][]string
	
	// lene := 0
	for i := 0; i < len(tab); i++ {

		Split := strings.Split(tab[i][0], "-")
		ant := Split[0]
		if i ==1 {
			//fmt.Println()
			fmt.Println(checkpathee)
			fmt.Println(checkformis)
			fmt.Println(valid(ant,checkformis) )
			fmt.Println(valid(ExtraitP(tab[i]),checkpathee))

		}

		
		

		if (valid(ExtraitP(tab[i]),checkpathee) && i != len(tab)-1) || valid(ant,checkformis) {
			//
			continue
		} else {
			//
		    checkant(ant, &checkformis)
			checkpathe(&checkpathee, ExtraitP(tab[i]))
			res = append(res, tab[i])

			// lene = len(tab[i])
		}
	}
	// fmt.Println(checkformis)
	// fmt.Println(checkpathee)
	return res
}

func valid(str string, tab []string) bool {
	for _, val := range tab {
		if str == val {
			return true
		}
	}
	return false
}

func ExtraitP(T []string) string {
	actuelPath := ""
	for i := 0; i < len(T); i++ {
		Split := strings.Split(T[i], "-")[1]
		actuelPath += Split

	}
	return actuelPath
}

func checkpathe(tab1 *[]string, path string ) bool {
	for _, val := range *tab1 {
		if val == path {
			return true
		}
	}

	*tab1 = append(*tab1, path)

	return false
}

func checkant(ant string, Tab *[]string) bool {
	// fmt.Println(Tab)
	for _, val := range *Tab {
		if ant == val {
			return true
		}
	}
	*Tab = append(*Tab, ant)
	return false
}

func TAbloOfAnts(ants int) []string {
	var T []string
	for i := 1; i <= ants; i++ {
		T = append(T, strconv.Itoa(i))
	}
	return T
}
