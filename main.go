package main

import (
	"context"
	"github.com/skkugoon/strattonight/data"
	"github.com/skkugoon/strattonight/ui"
)

func main() {
	ctx := context.Background()

	stratton := data.WebsocketClient(false, ctx)

	ui.CommandInterface(&stratton)
}
