package ui

import (
	"bufio"
	"fmt"
	"github.com/skkugoon/strattonight/data"
	"log"
	"os"
	"strings"
)

func CommandInterface(live *data.Stratton) {
	reader := bufio.NewReader(os.Stdin)

	// Print message retrieved from websocket
	go live.Static.ReadFromSocket()
	go live.Stream.ReadFromSocket()

	// Running an infinite loop
	for {
		fmt.Print("stratton >> ")
		text, _ := reader.ReadString('\n')

		// Lowercase + delete whitespace with trimmer
		text = strings.ToLower(strings.TrimSpace(text))

		switch text {
		case "test-command":
			foo()
		case "static-ping":
			live.Ping()
		case "stream-test-init":
			// TODO: Get input from user
			live.RequestStream()
		case "stream-test-exit":
			// TODO: Get input from user - check from the database.
			live.RemoveStream()

		case "exit", "quit":
			// Close all websocket clients with cancel func
			live.Static.Cancel()
			live.Static.Close()

			//live.Stream.Cancel()
			live.Stream.Close()
			return
		default:
			fmt.Println("unknown command:", text)
		}
	}
}

func foo() {
	log.Println("testing command interface")
}
