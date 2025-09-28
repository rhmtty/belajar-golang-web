package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "x-cookie-name"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("X-Cookie-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000/?name=Rohmat", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookies := rec.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Cookie-Name"
	cookie.Value = "Rhmt"
	req.AddCookie(cookie)

	rec := httptest.NewRecorder()

	GetCookie(rec, req)
	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
