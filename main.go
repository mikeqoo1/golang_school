package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	D "golang_school/lib/db"
	E "golang_school/lib/egg"
	F "golang_school/lib/file"
	C "golang_school/lib/gochannels"
	//G "golang_school/lib/gtkkk"
	L "golang_school/lib/leetcode"
	B "golang_school/lib/string"
	A "golang_school/lib/time"
	Tree "golang_school/lib/tree"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"time"

	"go.uber.org/zap"
)

var (
	//Version 版本
	Version string

	//Build 編譯版本
	Build string

	pool *sql.DB
)

func init() {
	pool = D.InitDB()
	F.InitLog()
	F.LOG.Info("log 初始化成功")
}

// MambaOut R.I.P #24 #8
type MambaOut struct {
	bodyID     string //size 1
	exCode     string //size 1
	msgType    string //size 1
	typeID     string //size 1
	connID     string //size 1
	pvcID      string //size 1
	rtnState   string //size 1
	data       []byte
	therestStr []byte
}

// Kobe Bryant If you're afraid to fail, then you're probably going to fail.
var Kobe MambaOut

// MambaMentality Mamba Mentality Never Ends
func MambaMentality(originalsdata []byte) MambaOut {
	originalsdata = append(originalsdata, Kobe.therestStr...) //接上之前的剩餘
	index := bytes.IndexByte(originalsdata, byte('\n'))
	fmt.Println(index, len(originalsdata))
	fmt.Println("1. originalsdata=", originalsdata)
	if index == -1 {
		Kobe.therestStr = originalsdata
	}
	dataLEN, err := strconv.Atoi(string(originalsdata[8:12])) //size 4
	if err != nil {
		fmt.Println("代表第一個出現的不是韻玲姐給的換行符號, 是內容的資料, 跳過繼續找")
	}
	fmt.Println("資料長度", dataLEN)

	if dataLEN+32 > len(originalsdata) {
		fmt.Println("代表資料長度大於現在這筆資料的長度, 這樣資料一定不完整, 當剩餘除處理")
		Kobe.therestStr = originalsdata
	} else {
		Kobe.bodyID = string(originalsdata[1:2])
		Kobe.exCode = string(originalsdata[2:3])
		Kobe.msgType = string(originalsdata[3:4])
		Kobe.typeID = string(originalsdata[4:5])
		Kobe.connID = string(originalsdata[5:6])
		Kobe.pvcID = string(originalsdata[6:7])
		Kobe.rtnState = string(originalsdata[7:8])
		Kobe.data = originalsdata[32 : 32+dataLEN]

		fmt.Println("起始(string)", originalsdata[0:1], string(originalsdata[0:1]))
		fmt.Println("BodyID(string)", Kobe.bodyID)
		fmt.Println("ExCode(string)", Kobe.exCode)
		fmt.Println("MsgType(string)", Kobe.msgType)
		fmt.Println("TypeID(string)", Kobe.typeID)
		fmt.Println("ConnID(string)", Kobe.connID)
		fmt.Println("PvcID(string)", Kobe.pvcID)
		fmt.Println("rtnState(string)", Kobe.rtnState)
		fmt.Println("brokID(string)", originalsdata[12:20]) //size 8
		fmt.Println("wtmpID(string)", originalsdata[20:32]) //size 12
		fmt.Println("bodyData(string)", Kobe.data)
	}
	return Kobe
}

func calcSum(arr []int, res chan int) {
	var sum = 0
	for _, v := range arr {
		sum += v
	}

	res <- sum
}

type NbaPlayer struct {
	Name   string
	Points int
	Min    int
}

type Players []NbaPlayer

// 取得長度
func (p Players) Len() int { return len(p) }

func (p Players) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// 繼承NbaPlayer的所有屬性和方法 所以SortByPoints也實現了 Len() 和 Swap() 方法
type SortByPoints struct{ Players }

// 根據得分大到小排序 (自己的排序邏輯)
func (p SortByPoints) Less(i, j int) bool {
	return p.Players[i].Points > p.Players[j].Points
}

type SortByMin struct{ Players }

// 按上場照時間大到小排序 (自己的排序邏輯)
func (p SortByMin) Less(i, j int) bool {
	return p.Players[i].Min > p.Players[j].Min
}

