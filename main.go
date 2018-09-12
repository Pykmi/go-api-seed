package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"bitbucket.org/pykmiteam/mock-api/datastore"
	"bitbucket.org/pykmiteam/mock-api/logger"
)

func main() {
	// set default commandline flags and parse them
	httphost := flag.String("host", "localhost", "HTTP hostname")
	httpport := flag.String("port", "3088", "HTTP port number")
	db := flag.String("db", "127.0.0.1:9000", "Database server host")

	flag.Parse()

	server := net.JoinHostPort(*httphost, *httpport)

	// split the database host from port number
	addr, p, _ := net.SplitHostPort(*db)
	port, err := strconv.Atoi(p)
	if err != nil {
		panic(err)
	}

	// create datastore
	StoreOpt := datastore.StoreOptions{Host: addr, Namespace: "test", Port: port}
	Store := datastore.New(StoreOpt)

	// create event logger
	EventLogger := logger.New()

	// start the server
	if err := startServer(server, EventLogger, Store); err != nil {
		log.Printf("%#v", err)
		return
	}
}

/**
 * Starts the HTTP server.
 */
func startServer(server string, EventLogger *logger.Logger, Store *datastore.Store) error {
	log.Println("Server started on at: ", server)

	// create http routes
	APIrouter := setupRoutes(EventLogger, Store)

	// start listening for the client connections
	err := http.ListenAndServe(server, APIrouter)
	if err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}
