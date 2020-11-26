package game

type arc struct {
	c1 int
	c2 int
}

type arcQueue struct {
	arcs []arc
}

func (q *arcQueue) add(a arc) {
	q.arcs = append(q.arcs, a)
}

func (q *arcQueue) pop() arc {
	if len(q.arcs) >= 1 {
		retArc := q.arcs[0]
		q.arcs = q.arcs[1:]
		return retArc
	}
	return arc{}
}

func (q *arcQueue) newQueue(c *csp) {
	n := len(c.queens)
	for _, q1 := range c.queens[0 : n-1] {
		for _, q2 := range c.queens[1:n] {
			a := arc{q1.col, q2.col}
			q.add(a)
		}
	}
}
