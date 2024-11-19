package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/soupaulodev/chat-server/utils"
)

type Client struct {
	connection net.Conn
	Nickname string
	server *Server
	ch chan string
}

func NewClient(connection net.Conn, server *Server) *Client {
	nickname := fmt.Sprintf("User%d", connection.RemoteAddr().(*net.TCPAddr).Port)
	return &Client{
		connection: connection,
		Nickname: nickname,
		server: server,
		ch: make(chan string),
	}
}

func (c *Client) Start() {
	go c.readInput()
	go c.listenToChannel()
	c.SendMessage("Welcome to the chat! " + c.Nickname)
}

func (c *Client) readInput() {
	scanner := bufio.NewScanner(c.connection)
    for scanner.Scan() {
        text := scanner.Text()

        if strings.HasPrefix(text, "/") {
            c.handleCommand(text)
            continue
        }

        if err := utils.ValidateMessage(text); err != nil {
            c.SendMessage("Error: " + err.Error())
            continue
        }

        c.server.messages <- fmt.Sprintf("[%s]: %s", c.Nickname, text)
    }

	c.server.removeClient(c)
    c.connection.Close()
}

func (c *Client) listenToChannel() {
	for msg := range c.ch {
		fmt.Fprintln(c.connection, msg)
	}
}

func (c *Client) SendMessage(msg string) {
	c.ch <- msg
}