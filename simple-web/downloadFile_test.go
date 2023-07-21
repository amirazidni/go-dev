package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	// more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
