package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dalyndalton/MWrap/wrapper"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var w *wrapper.Wrapper = nil

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		log.Println(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))
		w.SendMessage(string(p))
	}
}

func listener(conn *websocket.Conn) {
	msg_c := make(chan string, 1)
	go w.DisplayLogs(msg_c)
	for {
		message := <-msg_c
		log.Println(message)
		conn.WriteMessage(1, []byte(message))
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	go listener(ws)
	go reader(ws)

}

func setupRoutes(w *wrapper.Wrapper) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func SetupWebsocket(wrap *wrapper.Wrapper) {
	w = wrap
	fmt.Println("✅ Serving socket on :8080/ws")

	setupRoutes(w)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
