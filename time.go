package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("輸入時間 格式範例:2006-01-02")
	timeStr := time.Now().Format("2006-01-02") //現在時間
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')   //輸入時間
	input = strings.Replace(input, "\n", "", -1) //去除換行
	if err != nil {
		fmt.Println("輸入錯誤", err)
	}
	A := strings.Split(input, "-")
	B := strings.Split(timeStr, "-")
	var day int
	fmt.Println("len===>", len(input), A[2])
	fmt.Println("len===>", len(timeStr), B[2])
	for i := 0; i < len(B); i++ {
		aint, err := strconv.Atoi(A[i])
		if err != nil {
			fmt.Println(err)
		}
		bint, _ := strconv.Atoi(B[i])
		if i == 0 {
			day += (bint - aint) * 365
		} else if i == 1 {
			day += (bint - aint) * 31
		} else {
			day += (bint - aint)
		}
	}
	fmt.Println("現在時間-輸入時間")
	fmt.Println("剩餘天數", day)
}
