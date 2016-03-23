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

func TestServeHTTP(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(WriteData))
	defer srv.Close()
	url, err := url.Parse(srv.URL)
	if err != nil {
		panic("error making url")
	}
	db := &mockDB{
		hostPort: url.Host,
	}
	deflector := New(db)

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic("error making request")
	}

	w := httptest.NewRecorder()
	deflector.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Got bad status code in test request: %d\n", w.Code)
	}

	if w.Body.String() != "it works" {
		t.Fatalf("Got bad body in test request: %q\n", w.Body.String())
	}
}
