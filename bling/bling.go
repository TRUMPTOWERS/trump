package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	name := flag.String("name", "", "host name")
	port := flag.Int("port", 0, "port")

	flag.Parse()

	if *name == "" {
		log.Fatalln("no name given")
	}

	if *port == 0 {
		log.Fatalln("no port given")
	}

	nameStr := *name
	portStr := strconv.Itoa(*port)

	if err != nil {
		log.Fatalln("malformed port number")
	}

	res, err := http.PostForm("http://donald.drumpf:2016/register",
		url.Values{"name": {nameStr}, "port": {portStr}})

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("error: response returned with code %d\n", res.StatusCode)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
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
