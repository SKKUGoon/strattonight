package ws

// Websocket needs to receive data subscription messages.
// Binance's subscription message takes 3 parts.
// `BinanceMessage` class stores them in struct for easy manipulation
// `BinanceMessageFmt` is derived from `BinanceMessage` struct

import (
	"encoding/json"
	"github.com/skkugoon/strattonight/asset"
	"log"
	"math/rand"
	"time"
)

type MsgConstructor interface {
	BinanceMessage
}

type BinanceMessage struct {
	Method string        `json:"method"`
	Params []ParamTicker `json:"params"`
	Id     int           `json:"id"`
}

type MsgFormat interface {
	BinanceMessageFmt
}

type BinanceMessageFmt struct {
	// JSON Message struct that's actually sent to Binance
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int      `json:"id"`
}

type MsgInterface interface {
	SubscribeByte() (int, []byte, error)
	UnsubscribeByte() (int, []byte, error)
}

func (bm *BinanceMessage) createTickerMessage(info SingleTickerMethod, assets ...asset.Asset) *BinanceMessage {
	for _, i := range assets {
		p := ParamTicker{
			AssetTicker: i,
			CallFor:     info,
		}
		bm.Params = append(bm.Params, p)
	}
	return bm
}

func (bm *BinanceMessage) createMarketMessage(marketInfos ...WholeMarketMethod) *BinanceMessage {
	for _, i := range marketInfos {
		p := ParamTicker{
			MarketTicker: i,
		}
		bm.Params = append(bm.Params, p)
	}
	return bm
}

func (bm *BinanceMessage) Subscribe() *BinanceMessageFmt {
	bm.Method = "SUBSCRIBE"

	rand.NewSource(time.Now().UnixNano())
	randomInt := rand.Intn(51)
	bm.Id = randomInt

	form := BinanceMessageFmt{
		Id:     bm.Id,
		Method: bm.Method,
		Params: func() []string {
			var result []string
			for _, i := range bm.Params {
				result = append(result, i.ToString())
			}
			return result
		}(),
	}
	log.Printf("subscription request sent: %v", form)
	return &form
}

func (bm *BinanceMessage) Unsubscribe() *BinanceMessageFmt {
	bm.Method = "UNSUBSCRIBE"

	form := BinanceMessageFmt{
		Id:     bm.Id,
		Method: bm.Method,
		Params: func() []string {
			var result []string
			for _, i := range bm.Params {
				result = append(result, i.ToString())
			}
			return result
		}(),
	}

	return &form
}

// SubscribeByte Returns websocket subscription message id, json message and error.
func (bm *BinanceMessage) SubscribeByte() (int, []byte, error) {
	bm.Method = "SUBSCRIBE"

	// Generate a random integer between 0 and 50 (inclusive)
	rand.NewSource(time.Now().UnixNano())
	randomInt := rand.Intn(51)
	bm.Id = randomInt

	b, err := json.Marshal(bm)
	return bm.Id, b, err
}

// UnsubscribeByte Returns websocket unsubscription message id, json message and error.
func (bm *BinanceMessage) UnsubscribeByte() (int, []byte, error) {
	bm.Method = "UNSUBSCRIBE"

	b, err := json.Marshal(bm)
	return bm.Id, b, err
}
