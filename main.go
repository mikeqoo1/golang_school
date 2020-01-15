package main

import (
	"bytes"
	"database/sql"
	"fmt"
	D "golang_school/lib/db"
	E "golang_school/lib/egg"
	C "golang_school/lib/gochannels"
	L "golang_school/lib/leetcode"
	B "golang_school/lib/string"
	A "golang_school/lib/time"
	"strconv"

	"rsc.io/quote"
)

var (
	//Version 版本
	Version string

	//Build 編譯版本
	Build string

	//IsLeetCode 是不是跑LeetCode
	IsLeetCode string

	pool *sql.DB
)

func init() {
	pool = D.InitDB()
}

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build time: ", Build)
	fmt.Println("GOMod Test", quote.Hello())
	if IsLeetCode == "TURE" {

		fmt.Println("LeetcodeNo.3")
		fmt.Println(B.LengthOfLongestSubstring("aab"))
		fmt.Println(B.LengthOfLongestSubstring("aab"))
		fmt.Println(B.LengthOfLongestSubstring("bbbbb"))
		fmt.Println(B.LengthOfLongestSubstring("pwwkew"))

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

		a := []int{-1, 0, 1, 2, -1, -4}
		b := L.ThreeSum(a)
		fmt.Println(b)

		fmt.Println(L.LetterCombinations("56"))
	} else {
		fmt.Println("時間格式")
		A.TimetoFormat()
		//A.SubDay()
		fmt.Println("字串處理(向左補齊 向右補齊)")
		B.PadLeft("123", 5, "0")
		B.PadRight("123", 5, "0")
		fmt.Println("Json格式轉換Struct")
		B.TestJsonToStruct()
		fmt.Println("Struct轉換Json格式")
		B.TestStructToJson()
		fmt.Println("Goroutine")
		C.TestChannel()
		byteA := []byte{255, 255, 177, 224} //-20000
		val, _ := B.BytesToIntS(byteA)
		fmt.Println(val)
		strint := strconv.FormatInt(int64(val), 10)
		fmt.Println(strint)

		egg := E.NewWaterEgg()
		egg.SetTestSyncMapValue("房東與貓", "往往")
		a := egg.GetTestSyncMapValue("房東與貓")
		fmt.Println(a.Name)
		fmt.Println(a.Number)
		a = egg.GetTestSyncMapValue("湖泊樂隊")
		fmt.Println(a.Name)
		fmt.Println(a.Number)
	}

	//var, nnn, =, "2019-04-23, 08:42:31.90"
	//fmt.Println("時間:", nnn)
	//fmt.Println("時間長度", len(nnn))
	//fmt.Println("時間[11:22]", nnn[11:22])
	//fmt.Println("時間轉換", strings.Replace(strings.Replace(nnn[11:22], ".", "", 1), ":", "", 2))

	//切包測試

	//[4096]byte{170, 2, 50, 84, 50, 48, 2, 255, 48, 49, 51, 51, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 130, 0, 0, 0, 55, 93, 240, 173, 10, 2, 23, 102, 1, 4, 1, 249, 0, 52, 1, 4, 1, 4, 65, 48, 48, 55, 49, 0, 0, 0, 44, 48, 48, 48, 48, 48, 49, 50, 49, 2, 84, 88, 70, 76, 57, 47, 65, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 255, 255, 244, 72, 0, 3, 0, 150, 50, 124, 49, 1, 2, 0, 79, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 237, 64, 0, 1, 0, 0, 0, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 93, 240, 172, 214, 1, 142, 93, 240, 173, 10, 2, 23, 4, 0, 0, 5, 218, 0, 0, 0, 55, 1, 146, 10}
	buf1 := "<AA><FF>1r2A<FF><FF>0004845T0000<FF><FF><FF><FF><FF><FF><FF><FF><FF><FF><FF><FF>T\u0002O\u0002\n"

	for i := 0; i < 1; i++ {
		index := bytes.IndexByte([]byte(buf1), byte('\n'))
		fmt.Println(index)
		fmt.Println([]byte(buf1)[8:12])
		fmt.Println(strconv.Atoi(string(buf1[8:12])))
	}
	//回報
	//buf2 := [4096]byte{170, 2, 50, 84, 50, 48, 2, 255, 48, 49, 51, 51, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 130, 0, 0, 0, 55, 93, 240, 173, 10, 2, 23, 102, 1, 4, 1, 249, 0, 52, 1, 4, 1, 4, 65, 48, 48, 55, 49, 0, 0, 0, 44, 48, 48, 48, 48, 48, 49, 50, 49, 2, 84, 88, 70, 76, 57, 47, 65, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 255, 255, 244, 72, 0, 3, 0, 150, 50, 124, 49, 1, 2, 0, 79, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 237, 64, 0, 1, 0, 0, 0, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 93, 240, 172, 214, 1, 142, 93, 240, 173, 10, 2, 23, 4, 0, 0, 5, 218, 0, 0, 0, 55, 1, 146, 10}

	// fmt.Println("起始(string)", buf1[0:1], string(buf1[0:1]))
	// fmt.Println("BodyID(string)", buf1[1:2], string(buf1[1:2]))
	// fmt.Println("ExCode(string)", buf1[2:3], string(buf1[2:3]))
	// fmt.Println("MsgType(string)", buf1[3:4], string(buf1[3:4]))
	// fmt.Println("TypeID(string)", buf1[4:5], string(buf1[4:5]))
	// fmt.Println("ConnID(string)", buf1[5:6], string(buf1[5:6]))
	// fmt.Println("PvcID(string)", buf1[6:7], string(buf1[6:7]))
	// fmt.Println("rtnState(string)", buf1[7:8])
	// fmt.Println(buf1[8:12])
	// fmt.Println("[8:12]轉數字")
	// aaa, _ := strconv.Atoi(string(buf1[8:12]))
	// fmt.Println(strconv.Atoi(string(buf1[8:12])))
	// fmt.Println("brokID(string)", buf1[12:20])
	// fmt.Println("wtmpID(string)", buf1[20:32])
	// fmt.Println("bodyData(string)", buf1[32:32+aaa])

	//fmt.Println(bytes.IndexByte([]byte("kch\nicken\n"), byte('\n')))
	//fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g')))
}
