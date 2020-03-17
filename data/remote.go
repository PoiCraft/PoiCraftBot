package data

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

var wsConn *websocket.Conn
var err error

func openWebSocket() error {
	wsConn, err = websocket.Dial("ws://localhost:8080/ws", "", "http://localhost:8080")
	return err
}
func WebSocketExec(cmdString string) string {
	message := []byte(cmdString)
	_, err = wsConn.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	Logger.Infof("Send: %s\n", message)

	var msg = make([]byte, 512)
	m, err := wsConn.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])
	return string(msg[:m])
}
