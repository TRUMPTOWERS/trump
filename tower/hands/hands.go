package hands

import "net"

type DB struct {
	db map[string]address
}

type address struct {
	ip   net.IP
	port int
}

func New() *DB {
	return &DB{}
}

func (db *DB) Set(domain string, ip net.IP, port int) {

}

func (db *DB) Get(domain string) (net.IP, int) {

}
