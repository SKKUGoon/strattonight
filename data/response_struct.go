package data

// ResponseAggTrade
// Example response should be
// ```
//
//	{
//	     "e":"aggTrade",
//	     "E":1685804889759,
//	     "a":1754196180,
//	     "s":"BTCUSDT",
//	     "p":"27177.80",
//	     "q":"0.005",
//	     "f":3776781864,
//	     "l":3776781864,
//	     "T":1685804889603,
//	     "m":false
//	}
//
// ```
// Is Buyer the Market Maker ?
// Source: https://www.bitpanda.com/academy/en/lessons/what-is-a-market-maker/
// Source: https://dev.binance.vision/t/a-little-confusion-about-is-buyer-market-maker-part-of-trade-stream/5660
//
// Market makers earn through trading in spreads.
// MM buys from seller at (maker) price    displayed as  -> true
// and delivers to buyer at (maker) price  displayed as  -> false
// Therefore `true` can be interpreted as `selling order`, and vice versa
// ```
//
//	interaction    |   buyer (maker)  |  buyer (taker)
//
// ----------------|------------------|------------------
// seller (maker)  |    no match      |  false
// seller (taker)  |    true          |  no match
// ```
type ResponseAggTrade struct {
	E  string `json:"e"` // Event type
	E1 int64  `json:"E"` // Event time
	A  int    `json:"a"` // Aggregate trade ID
	S  string `json:"s"` // Symbol
	P  string `json:"p"` // Price
	Q  string `json:"q"` // Quantity
	F  int64  `json:"f"` // First trade ID
	L  int64  `json:"l"` // Last trade ID
	T  int64  `json:"T"` // Trade time
	M  bool   `json:"m"` // Is the buyer the market maker?
}

// ResponseMarkPrice
// Example response should be.
// ```
//
//	{
//	     "e":"markPriceUpdate",
//	     "E":1685804340019,
//	     "s":"BTCUSDT",
//	     "p":"27157.85214815",
//	     "P":"27182.13215031",
//	     "i":"27173.52933333",
//	     "r":"0.00004928",
//	     "T":1685808000000
//	}
//
// ```
type ResponseMarkPrice struct {
	E  string `json:"e"` // Event type
	E1 int64  `json:"E"` // Event time
	S  string `json:"s"` // Symbol
	P  string `json:"p"` // markPrice
	P1 string `json:"P"` // Estimated Settle Price
	I  string `json:"i"` // Index price
	R  string `json:"r"` // Funding rate
	T  int64  `json:"T"` // Next funding time
}

// ResponseDepth
// Example response should be.
// ```
//
//	{
//	     "e":"depthUpdate",
//	     "E":1685804359113,
//	     "T":1685804359105,
//	     "s":"BTCUSDT",
//	     "U":2908041773679,
//	     "u":2908041779294,
//	     "pu":2908041773669,
//	     "b":[["2715.60","0.003"], ...],
//	     "a":[["27161.40","0.463"], ...],
//	}
//
// ```
type responseDepth struct {
	E  string     `json:"e"` // Event type
	E1 int64      `json:"E"` // Event time
	T  int64      `json:"T"` // Transaction time
	S  string     `json:"s"` // Symbol
	U  int64      `json:"U"`
	U1 int64      `json:"u"`
	Pu int64      `json:"pu"`
	B  [][]string `json:"b"` // Bids to be updated [ Price level, Quantity ]
	A  [][]string `json:"a"` // Asks to be updated [ Price level, Quantity ]
}

// ResponseMiniTicker
// Example response would be
// ```
//
//	{
//	     "e":"24hrMiniTicker",
//	     "E":1685804417219,
//	     "s":"BTCUSDT",
//	     "c":"27162.20",
//	     "o":"26952.00",
//	     "h":"27290.00",
//	     "l":"26949.50",
//	     "v":"174598.979",
//	     "q":"4738831280.77",
//	}
//
// ```
type ResponseMiniTicker struct {
	E  string `json:"e"`
	E1 int64  `json:"E"`
	S  string `json:"s"` // Symbol
	C  string `json:"c"` // Close price
	O  string `json:"o"` // Open price
	H  string `json:"h"` // High price
	L  string `json:"l"` // Low price
	V  string `json:"v"` // Total traded base asset volume (traded itself)
	Q  string `json:"q"` // Total traded quote asset volume (USDT base volume)
}

// ResponseBookTicker
// Push update to the best bid or ask price or quantity
// Example response would be
// ```
//
//	{
//	    "e":"bookTicker",
//	    "u":2917970627001,
//	    "s":"BTCUSDT",
//	    "b":"25850.10",
//	    "B":"1.003",
//	    "a":"25850.20",
//	    "A":"13.762",
//	    "T":1686061192228,
//	    "E":1686061192233
//	}
//
// ```
type ResponseBookTicker struct {
	E  string `json:"e"`
	U  int64  `json:"u"` // Order book update ID
	S  string `json:"s"` // Asset name
	B  string `json:"b"` // best bid price
	B1 string `json:"B"` // best bid quantity
	A  string `json:"a"` // best ask price
	A1 string `json:"A"` // best ask quantity
	T  int64  `json:"T"` // transaction time
	E1 int64  `json:"E"` // event time
}

// ResponseMarketForceOrder
// Liquidation order snapshot. Only the latest one liquidation order within 1000ms will be pushed.
// Example response would be
// ```
//
//	{
//	     "e":"forceOrder",
//	     "E":1685859893251,
//	     "o":{
//	          "s":"OCEANUSDT",
//	          "S":"SELL",
//	          "o":"LIMIT",
//	          "f":"IOC",
//	          "q":"945",
//	          "p":"0.40173",
//	          "ap":"0.40990",
//	          "X":"FILLED",
//	          "l":"64",
//	          "z":"945",
//	          "T":1685859893245
//	         }
//	}
//
// ```
type ResponseMarketForceOrder struct {
	E  string `json:"e"`
	E1 int64  `json:"E"` // Event Time
	O  struct {
		S  string `json:"s"`  // Symbol
		S1 string `json:"S"`  // Side
		O  string `json:"o"`  // Order Type = LIMIT
		F  string `json:"f"`  // Time in Force = IOC (Immediate or cancel)
		Q  string `json:"q"`  // Original Quantity
		P  string `json:"p"`  // Price
		Ap string `json:"ap"` // Average Price
		X  string `json:"X"`  // Order Status
		L  string `json:"l"`  // Order Last filled quantity
		Z  string `json:"z"`  // Order filled accumulated quantity
		T  int64  `json:"T"`  // Order trade time
	} `json:"o"`
}
