package main

import (
	"fmt"
	"workFile"
)
func GTLN(a []int) int{
	max := a[0]
	for i:=range a{
		if max<a[i] {
			max = a[i]
		}
	}
	return max
}
func GTNN(a []int) int{
	min := a[0]
	for i:=range a{
		if min>a[i] {
			min = a[i]
		}
	}
	return min
}
func GTTB(a []int) float32{
	if len(a)==0 {
		return 0
	}
	sum :=0
	for i:=range a{
		sum+=a[i]
	}
	return float32(float32(sum)/float32(len(a)))
}
func main() {
	a := workFile.MyFile{
		Path: "file2.txt",
	}
	arrInput := a.GetDataArrayNumber()
	
	fmt.Println(arrInput)
	fmt.Println("Gia tri lon nhat: ",GTLN(arrInput))
	fmt.Println("Gia tri nho nhat: ",GTNN(arrInput))
	fmt.Println("Gia tri trung binh: ",GTTB(arrInput))

}
