package deflect

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Deflector is a handler that proxies based on a provided addrDB
type Deflector struct {
	db addrDB
}

type addrDB interface {
	Get(string) string
	Set(string, string)
}

// ServerHTTP creates a reverse proxy based on the stored hands, and passes
// the request to the proxy
func (d *Deflector) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hostToGet, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		http.Error(rw, "No/malformed name: "+r.Host, http.StatusBadRequest)
		return
	}
	hostPort := d.db.Get(hostToGet)
	log.Printf("hostPort %q found for request %q\n", hostPort, hostToGet)

	if hostPort == "" {
		http.NotFound(rw, r)
		return
	}

	proxyURL := &url.URL{Host: hostPort, Scheme: "http"}
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	proxy.ServeHTTP(rw, r)
}

// New creates a Deflector with a new hands DB
func New(db addrDB) http.Handler {
	return &Deflector{db: db}
}
