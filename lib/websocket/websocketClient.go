package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.88.186:30001", "http service address")
var max int

func main() {
	flag.Parse()
	log.SetFlags(0)

	var ch = make(chan int, 1)

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/GsCmd186"}
	log.Printf("connecting to %s", u.String())

	var err error
	if len(os.Args) <= 1 {
		fmt.Println("沒有輸入參數, 改預設值50次")
		max = 50
	} else {
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("輸入參數有誤, 以10次來處理")
			max = 10
		}
	}
	fmt.Println("Max值", max)

	c, _, err2 := websocket.DefaultDialer.Dial(u.String(), nil)
	if err2 != nil {
		log.Fatal("dial:", err2)
	}
	defer c.Close()

	mj116 := make(chan int)
	defer close(mj116)

	var z = 0
	go keepread(c, z, ch)

	var count = 0
	go keepwrite(c, count, ch)

	select {
	case <-ch:
		fmt.Println("發送並接收完, 關閉連線")
		c.Close()
	}
}

func keepwrite(c *websocket.Conn, count int, ch chan int) {
	for {
		count++
		b := "{\"cmd\":\"action\", \"data\":\"1\"}"
		if count == max+1 {
			b = "{\"cmd\":\"action\", \"data\":\"2\"}"
		}
		err := c.WriteMessage(websocket.TextMessage, []byte(b))
		fmt.Printf("第%d次 msg = %s\n", count, b)
		if err != nil {
			log.Println("write:", err)
			c.Close()
		}
		if count == max+1 {
			break
		}
	}
}

func keepread(c *websocket.Conn, z int, ch chan int) {
	for {
		wstype, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.Close()
		}
		z++
		log.Printf("Type:%d,收到的第%d筆, recv: %s", wstype, z, message)
		if z == max+1 {
			ch <- 1
			break
		}
	}
}
