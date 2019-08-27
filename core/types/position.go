package types

import (
	"fmt"
)

type Position struct {
	X, Y, Z float64
}

func (p Position) IsEqual(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}

func (p Position) String() string {
	return fmt.Sprintf("x:%.2f,y:%.2f,z:%.2f", p.X, p.Y, p.Z)
}
