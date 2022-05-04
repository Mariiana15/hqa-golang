package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
)

type WebsocketHQA struct {
	Protocol string   `json:"protocol"`
	Token    string   `json:"token"`
	Hosts    []string `json:"hosts"`
}

var upgrader = websocket.Upgrader{ReadBufferSize: 1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 15000, // <-- add this line
}

type Message struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (ws *WebsocketHQA) NewWebSocketHQA() {

	path, _ := filepath.Abs("../configuration/config.json")
	file, _ := ioutil.ReadFile(path)
	var result map[string]interface{}
	json.Unmarshal([]byte(file), &result)
	byteData, _ := json.Marshal(result["websocket"])
	json.Unmarshal(byteData, &ws)
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {

	var wsHQA WebsocketHQA
	wsHQA.NewWebSocketHQA()
	upgrader.Subprotocols = append(upgrader.Subprotocols, "upgrader")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return Contains(wsHQA.Hosts, r.Header.Get("origin"))
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", msgDataCorrupt)

	}
	defer ws.Close()

	log.Println("Connected!")

	for {
		var message Message
		//	mt, message, _ := ws.ReadMessage()
		//	log.Printf("Message received: %s", message)
		err := ws.ReadJSON(&message)
		if err != nil {
			fmt.Printf("{\"error\": \"%v\"}", msgDataCorrupt)
			break
		}
		if message.Token == "" || message.Token != "key" {
			fmt.Printf("{\"error\": \"%v\"}", msgDataCorrupt)
			break
		}

	}
	return ws
}

func SendMessageJsonWS(ws *websocket.Conn, message string) {

	defer ws.Close()
	if err := ws.WriteJSON(message); err != nil {
		log.Printf("error occurred: %v", err)
	}
}

func SendMessageWS(ws *websocket.Conn, message string) {

	defer ws.Close()
	byteData, _ := json.Marshal(message)
	if err := ws.WriteMessage(1, byteData); err != nil {
		log.Printf("error occurred: %v", err)
	}
}
