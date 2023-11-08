package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/skkugoon/strattonight/ent"
	"github.com/skkugoon/strattonight/strategy"
	"log"
)

type Stratton struct {
	Setup *tradeSetup

	// Binance websocket connections
	Static *StrattonData
	Stream *StrattonData

	// Local(Cloud) database connections
	Local *ent.Client
}

type StrattonData struct {
	Conn *websocket.Conn

	// Go-routine controller
	Ctx    context.Context
	Cancel context.CancelFunc
}

type tradeSetup struct {
	Leverage float32 `json:"leverage"`
	Profit   float32 `json:"profit"`
	Slippage float32 `json:"slippage"`
	Fee      float32 `json:"fee"`
}

func (t *tradeSetup) TargetProfit() float32 {
	// Calculate target price to achieve profit assigned in tradeSetup
	return t.Profit/t.Leverage + t.Slippage/t.Leverage + (t.Fee)*t.Leverage
}

func (sd *Stratton) Ping() error {
	ping := messagePing()

	err := sd.Static.Conn.WriteJSON(ping)
	if err != nil {
		log.Printf("error writing message to websocket: %v\n", err)
		return err
	}
	return nil
}

func (sd *Stratton) DisplaySetup() error {
	str, err := json.Marshal(sd.Setup)
	if err != nil {
		log.Printf("failed to marshal setup json: %v\n", err)
		return err
	}
	log.Println(sd.Setup)
	log.Printf("%s", string(str))
	return nil
}

func (sd *Stratton) RequestStream() {
	msg, err := subscribeBuilder(true, depth20, "btcusdt")
	if err != nil {
		return
	}

	if err = sd.Stream.Conn.WriteJSON(msg); err != nil {
		log.Printf("failed to write subscribe message to stream: %v", err)
	}
}

func (sd *Stratton) RemoveStream() {
	msg, err := subscribeBuilder(false, depth20, "btcusdt")
	if err != nil {
		return
	}
	if err = sd.Stream.Conn.WriteJSON(msg); err != nil {
		log.Printf("failed to write unsubscribe message to stream: %v", err)
	}
}

func (sd *StrattonData) ReadFromSocket(margin float32) {
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
			//log.Println(depth)
			fmt.Println("=== ASK ===")
			strategy.DepthAnalysis(depth.A, margin, true) // Ask, Long - true
			fmt.Println("=== BID ===")
			strategy.DepthAnalysis(depth.B, margin, false) // Bid, Short - false
		}
	}
}

func (sd *StrattonData) Close() {
	err := sd.Conn.Close()
	if err != nil {
		log.Printf("graceful closing of websocket failed: %v", err)
	}
}
