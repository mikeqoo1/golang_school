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
	ans1 := L.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(ans1)
	ans2 := L.Reverse(-123)
	fmt.Println(ans2)
	ans3 := L.MyAtoi(" -456")
	fmt.Println(ans3)
	ans4 := L.IsPalindrome(121)
	fmt.Println(ans4)

	// var nnn = "2019-04-23 11:42:31.90"
	// fmt.Println(len(nnn))
	// fmt.Println(nnn[11:22])
	// fmt.Println(strings.Replace(strings.Replace(nnn[11:22], ".", "", 1), ":", "", 2))

	a := []int{-1, 0, 1, 2, -1, -4}
	b := L.ThreeSum(a)
	fmt.Println(b)
}
