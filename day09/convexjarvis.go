package main

import (
	"sort"
)

// Находим самую левую (и самую нижнюю) точку
func leftmost(points []Point2) Point2 {
	min := points[0]
	for _, p := range points {
		if p.X < min.X || (p.X == min.X && p.Y < min.Y) {
			min = p
		}
	}
	return min
}

// Векторное произведение (определяет ориентацию)
func orientation(p, q, r Point2) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
	if val == 0 {
		return 0 // коллинеарны
	}
	if val > 0 {
		return 1 // по часовой стрелке
	}
	return -1 // против часовой стрелки
}

func convexHullJarvis(points []Point2) []Point2 {
	n := len(points)
	if n < 3 {
		return points
	}

	// sort
	sort.Slice(points, func(i, j int) bool {
		pi := points[i]
		pj := points[j]
		return pi.X < pj.X || (pi.X == pj.X && pi.Y < pj.Y)
	})

	// Находим самую левую точку
	l := leftmost(points)

	hull := []Point2{}
	p := l

	for {
		hull = append(hull, p)

		q := points[0]
		if p == q {
			q = points[1]
		}

		// Ищем следующую точку оболочки
		for i := range n {
			// Если orientation(p, points[i], q) == -1, значит points[i] лежит левее
			if orientation(p, points[i], q) == -1 {
				q = points[i]
			}
		}

		p = q
		if p == l { // Вернулись к начальной точке
			break
		}
	}

	return hull
}
