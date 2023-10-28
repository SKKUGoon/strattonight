package data

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
)

type Stratton struct {
	Static StrattonData
	Stream StrattonData
}

type StrattonData struct {
	Conn *websocket.Conn

	// Go-routine controller
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (sd *StrattonData) Ping() {
	ping := messagePing()

	err := sd.Conn.WriteJSON(ping)
	if err != nil {
		log.Printf("error writing message to websocket: %v", err)
	}
}

func (sd *StrattonData) RequestStream() {

}

func (sd *StrattonData) RemoveStream() {

}

func (sd *StrattonData) ReadFromSocket() {
	for {
		select {
		case <-sd.Ctx.Done():
			log.Println("stopping stream reader")
			return
		default:
			_, message, err := sd.Conn.ReadMessage()
			if err != nil {
				log.Fatalf("failed reading from websocket: %v", err)
			}
			log.Println(string(message))
		}
	}
}

func (sd *StrattonData) Close() {
	err := sd.Conn.Close()
	if err != nil {
		log.Printf("graceful closing of websocket failed: %v", err)
	}
}
