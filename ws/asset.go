package ws

import (
	"fmt"
	"github.com/skkugoon/strattonight/asset"
	"log"
	"nhooyr.io/websocket/wsjson"
)

type SingleTickerMethod string
type WholeMarketMethod string

const (
	WHOLE_FRCE WholeMarketMethod = "!forceOrder@arr"
	WHOLE_MARK WholeMarketMethod = "!markPrice@arr"
	WHOLE_MINI WholeMarketMethod = "!miniTicker@arr"
	WHOLE_TICK WholeMarketMethod = "!ticker@arr"
)

const (
	AGGREGATE   SingleTickerMethod = "aggTrade"  // Market trade information. Aggregated for fills with same price.
	MARK        SingleTickerMethod = "markPrice" // Mark price and funding rate. 1s or 3s
	DEPTH5      SingleTickerMethod = "depth5"    //
	DEPTH10     SingleTickerMethod = "depth10"
	DEPTH20     SingleTickerMethod = "depth20"
	MINI_TICKER SingleTickerMethod = "miniTicker" // mini ticker statistics
	BOOK_TICKER SingleTickerMethod = "bookTicker" // Update to best bid or ask price or quantity.
	LIQUIDATION SingleTickerMethod = "forceOrder" // Liquidation order.
)

type ParamTicker struct {
	MarketTicker WholeMarketMethod
	AssetTicker  asset.Asset
	CallFor      SingleTickerMethod
}

func (ws *ParamTicker) ToString() string {
	if ws.MarketTicker == "" {
		repr := fmt.Sprintf("%s@%s", ws.AssetTicker.ToString(), ws.CallFor)
		return repr
	} else {
		repr := fmt.Sprintf("%s", ws.MarketTicker)
		return repr
	}
}

func dataSubWholeForceOrder(ctx BinanceContext, _ []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createMarketMessage(WHOLE_FRCE).Subscribe()

	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send whole market force message: %v\n", err)
	}
	return msg
}

func dataSubWholeMarkPrice(ctx BinanceContext, _ []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createMarketMessage(WHOLE_MARK).Subscribe()

	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send whole market mark price message: %v\n", err)
	}
	return msg
}

func DataSubWholeMiniTicker(ctx BinanceContext, _ []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createMarketMessage(WHOLE_MINI).Subscribe()

	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send whole market mark price message: %v\n", err)
	}
	return msg
}

func DataSubWholeTicker(ctx BinanceContext, _ []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createMarketMessage(WHOLE_TICK).Subscribe()

	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send whole market tick price message: %v\n", err)
	}
	return msg
}

func depths(meth SingleTickerMethod, arr []asset.Asset) (BinanceMessage, *BinanceMessageFmt) {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(meth, arr...).Subscribe()
	return msg, fmtMsg
}

func dataSubDepth5(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg, fmtMsg := depths(DEPTH5, arr)
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send depth5 subscription message: %v\n", err)
	}
	return msg
}

func dataSubDepth10(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg, fmtMsg := depths(DEPTH10, arr)
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send depth10 subscription message: %v\n", err)
	}
	return msg
}

func dataSubDepth20(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg, fmtMsg := depths(DEPTH20, arr)
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send depth 20 subscription message: %v\n", err)
	}
	return msg
}

func DataSubMarkPrice(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(MARK, arr...).Subscribe()
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send mark price subscription message: %v\n", err)
	}
	return msg
}

func DataSubMiniTicker(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(MINI_TICKER, arr...).Subscribe()
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send mini ticker subscription message: %v\n", err)
	}
	return msg
}

func dataSubBookTicker(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(BOOK_TICKER, arr...).Subscribe()
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send book ticker subscription message: %v\n", err)
	}
	return msg
}

func dataSubLiquidation(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(LIQUIDATION, arr...).Subscribe()
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send liquidation subscription message: %v\n", err)
	}
	return msg
}

func dataSubAggTrade(ctx BinanceContext, arr []asset.Asset) BinanceMessage {
	msg := BinanceMessage{}
	fmtMsg := msg.createTickerMessage(AGGREGATE, arr...).Subscribe()
	err := wsjson.Write(ctx, ctx.Conn, fmtMsg)
	if err != nil {
		log.Printf("failed to send aggTrade subscription message: %v\n", err)
	}
	return msg
}
