package deflect

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type mockDB struct {
	hostPort string
}

func (db *mockDB) Get(name string) string {
	return db.hostPort
}

func (db *mockDB) Set(name string, hostPort string) {
	panic("not implemented")
}

func WriteData(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("it works"))
}

func TestReverseProxyServeHTTP(t *testing.T) {
	t.Parallel()
	// Create test HTTP Server to proxy our requests from
	srv := httptest.NewServer(http.HandlerFunc(WriteData))
	defer srv.Close()
	url, err := url.Parse(srv.URL)
	if err != nil {
		panic("error making url")
	}

	// Create new mockDB that will always return our mock server
	db := &mockDB{
		hostPort: url.Host,
	}
	deflector := New(db)

	// Create mock request
	req, err := http.NewRequest("GET", "http://donald.drumpf:8080", nil)
	if err != nil {
		panic("error making request")
	}

	w := httptest.NewRecorder()
	deflector.ServeHTTP(w, req)

	// Ensure our reverse proxy is working
	if w.Code != http.StatusOK {
		t.Error(w.Body.String())
		t.Fatalf("Got bad status code in test request: %d\n", w.Code)
	}

	if w.Body.String() != "it works" {
		t.Fatalf("Got bad body in test request: %q\n", w.Body.String())
	}
}

func TestBadRequestOnNoHost(t *testing.T) {
	t.Parallel()

	db := &mockDB{}
	deflector := New(db)

	// Request that does not contain a host or port
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic("error making request")
	}

	w := httptest.NewRecorder()
	deflector.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Did not get StatusBadRequest: Got %d\n", w.Code)
	}
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	// Create Empty DB
	db := &mockDB{}
	deflector := New(db)

	// Request for a domain name not in the db
	req, err := http.NewRequest("GET", "http://donald.drumpf", nil)
	if err != nil {
		panic("error making request")
	}
	req.Host = "donald.drumpf"

	w := httptest.NewRecorder()

	// Receive 404 for not finding anything in DB
	deflector.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("Did not get StatusNotFound: Got %d\n", w.Code)
	}
}

func TestValidHost(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(WriteData))
	defer srv.Close()

	db := &mockDB{hostPort: srv.URL[7:]}
	deflector := New(db)

	req, err := http.NewRequest("GET", "donald.drumpf", nil)
	if err != nil {
		panic("error making request")
	}
	req.Host = "donald.drumpf"

	w := httptest.NewRecorder()
	deflector.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error(w.Body.String())
		t.Fatalf("Did not get StatusOK: Got %d\n", w.Code)
	}
}
