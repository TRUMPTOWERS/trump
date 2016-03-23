package main

import (
	"log"
	"net/http"

	"github.com/trumptowers/trump/tower/deflect"
	"github.com/trumptowers/trump/tower/hands"
	"github.com/trumptowers/trump/tower/theleastracist"
)

func main() {
	db := hands.New()
	reg := theleastracist.NewRegistrar(db)
	deflector := deflect.New(db)

	regMux := http.NewServeMux()

	regMux.Handle("/register", reg)

	go func() {
		log.Fatal(http.ListenAndServe(":80", deflector))
	}()
	log.Fatal(http.ListenAndServe(":2016", regMux))
}
