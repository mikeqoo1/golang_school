package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if op == 0 {
					fmt.Println("收到op = 0 異常處理")
					conn.Close()
					break
				}
				fmt.Println("收到的msg", string(msg))
				fmt.Println("op =", op)
				if err != nil {
					fmt.Println("ReadClientData ERROR msg", err.Error())
					conn.Close()
					// handle error
				}

				err = wsutil.WriteServerMessage(conn, op, msg)
				if op == 0 {
					fmt.Println("收到op = 0 異常處理")
					conn.Close()
					break
				}
				fmt.Println("送出的msg", string(msg))
				fmt.Println("op =", op)
				if err != nil {
					fmt.Println("WriteServerMessage ERROR msg", err.Error())
					conn.Close()
					// handle error
				}
			}
		}()
	}))
}

