package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

// JSON Config Struct
type configuration struct {
	PORT      string
	SERVERS   []string
	ALGORITHM string
}

// Global Variables
var platformSettings configuration
var activeConnections int
var rrTarget = 0
var urlContents = ""

func main() {

	//Prints to console that the program has started.
	fmt.Printf("Initialising RapidBalance\n")
	fmt.Printf("Starting server on port 80\n")

	// #-----------------------------------------------------------------------#
	// #                            JSON Handling                              #
	// #-----------------------------------------------------------------------#

	// Load Configuration File.
	jsonConfiguration, err := os.Open("config.json")
	if err != nil {
		// Something errored; print the error.
		fmt.Println(err)
	}
	// Initalise configuration struct; import configuration file into the struct.
	byteVal, _ := ioutil.ReadAll(jsonConfiguration)
	// Imports the JSON values into the "platformSettings" variable.
	json.Unmarshal([]byte(byteVal), &platformSettings)
	// Grab the values of the JSON configuration and place them into a map.
	for i := 0; i < len(platformSettings.SERVERS); i += 1 {
		serverPool.Add(platformSettings.SERVERS[i], 0)
	}
	fmt.Println("Configured Nodes:")
	fmt.Println(serverPool.GetKeys())
	// Set the "urlContents" variable to the first server in the config.
	urlContents = serverPool.GetUrl()

	// #-----------------------------------------------------------------------#
	// #                             HTTP Server                               #
	// #-----------------------------------------------------------------------#

	// Link to page used to authenticate traffic simulator.
	http.Handle("/loaderio-d5ffdb9edd0a9d365b84452391c9d86e/", http.StripPrefix("/loaderio-d5ffdb9edd0a9d365b84452391c9d86e/", http.FileServer(http.Dir("./static/"))))
	// Routes traffic to the handler functions.
	if platformSettings.ALGORITHM == "RoundRobin" {
		http.HandleFunc("/", roundRobinHandler)
	} else if platformSettings.ALGORITHM == "LeastConnection" {
		http.HandleFunc("/", leastConnectionsHandler)
	}
	// HTTP Server Paramaters.
	s := &http.Server{
		Addr:         ":" + platformSettings.PORT,
		ConnState:    connStateListener,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	// Start the HTTP Server.
	panic(s.ListenAndServe())
}

// #-----------------------------------------------------------------------#
// #                         Routing Handlers                              #
// #-----------------------------------------------------------------------#

func leastConnectionsHandler(res http.ResponseWriter, r *http.Request) {
	// Parse a server from urlContents to the url variable.
	url, _ := url.Parse(urlContents)
	fmt.Println(serverPool.GetValues())
	// Parse the URL through the reverse proxy.
	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(res, r)
}

func roundRobinHandler(res http.ResponseWriter, r *http.Request) {
	// Route to RoundRobin
	url := roundRobin()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	log.Printf("Routing to URL: %s", url.String())
	rProxy.ServeHTTP(res, r)

}

// #-----------------------------------------------------------------------#
// #                     Connection State Listener                         #
// #-----------------------------------------------------------------------#

func connStateListener(c net.Conn, cs http.ConnState) {
	if platformSettings.ALGORITHM == "LeastConnection" {
		switch cs {
		// Add an active connection to the serverPool map.
		case http.StateActive:
			urlContents = serverPool.GetUrl()
			serverPool.AddConn(urlContents)
			fmt.Println(serverPool.GetValues())
			log.Printf("Routing to URL: %s", urlContents)
		case http.StateIdle:
			// This code would be revised in future versions to
			// correctly remoce connections from the map.

			/*
				serverPool.RemoveConn(urlContents)
				fmt.Println(serverPool.GetValues())
				log.Printf("Removing from: %s", urlContents)
			*/
		}
	}
}
