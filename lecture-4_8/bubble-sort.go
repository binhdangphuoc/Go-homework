package main

import (
	"fmt"
	"workFile"
)


func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		checkChange := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] ^ arr[j+1]
				arr[j+1] = arr[j] ^ arr[j+1]
				arr[j] = arr[j] ^ arr[j+1]
				checkChange = false
			}
		}
		if checkChange {
			break
		}
	}
	return arr
}
func quickSortStep(arr []int, L int, R int) int {
	x := arr[L]
	i := L + 1
	j := R
	for i <= j {
		for j >= i && arr[j] >= x {
			j--
		}
		for i <= j && arr[i] < x {
			i++
		}

		if i < j {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
			j--
		}
	}
	tmp := arr[L]
	arr[L] = arr[j]
	arr[j] = tmp
	return j //tra ve so da dc sap xep dung vi tri
}
func quickSort(A []int, L int, R int) {
	if L < R {
		k := quickSortStep(A, L, R)
		quickSort(A, L, k-1)
		quickSort(A, k+1, R)
	}
}
func merge(arr []int, L int, M int, R int){
	arrTmp := arr
	k:=L
	i:=L
	j:=M+1
	for i<=M&&j<=R {
		if arr[i]>arr[j]{
			arrTmp[k]=arr[j]
			k++
			j--
		}else{
			arrTmp[k]=arr[i]
			k++
			i++
		}
	}
	for i<=M {
		arrTmp[k]=arr[i]
		k++
		i++
	}
	for j<=R {
		arrTmp[k]=arr[j]
		k++
		j++
	}
	for t:=L;t<=R;t++{
		arr[t]=arrTmp[t]
	}
}
func mergeSort(A []int, L int, R int){
	if L<R {
		M:=(L+R)/2
		mergeSort(A,L,M)
		mergeSort(A,M+1,R)
		merge(A,L,M,R)
	}
}
func main() {
	a := workFile.MyFile{
		Path: "file1.txt",
	}
	arrInput := a.GetDataArrayNumber()

	arrBubbleSort := arrInput
	fmt.Println("output Bubble sort: ", bubbleSort(arrBubbleSort))

	arrQuickSort := arrInput
	quickSort(arrQuickSort, 0, len(arrQuickSort)-1)
	fmt.Println("output Quick sort: ",arrQuickSort)

	arrMergeSort := arrInput
	mergeSort(arrMergeSort, 0, len(arrQuickSort)-1)
	fmt.Println("output Merge sort: ",arrMergeSort)

	

}
