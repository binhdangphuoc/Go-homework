package main

import (
	"fmt"
	"os"
	
)
func check(err error){
	if err!=nil{
		panic(err)
	}
}
func main() {
	f,err := os.Create("fileGhi.txt")
	check(err)
	f.WriteString("aaaaaaaa")
	defer f.Close()
	fmt.Println("oke")
}
