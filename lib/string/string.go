package string

import (
	"bytes"
	"fmt"
	"strconv"
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
