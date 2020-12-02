// Controller, Router, and launchpoint for webapp

package main

import (
	"log"
	"net/http"
	"strconv"

	"./game"
	"./web"
)

// Define Webpages
var index *web.View
var nosolution *web.View

var eightSetup *web.View
var eightDisplay *web.View

var twelveSetup *web.View
var twelveDisplay *web.View

var sixteenSetup *web.View
var sixteenDisplay *web.View

func main() {
	// View Construction
	index = web.NewView("bootstrap", "web/index.gohtml")
	nosolution = web.NewView("bootstrap", "web/no_solution.gohtml")
	eightSetup = web.NewView("bootstrap", "web/eight_setup.gohtml")
	eightDisplay = web.NewView("bootstrap", "web/solution_display.gohtml")
	twelveSetup = web.NewView("bootstrap", "web/twelve_setup.gohtml")
	twelveDisplay = web.NewView("bootstrap", "web/solution_display.gohtml")
	sixteenSetup = web.NewView("bootstrap", "web/sixteen_setup.gohtml")
	sixteenDisplay = web.NewView("bootstrap", "web/solution_display.gohtml")

	// Router Definitions
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/eight_setup", eightSetupHandler)
	http.HandleFunc("/eight_display", eightDisplayHandler)
	http.HandleFunc("/twelve_setup", twelveSetupHandler)
	http.HandleFunc("/twelve_display", twelveDisplayHandler)
	http.HandleFunc("/sixteen_setup", sixteenSetupHandler)
	http.HandleFunc("/sixteen_display", sixteenDisplayHandler)

	// Start Server
	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)

}

// processQSlice encodes user input for the csp.
func processQSlice(qs []string) []int {
	islice := make([]int, 0)
	for _, s := range qs {
		if len(s) == 0 {
			// Empty assignments are encoded
			// as moveable queens.
			islice = append(islice, -1)

		} else {
			// Convert string row assignment
			i, err := strconv.Atoi(s)

			if err != nil {
				log.Fatal(err)
			}
			islice = append(islice, i)
		}
	}
	return islice
}

// indexHandler renders the homepage.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

// eightSetupHandler renders the 8-queen form data.
func eightSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 8-Queens Setup")
	eightSetup.Render(w, nil)
}

// eightDisplayHandler parses the 8-queen form data and initiates a csp.
func eightDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Collect form data
	qslice := []string{
		r.Form.Get("setq1"),
		r.Form.Get("setq2"),
		r.Form.Get("setq3"),
		r.Form.Get("setq4"),
		r.Form.Get("setq5"),
		r.Form.Get("setq6"),
		r.Form.Get("setq7"),
		r.Form.Get("setq8"),
	}

	solns := game.NewGame(processQSlice(qslice))

	if string(solns) == "[]" {
		log.Println("Loading: No Solution")
		nosolution.Render(w, nil)
	} else {
		log.Println("Loading: 8-Queens Display")
		// Construct an interface to pass the solutions
		eightDisplay.Render(w, map[string]string{"solns": string(solns)})
	}
}

// twelveSetupHandler renders the 12-queen form data.
func twelveSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 12-Queens Setup")
	twelveSetup.Render(w, nil)
}

// twelveDisplayHandler parses the 12-queen form data and initiates a csp.
func twelveDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Collect form data
	qslice := []string{
		r.Form.Get("setq1"),
		r.Form.Get("setq2"),
		r.Form.Get("setq3"),
		r.Form.Get("setq4"),
		r.Form.Get("setq5"),
		r.Form.Get("setq6"),
		r.Form.Get("setq7"),
		r.Form.Get("setq8"),
		r.Form.Get("setq9"),
		r.Form.Get("setq10"),
		r.Form.Get("setq11"),
		r.Form.Get("setq12"),
	}

	solns := game.NewGame(processQSlice(qslice))

	if string(solns) == "[]" {
		log.Println("Loading: No Solution")
		nosolution.Render(w, nil)
	} else {
		log.Println("Loading: 12-Queens Display")
		// Construct an interface to pass the solutions
		twelveDisplay.Render(w, map[string]string{"solns": string(solns)})
	}
}

// sixteenSetupHandler renders the 16-queen form data.
func sixteenSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading: 16-Queens Setup")
	sixteenSetup.Render(w, nil)
}

// sixteenDisplayHandler parses the 16-queen form data and initiates a csp.
func sixteenDisplayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Collect form data
	qslice := []string{
		r.Form.Get("setq1"),
		r.Form.Get("setq2"),
		r.Form.Get("setq3"),
		r.Form.Get("setq4"),
		r.Form.Get("setq5"),
		r.Form.Get("setq6"),
		r.Form.Get("setq7"),
		r.Form.Get("setq8"),
		r.Form.Get("setq9"),
		r.Form.Get("setq10"),
		r.Form.Get("setq11"),
		r.Form.Get("setq12"),
		r.Form.Get("setq13"),
		r.Form.Get("setq14"),
		r.Form.Get("setq15"),
		r.Form.Get("setq16"),
	}

	solns := game.NewGame(processQSlice(qslice))

	if string(solns) == "[]" {
		log.Println("Loading: No Solution")
		nosolution.Render(w, nil)
	} else {
		log.Println("Loading: 16-Queens Display")
		// Construct an interface to pass the solutions
		sixteenDisplay.Render(w, map[string]string{"solns": string(solns)})
	}
}
