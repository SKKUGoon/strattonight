package user

type WssTradePacket[T any] struct {
	ID     string       `json:"id"`
	Method MethodsTrade `json:"method"`
	Params T            `json:"params"`
}

type PlaceNewOrderReq struct {
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	Type        string `json:"type"`
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	ApiKey      string `json:"apiKey"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

type TestNewOrderReq struct {
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	Type        string `json:"type"`
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	ApiKey      string `json:"apiKey"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

type OrderStatusReq struct {
	Symbol    string `json:"symbol"`
	OrderId   int64  `json:"orderId"`
	ApiKey    string `json:"apiKey"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

type CancelOrderReq struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderId string `json:"origClientOrderId"`
	ApiKey            string `json:"apiKey"`
	Signature         string `json:"signature"`
	Timestamp         int64  `json:"timestamp"`
}

type CancelReplaceOrderReq struct {
	Symbol                  string `json:"symbol"`
	CancelReplaceMode       string `json:"cancelReplaceMode"`
	CancelOrigClientOrderId string `json:"cancelOrigClientOrderId"`
	Side                    string `json:"side"`
	Type                    string `json:"type"`
	TimeInForce             string `json:"timeInForce"`
	Price                   string `json:"price"`
	Quantity                string `json:"quantity"`
	ApiKey                  string `json:"apiKey"`
	Signature               string `json:"signature"`
	Timestamp               int64  `json:"timestamp"`
}

type WssTradeRespPacket[T any] struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	WorkingTime         int64  `json:"workingTime"`
	Fills               []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int    `json:"tradeId"`
	} `json:"fills"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type AckNewOrder struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		OrderListId   int    `json:"orderListId"`
		ClientOrderId string `json:"clientOrderId"`
		TransactTime  int64  `json:"transactTime"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

type ResultNewOrder struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		Symbol                  string `json:"symbol"`
		OrderId                 int64  `json:"orderId"`
		OrderListId             int    `json:"orderListId"`
		ClientOrderId           string `json:"clientOrderId"`
		TransactTime            int64  `json:"transactTime"`
		Price                   string `json:"price"`
		OrigQty                 string `json:"origQty"`
		ExecutedQty             string `json:"executedQty"`
		CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
		Status                  string `json:"status"`
		TimeInForce             string `json:"timeInForce"`
		Type                    string `json:"type"`
		Side                    string `json:"side"`
		WorkingTime             int64  `json:"workingTime"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

type FullNewOrder struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		Symbol              string `json:"symbol"`
		OrderId             int64  `json:"orderId"`
		OrderListId         int    `json:"orderListId"`
		ClientOrderId       string `json:"clientOrderId"`
		TransactTime        int64  `json:"transactTime"`
		Price               string `json:"price"`
		OrigQty             string `json:"origQty"`
		ExecutedQty         string `json:"executedQty"`
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status              string `json:"status"`
		TimeInForce         string `json:"timeInForce"`
		Type                string `json:"type"`
		Side                string `json:"side"`
		WorkingTime         int64  `json:"workingTime"`
		Fills               []struct {
			Price           string `json:"price"`
			Qty             string `json:"qty"`
			Commission      string `json:"commission"`
			CommissionAsset string `json:"commissionAsset"`
			TradeId         int    `json:"tradeId"`
		} `json:"fills"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

type TestNewOrder struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

type OrderState struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
		Symbol                  string `json:"symbol"`
		OrderId                 int64  `json:"orderId"`
		OrderListId             int    `json:"orderListId"`
		ClientOrderId           string `json:"clientOrderId"`
		Price                   string `json:"price"`
		OrigQty                 string `json:"origQty"`
		ExecutedQty             string `json:"executedQty"`
		CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
		Status                  string `json:"status"`
		TimeInForce             string `json:"timeInForce"`
		Type                    string `json:"type"`
		Side                    string `json:"side"`
		StopPrice               string `json:"stopPrice"`
		IcebergQty              string `json:"icebergQty"`
		Time                    int64  `json:"time"`
		UpdateTime              int64  `json:"updateTime"`
		IsWorking               bool   `json:"isWorking"`
		WorkingTime             int64  `json:"workingTime"`
		OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}
