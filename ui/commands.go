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

	commands := map[string]Command{
		"test-command":     &TestCommand{},
		"static-ping":      &StaticPingCommand{live},
		"stream-test-init": &StreamInitTestCommand{live},
		"stream-test-exit": &StreamExitTestCommand{live},
	}

	// Running an infinite loop
	for {
		fmt.Print("stratton >> ")
		text, _ := reader.ReadString('\n')

		// Lowercase + delete whitespace with trimmer
		text = strings.ToLower(strings.TrimSpace(text))

		// Check if the command exists in the `commands` map
		if command, ok := commands[text]; ok {
			err := command.Execute()
			if err != nil {
				log.Printf("command execution error: %v", err)
				continue
			}
		} else {
			switch text {
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
}

func foo() {
	log.Println("testing command interface")
}
