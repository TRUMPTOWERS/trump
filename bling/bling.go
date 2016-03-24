package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	name := flag.String("name", "", "host name")
	port := flag.Int("port", 0, "port")

	flag.Parse()

	useArgs := false
	var argName string
	var argPort int

	if *name == "" && *port == 0 {

		if len(os.Args) != 3 {
			log.Fatalln("incorrect number of arguments (expected 2)")
		}

		useArgs = true
		argName = os.Args[1]
		var err error
		argPort, err = strconv.Atoi(os.Args[2])

		if err != nil {
			log.Fatalln("invalid port number")
		}
	} else if *name == "" {
		log.Fatalln("no name given")
	} else if *port == 0 {
		log.Fatalln("no port given")
	}

	var nameStr string
	var portStr string

	if useArgs {
		nameStr = argName
		portStr = strconv.Itoa(argPort)
	} else {
		nameStr = *name
		portStr = strconv.Itoa(*port)
	}

	res, err := http.PostForm("http://donald.drumpf:2016/register",
		url.Values{"name": {nameStr}, "port": {portStr}})

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("error: response returned with code %d\n", res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("error reading body")
			return
		}
		strBody := string(body)
		if strBody != "" {
			log.Printf("response body: %q\n", strBody)
		}
	}
	log.Println("OK")
}
