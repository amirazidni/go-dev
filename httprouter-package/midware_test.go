package golanghttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	midware := LogMiddleware{router}

	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	midware.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Middleware", string(body))
}
