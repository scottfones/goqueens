package main

import (
	"fmt"
	"log"
	"net/http"

	"./web"
)

// Define Webpages
var index *web.View

var eightSetup *web.View
var eightDisplay *web.View

var twelveSetup *web.View
var twelveDisplay *web.View

var sixteenSetup *web.View
var sixteenDisplay *web.View

func main() {
	index = web.NewView("bootstrap", "web/index.gohtml")
	eightSetup = web.NewView("bootstrap", "web/eight_setup.gohtml")
	eightDisplay = web.NewView("bootstrap", "web/eight_display.gohtml")
	twelveSetup = web.NewView("bootstrap", "web/twelve_setup.gohtml")
	twelveDisplay = web.NewView("bootstrap", "web/twelve_display.gohtml")
	sixteenSetup = web.NewView("bootstrap", "web/sixteen_setup.gohtml")
	sixteenDisplay = web.NewView("bootstrap", "web/sixteen_display.gohtml")

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/eight_setup", eightSetupHandler)
	http.HandleFunc("/eight_display", eightDisplayHandler)
	http.HandleFunc("/twelve_setup", twelveSetupHandler)
	http.HandleFunc("/twelve_display", twelveDisplayHandler)
	http.HandleFunc("/sixteen_setup", sixteenSetupHandler)
	http.HandleFunc("/sixteen_display", sixteenDisplayHandler)

	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func eightSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 8-Queens Setup")
	eightSetup.Render(w, nil)
}

func eightDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("Loading: 8-Queens Display")
	eightDisplay.Render(w, nil)
}

func twelveSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 12-Queens Setup")
	twelveSetup.Render(w, nil)
}

func twelveDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("Loading: 12-Queens Display")
	twelveDisplay.Render(w, nil)
}

func sixteenSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 16-Queens Setup")
	sixteenSetup.Render(w, nil)
}

func sixteenDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("Loading: 16-Queens Display")
	sixteenDisplay.Render(w, nil)
}
