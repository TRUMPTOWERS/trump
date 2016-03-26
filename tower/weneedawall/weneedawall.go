package weneedawall

import (
	"log"
	"net"
)

func New() *discovery {
	return &discovery{done: make(chan struct{})}
}

type discovery struct {
	conn *net.UDPConn
	done chan struct{}
}

func (d *discovery) Listen() error {
	conn, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4allsys,
		Port: 2016,
	})
	if err != nil {
		return err
	}
	d.conn = conn
	defer conn.Close()

	buf := make([]byte, 1024)
	// Read the broadcast messages
	for {
		select {
		case <-d.done:
			return nil
		default:
			n, remoteAddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				log.Printf("error: could not read multicast message: %q\n", err)
				continue
			}
			msg := string(buf[0:n])
			if msg != "TRUMPTOWER" {
				log.Printf("received unknown message: %q\n", msg)
				continue
			}
			go d.respond(remoteAddr)
		}
	}
}

func (d *discovery) respond(remoteAddr *net.UDPAddr) {
	log.Printf("responding to tower discovery: %v\n", remoteAddr)
	_, err := d.conn.WriteToUDP([]byte("TRUMP"), remoteAddr)
	if err != nil {
		log.Printf("error: could not write message to client: %q\n", err)
	}
}

func (d *discovery) Close() {
	d.conn.Close()
	d.done <- struct{}{}
}
