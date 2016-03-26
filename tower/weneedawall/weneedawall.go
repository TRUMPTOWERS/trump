package weneedawall

import (
	"log"
	"net"
)

type conn interface {
	ReadFromUDP([]byte) (int, *net.UDPAddr, error)
	WriteToUDP([]byte, *net.UDPAddr) (int, error)
	Close() error
}

// Discovery is the multicast server that listens for
// TRUMPTOWER discovery requests
type Discovery struct {
	Conn conn
}

func NewDiscovery(port int) (*Discovery, error) {
	conn, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4allsys,
		Port: 2016,
	})
	if err != nil {
		return nil, err
	}
	return &Discovery{conn}, nil
}

// ListenAndServe launches a Discovery Server and listens to discover requests
func ListenAndServe(port int) error {
	d, err := NewDiscovery(port)
	if err != nil {
		return err
	}
	return d.ListenAndServe()
}

// Listen listens for incoming discover requestions
func (d *Discovery) ListenAndServe() error {
	buf := make([]byte, 1024)
	// Read the broadcast messages
	for {
		n, remoteAddr, err := d.Conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}
		msg := string(buf[0:n])
		if msg != "TRUMPTOWER" {
			log.Printf("received unknown message: %q\n", msg)
			continue
		}
		go d.respond(remoteAddr)
	}
}

func (d *Discovery) respond(remoteAddr *net.UDPAddr) {
	log.Printf("responding to tower discovery: %v\n", remoteAddr)
	_, err := d.Conn.WriteToUDP([]byte("TRUMP"), remoteAddr)
	if err != nil {
		log.Printf("error: could not write message to client: %q\n", err)
	}
}

// Close terminates the discovery server connection
func (d *Discovery) Close() {
	d.Conn.Close()
}
