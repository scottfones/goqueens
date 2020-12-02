// backtracking search, mrv heuristic, and solution struct

package game

import "log"

// Solution holds game state upon successful assignment of a queen.
type Solution struct {
	Column int
	Value  int
	Queens []Queen
}

// mrv returns the queen with the most constrained domain.
// When multiple such queens exist, it works through the csp
// from left-to-right (lowest column value first).
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

// backtrack places the most constrained queen recursively until 
// all queens are placed, or no solution can be found. Copies of
// the csp are used to minimize the amount of changes that would
// need to be backed out. 
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

// backtrackingSearch initializes the recursice backtrack function
// by passing the csp and an empty slice of solutions
func backtrackingSearch(c *csp) []Solution {
	return backtrack(c, make([]Solution, 0))
}
