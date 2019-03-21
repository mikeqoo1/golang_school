package main

import (
	"database/sql"
	"fmt"
	D "golang_school/lib/db"
	C "golang_school/lib/gochannels"
	L "golang_school/lib/leetcode"
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
	fmt.Println("JsonToStruct")
	B.TestJsonToStruct()
	fmt.Println("StructToJson")
	B.TestStructToJson()
	fmt.Println("goroutine")
	fmt.Println("LeetcodeNo.3")
	fmt.Println(B.LengthOfLongestSubstring("aab"))
	fmt.Println(B.LengthOfLongestSubstring("aab"))
	fmt.Println(B.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(B.LengthOfLongestSubstring("pwwkew"))
	C.TestChannel()
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	Q := L.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(Q)
	P := L.Reverse(-123)
	fmt.Println(P)
	PP := L.MyAtoi(" -456")
	fmt.Println(PP)
	PPP := L.IsPalindrome(121)
	fmt.Println(PPP)
}
