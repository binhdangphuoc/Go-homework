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
	f,err := os.Create("file1.txt")
	check(err)
	f.Close()
	fmt.Println("oke")
}
