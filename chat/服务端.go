package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var Clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Username  string `json:"username"`
	Message  string `json:"message"`
}

func HandleConnections(i int,w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	if i == 2 && len(Clients) > 1 {
		_ = ws.WriteJSON("超出人数")
		return
	}

	// Register our new client
	Clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(Clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				_ = client.Close()
				delete(Clients, client)
			}
		}
	}
}