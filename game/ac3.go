package game

//import "log"

func revise(c *csp, col1, col2 int) bool {
	revised := false

	q1 := c.queens[col1-1]
	q2 := c.queens[col2-1]
	tmpDomain := make([]int, 0)

	for _, x1 := range q1.domain {
		for _, x2 := range q2.domain {
			if c.isInConstraintMap(q1, q2, x1, x2) {
				tmpDomain = append(tmpDomain, x1)
				break
			}
		}
	}

	if len(tmpDomain) != len(q1.domain) {
		//log.Println("Revise True")
		//log.Println("Orig: ", c.queens[col1-1].domain)
		//log.Println("New: ", tmpDomain)
		c.queens[col1-1].domain = tmpDomain
		revised = true
	}
	return revised
}

func ac3(c *csp) bool {
	//c.print(3)
	aq := arcQueue{}
	aq.newQueue(c)
	//log.Println("AQ: ", aq)

	for len(aq.arcs) > 0 {
		aqp := aq.pop()
		col1 := aqp.c1
		col2 := aqp.c2

		/*
		log.Println("Comparing Q", col1, " and Q", col2)
		log.Println("Q1 Domain: ", c.queens[col1-1].domain)
		log.Println("Q2 Domain: ", c.queens[col2-1].domain)
		cm, ok := c.conMap[cspPair{col1, col2}]
		if ok {
			log.Println("ConMap: ", cm)
		} else {
			log.Println("RevConMap: ", c.conMap[cspPair{col2, col1}])
		}
		*/

		if revise(c, col1, col2) {
			//log.Println("Revise: True")
			if len(c.queens[col1-1].domain) == 0 {
				//log.Println("AC3: False")
				//c.print(13)
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
