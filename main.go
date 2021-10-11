package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

func handleReq(w http.ResponseWriter, req *http.Request) {
	// handle onlly tunnel request
	// only http CONNECT methods allowed
	if req.Method != http.MethodConnect {
		log.Println(req.Method, req.RequestURI) // for debuging
		http.NotFound(w, req)
		return
	}

	// validate the tunned request destication is enabled
	found := false
	for _, dest := range allowedProviders {
		if req.RequestURI == dest {
			found = true
			break
		}
	}
	if found == false {
		log.Println("invalid destination", req.RequestURI)
		http.Error(w, "invalid destination", http.StatusBadRequest)
		return
	}

	tunnel(w, req)
}

// tunnel create a tunnel between the user and the
func tunnel(w http.ResponseWriter, req *http.Request) {
	log.Println("Handleing req")

	// Connect to destination
	dest, err := net.Dial("tcp", req.RequestURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// close the connection when function ends with or without errors
	defer dest.Close()

	// send status 200 code to client
	w.WriteHeader(http.StatusOK)

	// get the underlying net.Conn from the writer/source
	// to Read/write data from/to the connection.
	src, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// close the connection when function ends with or without errors
	defer src.Close()

	// run in baralel copy between src and dest
	// use wait group to wait for both to finish
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		_, err := io.Copy(dest, src)
		if err != nil {
			log.Println("copy dest, src err:", err)
		}
	}()
	go func() {
		defer wg.Done()

		_, err := io.Copy(src, dest)
		if err != nil {
			log.Println("copy src, dest err:", err)
		}
	}()
	// wait for copy between src & dest to be done
	wg.Wait()
}

func main() {
	// read command line flags and set default values
	loadConfig()

	handler := http.HandlerFunc(handleReq)

	err := http.ListenAndServe(listen, handler)
	if err == http.ErrServerClosed {
		log.Println("Gracefully stopping")
	} else {
		panic(err)
	}
}
