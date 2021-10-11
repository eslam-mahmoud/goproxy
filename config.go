package main

import (
	"log"
	"os"
	"strings"
)

// set default values to be used if nothing was set from env vars
const defaultListen = ":8081"
const defaultProvider = "api.giphy.com:443,google.com:443"

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
		providers = defaultProvider
	}
	allowedProviders = strings.Split(providers, ",")

	log.Println("starting app")
	log.Println("listen:", listen)
	log.Println("providers:", allowedProviders)
}
