package xnet

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UdpServer struct {
	udpAddr    *net.UDPAddr
	addr       string
	conn       *net.UDPConn
	connLock   sync.Locker
	ClientMap  map[string]*UdpClient
	mapLock    sync.Locker
	DataHandle func(req UdpDataHandleRequest)
}

type UdpClient struct {
	Id      string
	udpAddr *net.UDPAddr
	server  *UdpServer
}

type UdpDataHandleRequest struct {
	Server *UdpServer
	Client *UdpClient
}

func (s *UdpServer) Run() {
	tmp := strings.Split(s.addr, ":")
	if len(tmp) != 2 {
		panic(`as4awdps74 Run len(tmp) != 2`)
	}
	port, err := strconv.Atoi(tmp[1])
	if err != nil {
		panic(`7nx7b93jny`)
	}
	s.udpAddr = &net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(tmp[0]),
	}
	s.conn, err = net.ListenUDP("udp4", s.udpAddr)
	if err != nil {
		panic(`f4gsfejw3r ` + err.Error())
	}
	go s.GetThread()
}

func (s *UdpServer) GetThread() {
	for {
		data := make([]byte, 2048)
		s.connLock.Lock()
		n, remoteAddr, err := s.conn.ReadFromUDP(data)
		if err != nil || n == 0 {
			s.connLock.Unlock()
			fmt.Println("p4zu7zgexf", "ReadFromUDP", err)
			time.Sleep(time.Microsecond * 50)
			continue
		}
		s.connLock.Unlock()
		go func(udpAddr *net.UDPAddr) {
			id := GetClientId(udpAddr.String())
			s.mapLock.Lock()
			udpClient, ok := s.ClientMap[id]
			if !ok {
				fmt.Println(`xv3wqwdgw3`, `new client`, len(s.ClientMap)+1)
				udpClient = &UdpClient{
					Id:      id,
					udpAddr: udpAddr,
					server:  s,
				}
				s.ClientMap[id] = udpClient
			}
			s.mapLock.Unlock()
			//go or no go
			go s.DataHandle(UdpDataHandleRequest{
				Server: s,
				Client: udpClient,
			})
		}(remoteAddr)
		time.Sleep(time.Microsecond * 50)
	}
}
