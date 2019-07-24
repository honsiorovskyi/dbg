package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}

func main() {
	addr, ok := os.LookupEnv("LISTEN")
	if !ok {
		addr = ":8080"
	}

	http.HandleFunc("/", handle)
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
