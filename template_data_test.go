package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Rhmt",
		"Address": map[string]any{
			"Street": "Jalan XX",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title, Name string
	Address
}

func TemplateDataStruct(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Rhmt",
		Address: Address{
			Street: "Perumahan Bandara Mas",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
