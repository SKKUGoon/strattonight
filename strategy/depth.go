package strategy

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"strconv"
)

// Among depth, find part of the order book
// that can be easily overrun by push of a volume.
// This leads to spike + liquidation of
// opposite derivative position

// Based minute wise trading volume, check absolute volume
// that can be moved within 10 seconds

// Should we take account the volatility?
// Then we need klines - TODO: Ask for Klines with API
// Logic - How much volumes were utilized to move how much of a price?
func sparseDepthSize(minVolumes []float32, klines []float32) {
	log.Println("recording minute volume")
	// When minute volumes are given

	log.Println("recording kline volatility with ohlc information")
	//
}

func DepthAnalysis(depths [][]string, targetProfit float32, longShort bool) {
	// [ price in string, quantity in string ]
	// Look for quantity that is relatively sparse
	//fmt.Println(len(depths))
	if len(depths) <= 0 {
		return
	}

	profitRate := 1 + targetProfit*func() float32 {
		if longShort {
			return 1
		} else {
			return -1
		}
	}()

	firstDepth, _ := strconv.ParseFloat(depths[0][0], 32)
	targetPrice := float32(firstDepth) * profitRate
	fmt.Println("target price is ", firstDepth, profitRate, targetPrice)
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Define liquidity weight as price * quantity
	// Look for absolute price range and set a minimum scalping profit target
	for _, bookDepth := range depths {
		// Define data
		price, _ := strconv.ParseFloat(bookDepth[0], 32)

		quantity, _ := strconv.ParseFloat(bookDepth[1], 32)

		weight := price * quantity

		fmt.Printf("%v: weight %v   ", price, weight)

		if longShort {
			if float32(price) > targetPrice {
				fmt.Printf("%s\n", green("v"))
			} else {
				fmt.Printf("\n")
			}
		} else {
			if float32(price) < targetPrice {
				fmt.Printf("%s\n", red("v"))
			} else {
				fmt.Printf("\n")
			}
		}
	}
}
