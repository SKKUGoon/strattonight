package ws

import (
	"context"
	"github.com/skkugoon/strattonight/asset"
)

var StreamCallMap = map[string]func([]asset.Asset, context.Context){
	"depth5":           streamDepth5,
	"depth10":          streamDepth10,
	"depth20":          streamDepth20,
	"aggTrade":         streamAggTrade,
	"bookTicker":       streamBookTicker,
	"forceOrder":       streamLiquidation,
	"marketForceOrder": streamWholeForceOrder,
	"marketMarkPrice":  streamWholeMarkPrice,
}

// Use Market depth information to find out dangling liquidity in the market
// If the market depth disappears rapidly, it usually means that short wave is coming.

func streamDepth5(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseDepth](a, btx, dataSubDepth5, processDepthDisplay)
}

func streamDepth10(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseDepth](a, btx, dataSubDepth10, processDepthDisplay)
}

func streamDepth20(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseDepth](a, btx, dataSubDepth20, processDepthIo)
}

// Use AggTrade to find out buy order or sell order
// Make transaction strength with this data stream

func streamAggTrade(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseAggTrade](a, btx, dataSubAggTrade, nil)
}

// Use book ticker to find out best bidding price and best asking price
// These are the prices that will determine market price order

func streamBookTicker(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseBookTicker](a, btx, dataSubBookTicker, nil)
}

// Liquidation order can be lagging signal of a derivative squeeze.

func streamLiquidation(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseMarketForceOrder](a, btx, dataSubLiquidation, processForceOrderDisplay)
}

func streamWholeForceOrder(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseMarketForceOrder](a, btx, dataSubWholeForceOrder, processForceOrderDisplay)
}

func streamWholeMarkPrice(a []asset.Asset, ctx context.Context) {
	btx := dataStream(ctx)
	go dataStreamReq1D[ResponseMarkPrice](a, btx, dataSubWholeMarkPrice, nil)
}
