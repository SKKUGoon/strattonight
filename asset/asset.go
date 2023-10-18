package asset

import (
	"fmt"
	"strings"
)

type Asset struct {
	Ticker   string
	Currency string
}

func (a *Asset) ToString() string {
	if a.Currency == "" {
		// Default currency is USDT
		a.Currency = "usdt"
	}

	// Default currency display is in lowercase
	repr := fmt.Sprintf(
		"%s%s",
		strings.ToLower(a.Ticker),
		strings.ToLower(a.Currency),
	)
	return repr
}

func ParseAsset(str string) []Asset {
	var result []Asset
	// Divide by ","
	assetStr := strings.Split(str, ",")
	for _, a := range assetStr {
		isQuote := strings.Split(a, "/")
		if len(isQuote) > 1 {
			// There's pre-designated quote currency
			result = append(result, Asset{Ticker: isQuote[0], Currency: isQuote[1]})
		} else {
			result = append(result, Asset{Ticker: isQuote[0]})
		}
	}
	return result
}
