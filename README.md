# goqueens
Find solutions to n-Queens Constraint Satisfaction Problems via AC-3 and Backtracking search. 

MVC pattern implemented in Golang:
* main.go
    * Controller connecting the web (View) package with the game (Model) package
* web package
    * Views generated with ["html/template"](https://golang.org/pkg/html/template/) and [Material Design Lite](https://getmdl.io/)
* game package
    * Model consists of a Constraint Satisfaction Problem (CSP) with a slice of n queens and a map of constraints
        * Each queen tracks its column, row, and valid domain assignments. It also stores a boolean, `moveable`, indicating whether the user has restricted its location.
    * A Backtracking search is used to find solutions to the CSP
        * The queen with the most constrained domain is selected for assignment by an MRV heuristic. 
        * AC-3 is used to maintain consistency and generate inferences.
        * Solutions are returned to the controller as JSON for display.

To Run:
1. Download [Go](https://golang.org/dl/)
2. `cd` into `goqueens` project directory
3. Execute `go run main.go` 
4. Browse to [localhost:3000](localhost:3000)