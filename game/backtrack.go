package game

import "log"

type Solution struct {
	Column int
	Value  int
	Queens []Queen
}

func mrv(c *csp) Queen {
	var minq Queen

	for i := range c.queens {
		q := c.queens[i]

		if q.Row != -1 {
			continue

		} else if minq.Row == 0 {
			minq = q

		} else if len(q.Domain) < len(minq.Domain) {
			minq = q
		}
	}
	return minq
}

func backtrack(c *csp, assignments []Solution) []Solution {
	failureState := make([]Solution, 0)

	if c.isAssigned() {
		return assignments
	}

	pivotQueen := mrv(c)
	log.Println("Looking at Queen:", pivotQueen.Col)
	for _, val := range pivotQueen.Domain {
		log.Println("Domain Value:", val)

		tcsp := c.getCopy()
		tq := &tcsp.queens[pivotQueen.Col-1]
		tq.Row = val
		tq.Domain = []int{val}

		if ac3(&tcsp) {
			s := Solution{pivotQueen.Col, val, tcsp.queens}
			assignments = append(assignments, s)
			tresult := backtrack(&tcsp, assignments)

			if len(tresult) > 0 {
				log.Println("Solution Length:", len(tresult))
				return tresult
			}
			assignments = assignments[:len(assignments)-1]
		}
	}
	log.Println("Solution not found")
	return failureState
}

func backtrackingSearch(c *csp) []Solution {
	return backtrack(c, make([]Solution, 0))
}
