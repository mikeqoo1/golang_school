package channels

import (
	"fmt"
	"time"
)

//ABC 變數使用
var ABC = make(chan []byte, 5)
var c = 10

//TestChannel 測試通道使用
func TestChannel() {
	go blackpink("5", "6", "7", "8123450")
	go blackpink("1", "2", "3", "4987659")
	go blackpink("A", "B", "C", "DXYZP")
	go blackpink("P", "I", "N", "KNBAA")
	go blackpink("0", "0", "0", "0")
	time.Sleep(time.Microsecond * 100)
	time.Sleep(time.Microsecond * 100)
	go bigbang()
	go bigbang()
	go bigbang()
	//go bigbang()
	//go bigbang()
	time.Sleep(time.Microsecond * 100)
	time.Sleep(time.Microsecond * 100)
}

func blackpink(J string, lisa string, rose string, jisoo string) {
	temp := J + lisa + rose + jisoo
	testpush([]byte(temp))
}

func testpush(A []byte) {
	ABC <- A
}

func bigbang() {
	B := testpop()
	fmt.Println(B)
}

func testpop() []byte {
	return <-ABC
}
