package main

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

var templates = map[string]*template.Template{}

func init() {
	templateDir := os.DirFS("templates")
	matches, err := fs.Glob(templateDir, `*.html`)
	if err != nil {
		log.Fatal(err)
	}
	for _, templateFileName := range matches {
		file, err := templateDir.Open(templateFileName)
		if err != nil {
			log.Fatal(err)
		}
		contents, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		templateObject, err := template.New(strings.TrimSuffix(templateFileName, `.html`)).Parse(string(contents))
		if err != nil {
			log.Fatal(err)
		}
		templates[templateObject.Name()] = templateObject
	}

	for _, template := range templates {
		fmt.Println(template.Name())
	}
}

func main() {
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":5000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	template, ok := templates[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	template.Execute(w, nil)
}
