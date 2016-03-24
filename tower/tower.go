package main

import (
	"log"
	"net/http"

	"github.com/TRUMPTOWERS/trump/tower/deflect"
	"github.com/TRUMPTOWERS/trump/tower/doasitellthem"
	"github.com/TRUMPTOWERS/trump/tower/hands"
	"github.com/TRUMPTOWERS/trump/tower/theleastracist"
)

func main() {
	db := hands.New()
	reg := theleastracist.NewRegistrar(db)
	deflector := deflect.New(db)
	regMux := http.NewServeMux()
	regMux.Handle("/register", reg)

        api := doasitellthem.NewServeMux(db)

	go func() {
		log.Fatal(http.ListenAndServe(":8081", deflector))
	}()
        go func() {
                log.Fatal(http.ListenAndServe(":8082", api))
        }()

	log.Fatal(http.ListenAndServe(":2016", regMux))
}
