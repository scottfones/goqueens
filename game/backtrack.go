package game

import "log"

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

func isConsistent(c *csp, tq *queen) bool {
	for idx := range c.queens {
		cq := c.queens[idx]

		if cq.col == tq.col {
			continue
		} else if cq.row == -1 {
			continue
		} else if tq.isConflict(cq, tq.row, cq.row) {
			return false
		}
	}
	return true
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
	//c.print(2)
	if c.isAssigned() {
		return assignments
	}

	pivotQueen := mrv(c)

	//log.Println("FailLen: ", len(failureState))

	log.Println("Looking at Queen:", pivotQueen.col)
	for _, val := range pivotQueen.domain {
		tcsp := c.getCopy()
		tq := &tcsp.queens[pivotQueen.col-1]
		tq.row = val
		tq.domain = []int{val}

		log.Println("Domain Value:", val)
		//c.print(4)
		//tcsp.print(6)
		//panic("stop")

		//if isConsistent(&tcsp, tq) {
		//	log.Println("CON TRUE")
		//}

		if ac3(&tcsp) {
			//log.Println("AC3 True")
			//pivotQueen.row = val
			//pivotQueen.domain = []int{val}

			addInferences(&tcsp)
			//tcsp.print(3)
			//tcsp.print(1)
			assignments = append(assignments, tcsp)
			tresult := backtrack(&tcsp, assignments)
			//tcsp.print(13)

			log.Println("TRESULT: ", tresult)
			//curQueens := c.queens
			//retQueens := tresult[len(tresult)-1].queens

			//for i := range curQueens {
			//	if curQueens[i].isDomainEqual(retQueens[i].domain) {
			//		tresult = make([]csp, 0)
			//		break
			//	}
			//}
			// if result is not failure
			log.Println("LEN TRE:", len(tresult))
			if len(tresult) > 0 {
				//tcsp.print(25)
				//ac3(&tcsp)
				//addInferences(c)
				return tresult
			}
			//panic("stop")
		}
	}
	/// return failure
	return failureState
}

func backtrackingSearch(c *csp) []csp {
	//assignments := []csp{csp{queens: c.queens}}
	return backtrack(c, make([]csp, 0))
}
