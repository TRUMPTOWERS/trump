package hands_test

import (
	"net"
	"testing"

	"github.com/trumptowers/trump/tower/hands"
)

func TestGetIP(t *testing.T) {
	mine := hands.New()
	thisIP := net.ParseIP("127.5.6.250")

	mine.Set("trump", thisIP, 2016)

	gotIP, _ := mine.Get("trump")

	if !gotIP.Equal(thisIP) {
		t.Fatalf("Mismatched IP, expected %q, got %q\n", thisIP, gotIP)
	}
}

func TestGetPort(t *testing.T) {
	mine := hands.New()
	thisIP := net.ParseIP("127.5.6.250")
	mine.Set("trump", thisIP, 2016)

	_, gotPort := mine.Get("trump")

	if gotPort != 2016 {
		t.Fatalf("Mismatched port, expected %d, got %d\n", 2016, gotPort)
	}
}
