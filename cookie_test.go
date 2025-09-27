package belajargolangweb

import (
	"fmt"
	"net/http"
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
	cookie, err := req.Cookie("x-cookie-name")
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
