package game

/*Queen defines the fields necessary for the CSP.*/
type queen struct {
	col      int   `json:"col"`
	row      int   `json:"row"`
	moveable bool  `json:"moveable"`
	domain   []int `json:"-"`
}

func (q *queen) createDomain(n int) {
	dom := make([]int, n)

	for i := 0; i < n; i++ {
		dom[i] = i + 1
	}

	q.domain = dom
}

func (q *queen) isConflict(q2 *queen, r1, r2 int) bool {
	lineFit := (r1 - r2) / (q.col - q2.col)

	switch lineFit {
	case -1, 0, 1:
		return true
	}
	return false
}
