package hands

import (
	"sort"
	"testing"
	"time"
)

func TestGetHostPort(t *testing.T) {
	t.Parallel()
	mine := New()
	thisIP := "127.5.6.250:2016"

	mine.Set("trump", thisIP)

	gotIP := mine.Get("trump")

	if gotIP != thisIP {
		t.Fatalf("Mismatched HostPort, expected %q, got %q\n", thisIP, gotIP)
	}
}

func TestGetTimestamp(t *testing.T) {
	t.Parallel()
	mine := New()

	mine.Set("trump", "127.5.6.250:2016")

	timestamp := mine.GetTimestamp("trump")
	if time.Since(timestamp).Seconds() > 1 {
		t.Fatalf("Incorrect timestamp")
	}
}

func TestGetExpired(t *testing.T) {
	t.Parallel()

	mine := &DB{
		db:       make(map[string]dbEntry),
		duration: time.Duration(50 * time.Millisecond),
	}

	mine.Set("trump", "127.5.6.250:2016")
	time.Sleep(100 * time.Millisecond)

	gotHost := mine.Get("trump")

	if gotHost != "" {
		t.Fatalf("data persists after expiration")
	}
}

func TestGetAll(t *testing.T) {
	t.Parallel()
	mine := New()

	mine.Set("trump", "127.5.6.250:2016")
	mine.Set("drumpf", "127.0.0.1:2020")

	all := mine.GetAll()

	sort.Strings(all)
	if all[0] != "drumpf" || all[1] != "trump" {
		t.Fatalf("failed to retreive all entries: %v", all)
	}
}

func TestGetAllExpired(t *testing.T) {
	t.Parallel()

	mine := &DB{
		db:       make(map[string]dbEntry),
		duration: time.Duration(50 * time.Millisecond),
	}

	mine.Set("trump", "127.5.6.250:2016")
	time.Sleep(100 * time.Millisecond)
	mine.Set("drumpf", "127.0.0.1:2020")

	gotHosts := mine.GetAll()

	if len(gotHosts) != 1 || gotHosts[0] != "drumpf" {
		t.Fatalf("data persists after expiration")
	}
}

func TestGetAllEmpty(t *testing.T) {
	t.Parallel()

	mine := &DB{
		db: make(map[string]dbEntry),
	}

	getHosts := mine.GetAll()
	if getHosts == nil {
		t.Fatal("GetAll should return an empty slice, not nil")
	}
}

func TestNotExist(t *testing.T) {
	t.Parallel()
	mine := New()

	gotHostPort := mine.Get("trump")
	if gotHostPort != "" {
		t.Fatalf("Got non-zero result for non-existant value: %q\n", gotHostPort)
	}
}

func TestOverwrite(t *testing.T) {
	t.Parallel()

	mine := New()
	thisIP := "127.5.6.250:2016"
	newIP := "127.0.0.1:2020"
	mine.Set("trump", thisIP)
	mine.Set("trump", newIP)

	gotHostPort := mine.Get("trump")

	if gotHostPort != newIP {

		t.Fatalf("Got wrong values for overwritten entry: %q\n", gotHostPort)
	}
}
