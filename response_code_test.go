package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode_Invalid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCode_Valid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000/?name=Rohmat", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
