package time

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//SubDay 日期相減
func SubDay() {
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

//TimetoFormat 顯示time Format
func TimetoFormat() {
	p := fmt.Println

	t := time.Now()
	p("No Format =", t)
	p("RFC3339 =", t.Format(time.RFC3339))
	p("RFC3339Nano =", t.Format(time.RFC3339Nano))
	sad := fmt.Sprintf("%02d%02d%02d", t.Hour(), t.Minute(), t.Second())
	//BBB := t.Minute()
	//CCC := t.Second()
	p("AAA =", sad)

	//拿n天前
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -2) //要3天前的時間, 把-2改-3
	p("n天前RFC3339 =", oldTime.Format(time.RFC3339))
}

//TimeCmp 比較時間
func TimeCmp() {
	現在時間 := time.Now()
	中午 := "12:00PM"
	晚上 := "12:00AM"
	fmt.Println(現在時間.Format(time.Kitchen))
	nowtimestr, err := time.Parse(time.Kitchen, 現在時間.Format(time.Kitchen))
	noontimestr, _ := time.Parse(time.Kitchen, 中午)
	nighttimestr, _ := time.Parse(time.Kitchen, 晚上)
	fmt.Println("現在時間=", nowtimestr)
	fmt.Println("中午=", noontimestr)
	fmt.Println("晚上=", nighttimestr)
	if err == nil && nowtimestr.After(noontimestr) {
		fmt.Println("現在時間 > 中午")
	} else if nowtimestr.After(nighttimestr) {
		fmt.Println("現在時間 > 晚上")
	} else if nowtimestr.Before(noontimestr) {
		fmt.Println("現在時間 < 中午")
	} else if nowtimestr.Before(nighttimestr) {
		fmt.Println("現在時間 < 晚上")
	}
}

//Duration 持續時間的類型比對
func Duration() {
	//Duration類型轉int64 float64
	var durationseconds time.Duration = (1250 * 10) * time.Millisecond
	var durationminute time.Duration = 2 * time.Minute
	var float64seconds float64 = durationseconds.Seconds()
	var float64minutes float64 = durationminute.Minutes()
	fmt.Printf("Seconds [%.3f]\nMinutes [%.2f]\n", float64seconds, float64minutes)

	//ms轉int64 float64
	var duration500milliseconds time.Duration = 500 * time.Millisecond
	var castToInt64 int64 = duration500milliseconds.Nanoseconds() / 1e6
	var castToFloat64 float64 = duration500milliseconds.Seconds() * 1e3
	fmt.Printf("Duration [%v]\ncastToInt64 [%d]\ncastToFloat64 [%.0f]\n", duration500milliseconds, castToInt64, castToFloat64)

	//應用
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond
	startingTime := time.Now().UTC()
	time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()
	var duration time.Duration = endingTime.Sub(startingTime)
	if duration >= waitFiveHundredMillisections {
		fmt.Printf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", waitFiveHundredMillisections, duration, duration.Nanoseconds()/1e6, duration.Seconds())
	}
}
