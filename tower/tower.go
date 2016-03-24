package main

import (
	"log"
	"net/http"

	"github.com/TRUMPTOWERS/trump/tower/deflect"
	"github.com/TRUMPTOWERS/trump/tower/hands"
	"github.com/TRUMPTOWERS/trump/tower/theleastracist"
	"github.com/conslo/pac"
)

func main() {
	db := hands.New()
	reg := theleastracist.NewRegistrar(db)
	deflector := deflect.New(db)

	regMux := http.NewServeMux()

	regMux.Handle("/register", reg)
	regMux.Handle("/pac", pac.New(pac.Me, "", 8081, "drumpf"))

	go func() {
		log.Fatal(http.ListenAndServe(":8081", deflector))
	}()
	log.Fatal(http.ListenAndServe(":2016", regMux))
}
