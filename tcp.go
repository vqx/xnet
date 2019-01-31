package xnet

import (
	"net"
	"sync"
	"time"
)

type TcpServer struct {
	Addr       string
	conn       *net.Conn
	listener   net.Listener
	ClientMap  map[string]*TcpClient
	lock       sync.Mutex
	DataHandle func(*TcpServer, []byte)
}

type TcpClient struct {
	conn   net.Conn
	lock   sync.Mutex
	server *TcpServer
}

func NewTcpServer(addr string, DataHandle func(*TcpServer, []byte)) *TcpServer {
	result := &TcpServer{
		Addr:       addr,
		ClientMap:  map[string]*TcpClient{},
		DataHandle: DataHandle,
	}
	return result
}

func (s *TcpServer) Run() {
	var err error
	s.listener, err = net.Listen("tcp4", s.Addr)
	if err != nil {
		panic(`tzh6dhq9kf net.Listen ` + err.Error())
	}
	go s.ListenThread()
	for {
		time.Sleep(time.Second)
	}
}
func (s *TcpServer) ListenThread() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		tmpClient := &TcpClient{
			conn: conn,
		}
		s.lock.Lock()
		s.ClientMap[GetClientId(conn.RemoteAddr().String())] = tmpClient
		s.lock.Unlock()
		go tmpClient.GetDataThread()
	}
}

func (c *TcpClient) GetDataThread() {
	for {
		data, err := c.Get(2048)
		if err != nil {
			c.lock.Lock()
			c.conn.Close()
			c.lock.Unlock()
			break
		}
		c.server.DataHandle(c.server, data)
	}
}

func (c *TcpClient) Send(data []byte) (n int, err error) {
	c.lock.Lock()
	n, err = c.conn.Write(data)
	c.lock.Unlock()
	return
}

func (c *TcpClient) Get(size int) (data []byte, err error) {
	data = make([]byte, size)
	var n int
	c.lock.Lock()
	n, err = c.conn.Read(data)
	c.lock.Unlock()
	return data[:n], err
}
