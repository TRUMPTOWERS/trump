package doasitellthem

import (
	"encoding/json"
	"net/http"
)

type addrDB interface {
	Get(string) string
	GetAll() []string
	Set(string, string)
}

// Bling are the entries
type Bling struct {
	Name string `json:"name"`
}

// GetAll is a handler that returns a list of all strings in the db
type GetAll struct {
	db addrDB
}

// NewGetAll constructs a GetAll from a provided addrDB
func NewGetAll(db addrDB) *GetAll {
	return &GetAll{db: db}
}

// NewHandler creates a new Mux that serves both the frontend and
// corresponding API
func NewHandler(db addrDB) http.Handler {
	mux := http.NewServeMux()
	getAll := NewGetAll(db)
	// Handle Route registration here
	mux.Handle("/getAll", getAll)
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	return mux
}

// ServeHTTP is the handler for a GetAll
func (get *GetAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "json")
	encoder := json.NewEncoder(rw)
	names := get.db.GetAll()
	var blings []Bling
	for _, s := range names {
		blings = append(blings, Bling{
			Name: s,
		})
	}
	encoder.Encode(blings)
}
