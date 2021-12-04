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
	moo()
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":5000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	err := templates.ExecuteTemplate(w, "layout.html", struct {
		AllTabs   []string
		ActiveTab string
	}{
		AllTabs:   []string{`inbox`, `laterbox`, `keepbox`, `history`, `account`},
		ActiveTab: name,
	})
	if err != nil {
		log.Fatal(err)
	}
}
