package game

import "log"

type solution struct {
}

func mrv(c *csp) queen {
	var minq queen

	for i := range c.queens {
		q := c.queens[i]

		if q.row != -1 {
			continue

		} else if minq.row == 0 {
			minq = q

		} else if len(q.domain) < len(minq.domain) {
			minq = q
		}
	}

	return minq
}

func addInferences(c *csp) {
	for idx := range c.queens {
		q := c.queens[idx]

		if len(q.domain) == 1 {
			c.queens[idx].row = q.domain[0]
		}
	}
}

func backtrack(c *csp, assignments []csp) []csp {
	failureState := make([]csp, 0)

	if c.isAssigned() {
		return assignments
	}

	pivotQueen := mrv(c)
	log.Println("Looking at Queen:", pivotQueen.col)
	for _, val := range pivotQueen.domain {
		tcsp := c.getCopy()
		tq := &tcsp.queens[pivotQueen.col-1]
		tq.row = val
		tq.domain = []int{val}

		log.Println("Domain Value:", val)
		if ac3(&tcsp) {

			addInferences(&tcsp)
			assignments = append(assignments, tcsp)
			tresult := backtrack(&tcsp, assignments)

			if len(tresult) > 0 {
				log.Println("Solution Length:", len(tresult))
				return tresult
			}
		}
	}
	log.Println("Solution not found")
	return failureState
}

func backtrackingSearch(c *csp) []csp {
	return backtrack(c, make([]csp, 0))
}
