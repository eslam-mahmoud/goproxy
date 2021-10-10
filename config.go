package main

// import (
// 	"flag"
// 	"log"
// 	"strings"
// )

// const defaultListen = ":8080"
// const defaultProvider = "api.giphy.com:443"

// var (
// 	allowedProviders = []string{}
// 	listen           = ""
// )

// func loadConfig() {
// 	// return pointer
// 	listen = *flag.String("listen", defaultListen, "The address and port to listen for http connections on")
// 	providers := *flag.String("providers", defaultProvider, "comma separated string of enabled providers")

// 	// parses from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
// 	flag.Parse()
// 	allowedProviders = strings.Split(providers, ",")
// 	log.Println("starting app")
// 	log.Println("listen:", listen)
// 	log.Println("providers:", allowedProviders)

// 	// RUN
// 	// go run main.go -listen=":8080" -providers="api.giphy.com:443,google.com:80"
// }

import (
	"log"
	"os"
	"strings"
)

// set default values to be used if nothing was set from env vars
const defaultListen = ":8080"
const defaultProvider = "api.giphy.com:443"

var (
	allowedProviders = []string{}
	listen           = ""
)

//loadConfig read env vars to get config or use default values
func loadConfig() {
	// return pointer
	listen = os.Getenv("LISTEN")
	if listen == "" {
		listen = defaultListen
	}

	providers := os.Getenv("PROVIDERS")
	if providers == "" {
		allowedProviders = append(allowedProviders, defaultProvider)
	} else {
		allowedProviders = strings.Split(providers, ",")
	}

	log.Println("starting app")
	log.Println("listen:", listen)
	log.Println("providers:", allowedProviders)
}
