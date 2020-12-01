package game

import (
	"log"
)

type cspPair struct {
	p1 int
	p2 int
}

// csp is the top level game structure.
// It holds the overall game state.
type csp struct {
	queens []queen
	conMap map[cspPair][]cspPair
}

func (c *csp) createQueens(initQs []int) {
	for i, v := range initQs {
		if v == -1 {
			// If queen is movable
			c.queens = append(c.queens, queen{col: i + 1, row: -1, moveable: true})
			c.queens[i].createDomain(len(initQs))

		} else {
			// If queen is restricted
			c.queens = append(c.queens, queen{i + 1, v, false, []int{v}})
		}
	}
}

func (c *csp) createConstraints() {
	n := len(c.queens)
	c.conMap = make(map[cspPair][]cspPair)

	// Iterate from the first queen to the second to last
	for i, p1 := range c.queens[0 : n-1] {
		// Iterate from the second queen to the last
		for _, p2 := range c.queens[i+1 : n] {
			kp := cspPair{p1.col, p2.col} // Map Key
			assignSlice := []cspPair{}    // Map Value

			for _, r1 := range p1.domain {
				for _, r2 := range p2.domain {
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

func (c *csp) isAssigned() bool {
	for _, q := range c.queens {
		if q.row == -1 {
			return false
		}
	}
	return true
}

func (c *csp) isInConstraintMap(q1, q2 queen, x1, x2 int) bool {
	if q2.col < q1.col {
		return c.isInConstraintMap(q2, q1, x2, x1)
	}
	constraints := c.conMap[cspPair{q1.col, q2.col}]

	for _, conPair := range constraints {
		if conPair.p1 == x1 && conPair.p2 == x2 {
			return true
		}
	}
	return false
}

func (c *csp) getCopy() csp {
	qcopy := make([]queen, 0)
	for i := range c.queens {
		qcopy = append(qcopy, c.queens[i])
	}
	return csp{qcopy, c.conMap}
}

func (c *csp) getQueen(col int) queen {
	idx := col - 1
	return c.queens[idx]
}

func (c *csp) print(i int) {
	log.Println("Game State: ", i)
	log.Println("\tSize: ", len(c.queens))
	log.Println()
	for _, q := range c.queens {
		log.Println("\tQueen ", q.col, ":")
		log.Println("\t\tColumn: ", q.col)
		log.Println("\t\tRow: ", q.row)
		log.Println("\t\tMovable: ", q.moveable)
		log.Println("\t\tDomain: ", q.domain)
		log.Println()
	}
	log.Println()
	log.Println("Constraint Maps:")
	n := len(c.queens)
	for i, p1 := range c.queens[0 : n-1] {
		for _, p2 := range c.queens[i+1 : n] {
			ap := cspPair{p1.col, p2.col}
			output, ok := c.conMap[ap]

			if ok {
				log.Println("\tQ(", p1.col, ",", p2.col, "):", output)
			}
		}
	}
}

/*NewGame creates a csp of size n with initial queen
placements according to map initQ*/
func NewGame(initQs []int) {
	var game = &csp{}
	game.createQueens(initQs)
	game.createConstraints()
	if ac3(game) {
		addInferences(game)
		game.print(0)
		sol := backtrackingSearch(game)
		for idx, g := range sol {
			log.Println("\n\nPrinting Solution ", idx)
			g.print(idx)
		}
		log.Println("Solution Length:", len(sol))
	} else {
		log.Println("Input is inconsistant with constraints.")
	}
}
