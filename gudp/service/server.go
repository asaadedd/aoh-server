package service

import (
	"log"
	"net"
)
const ProtocolId = 1

type ServerConfig struct {
	maxClientsNumber int
}

type Server struct {
	config *ServerConfig
	clients map[net.Addr]*ConnectedClient
	numberOfConnectedClients int
}

func (server *Server) isServerFull() bool {
	return server.numberOfConnectedClients >= server.config.maxClientsNumber
}

func (server *Server) start(port int) {
	listeningAddress, err := net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return
	}
	conn, err := net.ListenUDP("udp",listeningAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}
		if client, ok := server.clients[addr]; ok && server.isServerFull() {
			client.sendMessage(buf[:n])
		} else {
			server.clients[addr] = NewClient(addr, server)
		}
	}
}

func NewServer(config *ServerConfig) *Server {
	return &Server{config, make(map[net.Addr]*ConnectedClient), 0}
}