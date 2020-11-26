package game

import "log"

func revise(c *csp, col1, col2 int) bool {
	revised := false

	q1 := c.getQueen(col1)
	q2 := c.getQueen(col2)
	tmpDomain := make([]int, 0)

	for idx1 := 0; idx1 < len(q1.domain); {
		for idx2 := 0; idx2 < len(q2.domain); {
			if !(q1.isConflict(q2, q1.domain[idx1], q2.domain[idx2])) {
				tmpDomain = append(tmpDomain, q1.domain[idx1])
				idx2 = len(q2.domain)
				continue
			}
			idx2++
		}
		idx1++
	}

	if !(q1.isDomainEqual(tmpDomain)) {
		log.Println("revised: ", tmpDomain)
		revised = true
		q1.domain = tmpDomain
	}

	return revised
}

func ac3(c *csp) bool {
	aq := arcQueue{}
	aq.newQueue(c)
	log.Println("AQ: ", aq)

	for len(aq.arcs) > 0 {
		aqp := aq.pop()
		col1 := aqp.c1
		col2 := aqp.c2

		if revise(c, col1, col2) {
			qn := c.getQueen(col1)
			if len(qn.domain) == 0 {
				return false
			}

			for _, qneighbor := range c.queens {
				switch qneighbor.col {
				case col1, col2:
					continue
				}
				aq.add(arc{col1, qneighbor.col})
			}
		}
		//log.Println(len(aq.arcs))
	}
	return true
}
