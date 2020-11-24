package game

import (
	"log"
)

type assignment struct {
	r1 int
	r2 int
}

type constraint struct {
	q1       *queen
	q2       *queen
	assSlice []assignment
}

/*csv is the top level game structure.
It holds the overall game state. */
type csp struct {
	size     int
	queens   []queen
	conSlice []constraint
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

	for i := 0; i < n-1; i++ {
		j := i + 1
		for j < n {
			p1 := c.queens[i]
			p2 := c.queens[j]
			con := constraint{q1: &p1, q2: &p2}

			for r1 := range p1.domain {
				for r2 := range p2.domain {
					log.Println("Comparing: (", p1.col, ",", r1, ") and (", p2.col, ",", r2, ")")
					if !(p1.isConflict(&p2, r1, r2)) {
						log.Println("No Conflict")
						a := assignment{r1, r2}
						con.assSlice = append(con.assSlice, a)
					}
				}
			}
			c.conSlice = append(c.conSlice, con)
			j++
		}
	}
}

/*NewGame creates a csv of size n with initial queen
placements according to map initQ*/
func NewGame(initQs []int) {
	var game = csp{size: len(initQs)}
	game.createQueens(initQs)
	game.createConstraints()
	log.Println(game)
}
