// queue to hold arc key pairs for constraint map.
// Used by ac3 and revise. Created to more closely
// adhere to the textbooks backtracking implementation.

package game

type arc struct {
	c1 int
	c2 int
}

// arcQueue is a slice of arcs.
type arcQueue struct {
	arcs []arc
}

// add appends an arc to the end of the queue
func (q *arcQueue) add(a arc) {
	q.arcs = append(q.arcs, a)
}

// pop returns the oldest arc in the queue.
// If empty, it returns an empty arc.
func (q *arcQueue) pop() arc {
	if len(q.arcs) >= 1 {
		retArc := q.arcs[0]
		q.arcs = q.arcs[1:]
		return retArc
	}
	return arc{}
}

// newQueue initializes a queue with all possible
// arc combinations except when q1 == q2.
func (q *arcQueue) newQueue(c *csp) {
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
