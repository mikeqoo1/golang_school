package string

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

//PadLeft 左邊補齊
func PadLeft(str string, length int, pad string) string {
	fmt.Println(str)
	templen := strconv.FormatInt(int64(length), 10)
	format := "%" + pad + templen + "s"
	fmt.Println(format)
	str = fmt.Sprintf(format, str)
	fmt.Println(str)
	return str
}

//PadRight 右邊補齊
func PadRight(str string, length int, pad string) string {
	var buf bytes.Buffer
	fmt.Println(str)
	buf.WriteString(str)
	tempi := length - len(str)
	for i := 0; i < tempi; i++ {
		buf.WriteString(pad)
	}
	fmt.Println(buf.String())
	return buf.String()
}

//LengthOfLongestSubstring 沒有重複字元的最長字串
func LengthOfLongestSubstring(s string) int {
	//暴力for loop, 超時就不貼了

	//法2的想法: 假設xbyauxalopmnbu
	//a重複: 不是xbyaux最長就是, uxalopmnbu最長
	//從頭開始, 找到重複的地方就把start指標指到重複之前的"u"再開始找
	m := [256]int{} // ASCII 0~255
	for i := 0; i < 256; i++ {
		m[i] = -1
	}

	length, start := 0, -1

	for i := 0; i < len(s); i++ {
		// 如果該字元被找過, 則紀錄該字元已經紀錄的最大位置
		//fmt.Println("0.index =", m[s[i]])
		if m[s[i]] > start {
			//fmt.Println("ASCII =", s[i])
			//fmt.Println("1.index =", m[s[i]])
			start = m[s[i]]
		}
		// 將存在的字元紀錄為key, 字元在字串中的位置紀錄為value
		m[s[i]] = i //value
		//fmt.Println("2.index =", m[s[i]])
		// 當前位置 - 已經紀錄的最大位置
		if i-start > length {
			length = i - start
		}
		//fmt.Println("start =", start)
	}
	return length
}

//GoRUNS 解析 rune 型態
func GoRUNS() {
	var str = "hello, 你好嗎?"
	//golang中string底層是用byte實作的，len 實際是在用字元長度計算大小 所以一中文占3個字元長度,大小就是12了
	fmt.Println("len(str):", len(str))

	//以下2種方法都可以得到str的字串的長度

	//golang中的unicode/utf8包提供了用utf-8獲取長度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//用rune型態來unicode字串
	fmt.Println("rune:", len([]rune(str)))

	//byte 等於int8， 大致用来處理ascii字元
	//rune 等於int32, 大致用来處理unicode或utf-8字串
}
