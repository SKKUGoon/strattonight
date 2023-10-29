```go
package ws

// Process data received from websocket call
// Each data stream possesses unique processing method
// 1d processing means 1 dimensional process
// where only summation, and multiplications are utilized with
// no shared variable among each websocket messages

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func processDepthDisplay(d ResponseDepth) error {
	name := d.S
	buy := d.B
	ask := d.A
	if len(buy) == 0 && len(ask) == 0 {
		return errors.New("no data received through depth call")
	}

	// price : price * quantity (hanging volume)
	buyDepth := map[float64]float64{}
	askDepth := map[float64]float64{}

	for _, b := range buy {
		p, _ := strconv.ParseFloat(b[0], 64)
		q, _ := strconv.ParseFloat(b[1], 64)
		buyDepth[p] = p * q
	}

	for _, a := range ask {
		p, _ := strconv.ParseFloat(a[0], 64)
		q, _ := strconv.ParseFloat(a[1], 64)
		askDepth[p] = p * q
	}

	log.Printf("%s buy depth: %v\n", name, buyDepth) // Lower order book
	log.Printf("%s ask depth: %v\n", name, askDepth) // Upper order book
	fmt.Println()
	return nil
}

func processForceOrderDisplay(d ResponseMarketForceOrder) error {
	forced := d.O
	log.Printf("%s Forced %s order. Quantity %s\n", forced.S, forced.S1, forced.Q)
	return nil
}

func processDepthIo(d ResponseDepth) error {
	name := d.S
	bid := d.B
	ask := d.A
	if len(bid) == 0 && len(ask) == 0 {
		return errors.New("no data received through depth call")
	}

	// Hanging liquidity means, market makers are anticipating
	var buyHanging float64 = 0
	var askHanging float64 = 0

	for _, b := range bid {
		p, _ := strconv.ParseFloat(b[0], 64)
		q, _ := strconv.ParseFloat(b[1], 64)
		buyHanging += p * q
	}

	for _, a := range ask {
		p, _ := strconv.ParseFloat(a[0], 64)
		q, _ := strconv.ParseFloat(a[1], 64)
		askHanging += p * q
	}

	log.Printf("%s bid hanging :%v | ask hanging :%v\n", name, buyHanging, askHanging)
	fmt.Println()
	return nil
}

```