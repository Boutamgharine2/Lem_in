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

func main() {

	

	var Liason []string
	var Romme []string
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
	//insect := str1[0]
	for i := 1; i < len(str1)-1; i++ {

		if strings.Contains(str1[i], "-") {
			Liason = append(Liason, str1[i])
		}
		Romm = append(Romm, anouar.Roms(str1[i]))

	}
	for i := 0; i < len(Romm); i++ {
		if Romm[i] != "" {
			Romme = append(Romme, Romm[i])
		}
	}
	// fmt.Println(insect)
	// fmt.Println(Romme)
	 fmt.Println(Liason)
	test:=&Graph{}
	for i:=0; i < len(Romme) ; i++ {
		test.AddVertex(Romme[i])
	}
	for i:=0 ;i<len(Liason);i++ {
		Tab := Supartion(Liason[i])
		test.AddEdge(Tab[0],Tab[1])
	}

	test.Print()
}


func Supartion(s string) []string { 
	var T [] string
	c:=""
	for i:=0 ;i<len(s);i++ {
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
