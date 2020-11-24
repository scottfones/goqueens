package game

import (
	"log"
)

type assignment struct {
	r1 int
	r2 int
}

type constraint struct {
	q1       int
	q2       int
	assSlice []assignment
}

/*csv is the top level game structure.
It holds the overall game state. */
type csp struct {
	size     int          `json:"size"`
	queens   []queen      `json:"queens"`
	conSlice []constraint `json:"constraints"`
}

func (c *csp) createQueens(initQs []int) {
	for i, v := range initQs {
		if v == -1 {
			c.queens = append(c.queens, queen{col: i + 1, row: -1, moveable: true})
			c.queens[i].createDomain(c.size)

		} else {
			c.queens = append(c.queens, queen{i + 1, v + 1, false, []int{v + 1}})
		}
	}
}

func (c *csp) createConstraints() {
	n := c.size

	for i, p1 := range c.queens[0 : n-1] {
		for _, p2 := range c.queens[i+1 : n] {
			con := constraint{q1: p1.col, q2: p2.col}
			//log.Println("Q1: ", p1)
			//log.Println("Q2: ", p2)

			for _, r1 := range p1.domain {
				for _, r2 := range p2.domain {
					//log.Println("Comparing: (", p1.col, ",", r1, ") and (", p2.col, ",", r2, ")")
					if !(p1.isConflict(&p2, r1, r2)) {
						//log.Println("No Conflict")
						a := assignment{r1, r2}
						con.assSlice = append(con.assSlice, a)
					}
				}
			}
			//log.Println("Assigning: ", con.assSlice)
			c.conSlice = append(c.conSlice, con)
		}
	}
}

func (c *csp) revise(q1, q2 *queen) bool {
	revise := false

	return revise
}

func (c *csp) print(i int) {
	log.Println("Game State: ", i)
	log.Println("\tSize: ", c.size)
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
	log.Println("Constraints:")
	for _, con := range c.conSlice {
		//log.Println("(", con, ", ", con[1], "): ", con.assSlice)
		log.Println(con)
	}
}

/*NewGame creates a csv of size n with initial queen
placements according to map initQ*/
func NewGame(initQs []int) {
	var game = csp{size: len(initQs)}
	game.createQueens(initQs)
	game.createConstraints()
	game.print(0)
	//jsonStr, _ := json.Marshal(game)
	//log.Println(string(jsonStr))
	//log.Println(game)

}
