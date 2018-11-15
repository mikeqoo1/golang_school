package main

import (
	"fmt"
	C "golang_school/lib/gochannels"
	B "golang_school/lib/string"
	A "golang_school/lib/time"
)

func init() {

}

func main() {
	fmt.Println("時間格式")
	A.TimetoFormat()
	//A.SubDay()
	fmt.Println("字串處理")
	B.PadLeft("123", 5, "0")
	B.PadRight("123", 5, "0")
	fmt.Println("goroutine")
	C.TestChannel()
}
