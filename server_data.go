package telnet

import (
	"net"
	"sync"
)

type ClientData struct {
	username     string
	messageCount int
	Conn         net.Conn
}

type ServerData struct {
	Clients map[string]ClientData
	mutex   sync.RWMutex
}

func InitServerData() ServerData {
	return ServerData{Clients: make(map[string]ClientData)}
}

func (s *ServerData) Count(name string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	clientData := s.Clients[name]
	clientData.messageCount++
	s.Clients[name] = clientData
	return clientData.messageCount
}

func (s *ServerData) initClient(username string, conn net.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	key := conn.RemoteAddr().String()
	if _, ok := s.Clients[key]; !ok {
		s.Clients[key] = ClientData{
			username:     "",
			messageCount: 1,
			Conn:         conn,
		}
	}
}

func (s *ServerData) removeClient(conn net.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	key := conn.RemoteAddr().String()
	if _, ok := s.Clients[key]; ok {
		delete(s.Clients, key)
	}
}
