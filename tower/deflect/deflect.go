package deflect

import (
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
	hostPort := d.db.Get(r.URL.Host)

	if hostPort == "" {
		http.NotFound(rw, r)
		return
	}

	proxyUrl := &url.URL{Host: hostPort, Scheme: "http"}
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	proxy.ServeHTTP(rw, r)
}

// New creates a Deflector with a new hands DB
func New(db addrDB) http.Handler {
	return &Deflector{db: db}
}
