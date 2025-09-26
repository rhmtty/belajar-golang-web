package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:8000", nil)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("x-Powered-By", "Aing Maung")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:8000", nil)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-powered-by"))
}
