package server

import (
	"strings"

	"github.com/soupaulodev/chat-server/utils"
)

func (c *Client) handleCommand(cmd string) {
	parts := strings.SplitN(cmd, " ", 2)
	switch parts[0] {
		case "/list":
			c.listClients()
		case "/nick":
			if len(parts) > 1 {
				c.changeNickname(parts[1])
			} else {
				c.SendMessage("Usage: /nick <new_nickname>")
			}
		default:
			c.SendMessage("Unknown command: " + parts[0])
	}
}

func (c *Client) listClients() {
	clients := []string{}

	c.server.mut.Lock()

	for nickname := range c.server.clients {
		clients = append(clients, nickname)
	}

	c.server.mut.Unlock()
	c.SendMessage("Online Users: " + strings.Join(clients, ", "))
}

func (c *Client) changeNickname(newNick string) {
	if err := utils.ValidateNickname(newNick); err != nil {
        c.SendMessage("Erro: " + err.Error())
        return
    }

    c.server.mut.Lock()
    defer c.server.mut.Unlock()

    if _, exists := c.server.clients[newNick]; exists {
        c.SendMessage("Error: nickname already in use")
        return
    }

    delete(c.server.clients, c.Nickname)
    c.Nickname = newNick
    c.server.clients[newNick] = c

    c.SendMessage("Changed nickname to: " + newNick)
    c.server.messages <- newNick + " Now is known as " + newNick
}