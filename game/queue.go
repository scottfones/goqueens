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
	//n := len(c.queens)
	for _, q1 := range c.queens {
		for _, q2 := range c.queens {
			if q1.Col == q2.Col {
				continue
			}
			a := arc{q1.Col, q2.Col}
			q.add(a)
		}
	}
}
