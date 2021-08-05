package main

import (
	"fmt"
	
	"os"
	"strconv"
	"strings"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getData(path string) []int {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	var stringRead = make([]byte, 1024)
	f.Read(stringRead)
	array1 := strings.Split(string(stringRead), " ")
	l := 0
	for array1[len(array1)-1][l] >= 48 && array1[len(array1)-1][l] <= 58 {
		l++
	}
	array1[len(array1)-1] = array1[len(array1)-1][:l]
	// fmt.Println(array1)
	var array2 []int
	for i := range array1 {
		tmp, _ := strconv.Atoi(array1[i])
		array2 = append(array2, tmp)
	}
	return array2
}
func ktSoTrongFile(x int, path string) bool{
    arr := getData(path)
	for i:=range arr{
        if(x==arr[i]) {
			return true
		}
    }
	return false
}
func main(){
	pathFile := "file2.txt"
	n := 5
	
	if ktSoTrongFile(n, pathFile){
		fmt.Println(n," co trong file ",pathFile)
	}else{
		fmt.Println(n," ko co trong file ",pathFile)
	}

}