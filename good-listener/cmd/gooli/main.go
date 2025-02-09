package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ivanrusanov/tools/good-listener/internal"
)

type endpoints []string

func (e *endpoints) String() string {
	return fmt.Sprintf("%s", *e)
}

func (e *endpoints) Set(value string) error {
	*e = append(*e, value)
	return nil
}

type args struct {
	port      int
	endpoints endpoints
}

func readInput() args {
	port := flag.Int("port", 8080, "Port to listen to")

	var endpoints endpoints
	flag.Var(&endpoints, "endpoint", "Endpoints to serve")

	flag.Parse()

	if err := internal.CheckPortCorrect(*port); err != nil {
		log.Fatal(err)
	}

	if err := internal.CheckEndpointsCorrect(endpoints); err != nil {
		log.Fatal(err)
	}

	return args{
		port:      *port,
		endpoints: endpoints,
	}
}

func main() {
	input := readInput()

	for _, ep := range input.endpoints {
		http.HandleFunc(ep, internal.LoggingHandler)
	}

	log.Printf("Starting server on port %d...", input.port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", input.port), nil)
	if err != nil {
		log.Printf("Something went wrong: %s", err)
	}
}
