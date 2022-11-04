// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// type Message struct {
// 	Greeting string `json:"greeting"`
// }

// var (
// 	wsUpgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}
// 	wsConn *websocket.Conn
// )

// func WsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
// 		// check the http.reques
// 		// make sure it's ok to acess
// 		return true
// 	}

// 	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Printf("could not upgrade : %s\n", err.Error())
// 		return
// 	}

// 	defer wsConn.Close()

// 	// event loop
// 	for {
// 		var msg Message

// 		err := wsConn.ReadJSON(&msg)
// 		if err != nil {
// 			fmt.Printf("error reading json: %s\n", err.Error())
// 			break
// 		}

// 		fmt.Printf("Message Received: %s\n", msg.Greeting)
// 	}
// }

// func main() {
// 	router := mux.NewRouter()

// 	log.Fatal(http.ListenAndServe(":9100", router))

// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Greeting string `json:"greeting"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// check the http.Request
		// make sure it's OK to access
		return true
	}
	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}

	defer wsConn.Close()

	// event loop
	for {
		var msg Message

		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}

		fmt.Printf("Message Received: %s\n", msg.Greeting)
		SendMessage("Hello, Client!:YOur Message" + msg.Greeting)
	}
}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/socket", WsEndpoint)

	log.Fatal(http.ListenAndServe(":9100", router))

}
