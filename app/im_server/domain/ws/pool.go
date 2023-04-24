package ws

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

// ConnPool  connection pool
var Pool *ConnPool

func init() {
	Pool = NewConnPool()
}

type ConnPool struct {
	clients *sync.Map
}

func NewConnPool() *ConnPool {
	return &ConnPool{
		clients: &sync.Map{},
	}
}

func (p *ConnPool) Add(c *ClientConn) {
	p.clients.Store(c.Id, c)
}

func (p *ConnPool) Remove(c *ClientConn) {
	p.clients.Delete(c.Id)
	err := c.Conn.Close()
	if err != nil {
		log.Errorf("close error: %v", err)
	}
}
