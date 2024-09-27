package main

import (
	// anouar "anouar/fonction"
	"fmt"
	// "log"
	// "os"
	// "strings"
)
type Graph struct {
	Vertices  []*Vertix
}

type Vertix struct {
    nam int
	adjacenlist []*Vertix
}


//AddVertex  ajouter une noude au graph

func (g *Graph) AddVertex(key int) {
	if !Contain(g.Vertices,key) { 
	g.Vertices = append( g.Vertices,&Vertix{nam:key})
	} else {
		err := fmt.Errorf("la noude %v est deja exist! ",key)
		fmt.Println(err)
	} 
}
func (G *Graph) Print() {
	for _,v := range G.Vertices {
		fmt.Print("\nvertice : %v ",v.nam)
		for _,v := range v.adjacenlist {
			fmt.Print(" %v ",v.nam)
		}
	}
}

// AddEdge add an edge to the graphe

func (g *Graph) AddEdge (from,to int) {

	// Get Vertex 
	FromVertex := g.GetVertex(from)
	ToVertex := g.GetVertex(to)

	//check err 
	if FromVertex == nil || ToVertex == nil {
		err := fmt.Errorf("invalid Edge %v==>%v",from,to)
		fmt.Println(err)
	} else if Contain(FromVertex.adjacenlist,to) {
		err := fmt.Errorf("%v==>%v  est deja exist!",from,to)
		fmt.Println(err)


	}else {
		FromVertex.adjacenlist = append(FromVertex.adjacenlist, ToVertex)
	}




}
func (G *Graph) GetVertex(k int) *Vertix {
	for i,v  := range  G.Vertices {
		if k == v.nam {

			 return G.Vertices[i]

		}
	}
	return nil

}
func Contain(s []*Vertix,key int)bool {

	for _,v := range s {
		if key == v.nam {
			return true
		}
	}
	return false 
}
func main () {
	test:=&Graph{}
	for i:=0 ; i < 5 ; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1,3)
	test.AddEdge(1,3)
	test.AddEdge(1,3)
	test.AddEdge(1,9)
	test.AddEdge(1,96)
	test.AddVertex(0)
	test.AddVertex(1)
	test.AddVertex(10)
	test.AddVertex(3)
	test.AddVertex(0)
	test.Print()
}




// func main() {
	
// 	var Liason []string
// 	var Romme []string
// 	var Romm [] string
// 	v:= os.Args
// 	if len(v)!=2 {
// 		log.Fatal("invalid Arguments!")
// 	}
// 	file,err:=os.ReadFile(v[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	str:= string(file)
// 	str1:=strings.Split( str,"\n")
// 	insect:=str1[0]
// 	for i:=1;i<len(str1)-1;i++ {

// 		if strings.Contains(str1[i],"-"){
// 			Liason=append(Liason, str1[i])
// 		}
// 		Romm=append(Romm,anouar.Roms(str1[i]))
		
// 	}
// 	for i:=0;i<len(Romm);i++ {
// 		if Romm[i]!= "" {
// 			Romme=append(Romme, Romm[i])
// 		}
// 	}
// 	fmt.Println(insect)
// 	fmt.Println(Romme)
// 	fmt.Println(Liason)


	
	
	



// }
