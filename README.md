# Welcome to Strattonight

## Strattonight

Project was named after famous fraudulent investing firm Stratton Oakmont. 
Using gRPC and other techniques, this repository will be a shelter for all
investors who wants to test their strategy for limited amount of time.

## Architecture

With use of gRPC and protobuf, Strattonight supports multi-language for all
strattonights who are willing to test their strategy. 

There will be a function that generates order types. 
The function will take these parameters:

1. `strategy_name`: `string`
2. `asset_whitelist`: `[]string`: asset for trading
3. `asset_blacklist`: `[]string`: asset for not trading. Both `asset whitelist` and `asset blacklist` can co-exist, 
but at least one of two must exist.
4. `data_to_watch`: `map[string]string`: Key - names of data for monitoring. Value - threshold like condition.

And For added concern for concurrency golang's `context.Context` will be used.

## Calling Data

```bash
go run main.go -name=depth5 -asset=ETH,BTC/BUSD  
```

Flags `name` and `asset` are used to demonstrate message queue calls later.
