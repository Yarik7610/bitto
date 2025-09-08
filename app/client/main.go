package client

import "github.com/codecrafters-io/bittorrent-starter-go/app/command"

type Client interface {
	HandleCommand(cmd string, args []string)
}

type client struct {
	commandController command.Controller
}

func New() Client {
	return &client{commandController: command.NewController()}
}

func (c *client) HandleCommand(cmd string, args []string) {
	c.commandController.HandleCommand(cmd, args)
}
