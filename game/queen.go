package game

/*Queen defines the fields necessary for the CSP.*/
type Queen struct {
	Col      int
	Row      int
	Moveable bool
	Domain   []int
}

func (q *Queen) createDomain(n int) {
	dom := make([]int, n)

	for i := 0; i < n; i++ {
		dom[i] = i + 1
	}
	q.Domain = dom
}

func (q *Queen) isDomainEqual(tmpDom []int) bool {
	if len(q.Domain) != len(tmpDom) {
		return false
	}

	for i, x := range q.Domain {
		if x != tmpDom[i] {
			return false
		}
	}
	return true
}

func (q *Queen) isConflict(q2 Queen, r1, r2 int) bool {
	lineFit := (float64(r1) - float64(r2)) / float64(q.Col-q2.Col)

	switch lineFit {
	case -1, 0, 1:
		return true
	}
	return false
}
