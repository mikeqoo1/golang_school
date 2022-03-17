package gtkkk

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/gotk3/gotk3/gtk"
)

const windowName = "window"

const textViewName = "textview"

const entryName = "entry"

const uIMain = "lib/gtkkk/calculator.glade"

var all string //運算式

func Run() {
	gtk.Init(nil)

	bldr, err := getBuilder(uIMain)
	if err != nil {
		fmt.Println("取得glade檔失敗")
		panic(err)
	}

	window, err := getWindow(bldr)
	if err != nil {
		fmt.Println("取得window物件失敗")
		panic(err)
	}

	window.SetTitle("簡易計算機")
	window.SetDefaultSize(300, 300)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		fmt.Println("關閉視窗失敗")
		panic(err)
	}

	textview, err := getTextView(bldr)
	if err != nil {
		fmt.Println("取得textview物件失敗")
		panic(err)
	}
	textviewbuffer, err := textview.GetBuffer()
	if err != nil {
		fmt.Println("取得textview buffer失敗")
		panic(err)
	}

	entry, err := getEntry(bldr)
	if err != nil {
		fmt.Println("取得entry物件失敗")
		panic(err)
	}
	entry.SetText("請輸入運算式")

	one, _ := getButton(bldr, "one")
	one.Connect("clicked", func() {
		setsymbol("1", textviewbuffer)
	})
	two, _ := getButton(bldr, "two")
	two.Connect("clicked", func() {
		setsymbol("2", textviewbuffer)
	})
	three, _ := getButton(bldr, "three")
	three.Connect("clicked", func() {
		setsymbol("3", textviewbuffer)
	})
	four, _ := getButton(bldr, "four")
	four.Connect("clicked", func() {
		setsymbol("4", textviewbuffer)
	})
	five, _ := getButton(bldr, "five")
	five.Connect("clicked", func() {
		setsymbol("5", textviewbuffer)
	})
	six, _ := getButton(bldr, "six")
	six.Connect("clicked", func() {
		setsymbol("6", textviewbuffer)
	})
	seven, _ := getButton(bldr, "seven")
	seven.Connect("clicked", func() {
		setsymbol("7", textviewbuffer)
	})
	eight, _ := getButton(bldr, "eight")
	eight.Connect("clicked", func() {
		setsymbol("8", textviewbuffer)
	})
	nine, _ := getButton(bldr, "nine")
	nine.Connect("clicked", func() {
		setsymbol("9", textviewbuffer)
	})
	zero, _ := getButton(bldr, "zero")
	zero.Connect("clicked", func() {
		setsymbol("0", textviewbuffer)
	})
	add, _ := getButton(bldr, "add")
	add.Connect("clicked", func() {
		setsymbol("+", textviewbuffer)
	})
	sub, _ := getButton(bldr, "sub")
	sub.Connect("clicked", func() {
		setsymbol("-", textviewbuffer)
	})
	mul, _ := getButton(bldr, "mul")
	mul.Connect("clicked", func() {
		setsymbol("*", textviewbuffer)
	})
	div, _ := getButton(bldr, "div")
	div.Connect("clicked", func() {
		setsymbol("/", textviewbuffer)
	})
	clean, _ := getButton(bldr, "clean")
	clean.Connect("clicked", func() {
		setsymbol("claen", textviewbuffer)
	})
	enter, _ := getButton(bldr, "enter")
	enter.Connect("clicked", func() {
		entrydata, err := entry.GetText()
		if err != nil {
			fmt.Println("entry text失敗")
			panic(err)
		}
		if entrydata != "請輸入運算式" { //把entry的資料跟按鈕的結合
			fmt.Println("all=", all)
			fmt.Println("entrydata=", entrydata)
			all += entrydata
			fmt.Println("all=", all)
		}
		setsymbol("ans", textviewbuffer)

	})

	window.ShowAll()

	gtk.Main()
}

