package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// join folder
		filepath := path.Join("views", "index.html")

		// parsing template views/index.html
		tmpl, err := template.ParseFiles(filepath) //return *template.Template and error

		// response error if any
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// prepare interface data
		data := map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		// execute template (tmpl) with data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// routing static asset
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Listening at port 9000")
	http.ListenAndServe("localhost:9000", nil)
}
