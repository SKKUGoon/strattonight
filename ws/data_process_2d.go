package ws

import (
	"context"
	"log"
	"strconv"
	"time"
)

func processDepthIoSto(d ResponseDepth, askSto, bidSto *RaceSafeStore) {
	bid := d.B
	ask := d.A
	if len(bid) == 0 && len(ask) == 0 {
		return
	}

	var (
		bidHanging float64 = 0
		askHanging float64 = 0
	)
	for _, b := range bid {
		p, _ := strconv.ParseFloat(b[0], 64)
		q, _ := strconv.ParseFloat(b[1], 64)
		bidHanging += p * q
	}

	for _, a := range ask {
		p, _ := strconv.ParseFloat(a[0], 64)
		q, _ := strconv.ParseFloat(a[1], 64)
		askHanging += p * q
	}

	askSto.NumberAdd(askHanging)
	bidSto.NumberAdd(bidHanging)
}

func processDepthTimer(ctx context.Context, t time.Duration, askSto, bidSto *RaceSafeStore) {
	ticking := time.NewTicker(t)
	for {
		select {
		case <-ctx.Done():
			log.Println("process stopped upon request by context")
			return
		case <-ticking.C:
			log.Printf("ticker process after %v seconds\n", t)
			log.Printf("Ask Store: %v\n", askSto.Number)
			log.Printf("Bid Store: %v\n", bidSto.Number)
			askSto.NumberSet(0)
			bidSto.NumberSet(0)
			return
		}
	}
}
