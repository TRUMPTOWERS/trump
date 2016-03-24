package doasitellthem_test

import (
        "encoding/json"
        "net/http"
	"net/http/httptest"
	"testing"
        "github.com/TRUMPTOWERS/trump/tower/doasitellthem"
)

type mockDB struct {
    names []string
}

func (db *mockDB) GetAll() []string {
    return []string {"trump", "drumpf"}
}

func (db *mockDB) Get(name string) string {
    panic("not implemented")
}

func (db *mockDB) Set(name string, hostPort string) {
    panic("not implemented")
}

func TestGetAll(t *testing.T) {
    t.Parallel()

    db := &mockDB{}
    get := doasitellthem.NewGetAll(db)

    req, err := http.NewRequest("GET", "", nil)
    if err != nil {
        panic("error making request")
    }

    w := httptest.NewRecorder()
    get.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Error(w.Body.String())
        t.Fatalf("Got bad status code in test request: %d\n", w.Code)
    }

    var blings []doasitellthem.Bling
    json.Unmarshal(w.Body.Bytes(), &blings)

    if !(blings[0].Name == "trump" && blings[1].Name == "drumpf") {
       t.Fatalf("Got incorrect data") 
    }
}
