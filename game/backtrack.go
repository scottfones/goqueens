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

		// If row has been assigned
		if q.Row != -1 {
			continue

		// If minq hasn't been assigned
		} else if minq.Row == 0 {
			minq = q

		// Update to more constrained domain
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

		// Move forward with copy to ease backtracking
		tcsp := c.getCopy()
		tq := &tcsp.queens[pivotQueen.Col-1]
		tq.Row = val
		tq.Domain = []int{val}

		if ac3(&tcsp) {
			// Construct solution and append
			s := Solution{pivotQueen.Col, val, tcsp.queens}
			assignments = append(assignments, s)

			tresult := backtrack(&tcsp, assignments)

			// If tresult is not the fail state
			if len(tresult) > 0 {
				log.Println("Solution Length:", len(tresult))
				return tresult
			}
			// If fail state, remove from assignments
			assignments = assignments[:len(assignments)-1]
		}
	}
	log.Println("Solution not found")
	return failureState
}

func backtrackingSearch(c *csp) []Solution {
	return backtrack(c, make([]Solution, 0))
}
