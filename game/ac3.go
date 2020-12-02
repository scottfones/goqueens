// Implementation of AC-3 and the revise heuristic
// as described in the textbook.

package game

// revise compares the domains of two queens and
// removes inconsistencies determined by the 
// constraint map. Returns true if the queen
// located in `col1` has its domain reduced.
func revise(c *csp, col1, col2 int) bool {
	revised := false

	// Assigned to avoid pointer issues
	q1 := c.queens[col1-1]
	q2 := c.queens[col2-1]
	tmpDomain := make([]int, 0)

	// Look up all combinations of domain values.
	// Validated values of `x1` are held in `tmpDomain`.
	for _, x1 := range q1.Domain {
		for _, x2 := range q2.Domain {
			if c.isInConstraintMap(q1, q2, x1, x2) {
				tmpDomain = append(tmpDomain, x1)
				break
			}
		}
	}

	// If validated domain differs from input domain
	if len(tmpDomain) != len(q1.Domain) {
		c.queens[col1-1].Domain = tmpDomain
		revised = true
	}
	return revised
}

// ac3 implements AC-3 as described in the textbook.
// Work through a FIFO queue of constraint map keys 
// by validating/revising domains wrt the queen 
// located at `col1`. If the domain is revised, adds
// all possible neighbors to the queue for revalidation.
// Returns true if every queen has atleast one element
// in its domain.
func ac3(c *csp) bool {
	aq := arcQueue{}
	aq.newQueue(c)

	// While the queue isn't empty
	for len(aq.arcs) > 0 {
		aqp := aq.pop()
		col1 := aqp.c1
		col2 := aqp.c2

		if revise(c, col1, col2) {
			if len(c.queens[col1-1].Domain) == 0 {
				return false
			}

			// Add all `col1` neighbors to the queue
			for _, qneighbor := range c.queens {
				switch qneighbor.Col {
				case col1, col2:
					// Skip self and just completed
					continue
				}
				aq.add(arc{col1, qneighbor.Col})
			}
		}
	}
	return true
}
