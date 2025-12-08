package main

import "slices"

type Chain struct {
	Connections []Connection
}

func (ch *Chain) AddConnection(conn Connection) {
	ch.Connections = append(ch.Connections, conn)
}

func (ch *Chain) Contains(conn Connection) bool {
	return slices.ContainsFunc(ch.Connections, conn.IsTouch)
}

func (ch *Chain) ContainsPoint(point Point3) bool {
	for _, c := range ch.Connections {
		if c.Contains(point) {
			return true
		}
	}
	return false
}

func (ch *Chain) PointCount() int {
	seen := make(map[Point3]bool)
	for _, conn := range ch.Connections {
		seen[conn.P1] = true
		seen[conn.P2] = true
	}
	return len(seen)
}
