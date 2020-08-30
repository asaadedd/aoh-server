package service

import (
	"fmt"
	"github.com/asaadedd/aoh-server/gudp/shared"
	"net"
)

type ConnectedClient struct {
	status   status.Status
	address  net.Addr
	messages chan []byte
	conn     *Server
}

func (client *ConnectedClient) init() {
	go func() {
		for message := range client.messages {
			fmt.Printf("message = %s \n", message)
		}
	}()
}

func (client *ConnectedClient) sendMessage(message []byte) {
	client.messages <- message
}

func NewClient(address net.Addr, server *Server) (c *ConnectedClient) {
	client := ConnectedClient{status.CONNECTING, address, make(chan []byte), server}

	client.init()

	return &client
}
