package hands_test

import (
	"testing"
	"time"

	"github.com/trumptowers/trump/tower/hands"
)

func TestGetHostPort(t *testing.T) {
	t.Parallel()
	mine := hands.New()
	thisIP := "127.5.6.250:2016"

	mine.Set("trump", thisIP)

	gotIP := mine.Get("trump")

	if gotIP != thisIP {
		t.Fatalf("Mismatched HostPort, expected %q, got %q\n", thisIP, gotIP)
	}
}

func TestGetTimestamp(t *testing.T) {
	t.Parallel()
	mine := hands.New()

	mine.Set("trump", "127.5.6.250:2016")

	timestamp := mine.GetTimestamp("trump")
	if time.Since(timestamp).Seconds() > 1 {
		t.Fatalf("Incorrect timestamp")
	}
}

func TestGetExpired(t *testing.T) {
	t.Parallel()
}

func TestGetAll(t *testing.T) {
	t.Parallel()
	mine := hands.New()

	mine.Set("trump", "127.5.6.250:2016")
	mine.Set("drumpf", "127.0.0.1:2020")

	all := mine.GetAll()

	if !(all[0] == "trump" && all[1] == "drumpf") {
		t.Fatalf("failed to retreive all entries")
	}
}

func TestNotExist(t *testing.T) {
	t.Parallel()
	mine := hands.New()

	gotHostPort := mine.Get("trump")
	if gotHostPort != "" {
		t.Fatalf("Got non-zero result for non-existant value: %q\n", gotHostPort)
	}
}

func TestOverwrite(t *testing.T) {
	t.Parallel()

	mine := hands.New()
	thisIP := "127.5.6.250:2016"
	newIP := "127.0.0.1:2020"
	mine.Set("trump", thisIP)
	mine.Set("trump", newIP)

	gotHostPort := mine.Get("trump")

	if gotHostPort != newIP {

		t.Fatalf("Got wrong values for overwritten entry: %q\n", gotHostPort)
	}
}
