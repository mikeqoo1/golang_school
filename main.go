package main

import (
	B "golang_school/lib/string"
	A "golang_school/lib/time"
)

func init() {

}

func main() {
	A.TimetoFormat()
	//A.SubDay()
	B.PadLeft("123", 5, "0")
	B.PadRight("123", 5, "0")
}
