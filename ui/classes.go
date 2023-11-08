package ui

import "github.com/skkugoon/strattonight/data"

// Command interface
// As the functions supported in `stratton` increases,
// it will be difficult to maintain code with single `switch` phrase
// Use interface to combine the functions
type Command interface {
	Execute() error
}

type TestCommand struct{}

func (c *TestCommand) Execute() error {
	foo()
	return nil
}

type StaticPingCommand struct {
	live *data.Stratton
}

func (c *StaticPingCommand) Execute() error {
	err := c.live.Ping()
	return err
}

type DisplaySetupCommand struct {
	live *data.Stratton
}

func (c *DisplaySetupCommand) Execute() error {
	err := c.live.DisplaySetup()
	return err
}

type StreamInitTestCommand struct {
	live *data.Stratton
}

func (c *StreamInitTestCommand) Execute() error {
	c.live.RequestStream()
	return nil
}

type StreamExitTestCommand struct {
	live *data.Stratton
}

func (c *StreamExitTestCommand) Execute() error {
	c.live.RemoveStream()
	return nil
}
