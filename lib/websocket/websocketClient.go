package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.88.186:30001", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/GsCmd186"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	mj116 := make(chan int)
	defer close(mj116)

	var z = 0
	go keepread(c, z)

	var count = 0
	go keepwrite(c, count)

	select {}
}

func keepwrite(c *websocket.Conn, count int) {
	for {
		count++
		b := "{\"cmd\":\"action\", \"data\":\"1\"}"
		if count == 51 {
			b = "{\"cmd\":\"action\", \"data\":\"2\"}"
		}
		err := c.WriteMessage(websocket.TextMessage, []byte(b))
		fmt.Printf("第%d次 msg = %s\n", count, b)
		if err != nil {
			log.Println("write:", err)
			c.Close()
		}
		if count == 51 {
			break
		}
	}
}

func keepread(c *websocket.Conn, z int) {
	for {
		wstype, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.Close()
		}
		z++
		log.Printf("Type:%d,收到的第%d筆, recv: %s", wstype, z, message)
	}
}
