package ws

import (
	"context"
	"encoding/json"
	"github.com/skkugoon/strattonight/asset"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"time"
)

const BINANCE = "wss://fstream.binance.com/ws"

type BinanceContext struct {
	context.Context
	Conn *websocket.Conn
}

// dataStream
// Parent context `ctx` is connected to each websocket dial connection
// If the parent context's cancel() is called, it will exit simultaneously
func dataStream(ctx context.Context) BinanceContext {
	dialOpts := &websocket.DialOptions{
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		CompressionMode: websocket.CompressionDisabled,
	}

	conn, _, err := websocket.Dial(ctx, BINANCE, dialOpts)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	bctx := BinanceContext{ctx, conn}
	return bctx
}

func dataStreamReq1D[T any](
	req []asset.Asset,
	ctx BinanceContext,
	f func(BinanceContext, []asset.Asset) BinanceMessage,
	proc func(T) error,
) {
	_ = f(ctx, req)                     // Send websocket request
	go dataStreamReader1D[T](ctx, proc) // Infinite loop with stop context
}

func dataStreamReader1D[T any](ctx BinanceContext, process func(T) error) {
	defer func() {
		err := ctx.Conn.Close(websocket.StatusInternalError, "")
		if err != nil {
			log.Printf("dataStreamReader1D websocket conn closing error: %v\n", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("dataStreamReader1D exit gracefully called")
			return

		default:
			var recvParse T
			// After 10 error occurs, break out of loop
			errorCount := 0
			for errorCount < 10 {
				_, msgBytes, err := ctx.Conn.Read(ctx)
				if err != nil {
					log.Printf("failed to read message: %v\n", err)
					errorCount += 1
				}

				err = json.Unmarshal(msgBytes, &recvParse)
				if err != nil {
					log.Printf("marshaling error: %v\n", err)
					errorCount += 1
				}

				// Display data
				if process != nil {
					err = process(recvParse)
					if err != nil {
						// no data error
						errorCount += 1
					}
				} else {
					log.Printf("Structuralized %v\n", recvParse)
				}
			}
		}
	}
}

func dataStreamReq2D[T any](
	req []asset.Asset,
	ctx BinanceContext,
	f func(BinanceContext, []asset.Asset) BinanceMessage,
) {
	_ = f(ctx, req)
	dataStreamReader2D(ctx)
}

// For now depth only
func dataStreamReader2D(ctx BinanceContext) {
	defer func() {
		err := ctx.Conn.Close(websocket.StatusInternalError, "")
		if err != nil {
			log.Printf("dataStreamReader2D websocket conn closing error: %v\n", err)
		}
	}()

	d1 := RaceSafeStore{}
	d2 := RaceSafeStore{}

	// Read from mutex values
	ticking := 5 * time.Second
	go processDepthTimer(ctx, ticking, &d1, &d2)

	for {
		select {
		case <-ctx.Done():
			log.Println("dataStream2D exit gracefully called")
			return
		default:
			var recvParse ResponseDepth
			// After 10 error counts, break out of loop
			errorCount := 0
			for errorCount < 10 {
				_, msgBytes, err := ctx.Conn.Read(ctx)
				if err != nil {
					log.Printf("failed to read message: %v\n", err)
					errorCount += 1
				}

				err = json.Unmarshal(msgBytes, &recvParse)
				if err != nil {
					log.Printf("marshaling error: %v\n", err)
				}

				processDepthIoSto(recvParse, &d1, &d2)
			}
		}
	}
}