// getBuilder returns *gtk.getBuilder loaded with glade resource (if resource is given)
func getBuilder(filename string) (*gtk.Builder, error) {
	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	if filename != "" {
		err = b.AddFromFile(filename)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

// getWindow returns *gtk.Window object from the glade resource
func getWindow(b *gtk.Builder) (*gtk.Window, error) {
	obj, err := b.GetObject(windowName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window) //斷言判斷是不是gtk物件
	if !ok {
		return nil, err
	}

	return window, nil
}

// getButton returns *gtk.Button object from the glade resource
func getButton(b *gtk.Builder, id string) (*gtk.Button, error) {
	obj, err := b.GetObject(id)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return button, nil
}

// getTextView 取得TextView Obj
func getTextView(b *gtk.Builder) (*gtk.TextView, error) {
	obj, err := b.GetObject(textViewName)
	if err != nil {
		return nil, err
	}

	textview, ok := obj.(*gtk.TextView)
	if !ok {
		return nil, err
	}

	return textview, nil
}

// getEntry 取得Entry Obj
func getEntry(b *gtk.Builder) (*gtk.Entry, error) {
	obj, err := b.GetObject(entryName)
	if err != nil {
		return nil, err
	}

	entry, ok := obj.(*gtk.Entry)
	if !ok {
		return nil, err
	}

	return entry, nil
}

func setsymbol(a string, text *gtk.TextBuffer) {
	if a == "claen" {
		all = ""
		text.SetText(all)
	} else if a == "ans" {
		postfix := infix2ToPostfix(all)
		fmt.Printf("後綴表達式：%s\n", postfix)
		ansmunber := calculate(postfix)
		fmt.Println(ansmunber)
		ansstring := strconv.Itoa(ansmunber)
		all = ansstring
		//all += "=" + ansstring
		text.SetText(all)
		//all = ansstring
	} else {
		all += a
		text.SetText(all)
	}
}

//https://github.com/wuYin/blog/blob/master/codes/calculate-math-statement-by-go-stack/main.go

func infix2ToPostfix(exp string) string {
	stack := ItemStack{}
	postfix := ""
	expLen := len(exp)

	// 遍尋表達算式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue
		case "(":
			// 左括號直接塞進堆疊
			stack.Push("(")
		case ")":
			// 右括號則彈出元素直到遇到左括號
			for !stack.IsEmpty() {
				preChar := stack.Top()
				valStr := fmt.Sprintf("%v", preChar)
				if valStr == "(" {
					stack.Pop() // 彈出 "("
					break
				}
				postfix += valStr
				stack.Pop()
			}

			// 數字直接輸出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一個整數 由於執行了一個多餘的 j++ 需要減1

		default:
			// 操作符:遇到高優先级的運算符號 不斷彈出 直到遇見更低優先级的運算符號
			for !stack.IsEmpty() {
				top := stack.Top()
				valStr := fmt.Sprintf("%v", top)
				if top == "(" || isLower(valStr, char) {
					break
				}
				postfix += valStr
				stack.Pop()
			}
			// 低優先级的運算符號塞進堆疊
			stack.Push(char)
		}
	}

	// 堆疊不是空的話全部輸出
	for !stack.IsEmpty() {
		valStr := fmt.Sprintf("%v", stack.Pop())
		postfix += valStr
	}

	return postfix
}

func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的後綴表達式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

func calculate(postfix string) int {
	stack := ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 數字:直接push
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else {
			// 操作: 取出數字計算 把结果塞回去
			str1 := fmt.Sprintf("%v", stack.Pop())
			num1, _ := strconv.Atoi(str1)
			fmt.Println(num1)
			str2 := fmt.Sprintf("%v", stack.Pop())
			num2, _ := strconv.Atoi(str2)
			fmt.Println(num2)
			fmt.Println(nextChar)
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	sss := fmt.Sprintf("%v", stack.Top())
	result, _ := strconv.Atoi(sss)
	return result
}