func main() {
	// CPU 跟 記憶體 的效能分析檔
	cpuFile, _ := os.OpenFile("cpu.prof", os.O_CREATE|os.O_RDWR, 0644)
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()
	memFile, _ := os.OpenFile("mem.prof", os.O_CREATE|os.O_RDWR, 0644)
	defer memFile.Close()
	pprof.WriteHeapProfile(memFile)

	// F.LOG.Info("星夜裡的人",
	// 	zap.String("url", "https://www.youtube.com/watch?v=RnsVmZDQaJA"),
	// 	zap.Int("時間", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
	F.LOG.Info(" ", zap.String("Version: ", Version))
	F.LOG.Info(" ", zap.String("Build time: ", Build))
	F.LOG.Info("GoMod")

	argsWithProg := os.Args

	if len(argsWithProg) == 1 { //跑全部
		fmt.Println("請帶參數 1跑leetcode 2跑樹狀展示 3跑有的沒的 4跑面試 5跑GTK 6跑排序")
	} else {
		if os.Args[1] == "1" {
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

			fmt.Println(L.FourSum(a, 0))
			fmt.Println(L.IsValid("{[]}"))
			//ListNode
			fmt.Println("ListNode 相關")
			var l1 L.ListNode
			l1.Val = 1
			l1.Next = new(L.ListNode)
			fmt.Println("1.", l1)

			l1.Next.Val = 2
			l1.Next.Next = new(L.ListNode)
			fmt.Println("2.", l1.Next)

			l1.Next.Next.Val = 3
			l1.Next.Next.Next = new(L.ListNode)
			fmt.Println("3.", l1.Next.Next)
			fmt.Printf(">>>>展示樹狀結構<<<<\n")
			Tree.ThisTreeDemo()
			l1.Next.Next.Next.Val = 4
			l1.Next.Next.Next.Next = new(L.ListNode)
			fmt.Println("4.", l1.Next.Next.Next)

			l1.Next.Next.Next.Next.Val = 5
			fmt.Println("5.", l1.Next.Next.Next.Next)

			fmt.Println("原本l1", l1, l1.Next, l1.Next.Next, l1.Next.Next.Next, l1.Next.Next.Next.Next)
			fmt.Println("leetcode 19")
			l5 := L.RemoveNthFromEnd(&l1, 2)
			fmt.Println("過了leetcode 19 list", l5, l5.Next, l5.Next.Next, l5.Next.Next.Next, l5.Next.Next.Next.Next)

			numberarray := []int{1, 2, 3, 4, 5, 6, 7, 8}
			lkk := &L.ListNode{Val: numberarray[0], Next: nil}
			p := lkk
			for _, r := range numberarray[1:] {
				node := &L.ListNode{Val: r, Next: nil}
				p.Next = node
				p = p.Next
			}
			k := 3
			res := L.ReverseKGroup(lkk, k)
			for res != nil {
				if res.Next != nil {
					fmt.Printf("%d->", res.Val)
				} else {
					fmt.Printf("%d", res.Val)
				}
				res = res.Next
			}
			L.RemoveDuplicates(a)
		} else if os.Args[1] == "2" {
			fmt.Printf(">>>>展示樹狀結構<<<<\n")
			Tree.ThisTreeDemo()
		} else if os.Args[1] == "3" {
			//S.SendEmail()
			fmt.Println("時間格式")
			A.TimetoFormat()
			//A.SubDay()
			A.TimeCmp()
			fmt.Println("字串處理(向左補齊 向右補齊)")
			B.PadLeft("123", 5, "0")
			B.PadRight("123", 5, "0")
			fmt.Println("Json格式轉換Struct")
			B.TestJSONToStruct()
			fmt.Println("Struct轉換Json格式")
			B.TestStructToJSON()
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
			//切包測試
			originalsdata := []byte{170, 2, 50, 84, 50, 48, 2, 255, 48, 49, 51, 51, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 130, 0, 0, 0, 55, 93, 240, 173, 10, 2, 23, 102, 1, 4, 1, 249, 0, 52, 1, 4, 1, 4, 65, 48, 48, 55, 49, 0, 0, 0, 44, 48, 48, 48, 48, 48, 49, 50, 49, 2, 84, 88, 70, 76, 57, 47, 65, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 255, 255, 244, 72, 0, 3, 0, 150, 50, 124, 49, 1, 2, 0, 79, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 237, 64, 0, 1, 0, 0, 0, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 93, 240, 172, 214, 1, 142, 93, 240, 173, 10, 2, 23, 4, 0, 0, 5, 218, 0, 0, 0, 55, 1, 146, 10}
			Kobe = MambaMentality(originalsdata)
		} else if os.Args[1] == "4" {
			//用goroutine 實作 1加到100
			var res = make(chan int)
			var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			go calcSum(arr[:5], res)
			go calcSum(arr[5:], res)
			sum := <-res
			sum += <-res
			fmt.Println("sum =", sum)

			ctx, cancel := context.WithCancel(context.Background())
			go watch(ctx, "【監控1】")
			go watch(ctx, "【監控2】")
			go watch(ctx, "【監控3】")

			time.Sleep(10 * time.Second)
			fmt.Println("可以了，通知監控停止")
			cancel()
			//為了檢測監控過是否停止，如果没有監控輸出，就表示停止了
			time.Sleep(5 * time.Second)

			// 結果 打印後 打印中 打印前 panic: 觸發異常
			defer func() { fmt.Println("打印前") }()
			defer func() { fmt.Println("打印中") }()
			defer func() { fmt.Println("打印後") }()
			panic("觸發異常")
		} else if os.Args[1] == "5" {
			fmt.Println("GTK")
			//G.Run()
		} else if os.Args[1] == "6" {
			playerlist := Players{
				{
					Name:   "Kobe Bryant",
					Points: 33643,
					Min:    48643,
				},
				{
					Name:   "Michael Jordan",
					Points: 32292,
					Min:    41010,
				},
				{
					Name:   "LeBron James",
					Points: 35516,
					Min:    50276,
				},
				{
					Name:   "Kareem Abdul-Jabbar",
					Points: 38387,
					Min:    57446,
				},
				{
					Name:   "Karl Malone",
					Points: 36928,
					Min:    54852,
				},
			}

			fmt.Println("得分排序前")
			for _, play := range playerlist {
				fmt.Println(play.Name, ":", play.Points, "分")
			}
			sort.Sort(SortByPoints{playerlist})
			fmt.Println("得分排序後")
			for _, play := range playerlist {
				fmt.Println(play.Name, ":", play.Points, "分")
			}

			fmt.Println("上場時間排序前")
			for _, play := range playerlist {
				fmt.Println(play.Name, ":", play.Min, "分鐘")
			}
			sort.Sort(SortByMin{playerlist})
			fmt.Println("上場時間排序後")
			for _, play := range playerlist {
				fmt.Println(play.Name, ":", play.Min, "分鐘")
			}
		} else {
			fmt.Println("1跑leetcode 2跑樹狀展示 3跑有的沒的 4跑面試 5跑GTK 6跑排序")
		}
		//https://docs.nats.io/
		//待補
	}
	fmt.Println("\nHello Word!")
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "監控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine監控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
