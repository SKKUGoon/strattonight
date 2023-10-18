package user

// Methods for enum values

type MethodsNoParam string
type MethodsParamSymbol string
type MethodsTrade string
type MethodsAccount string
type MethodsUserDS string

// Constant types

type SymbolStatus string
type AccountPermissions string
type OrderStatus string
type OrderTypes string
type TimeInForce string

const (
	Ping       MethodsNoParam = "ping"
	ServerTime MethodsNoParam = "time"
)

const ExchangeInfo MethodsParamSymbol = "exchangeInfo"

const (
	OrderPlace        MethodsTrade = "order.place"
	TestOrderPlace    MethodsTrade = "order.test"
	OrderStat         MethodsTrade = "order.status"
	OrderCancel       MethodsTrade = "order.cancel"
	CurrentOpen       MethodsTrade = "openOrders.status"
	CurrentOpenCancel MethodsTrade = "openOrders.cancelAll"
)

const (
	DataStart MethodsUserDS = "userDataStream.start"
	DataPing  MethodsUserDS = "userDataStream.ping"
	DataStop  MethodsUserDS = "userDataStream.stop"
)

const (
	AccountStat MethodsAccount = "account.status"
)

const (
	PreTrading   SymbolStatus = "PRE_TRADING"
	Trading      SymbolStatus = "TRADING"
	PostTrading  SymbolStatus = "POST_TRADING"
	EndOfDay     SymbolStatus = "END_OF_DAY"
	Halt         SymbolStatus = "HALT"
	AuctionMatch SymbolStatus = "AUCTION_MATCH"
	Break        SymbolStatus = "BREAK"
)

const (
	Spot      AccountPermissions = "SPOT"
	Margin    AccountPermissions = "MARGIN"
	Leveraged AccountPermissions = "LEVERAGED"
)

const (
	New           OrderStatus = "NEW"
	PartialFilled OrderStatus = "PARTIALLY_FILLED"
	Filled        OrderStatus = "FILLED"
	Canceled      OrderStatus = "CANCELED"
	PendingCancel OrderStatus = "PENDING_CANCEL"
	Rejected      OrderStatus = "REJECTED"
	Expired       OrderStatus = "EXPIRED"
)

const (
	Limit           OrderTypes = "LIMIT"
	LimitMaker      OrderTypes = "LIMIT_MAKER"
	Market          OrderTypes = "MARKET"
	StopLoss        OrderTypes = "STOP_LOSS"
	StopLossLimit   OrderTypes = "STOP_LOSS_LIMIT"
	TakeProfit      OrderTypes = "TAKE_PROFIT"
	TakeProfitLimit OrderTypes = "TAKE_PROFIT_LIMIT"
)

// OrderCheckFor
// Each `OrderTypes` needs mandatory parameters.
// Or conditions are joined with ","
var OrderCheckFor = map[OrderTypes][]string{
	Limit:           {"timeInForce", "price", "quantity"},
	LimitMaker:      {"price", "quantity"},
	Market:          {"quantity,quoteOrderQty"},
	StopLoss:        {"quantity", "stopPrice,trailingDelta"},
	StopLossLimit:   {"timeInForce", "price", "quantity", "stopPrice,trailingDelta"},
	TakeProfit:      {"quantity", "stopPrice,trailingDelta"},
	TakeProfitLimit: {"timeInForce", "price", "quantity", "stopPrice,trailingDelta"},
}

const (
	// GTC
	// Good till canceled.
	// Order will remain on the book until you cancel it or filled
	GTC TimeInForce = "GTC"

	// IOC
	// Immediate or Cancel
	// Order will be filled for as much as possible. The unfilled quantity immediately expires
	IOC TimeInForce = "IOC"

	// FOK
	// Fill of Kill.
	// Order will expire unless it cannot be immediately filled for the entire quantity
	FOK TimeInForce = "FOK"
)
