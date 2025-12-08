package main

import (
	"fmt"
	"slices"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func puzzle1(points []Point3, limit int) int {
	connections := calcDistancesMap(points)
	// sort by distance
	slices.SortFunc(connections, func(m, n Connection) int {
		if m.D > n.D {
			return 1
		}
		if m.D < n.D {
			return -1
		}
		return 0
	})

	connections = connections[:limit]

	chains := make([]Chain, 0)
	for _, conn := range connections {
		AddChain(conn, &chains)
	}
	var counts []int
	for _, ch := range chains {
		local := ch.PointCount()
		counts = append(counts, local)
	}
	slices.SortFunc(counts, func(x, y int) int {
		if x < y {
			return 1
		} else if x > y {
			return -1
		}
		return 0
	})
	return common.Reduce(func(acc, x int) int { return acc * x }, 1, counts[:3])
}

func AddChain(conn Connection, chains *[]Chain) {
	var indices []int

	// look for chains that contain new connection points
	for j, ch := range *chains {
		if ch.Contains(conn) ||
			ch.ContainsPoint(conn.P1) ||
			ch.ContainsPoint(conn.P2) {
			indices = append(indices, j)
		}
	}

	if len(indices) == 0 {
		// none of the chains fit -> create a new chain
		newChain := Chain{}
		newChain.AddConnection(conn)
		*chains = append(*chains, newChain)
	} else {
		// take the first suitable chain
		target := &(*chains)[indices[0]]
		target.AddConnection(conn)

		// If there are other suitable chains -> merge them
		for _, idx := range indices[1:] {
			target.Connections = append(target.Connections,
				(*chains)[idx].Connections...)
		}

		// delete unnecessary chains
		for j := len(indices) - 1; j >= 1; j-- {
			idx := indices[j]
			*chains = append((*chains)[:idx], (*chains)[idx+1:]...)
		}
	}
}

func calcDistancesMap(points []Point3) []Connection {
	connectionsMap := make(map[string]Connection)

	for j := range points {
		for k := j + 1; k < len(points); k++ {
			d := points[j].GetDistance2(points[k])
			if d == 0 {
				continue
			}
			key := makeKey(points[j], points[k])
			connectionsMap[key] = Connection{P1: points[j], P2: points[k], D: d}
		}
	}

	// map to slice
	connections := make([]Connection, 0, len(connectionsMap))
	for _, conn := range connectionsMap {
		connections = append(connections, conn)
	}
	return connections
}

func makeKey(p1, p2 Point3) string {
	// "A-B" == "B-A"
	if p1.X < p2.X ||
		(p1.X == p2.X && p1.Y < p2.Y) ||
		(p1.X == p2.X && p1.Y == p2.Y && p1.Z < p2.Z) {
		return fmt.Sprintf("%v-%v", p1, p2)
	}
	return fmt.Sprintf("%v-%v", p2, p1)
}
