package main

import (
	"fmt"
	"workFile"
)

func ktSoTrongFile(x int, arr []int) bool{
    
	for i:=range arr{
        if(x==arr[i]) {
			return true
		}
    }
	return false
}
func main(){
	a:=workFile.MyFile{
		Path: "file1.txt",
	}
	arrInput := a.GetDataArrayNumber()
	
	var n int
	fmt.Print("Ban muon tim n= ")
	fmt.Scanf("%d", &n)
	
	if ktSoTrongFile(n, arrInput){
		fmt.Println(n," co trong file ",a.Path)
	}else{
		fmt.Println(n," ko co trong file ",a.Path)
	}

}