package strategy

import "fmt"

// Among depth, find part of the order book
// that can be easily overrun by push of a volume.
// This leads to spike + liquidation of
// opposite derivative position

func DepthAnalysis(depths [][]string) {
	// [ price in string, quantity in string ]
	// Look for quantity that is relatively sparse
	fmt.Println(len(depths))
}
