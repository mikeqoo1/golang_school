package egg

import (
	"fmt"
	"sync"
)

//Test 水煮蛋最小單位
type Test struct {
	Name   string
	Number string
}

//WaterEgg 水煮蛋裡面包同步的map
type WaterEgg struct {
	TestSyncMap sync.Map
}

//NewWaterEgg 產生一顆水煮蛋
func NewWaterEgg() *WaterEgg {
	//var ordInfoMap OrderInfo
	egg := new(WaterEgg)
	egg.TestSyncMap.Store("0", new(Test))
	return egg
}

//SetTestSyncMapValue 意義很明顯了
func (egg *WaterEgg) SetTestSyncMapValue(name string, number string) {
	var basic Test
	basic.Name = name
	basic.Number = number
	egg.TestSyncMap.Store(name, basic)
}

//GetTestSyncMapValue 意義很明顯了
func (egg *WaterEgg) GetTestSyncMapValue(name string) Test {
	var basic Test
	basic.Name = "水煮蛋"
	basic.Number = "浪子永不回頭"
	value, ok := egg.TestSyncMap.Load(name)
	if ok {
		switch value.(type) {
		case Test:
			return value.(Test)
		default:
			fmt.Println("unknown type")
		}
	}
	return basic
}
