package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var layoutDir string = "web/templates"
var webpage *template.Template

func main() {
	var err error
	webpage, err = template.ParseFiles(layoutFiles()...)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", handler)

	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	webpage.ExecuteTemplate(w, "webpage", nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "/*.gohtml")
	if err != nil {
		log.Println(err)
	}
	return files
}
