package game

/*Queen defines the fields necessary for the CSP.*/
type queen struct {
	col      int
	row      int
	moveable bool
	domain   []int
}

func (q *queen) createDomain(n int) {
	dom := make([]int, n)

	for i := 0; i < n; i++ {
		dom[i] = i + 1
	}
	q.domain = dom
}

func (q *queen) isDomainEqual(tmpDom []int) bool {
	if len(q.domain) != len(tmpDom) {
		return false
	}

	for i, x := range q.domain {
		if x != tmpDom[i] {
			return false
		}
	}
	return true
}

func (q *queen) isConflict(q2 queen, r1, r2 int) bool {
	lineFit := (float64(r1) - float64(r2)) / float64(q.col-q2.col)

	switch lineFit {
	case -1, 0, 1:
		return true
	}
	return false
}
