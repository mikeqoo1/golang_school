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
	Q := findMedianSortedArrays(nums1, nums2)
	fmt.Println(Q)
}

//findMedianSortedArrays 尋找中位數
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	nums3 := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i = i + 1
		} else {
			nums3 = append(nums3, nums2[j])
			j = j + 1
		}
	}
	if i < len(nums1) {
		for ; i < len(nums1); i++ {
			nums3 = append(nums3, nums1[i])
		}
	}
	if j < len(nums2) {
		for ; j < len(nums2); j++ {
			nums3 = append(nums3, nums2[j])
		}
	}
	lenSums3 := len(nums3)
	if lenSums3%2 == 0 {
		return (float64(nums3[lenSums3/2-1]) + float64(nums3[lenSums3/2])) / 2.0
	} else {
		return float64(nums3[(lenSums3+1)/2-1])
	}
}
