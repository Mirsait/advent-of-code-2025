package main

import (
	"fmt"
	"sort"
)

type Direction int

const (
	None = iota
	Right
	Left
	Up
	Down
)

func orthoConvexHull(points []Point2) []Point2 {

	sort.Slice(points, func(i, j int) bool {
		pi := points[i]
		pj := points[j]
		return pi.Y < pj.Y || (pi.Y == pj.Y && pi.X < pj.X)
	})

	// left-bottom point
	start := points[0]

	hull := []Point2{start}
	current := start
	direction := Right

	for {
		var nextPoint Point2
		var nextDirection Direction = None
		for _, p := range points {
			if p == current {
				continue
			}

			dx := p.X - current.X
			dy := p.Y - current.Y

			switch direction {
			case Right: // right?
				if dy == 0 && dx > 0 {
					if nextPoint == (Point2{}) || p.X < nextPoint.X {
						nextPoint = p
						nextDirection = Direction(direction)
					}
				}
			case Up:
				if dx == 0 && dy > 0 {
					if nextPoint == (Point2{}) || p.Y < nextPoint.Y {
						nextPoint = p
						nextDirection = Direction(direction)
					}
				}
			case Left:
				if dy == 0 && dx < 0 {
					if nextPoint == (Point2{}) || p.X > nextPoint.X {
						nextPoint = p
						nextDirection = Direction(direction)
					}
				}
			case Down:
				if dx == 0 && dy < 0 {
					if nextPoint == (Point2{}) || p.Y > nextPoint.Y {
						nextPoint = p
						nextDirection = Direction(direction)
					}
				}
			}

		}
		// if no nextpoint to current direction
		if nextPoint == (Point2{}) {
			if direction == Right { // right -> up
				for _, p := range points {
					if p == current {
						continue
					}
					if p.X == current.X && p.Y > current.Y {
						if nextPoint == (Point2{}) || p.X < nextPoint.X {
							nextPoint = p
							nextDirection = Up
						}
					}
				}
				fmt.Println("right -> up")
			} else if direction == Up { // up -> left or right
				for _, p := range points {
					if p == current {
						continue
					}
					if p.Y == current.Y {
						if p.X < current.X {
							if nextPoint == (Point2{}) || p.Y < nextPoint.Y {
								nextPoint = p
								nextDirection = Left
							}
						} else if p.X > current.X {
							if nextPoint == (Point2{}) || p.Y > nextPoint.Y {
								nextPoint = p
								nextDirection = Right
							}
						}
					}
				}
				fmt.Println("up -> left")
			} else if direction == Left { // left -> Down
				for _, p := range points {
					if p == current {
						continue
					}
					if p.X == current.X && p.Y < current.Y {
						if nextPoint == (Point2{}) || p.X > nextPoint.X {
							nextPoint = p
							nextDirection = Down
						}
					}
				}
				fmt.Println("left -> down")
			} else if direction == Down { // down -> right or left
				for _, p := range points {
					if p == current {
						continue
					}
					if p.Y == current.Y {
						if p.X > current.X {
							if nextPoint == (Point2{}) || p.Y > nextPoint.Y {
								nextPoint = p
								nextDirection = Right
							}
						} else if p.X < current.X {
							if nextPoint == (Point2{}) || p.Y < nextPoint.Y {
								nextPoint = p
								nextDirection = Left
							}
						}
					}
				}
				fmt.Println("down -> right")
			}
		}

		if nextPoint != (Point2{}) {
			if nextPoint == start && len(hull) > 1 {
				break
			}

			hull = append(hull, nextPoint)
			current = nextPoint
			direction = int(nextDirection)
		} else {
			break
		}
	} // for

	hull = append(hull, start)
	fmt.Println("hull:", hull)
	return hull
}

type Edge struct{ P1, P2 Point2 }

func hullToEdges(hull []Point2) []Edge {
	n := len(hull)
	edges := make([]Edge, 0)
	for j := 0; j < n-2; j++ {
		edges = append(edges, Edge{P1: hull[j], P2: hull[j+1]})
	}
	return edges
}

func getAreaFromEdges(edges []Edge) int {
	maxArea := 0
	n := len(edges) - 1
	for j := 0; j < n-1; j++ {
		for k := j + 1; k < n; k++ {
			e1 := edges[j]
			e2 := edges[k]
			if isRectangleSides(e1.P1, e1.P2, e2.P1, e2.P2) {
				d := areaByEdge(e1, e2)
				if d > maxArea {
					maxArea = d
				}
			}
		}
	}
	return maxArea
}

// Вектор из двух точек
func vec(a, b Point2) Point2 {
	return Point2{b.X - a.X, b.Y - a.Y, ""}
}

// Скалярное произведение
func dot(u, v Point2) int {
	return u.X*v.X + u.Y*v.Y
}

func cross(u, v Point2) int {
	return u.X*v.Y - u.Y*v.X
}

func length2(u Point2) int {
	return u.X*u.X + u.Y*u.Y
}

// Проверка: являются ли отрезки сторонами прямоугольника
func isRectangleSides(P1, P2, P3, P4 Point2) bool {
	A := vec(P1, P2)
	B := vec(P3, P4)

	// Проверка ненулевой длины
	if (A.X == 0 && A.Y == 0) || (B.X == 0 && B.Y == 0) {
		return false
	}

	// Случай соседних сторон (общая вершина + перпендикулярность)
	common := (P1 == P3) || (P1 == P4) || (P2 == P3) || (P2 == P4)
	if common && dot(A, B) == 0 {
		return true
	}

	// Случай противоположных сторон (параллельность + равные длины + нет общей вершины)
	if !common && cross(A, B) == 0 && length2(A) == length2(B) {
		return true
	}

	return false
}

func areaByEdge(e1, e2 Edge) int {
	// TODO
	panic("not implemented")
}

func isInside(point Point2) {
	// TODO
	panic("not implemented")
}
