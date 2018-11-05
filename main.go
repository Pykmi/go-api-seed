package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pykmi/go-api-seed/datastore"
	"github.com/pykmi/go-api-seed/logger"
)

func main() {
	// set default commandline flags and parse them
	httphost := flag.String("host", "localhost", "HTTP hostname")
	httpport := flag.String("port", "3088", "HTTP port number")

	dbhost := flag.String("db-host", "0.0.0.0", "Database server host")
	dbport := flag.String("db-port", "8081", "Database server port")

	dbuser := flag.String("db-user", "pykmi", "Database username")
	dbpass := flag.String("db-pass", "okilzw", "Database password")

	dbname := flag.String("db", "pykmi-dev-db", "Mongo database name")

	flag.Parse()

	server := net.JoinHostPort(*httphost, *httpport)

	// create datastore
	StoreOpt := datastore.StoreOptions{
		Host: *dbhost,
		Port: *dbport,
		User: *dbuser,
		Pass: *dbpass,
		Database: *dbname,
	}
	
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
