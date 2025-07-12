package main

import "fmt"

func maps(){
x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])
}
func mapsFunction(){
	elements := map[string]string{
"H": "Hydrogen",
"He": "Helium",
"Li": "Lithium",
"Be": "Beryllium",
"B": "Boron",
"C": "Carbon",
"N": "Nitrogen",
"O": "Oxygen",
"F": "Fluorine",
"Ne": "Neon",
}
fmt.Println(elements["Li"])
}
func main(){
	maps()
	mapsFunction()
}