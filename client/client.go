package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const serverURL = "ws://localhost:8080/ws"
const userID = "user123" // Replace this with the desired user ID

func main() {
	socket, _, err := websocket.DefaultDialer.Dial(serverURL, nil)

	if err != nil {
		log.Fatal("WebSocket connection error:", err)
		return
	}

	defer socket.Close()

	if err := socket.WriteMessage(websocket.TextMessage, []byte(userID)); err != nil {
		log.Fatal("Error sending userID:", err)
		return
	}

	sendLocation := func(la, lo float64) {
		locationData := UserLocation{
			Latitude:  la,
			Longitude: lo,
		}

		data, err := json.Marshal(locationData)
		if err != nil {
			log.Println("JSON marshal error:", err)
			return
		}

		err = socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Error sending location:", err)
		}
	}

	for {
		latitude := 37.7749
		longitude := -122.4194
		sendLocation(latitude, longitude)

		time.Sleep(5 * time.Second)
	}
}
