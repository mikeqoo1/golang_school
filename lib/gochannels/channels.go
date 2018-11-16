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
	b2 := []byte{0, 130, 0, 0, 0, 1, 91, 238, 40, 232, 3, 79, 102, 1, 4, 2, 88, 0, 48, 1, 4, 1, 4,
		66, 69, 48, 48, 49, 0, 0, 0, 1, 66, 48, 48, 48, 48, 49, 50, 49, 2, 84, 88, 70, 67, 57, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 0, 16, 5, 144, 0, 1, 0, 132, 38,
		102, 49, 1, 2, 0, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 91, 238, 40, 232, 3, 79, 91, 238, 40, 232, 3, 79, 4, 0, 0,
		24, 81, 0, 0, 0, 1, 1, 44}
	go testpush(b2)
	time.Sleep(time.Microsecond * 100)
	time.Sleep(time.Microsecond * 100)
	go bigbang()
	go bigbang()
	go bigbang()
	go bigbang()
	go bigbang()
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
	fmt.Println(string(B))
}

func testpop() []byte {
	return <-ABC
}
