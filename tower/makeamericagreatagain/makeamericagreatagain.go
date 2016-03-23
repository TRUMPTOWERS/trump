package theleastracist

import (
	"net"
	"net/http"
	"strconv"
)

type addrDB interface {
	Get(string) string
	Set(string, string)
}

type Registrar struct {
	db addrDB
}

func NewRegistrar(db addrDB) *Registrar {
	return &Registrar{db: db}
}

func (reg *Registrar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	port := r.FormValue("port")
	_, err := strconv.Atoi(port)
	if err != nil {
		http.Error(rw, "No/malformed port provided", 500)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(rw, "No/malformed host provided: "+r.RemoteAddr, 500)
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(rw, "No name provided", 500)
	}

	reg.db.Set(name, host+":"+port)
	rw.WriteHeader(http.StatusOK)
}
