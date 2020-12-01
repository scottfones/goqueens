package game

func revise(c *csp, col1, col2 int) bool {
	revised := false

	q1 := c.queens[col1-1]
	q2 := c.queens[col2-1]
	tmpDomain := make([]int, 0)

	for _, x1 := range q1.Domain {
		for _, x2 := range q2.Domain {
			if c.isInConstraintMap(q1, q2, x1, x2) {
				tmpDomain = append(tmpDomain, x1)
				break
			}
		}
	}

	if len(tmpDomain) != len(q1.Domain) {
		c.queens[col1-1].Domain = tmpDomain
		revised = true
	}
	return revised
}

func ac3(c *csp) bool {
	aq := arcQueue{}
	aq.newQueue(c)

	for len(aq.arcs) > 0 {
		aqp := aq.pop()
		col1 := aqp.c1
		col2 := aqp.c2

		if revise(c, col1, col2) {
			if len(c.queens[col1-1].Domain) == 0 {
				return false
			}

			for _, qneighbor := range c.queens {
				switch qneighbor.Col {
				case col1, col2:
					continue
				}
				aq.add(arc{col1, qneighbor.Col})
			}
		}
	}
	return true
}
