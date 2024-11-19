package server

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	addr string
	clients map[string]*Client
	mut sync.Mutex
	messages chan string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		clients: make(map[string]*Client),
		messages: make(chan string),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go s.broadcastMessages()

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		go s.handleNewClient(connection)
	}
}

func (s *Server) handleNewClient(connection net.Conn) {
	client := NewClient(connection, s)
	s.addClient(client)
	client.Start()
}

func (s *Server) addClient(client *Client) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.clients[client.Nickname] = client
	s.messages <- fmt.Sprintf("%s joined the chat", client.Nickname)
}

func (s *Server) removeClient(client *Client) {
	s.mut.Lock()
	defer s.mut.Unlock()
    delete(s.clients, client.Nickname)
    s.messages <- fmt.Sprintf("%s saiu do chat", client.Nickname)
}

func (s *Server) broadcastMessages() {
	for msg:= range s.messages {
		s.mut.Lock()
		for _, client := range s.clients {
			client.SendMessage(msg)
		}
		s.mut.Unlock()
	}
}