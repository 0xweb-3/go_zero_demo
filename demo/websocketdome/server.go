package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func serverWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("message:%s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", serverWs)
	fmt.Println("启动websocker。。。。。。")
	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}
