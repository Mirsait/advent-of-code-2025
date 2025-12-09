package main

import (
	"sort"
)

func puzzle2(points []Point3) int {
	connections := calcDistancesMap(points)

	sort.Slice(connections, func(j, k int) bool {
		return connections[j].D < connections[k].D
	})

	mst, _ := Kruskal(connections)
	last := mst[len(mst)-1]
	return last.P1.X * last.P2.X
}

type DSU struct {
	parent map[Point3]Point3
	rank   map[Point3]int
}

func NewDSU() *DSU {
	return &DSU{
		parent: make(map[Point3]Point3),
		rank:   make(map[Point3]int),
	}
}

func (d *DSU) MakeSet(p Point3) {
	if _, exists := d.parent[p]; !exists {
		d.parent[p] = p
		d.rank[p] = 0
	}
}

func (d *DSU) Find(p Point3) Point3 {
	// усли p - корень
	if d.parent[p] != p {
		// path compression
		d.parent[p] = d.Find(d.parent[p])
	}
	return d.parent[p]
}

func (d *DSU) Union(p1, p2 Point3) bool {
	root1 := d.Find(p1)
	root2 := d.Find(p2)

	if root1 == root2 {
		return false
	}

	// union by rank
	if d.rank[root1] < d.rank[root2] {
		d.parent[root1] = root2
	} else if d.rank[root1] > d.rank[root2] {
		d.parent[root2] = root1
	} else {
		d.parent[root2] = root1
		d.rank[root1]++
	}
	return true
}

func getUniquePoints(connections []Connection) []Point3 {
	pointSet := make(map[Point3]bool)
	uniquePoints := make([]Point3, 0)

	for _, conn := range connections {
		if !pointSet[conn.P1] {
			pointSet[conn.P1] = true
			uniquePoints = append(uniquePoints, conn.P1)
		}
		if !pointSet[conn.P2] {
			pointSet[conn.P2] = true
			uniquePoints = append(uniquePoints, conn.P2)
		}
	}

	return uniquePoints
}

func Kruskal(sortedConnections []Connection) ([]Connection, int) {
	// получаем все уникальные точки
	points := getUniquePoints(sortedConnections)

	// инициализируем DSU и создаем множества для всех точек
	dsu := NewDSU()
	for _, p := range points {
		dsu.MakeSet(p)
	}

	// строим MST
	mst := make([]Connection, 0, len(points)-1)
	totalWeight := 0

	for _, conn := range sortedConnections {
		// если точки в разных компонентах связности
		if dsu.Union(conn.P1, conn.P2) {
			mst = append(mst, conn)
			totalWeight += conn.D

			// если построили MST (n-1 ребро)
			if len(mst) == len(points)-1 {
				break
			}
		}
	}

	return mst, totalWeight
}
