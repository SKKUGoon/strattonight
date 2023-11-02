package data

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/skkugoon/strattonight/ent"
	"github.com/skkugoon/strattonight/strategy"
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

func (sd *Stratton) Ping() error {
	ping := messagePing()

	err := sd.Static.Conn.WriteJSON(ping)
	if err != nil {
		log.Printf("error writing message to websocket: %v", err)
		return err
	}
	return nil
}

func (sd *Stratton) RequestStream() {
	msg, err := subscribeBuilder(true, depth, "btcusdt")
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
			depth := responseDepth{}
			err := sd.Conn.ReadJSON(&depth)
			if err != nil {
				log.Fatalf("failed reading from websocket: %v", err)
			}
			// Testing for now
			log.Println(depth)
			strategy.DepthAnalysis(depth.A) // Ask
			strategy.DepthAnalysis(depth.B) // Bid
		}
	}
}

func (sd *StrattonData) Close() {
	err := sd.Conn.Close()
	if err != nil {
		log.Printf("graceful closing of websocket failed: %v", err)
	}
}
