package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	// prepare multiplexer
	mux := http.NewServeMux()

	// define handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// error if method not GET
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		// response body
		text := "hello world"

		// if instance_id is any, add to response body
		if instanceID != "" {
			text = text + ". from " + instanceID
		}

		w.Write([]byte(text))
	})
	// prepare server
	server := new(http.Server)

	// prepare multiplexer
	server.Handler = mux

	// listen to port
	server.Addr = "0.0.0.0:" + port

	// log start
	log.Println("server starting at", server.Addr)

	// start server
	err := server.ListenAndServe()

	// catch error
	if err != nil {
		log.Fatal(err.Error())
	}
}
