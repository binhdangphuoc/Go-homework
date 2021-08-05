package workFile

import (
	// "bytes"
	// "bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
)

type MyFile struct {
	Path string
}
type actionFile interface {
	Create()
	ReadF() string
	WriteF()
	DeleteF()
}

func (f MyFile) CreateF() {
	data, e := os.Create(f.Path)
	if e != nil {
		fmt.Println("Can not create file")
	}
	defer data.Close()
	fmt.Println("Create file : ", f.Path, " cuccess!")
}
func (f MyFile) ReadF() string {
	file, e := os.Open(f.Path)
	if e != nil {
		fmt.Println("Can not open file to read")
		return ""
	}
	defer file.Close()
	data := make([]byte, 1024)
	file.Read(data)
	var s string
	for i := range data {
		if data[i] == 0 {
			break
		}
		s += string(data[i])
	}
	return s
}
func (f MyFile) WriteF(s string) {

	file, e := os.OpenFile(f.Path, os.O_RDWR, 0644)
	if e != nil {
		fmt.Println("Can not open file to write")
		return
	}
	defer file.Close()
	_, e = file.WriteString(s)
	if e != nil {
		fmt.Println("Can not write to file")
		return
	}
	e = file.Sync()
	if e != nil {
		fmt.Println("File update uncuccess")
	}
	fmt.Println("Write to file cuccess")

}
func (f MyFile) DeleteF() {
	e := os.Remove(f.Path)
	if e != nil {
		fmt.Println("Can not delete file")
		return
	}
	fmt.Println("Delete file cuccess")
}
func (f MyFile) GetDataArrayNumber() []int {

	data := f.ReadF()
	data = strings.Trim(data, " ")
	
	arrData := strings.Split(data," ")

	var arrNumber []int
	for i:=range arrData{
		val,_ := strconv.Atoi(arrData[i])
		arrNumber =append(arrNumber, val)
	}
	return arrNumber
}
