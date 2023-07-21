package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "WELCOME"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", handlerIndex)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	// http.StripPrefix() = delete prefix from requested endpoint
	// http.Dir() = adjustment path parameter

	fmt.Println("server started at localhost:9000")
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
