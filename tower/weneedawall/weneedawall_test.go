package weneedawall_test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/TRUMPTOWERS/trump/tower/weneedawall"
)

func TestRespondsToDiscoverMessage(t *testing.T) {
	discovery, err := weneedawall.NewDiscovery(2016)
	if err != nil {
		log.Fatalf("could not create discovery server: %q\n", err)
	}
	go discovery.ListenAndServe()
	defer discovery.Close()

	// Listen to udp on ephmeral port
	conn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		t.Fatalf("error: could not establish udp listener: %q\n", err)
	}
	defer conn.Close()

	// Handle the response back from our trump tower
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			t.Fatalf("error reading response from conn: %q\n", err)
		}
		if msg := string(buf[0:n]); msg != "TRUMP" {
			t.Fatalf("error: unexpected response from broadcast: %q", msg)
		}
	}()

	// Send multicast message
	_, err = conn.WriteToUDP([]byte("TRUMPTOWER"), &net.UDPAddr{
		IP:   net.IPv4allsys,
		Port: 2016,
	})

	if err != nil {
		t.Fatalf("error writing broadcast message: %q\n", err)
	}
	<-done
}

func TestNotRespondingToUnknownMessage(t *testing.T) {
	discovery, err := weneedawall.NewDiscovery(2016)
	if err != nil {
		log.Fatalf("could not create discovery server: %q\n", err)
	}
	go discovery.ListenAndServe()
	defer discovery.Close()

	// Listen to udp on ephmeral port
	conn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		t.Fatalf("error: could not establish udp listener: %q\n", err)
	}
	defer conn.Close()

	// Handle the response back from our trump tower
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		buf := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		_, _, err := conn.ReadFromUDP(buf)
		if err == nil {
			t.Fatalf("error: should have received timeout due to no response\n")
		}
	}()

	// Send multicast message
	_, err = conn.WriteToUDP([]byte("FOO"), &net.UDPAddr{
		IP:   net.IPv4allsys,
		Port: 2016,
	})

	if err != nil {
		t.Fatalf("error writing broadcast message: %q\n", err)
	}
	<-done
}
