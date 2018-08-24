package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var port = 8081

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ws", ws)

	log.Printf("Server start. Port: %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("fatal: %v", err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ws error: %v", err)
		http.Error(w, "ws error", http.StatusBadRequest)
		return
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ws read message error: %v", err)
			http.Error(w, "ws read message error", http.StatusBadRequest)
			return
		}

		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Printf("ws write message error: %v", err)
			http.Error(w, "ws write message error", http.StatusBadRequest)
			return
		}
	}
}
