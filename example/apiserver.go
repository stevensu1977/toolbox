package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stevensu1977/toolbox/net"
)

var portPtr = flag.String("port", "9999", "api server listen port")
var hostPtr = flag.String("host", "", "api server listen host address")

func HandleAPI(w http.ResponseWriter, r *http.Request) {

	api := map[string]string{
		"API": "v1.0",
	}
	net.ServerJSON(w, api)
}

func main() {
	flag.Parse()

	address := *hostPtr + ":" + *portPtr

	http.HandleFunc("/api", HandleAPI)

	log.Printf("API Server start on %s ", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
