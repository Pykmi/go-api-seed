package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"bitbucket.org/pykmiteam/mock-api/logger"
)

func main() {
	// set default commandline flags and parse them
	httphost := flag.String("host", "localhost", "HTTP hostname")
	httpport := flag.String("port", "3088", "HTTP port number")

	flag.Parse()

	EventLogger := logger.New()

	server := net.JoinHostPort(*httphost, *httpport)

	// start the server
	if err := startServer(server, EventLogger); err != nil {
		log.Printf("%#v", err)
		return
	}
}

/**
 * Starts the HTTP server.
 */
func startServer(server string, EventLogger *logger.Logger) error {
	log.Println("Server started on at: ", server)

	// create http routes
	APIrouter := setupRoutes(EventLogger)

	// start listening for the client connections
	err := http.ListenAndServe(server, APIrouter)
	if err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}
