package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var port string

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s\t%s\n", req.Method, req.URL.Path)
		for k, v := range req.Header {
			for _, vv := range v {
				res.Header().Add(k, vv)
			}
		}
		w := io.MultiWriter(os.Stdout, res)
		io.Copy(w, req.Body)
	})

	flag.StringVar(&port, "port", ":6969", "the port to run the server on")
	flag.Parse()

	fmt.Printf("Starting ek at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
