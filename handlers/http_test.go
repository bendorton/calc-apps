package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/bendorton/calc-apps/external/should"
)

func TestHTTPServer_NotFound(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/bogus", http.StatusNotFound, "text/plain; charset=utf-8", "404 page not found\n")
}
func TestHTTPServer_MethodNotAllowed(t *testing.T) {
	assertHTTP(t, http.MethodPost, "/add?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodPost, "/subtract?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodPost, "/multiply?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodPost, "/divide?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
}
func TestHTTPServer_Add(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/add?a=1&b=2", http.StatusOK, "text/plain; charset=utf-8", "3")
	assertHTTP(t, http.MethodGet, "/add?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument a\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument b\n")
}
func TestHTTPServer_Subtract(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/subtract?a=2&b=1", http.StatusOK, "text/plain; charset=utf-8", "1")
	assertHTTP(t, http.MethodGet, "/subtract?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument a\n")
	assertHTTP(t, http.MethodGet, "/subtract?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument b\n")
}
func TestHTTPServer_Multiply(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/multiply?a=2&b=2", http.StatusOK, "text/plain; charset=utf-8", "4")
	assertHTTP(t, http.MethodGet, "/multiply?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument a\n")
	assertHTTP(t, http.MethodGet, "/multiply?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument b\n")
}
func TestHTTPServer_Divide(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/divide?a=4&b=2", http.StatusOK, "text/plain; charset=utf-8", "2")
	assertHTTP(t, http.MethodGet, "/divide?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument a\n")
	assertHTTP(t, http.MethodGet, "/divide?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "Invalid argument b\n")
}

func assertHTTP(t *testing.T, method, target string, expectedStatus int, expectedContentType, expectedResponse string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)
		response := httptest.NewRecorder()

		dumpRequest, _ := httputil.DumpRequest(request, true)
		t.Log("\n" + string(dumpRequest))

		NewRouter(nil).ServeHTTP(response, request)

		dumpResponse, _ := httputil.DumpResponse(response.Result(), true)
		t.Log("\n" + string(dumpResponse))

		should.So(t, expectedStatus, should.Equal, response.Code)
		should.So(t, expectedContentType, should.Equal, response.Header().Get("Content-Type"))
		should.So(t, expectedResponse, should.Equal, response.Body.String())
	})
}
