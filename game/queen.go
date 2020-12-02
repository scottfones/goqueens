// Struct and methods for each queen

package game

// Queen defines the fields necessary for the CSP.
// `Moveable` allows the View to encode whether the
// user restricted a queen's location.
type Queen struct {
	Col      int
	Row      int
	Moveable bool
	Domain   []int
}

// createDomain initializes each queen's domain to [1,n]
func (q *Queen) createDomain(n int) {
	dom := make([]int, n)

	for i := 0; i < n; i++ {
		dom[i] = i + 1
	}
	q.Domain = dom
}

// isConflict determines whether a queen, `q2`, assigned
// to row `r2` would be in conflict with the object queen,
// assigned to row `r1`. Returns true if the slope connecting
// the queens is -1, 0, or 1.
func (q *Queen) isConflict(q2 Queen, r1, r2 int) bool {
	// Calculate slope
	lineFit := (float64(r1) - float64(r2)) / float64(q.Col-q2.Col)

	switch lineFit {
	case -1, 0, 1:
		return true
	}
	return false
}
