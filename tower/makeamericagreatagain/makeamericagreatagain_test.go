package theleastracist

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct {
	hostPort string
	name     string
}

func (db *mockDB) Get(name string) string {
	panic("not implemented")
}

func (db *mockDB) Set(name string, hostPort string) {
	db.hostPort = hostPort
	db.name = name
}

func WriteData(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("it works"))
}

func TestServeHTTP(t *testing.T) {

	db := &mockDB{}
	reg := NewRegistrar(db)

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic("error making request")
	}
	req.RemoteAddr = "192.168.1.1:12351"

	q := req.URL.Query()
	q.Add("name", "trump")
	q.Add("port", "2016")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	reg.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error(w.Body.String())
		t.Fatalf("Got bad status code in test request: %d\n", w.Code)
	}

	if db.hostPort != "192.168.1.1:2016" {
		t.Fatalf("Got incorrect hostPort, expected \"192.168.1.1:2016\", got %q\n", db.hostPort)
	}

	if db.name != "trump" {
		t.Fatalf("Got incorrect name, expected \"trump\", got %q\n", db.name)
	}
}
