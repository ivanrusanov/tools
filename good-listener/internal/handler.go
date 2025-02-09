package internal

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func LoggingHandler(w http.ResponseWriter, r *http.Request) {
	d, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error while dumping request: ", err)
	}

	log.Println("Incomming request:", string(d))
}
