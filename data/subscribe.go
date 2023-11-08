package data

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"strings"
)

type SingleTickerMethod string
type WholeMarketMethod string

const (
	marketLiquidation WholeMarketMethod = "!forceOrder@arr"
	marketMarkPrice   WholeMarketMethod = "!markPrice@arr"
	marketMiniTicker  WholeMarketMethod = "!miniTicker@arr"
	marketTicker      WholeMarketMethod = "!ticker@arr"
)

const (
	aggTrade    SingleTickerMethod = "aggTrade"  // Market trade information. Aggregated for fills with same price.
	markPrice   SingleTickerMethod = "markPrice" // Mark price and funding rate. 1s or 3s
	depth       SingleTickerMethod = "depth"
	depth5      SingleTickerMethod = "depth5" //
	depth10     SingleTickerMethod = "depth10"
	depth20     SingleTickerMethod = "depth20"
	miniTicker  SingleTickerMethod = "miniTicker" // mini ticker statistics
	bookTicker  SingleTickerMethod = "bookTicker" // Update to best bid or ask price or quantity.
	liquidation SingleTickerMethod = "forceOrder" // Liquidation order.
)

type subscribeForm struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int      `json:"id"`
}

func subscribeBuilder(subscribe bool, method SingleTickerMethod, tickers ...string) (subscribeForm, error) {
	var streams []string
	for _, ticker := range tickers {
		streamName := fmt.Sprintf("%s@%s", strings.ToLower(ticker), method)
		streams = append(streams, streamName)
	}

	if len(streams) <= 0 {
		// No tickers given
		return subscribeForm{}, errors.New("no tickers given")
	}

	form := subscribeForm{
		Method: func(isSub bool) string {
			if isSub {
				return "SUBSCRIBE"
			} else {
				return "UNSUBSCRIBE"
			}
		}(subscribe),
		Params: streams,
		Id:     0,
	}

	// Report
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("%s to tickers %s\n", green(form.Method), green(strings.Join(streams, ", ")))

	return form, nil
}
