package main

import (
	"database/sql"
	"fmt"
	D "golang_school/lib/db"
	C "golang_school/lib/gochannels"
	B "golang_school/lib/string"
	A "golang_school/lib/time"
)

var pool *sql.DB

func init() {
	pool = D.InitDB()
}

func main() {
	fmt.Println("時間格式")
	A.TimetoFormat()
	//A.SubDay()
	fmt.Println("字串處理")
	B.PadLeft("123", 5, "0")
	B.PadRight("123", 5, "0")
	fmt.Println("goroutine")
	fmt.Println("LeetcodeNo.3")
	fmt.Println(B.LengthOfLongestSubstring("aab"))
	fmt.Println(B.LengthOfLongestSubstring("aab"))
	fmt.Println(B.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(B.LengthOfLongestSubstring("pwwkew"))
	C.TestChannel()
}
