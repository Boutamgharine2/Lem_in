package Lemin

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Handlfile() ([]string, []string, string) {
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
	insect := str1[0]

	for i := 1; i < len(str1); i++ {

		if strings.Contains(str1[i], "-") {
			Edges = append(Edges, str1[i])
		}
		Romm = append(Romm, Roms(str1[i]))

	}
	for i := 0; i < len(Romm); i++ {
		if Romm[i] != "" {
			vertexe = append(vertexe, Romm[i])
		}
	}
	return vertexe, Edges, insect
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

// func MoveAnts(numAnts int, paths [][]string) [][]string {
// 	var (
// 		res      []string
// 		resfinal [][]string
// 		matrix   [][]string
// 	)

// 	for i := 0; i < len(paths); i++ {
// 		for k := 0; k < numAnts; k++ {
// 			for j := 1; j < len(paths[i]); j++ {

//					restem := "L" + TAbloOfAnts(numAnts)[k] + "-" + paths[i][j]
//					res = append(res, restem)
//				}
//				matrix = append(matrix, res)
//				res = nil
//			}
//		}
//		fmt.Println(matrix)
//		resfinal = (HandlTab(matrix))
//		return resfinal
//	}
func MoveAnts(numAnts int, paths [][]string) [][]string {
	type path struct {
		i    int
		path []string
	}
	paths1 := []path{}
	for _, v := range paths {
		v = v[1:]
		paths1 = append(paths1, path{len(v), v})
	}
	result := make([][]string, numAnts)
	//fmt.Println(paths1, result, numAnts)
	for i := 1; i <= numAnts; i++ {
		minidx := 0
		for i1, v := range paths1 {
			if paths1[minidx].i >= v.i {
				minidx = i1
			}
		}
		pathCrossed := []string{}
		for i1 := len(paths1[minidx].path); i1 < paths1[minidx].i; i1++ {
			pathCrossed = append(pathCrossed, "")
		}
		for _, v := range paths1[minidx].path {
			pathCrossed = append(pathCrossed, fmt.Sprintf("L%d-%s", i, v))
		}
		result[i-1] = pathCrossed
		paths1[minidx].i++
	}
	for _, v := range result {
		fmt.Println(v)
	}
	return result
}

func HandlTab(tab [][]string) [][]string {
	var checkpathee []string
	var checkformis []string
	var res [][]string

	for i := 0; i < len(tab); i++ {

		Split := strings.Split(tab[i][0], "-")
		ant := Split[0]

		if (valid(ExtraitP(tab[i]), checkpathee) && i != len(tab)-1) || valid(ant, checkformis) {
			continue
		} else {
			checkant(ant, &checkformis)
			checkpathe(&checkpathee, ExtraitP(tab[i]))
			res = append(res, tab[i])

		}
	}

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

func checkpathe(tab1 *[]string, path string) bool {
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
