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

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "hello world")
	})
	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "hello world", string(body))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		fmt.Fprint(w, productId)
	})
	req := httptest.NewRequest("GET", "http://localhost:3000/products/P-002", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "P-002", string(body))
}

func TestRouterCatchAllParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("image")
		fmt.Fprint(w, productId)
	})
	req := httptest.NewRequest("GET", "http://localhost:3000/products/product-images.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "/product-images.png", string(body))
}
