package main

import (
	"fmt"
	"math"
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
func ktSoNguyenTo(x int) bool{
	for i:=2;i<=int(math.Sqrt(float64(x)));i++{
		if x%i==0 {
			return false
		}
	}
	return true
}
func main(){
	arrInput := getData("file2.txt")
	fmt.Println("arrInput: ",arrInput)

	arrDaySo := arrInput
	for i:=range arrDaySo{
		fmt.Println(arrDaySo[i],ktSoNguyenTo(arrDaySo[i]))
	}

}