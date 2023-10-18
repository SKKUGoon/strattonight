package user

type ReqNoParam struct {
	Id     string         `json:"id"`
	Method MethodsNoParam `json:"method"`
}

type RespNoParam struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		ServerTime int `json:"serverTime,omitempty"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

type ReqParamSymbol struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Params struct {
		Symbols []string `json:"symbols"`
	} `json:"params"`
}

type RespParamSymbol struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		Timezone   string `json:"timezone"`
		ServerTime int64  `json:"serverTime"`
		RateLimits []struct {
			RateLimitType string `json:"rateLimitType"`
			Interval      string `json:"interval"`
			IntervalNum   int    `json:"intervalNum"`
			Limit         int    `json:"limit"`
		} `json:"rateLimits"`
		ExchangeFilters []interface{} `json:"exchangeFilters"`
		Symbols         []struct {
			Symbol                     string   `json:"symbol"`
			Status                     string   `json:"status"`
			BaseAsset                  string   `json:"baseAsset"`
			BaseAssetPrecision         int      `json:"baseAssetPrecision"`
			QuoteAsset                 string   `json:"quoteAsset"`
			QuotePrecision             int      `json:"quotePrecision"`
			QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
			BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
			QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
			OrderTypes                 []string `json:"orderTypes"`
			IcebergAllowed             bool     `json:"icebergAllowed"`
			OcoAllowed                 bool     `json:"ocoAllowed"`
			QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
			AllowTrailingStop          bool     `json:"allowTrailingStop"`
			CancelReplaceAllowed       bool     `json:"cancelReplaceAllowed"`
			IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
			IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
			Filters                    []struct {
				FilterType string `json:"filterType"`
				MinPrice   string `json:"minPrice,omitempty"`
				MaxPrice   string `json:"maxPrice,omitempty"`
				TickSize   string `json:"tickSize,omitempty"`
				MinQty     string `json:"minQty,omitempty"`
				MaxQty     string `json:"maxQty,omitempty"`
				StepSize   string `json:"stepSize,omitempty"`
			} `json:"filters"`
			Permissions                     []string `json:"permissions"`
			DefaultSelfTradePreventionMode  string   `json:"defaultSelfTradePreventionMode"`
			AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes"`
		} `json:"symbols"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}
