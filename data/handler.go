package data

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/skkugoon/strattonight/ent"
	"log"
)

type Stratton struct {
	// Binance websocket connections
	Static StrattonData
	Stream StrattonData

	// Local(Cloud) database connections
	Local *ent.Client
}

type StrattonData struct {
	Conn *websocket.Conn

	// Go-routine controller
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (sd *Stratton) Ping() {
	ping := messagePing()

	err := sd.Static.Conn.WriteJSON(ping)
	if err != nil {
		log.Printf("error writing message to websocket: %v", err)
	}
}

func (sd *Stratton) RequestStream() {
	msg, err := subscribeBuilder(true, depth5, "btcusdt")
	if err != nil {
		return
	}

	if err = sd.Stream.Conn.WriteJSON(msg); err != nil {
		log.Printf("failed to write subscribe message to stream: %v", err)
	}
}

func (sd *Stratton) RemoveStream() {
	msg, err := subscribeBuilder(false, depth5, "btcusdt")
	if err != nil {
		return
	}
	if err = sd.Stream.Conn.WriteJSON(msg); err != nil {
		log.Printf("failed to write unsubscribe message to stream: %v", err)
	}
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
