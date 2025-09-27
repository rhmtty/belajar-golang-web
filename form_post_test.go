package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := req.PostForm.Get("firstName")
	lastName := req.PostForm.Get("lastName")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	reqBody := strings.NewReader("firstName=Rohmat&lastName=Y")
	req := httptest.NewRequest("POST", "http://localhost:8000", reqBody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	FormPost(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
