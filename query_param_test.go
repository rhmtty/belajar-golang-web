package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=Eko", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleQueryParam(writer http.ResponseWriter, req *http.Request) {
	firstName := req.URL.Query().Get("firstName")
	lastName := req.URL.Query().Get("lastName")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?firstName=Eko&lastName=Kurniawan", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParam(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParamValues(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleQueryValues(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=Rohmat&name=Tri&name=Y", nil)
	rec := httptest.NewRecorder()

	MultipleParamValues(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
