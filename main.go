package main

import (
	"fmt"
	"log"
	"net/http"

	"./web"
)

// Define Webpages
var index *web.View

var eightsetup *web.View
var eightdisplay *web.View

var twelvesetup *web.View
var twelvedisplay *web.View

func main() {
	index = web.NewView("bootstrap", "web/index.gohtml")
	eightsetup = web.NewView("bootstrap", "web/eight_setup.gohtml")
	eightdisplay = web.NewView("bootstrap", "web/eight_display.gohtml")
	twelvesetup = web.NewView("bootstrap", "web/twelve_setup.gohtml")
	twelvedisplay = web.NewView("bootstrap", "web/twelve_display.gohtml")

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/eight_setup", eightSetupHandler)
	http.HandleFunc("/eight_display", eightDisplayHandler)
	http.HandleFunc("/twelve_setup", twelveSetupHandler)
	http.HandleFunc("/twelve_display", twelveDisplayHandler)

	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func eightSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 8-Queens Setup")
	eightsetup.Render(w, nil)
}

func eightDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("Loading: 8-Queens Display")
	eightdisplay.Render(w, nil)
}

func twelveSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 12-Queens Setup")
	twelvesetup.Render(w, nil)
}

func twelveDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("Loading: 12-Queens Display")
	twelvedisplay.Render(w, nil)
}
