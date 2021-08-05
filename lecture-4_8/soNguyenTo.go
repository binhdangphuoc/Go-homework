package main

import (
	"fmt"
	"math"
	
	"workFile"
)

func ktSoNguyenTo(x int) bool{
	for i:=2;i<=int(math.Sqrt(float64(x)));i++{
		if x%i==0 {
			return false
		}
	}
	return true
}
func main(){
	a:= workFile.MyFile{Path:"file2.txt",}
	arrInput := a.GetDataArrayNumber()
	fmt.Println("arrInput: ",arrInput)

	
	for i:=range arrInput{
		fmt.Println(arrInput[i],ktSoNguyenTo(arrInput[i]))
	}

}