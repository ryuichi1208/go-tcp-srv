package client

import (
	"log"
	"net"
	"time"

	"github.com/mikioh/tcp"
	"github.com/mikioh/tcpopt"
)

type TcpSocketOption struct {
	addr   string
	port   int
	sndbuf int
	rcvbuf int
}

func Do() {
	c, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	tc, err := tcp.NewConn(c)
	if err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAlive(true)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveIdleInterval(3 * time.Minute)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveProbeInterval(30 * time.Second)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveProbeCount(3)); err != nil {
		log.Fatal(err)
	}
}
