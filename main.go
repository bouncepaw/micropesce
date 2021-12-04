package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

func init() {
	var err error
	var templatesDir = os.DirFS("templates")
	templates, err = template.ParseFS(templatesDir, "*")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseFlags()
	http.HandleFunc("/account", serveAccount)
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":5000", nil)
}

func defaultArgsMap() map[string]interface{} {
	return map[string]interface{}{
		"AllTabs": []string{`inbox`, `laterbox`, `keepbox`, `history`, `account`},
	}
}

type Feed struct {
	ID   uint
	Name string
	URL  string
}

func serveAccount(w http.ResponseWriter, rq *http.Request) {
	m := defaultArgsMap()
	m["ActiveTab"] = "account"
	m["Feeds"] = []Feed{
		{1, "moto moto", "gemini://jaba jaba"},
		{2, "cows", "gemini://cows.cows"},
		{3, "crows", "gemini://crows.crows"},
	}
	err := templates.ExecuteTemplate(w, "layout.html", m)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	m := defaultArgsMap()
	m["ActiveTab"] = name
	err := templates.ExecuteTemplate(w, "layout.html", m)
	if err != nil {
		log.Fatal(err)
	}
}
