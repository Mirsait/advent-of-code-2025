package main

type Point3 struct {
	X, Y, Z int
}

func (p *Point3) GetDistance2(other Point3) int {
	square := func(x int) int { return x * x }
	xx := square(p.X - other.X)
	yy := square(p.Y - other.Y)
	zz := square((p.Z - other.Z))
	return xx + yy + zz
}
