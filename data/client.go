package data

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"strconv"
)

const (
	mainStaticServer = "ws-api.binance.com:443"
	testStaticServer = "testnet.binance.vision"
	mainStreamServer = "fstream.binance.com"
	staticPath       = "/ws-api/v3"
	streamPath       = "/ws"
	scheme           = "wss"
)

func getTradeEnvironment() *tradeSetup {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("failed to load environment files: %v", err)
	}

	log.Printf("fetching trade environment\n")

	leverage, _ := strconv.ParseFloat(os.Getenv("LEVERAGE"), 32)
	targetProfit, _ := strconv.ParseFloat(os.Getenv("PROFIT"), 32)
	expectedSlippage, _ := strconv.ParseFloat(os.Getenv("SLIPPAGE"), 32)
	tradingFee, _ := strconv.ParseFloat(os.Getenv("FEE"), 32)

	return &tradeSetup{
		Leverage: float32(leverage),
		Profit:   float32(targetProfit),
		Slippage: float32(expectedSlippage),
		Fee:      float32(tradingFee),
	}
}

func WebsocketClient(test bool, ctx context.Context) Stratton {
	staticConn := staticClient(test, ctx)
	streamConn := streamClient(ctx)

	setups := getTradeEnvironment()

	return Stratton{
		Setup:  setups,
		Static: staticConn,
		Stream: streamConn,
	}
}

func streamClient(ctx context.Context) *StrattonData {
	c, cancel := context.WithCancel(ctx)

	// Generate websocket connection URL - for stream
	u := url.URL{
		Scheme: scheme,
		Host:   mainStreamServer,
		Path:   streamPath,
	}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dialing failed: %v", err)
	}

	// Single connection is valid for 24 hours.
	// Expect to be disconnected after 24 hour mark
	// Websocket server will send `ping frame` every 3 minutes
	// Gorilla's websocket automatically respond to received ping message with a pong/
	return &StrattonData{
		Conn:   conn,
		Ctx:    c,
		Cancel: cancel,
	}
}

func staticClient(test bool, ctx context.Context) *StrattonData {
	c, cancel := context.WithCancel(ctx)

	// Generate websocket connection URL
	u := url.URL{
		Scheme: scheme,
		Host: func(test bool) string {
			if test {
				return testStaticServer
			} else {
				return mainStaticServer
			}
		}(test),
		Path: staticPath,
	}
	log.Printf("connecting to %s", u.String())

	// Create websocket client
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dialing failed: %v", err)
	}

	// Single connection is valid for 24 hours.
	// Expect to be disconnected after 24 hour mark
	// Websocket server will send `ping frame` every 3 minutes
	// Gorilla's websocket automatically respond to received ping message with a pong/
	return &StrattonData{
		Conn:   conn,
		Ctx:    c,
		Cancel: cancel,
	}
}
