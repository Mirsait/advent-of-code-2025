package main

type Connection struct {
	P1, P2 Point3
	D      int // square
}

func (conn *Connection) IsEqual(other Connection) bool {
	return (conn.P1 == other.P1 && conn.P2 == other.P2) ||
		(conn.P1 == other.P2 && conn.P2 == other.P1)
}

func (conn *Connection) IsTouch(other Connection) bool {
	return conn.P1 == other.P1 || conn.P2 == other.P2 ||
		conn.P1 == other.P2 || conn.P2 == other.P1
}

func (conn *Connection) Contains(point Point3) bool {
	return conn.P1 == point || conn.P2 == point
}
