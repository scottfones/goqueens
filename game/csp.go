// Define csp struct and methods. 
// NewGame() exports entrypoint to model

package game

import (
	"encoding/json"
	"log"
)

type cspPair struct {
	p1 int
	p2 int
}

// csp is the top level game structure
// and holds the game state.
type csp struct {
	queens []Queen
	conMap map[cspPair][]cspPair // Map of key pair to valid row assignments
}

// createQueens from the slice generated in the view.
func (c *csp) createQueens(initQs []int) {
	for i, v := range initQs {
		if v == -1 {
			// If queen is movable
			c.queens = append(c.queens, Queen{Col: i + 1, Row: -1, Moveable: true})
			c.queens[i].createDomain(len(initQs))

		} else {
			// If queen is restricted
			c.queens = append(c.queens, Queen{i + 1, v, false, []int{v}})
		}
	}
}

// createConstraints via line fitting. If row
// assignments are valid, add to the conMap.
func (c *csp) createConstraints() {
	n := len(c.queens)
	c.conMap = make(map[cspPair][]cspPair) // Initialize empty conMap

	// Iterate from the first queen to the second to last
	for i, p1 := range c.queens[0 : n-1] {
		// Iterate from the second queen to the last
		for _, p2 := range c.queens[i+1 : n] {
			kp := cspPair{p1.Col, p2.Col} // Map key
			assignSlice := []cspPair{}    // Initialize empty map value

			for _, r1 := range p1.Domain {
				for _, r2 := range p2.Domain {
					if !(p1.isConflict(p2, r1, r2)) {
						// If no conflict add domain pair
						a := cspPair{r1, r2}
						assignSlice = append(assignSlice, a)
					}
				}
			}
			// Add entry to map
			c.conMap[kp] = assignSlice
		}
	}
}

// isAssgned returns true if every 
// queen has been assigned a row.
func (c *csp) isAssigned() bool {
	for _, q := range c.queens {
		if q.Row == -1 {
			return false
		}
	}
	return true
}

// isInConstraintMap returns true if `q1` at row `x1` and
// `q2` at row `x2` is found within the constraint map.
func (c *csp) isInConstraintMap(q1, q2 Queen, x1, x2 int) bool {
	if q2.Col < q1.Col {
		// Reverse queen order
		// Constraint map is generated in one "direction"
		return c.isInConstraintMap(q2, q1, x2, x1)
	}
	constraints := c.conMap[cspPair{q1.Col, q2.Col}]

	// Iterate through all pairs of valid row assignments
	for _, conPair := range constraints {
		if conPair.p1 == x1 && conPair.p2 == x2 {
			return true
		}
	}
	return false
}

// getCopy returns a copy of the csp to
// reduce pointer complications.
func (c *csp) getCopy() csp {
	qcopy := make([]Queen, 0)
	for i := range c.queens {
		qcopy = append(qcopy, c.queens[i])
	}
	return csp{qcopy, c.conMap}
}

// print writes formatted csp state to the log
func (c *csp) print(i int) {
	log.Println("Game State: ", i)
	log.Println("\tSize: ", len(c.queens))
	log.Println()
	printQueens(c.queens)
	log.Println()
	log.Println("Constraint Maps:")

	// Print constraint map in order
	n := len(c.queens)
	for i, p1 := range c.queens[0 : n-1] {
		for _, p2 := range c.queens[i+1 : n] {
			ap := cspPair{p1.Col, p2.Col}

			// If mapping exists, print
			output, ok := c.conMap[ap]
			if ok {
				log.Println("\tQ(", p1.Col, ",", p2.Col, "):", output)
			}
		}
	}
}

// printQueens formats queen field data and
// writes it to the log.
func printQueens(qs []Queen) {
	for _, q := range qs {
		log.Println("\tQueen ", q.Col, ":")
		log.Println("\t\tColumn: ", q.Col)
		log.Println("\t\tRow: ", q.Row)
		log.Println("\t\tMovable: ", q.Moveable)
		log.Println("\t\tDomain: ", q.Domain)
		log.Println()
	}
}

// NewGame creates a csp with initial queen
// state passed in from the view
func NewGame(initQs []int) []byte {
	var game = &csp{}
	game.createQueens(initQs)
	game.createConstraints()
	solns := make([]Solution, 0)

	// Early consistency check due to user input
	if ac3(game) {
		solns = backtrackingSearch(game)
		log.Println("Initial State:")
		game.print(0)
		for idx, s := range solns {
			log.Println("Solution Step", idx+1)
			log.Println("\tPivot Queen: ", s.Column)
			log.Println("\tAssigned Value: ", s.Value)
			log.Println("\tQueen States:")
			printQueens(s.Queens)
		}
		log.Println("Solution Length:", len(solns))
	} else {
		log.Println("Input is inconsistant with constraints.")
	}
	// Form JSON from solution structs
	jsonData, err := json.Marshal(solns)
	if err != nil {
		log.Fatal(err)
	}

	// Slice of solutions or empty slice
	return jsonData
}
