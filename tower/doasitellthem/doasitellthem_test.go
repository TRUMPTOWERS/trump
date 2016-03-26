package doasitellthem_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TRUMPTOWERS/trump/tower/doasitellthem"
)

type mockDB struct {
	names []string
}

func (db *mockDB) GetAll() []string {
	return db.names
}

func (db *mockDB) Get(name string) string {
	panic("not implemented")
}

func (db *mockDB) Set(name string, hostPort string) {
	panic("not implemented")
}

func TestNewHandler(t *testing.T) {
	db := mockDB{}
	handler := doasitellthem.NewHandler(&db)
	req, _ := http.NewRequest("GET", "/getAll", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("error: did not register getAll route correctly")
	}
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db := &mockDB{
		names: []string{"trump", "drumpf"},
	}
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

func TestGetNone(t *testing.T) {
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

	str := w.Body.String()
	if !strings.Contains(str, "[") || !strings.Contains(str, "]") {
		t.Fatalf("Result wasn't a json array, got %q", str)
	}
}
