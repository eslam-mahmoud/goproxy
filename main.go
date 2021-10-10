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
		// log.Println(req.Method, req.RequestURI) // for debuging
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
	// TODO check request done through postmans/insomnia VS CMD
	// we get requestURI as "/api.giphy.com/v1/gifs/search?q=morning&api_key=123" from api client apps
	// we get requestURI as "api.giphy.com:443" from CMD
	// if string(req.RequestURI[0]) == "/" {
	// log.Println(req.RequestURI) // for debuging
	// 	req.RequestURI = req.RequestURI[1:]
	// }

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

// How to start the app
/*
https://linux.die.net/man/1/curl
-p/--proxytunnel
    When an HTTP proxy is used (-x/--proxy), this option will cause non-HTTP protocols to attempt to tunnel through the proxy instead of merely using it to do HTTP-like operations.

	The tunnel approach is made with the HTTP proxy CONNECT request and requires that the proxy allows direct connect to the remote port number curl wants to tunnel through to.

-x/--proxy <proxyhost[:port]>
    Use the specified HTTP proxy. If the port number is not specified, it is assumed at port 1080.

    This option overrides existing environment variables that set the proxy to use. If there's an environment variable setting a proxy, you can set proxy to "" to override it.

    Note that all operations that are performed over a HTTP proxy will transparently be converted to HTTP. It means that certain protocol specific operations might not be available. This is not the case if you can tunnel through the proxy, as done with the -p/--proxytunnel option.

    Starting with 7.14.1, the proxy host can be specified the exact same way as the proxy environment variables, including the protocol prefix (http://) and the embedded user + password.

    If this option is used several times, the last one will be used.


curl -p --proxy 127.0.0.1:12345 https://api.giphy.com/v1/gifs/search?q=morning\&api_key=123\&limit=1 | jq
*/
