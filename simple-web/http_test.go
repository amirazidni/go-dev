package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye world")
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/bye", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// query params is case sensitive
// multiple parameter value, one param with multiple value
func MultiParamValue(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["name"]
	fmt.Fprint(w, strings.Join(query, " "))
}

func TestMultiParamValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Test&name=Multi&name=Value&name=params", nil)
	rec := httptest.NewRecorder()

	MultiParamValue(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
