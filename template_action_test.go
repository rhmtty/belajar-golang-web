package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Rhmt",
	})
}

func TestTemplateActionIf(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateActionIf(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]any{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateActionOperator(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]any{
		"Title": "Template Action Operator",
		"Hobbies": []string{
			"Game", "Sport", "Mancing",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateActionRange(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
func TemplateActionWith(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]any{
		"Title": "Template Action Operator",
		"Name":  "Rhmt",
		"Address": map[string]any{
			"Street": "Perumahan Bandara Mas",
			"City":   "Tangerang",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000", nil)
	rec := httptest.NewRecorder()

	TemplateActionWith(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
