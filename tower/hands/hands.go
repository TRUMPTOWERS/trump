package hands

import "sync"

// DB reprsents an interface to a domain:address database
type DB struct {
	db map[string]string
	sync.Mutex
}

// New creates a DB
func New() *DB {
	thisDB := &DB{}
	thisDB.db = make(map[string]string)
	return thisDB
}

// Set adds/replaces a value in the DB
func (db *DB) Set(domain string, hostPort string) {
	db.Lock()
	db.db[domain] = hostPort
	db.Unlock()
}

// Get retrives a value from the DB, or (nil,0) if it doesn't exist
func (db *DB) Get(domain string) string {
	db.Lock()
	thisAddress := db.db[domain]
	db.Unlock()
	return thisAddress
}
