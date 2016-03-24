package theleastracist

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

type addrDB interface {
	Get(string) string
	Set(string, string)
}

// Registrar is a handler that add entries to a addrDB
type Registrar struct {
	db addrDB
}

// NewRegistrar constructs a Registrar from a provided addrDB
func NewRegistrar(db addrDB) *Registrar {
	return &Registrar{db: db}
}

// ServeHTTP is the handler for a Registrar
func (reg *Registrar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	port := r.FormValue("port")
	_, err := strconv.Atoi(port)
	if err != nil {
		http.Error(rw, "No/malformed port provided", http.StatusBadRequest)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// In theory this should only be encountered with a badly formed request
		// ie: incorrectly written test
		http.Error(rw, "No/malformed host provided: "+r.RemoteAddr, http.StatusBadRequest)
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(rw, "No name provided", http.StatusBadRequest)
	}

	toSave := host + ":" + port
	reg.db.Set(name, host+":"+port)
	log.Printf("Saved %q for use at host %q\n", toSave, name)
	rw.WriteHeader(http.StatusOK)
}
