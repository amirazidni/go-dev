package golanghttprouter

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	req := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "goodbye cruel world", string(body))
}
