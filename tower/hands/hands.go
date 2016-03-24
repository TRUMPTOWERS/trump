package hands

import (
	"sync"
	"time"
)

type dbEntry struct {
	hostPort  string
	timestamp time.Time
}

// DB represents an interface to a domain:address database
type DB struct {
	db       map[string]dbEntry
	duration time.Duration
	sync.Mutex
}

// New creates a DB
func New() *DB {
	thisDB := &DB{
		db:       make(map[string]dbEntry),
		duration: time.Duration(24 * time.Hour),
	}
	return thisDB
}

// Set adds/replaces a value in the DB
func (db *DB) Set(domain string, hostPort string) {
	db.Lock()
	entry := dbEntry{hostPort: hostPort, timestamp: time.Now()}
	db.db[domain] = entry
	db.Unlock()
}

// Get retrives a value from the DB, or (nil,0) if it doesn't exist
func (db *DB) Get(domain string) string {
	db.Lock()
	thisAddress := db.db[domain].hostPort
	if time.Since(db.db[domain].timestamp) > db.duration {
		delete(db.db, domain)
		thisAddress = ""
	}
	db.Unlock()
	return thisAddress
}

// GetTimestamp retreives the time that an entry was added
func (db *DB) GetTimestamp(domain string) time.Time {
	db.Lock()
	time := db.db[domain].timestamp
	db.Unlock()
	return time
}

// GetAll retreives all entries in the database
func (db *DB) GetAll() []string {
	db.Lock()
	all := make([]string, 0)
	var expired []string

	for k, v := range db.db {
		if time.Since(v.timestamp) > db.duration {
			expired = append(expired, k)
		} else {
			all = append(all, k)
		}
	}

	for _, s := range expired {
		delete(db.db, s)
	}
	db.Unlock()
	return all
}
