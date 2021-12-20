package maps

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strings"
)

type Coordinate3D struct {
	X int
	Y int
	Z int
}

func NewCoordinate3D(s string) Coordinate3D {
	split := strings.Split(s, ",")
	return Coordinate3D{
		X: util.Str2Int(split[0]),
		Y: util.Str2Int(split[1]),
		Z: util.Str2Int(split[2]),
	}
}

func (c Coordinate3D) String() string {
	return fmt.Sprintf("(x=%d,y=%d,z=%d)", c.X, c.Y, c.Z)
}

func (c Coordinate3D) Diff(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: util.AbsInt(c.X - to.X),
		Y: util.AbsInt(c.Y - to.Y),
		Z: util.AbsInt(c.Z - to.Z),
	}
}

func (c Coordinate3D) ManhattanDistance(to Coordinate3D) int {
	d := c.Diff(to)
	return d.X + d.Y + d.Z
}

func (c Coordinate3D) Sub(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X - to.X,
		Y: c.Y - to.Y,
		Z: c.Z - to.Z,
	}
}

func (c Coordinate3D) Add(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X + to.X,
		Y: c.Y + to.Y,
		Z: c.Z + to.Z,
	}
}

type Rotation3D interface {
	Apply(c Coordinate3D) Coordinate3D
}

type RotationDirection struct {
	X bool
	Y bool
	Z bool
}

func (r RotationDirection) Apply(c Coordinate3D) Coordinate3D {
	if r.X {
		c.X = c.X * -1
	}
	if r.Y {
		c.Y = c.Y * -1
	}
	if r.Z {
		c.Z = c.Z * -1
	}

	return c
}

type RotationFacing struct {
	X string
	Y string
	Z string

	Direction RotationDirection
}

func (r RotationFacing) Apply(c Coordinate3D) Coordinate3D {
	cnew := Coordinate3D{}
	switch r.X {
	case "x":
		cnew.X = c.X
	case "y":
		cnew.X = c.Y
	case "z":
		cnew.X = c.Z
	}

	switch r.Y {
	case "x":
		cnew.Y = c.X
	case "y":
		cnew.Y = c.Y
	case "z":
		cnew.Y = c.Z
	}

	switch r.Z {
	case "x":
		cnew.Z = c.X
	case "y":
		cnew.Z = c.Y
	case "z":
		cnew.Z = c.Z
	}

	return r.Direction.Apply(cnew)
}

func (c Coordinate3D) ApplyRotation(x, y, z int) Coordinate3D {
	return Coordinate3D{X: x * c.X, Y: y * c.Y, Z: z * c.Z}
}

