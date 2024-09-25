package main

import (
	anouar "anouar/fonction"
	"fmt"
	"log"
	"os"
	"strings"
)





func main() {
	
	var Liason []string
	var Romme []string
	var Romm [] string
	v:= os.Args
	if len(v)!=2 {
		log.Fatal("invalid Arguments!")
	}
	file,err:=os.ReadFile(v[1])
	if err != nil {
		log.Fatal(err)
	}
	str:= string(file)
	str1:=strings.Split( str,"\n")
	insect:=str1[0]
	for i:=1;i<len(str1)-1;i++ {

		if strings.Contains(str1[i],"-"){
			Liason=append(Liason, str1[i])
		}
		Romm=append(Romm,anouar.Roms(str1[i]))
		
	}
	for i:=0;i<len(Romm);i++ {
		if Romm[i]!= "" {
			Romme=append(Romme, Romm[i])
		}
	}
	fmt.Println(insect)
	fmt.Println(Romme)
	fmt.Println(Liason)


	
	
	



}
