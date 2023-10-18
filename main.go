package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/skkugoon/strattonight/asset"
	"github.com/skkugoon/strattonight/ws"
	"time"
)

func wss() {
	// Define Flags
	var callName string
	var callAsset string
	flag.StringVar(&callName, "name", "", "call which stream")
	flag.StringVar(&callAsset, "asset", "", "call which asset divided by comma(,)")

	flag.Parse()

	fmt.Println(callName, callAsset)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(time.Second * 2) // Wait for graceful exit
	}()

	// Spawn independent goroutine
	ws.StreamCallMap[callName](asset.ParseAsset(callAsset), ctx)

	time.Sleep(time.Hour * 2)
}

func main() {

}
