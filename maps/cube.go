package maps

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
)

type Cuboid struct {
	From Coordinate3D
	To   Coordinate3D
}

func (c Cuboid) Coordinates() []Coordinate3D {
	var coordinates []Coordinate3D
	for x := c.From.X; x <= c.To.X; x++ {
		for y := c.From.Y; y <= c.To.Y; y++ {
			for z := c.From.Z; z <= c.To.Z; z++ {
				coordinates = append(coordinates, Coordinate3D{
					X: x, Y: y, Z: z,
				})
			}
		}
	}

	return coordinates
}

func (c Cuboid) Contains(co Cuboid) bool {
	return c.From.X <= co.From.X && c.To.X >= co.To.X &&
		c.From.Y <= co.From.Y && c.To.Y >= co.To.Y &&
		c.From.Z <= co.From.Z && c.To.Z >= co.To.Z
}

func (c Cuboid) String() string {
	return fmt.Sprintf("(x=%d..%d,y=%d..%d,z=%d..%d)",
		c.From.X, c.To.X, c.From.Y, c.To.Y, c.From.Z, c.To.Z)
}

func (c Cuboid) Size() int {
	return util.AbsInt(1 *
		(c.To.X - c.From.X) *
		(c.To.Y - c.From.Y) *
		(c.To.Z - c.From.Z))
}

func (c Cuboid) Subdivide(co Cuboid) ([]Cuboid, *Cuboid, []Cuboid) {
	if !c.IsOverlapping(co) {
		return []Cuboid{c}, nil, []Cuboid{co}
	}

	xvals := []int{
		util.MinInt(c.From.X, co.From.X),
		util.MaxInt(c.From.X, co.From.X),
		util.MinInt(c.To.X, co.To.X),
		util.MaxInt(c.To.X, co.To.X),
	}
	yvals := []int{
		util.MinInt(c.From.Y, co.From.Y),
		util.MaxInt(c.From.Y, co.From.Y),
		util.MinInt(c.To.Y, co.To.Y),
		util.MaxInt(c.To.Y, co.To.Y),
	}
	zvals := []int{
		util.MinInt(c.From.Z, co.From.Z),
		util.MaxInt(c.From.Z, co.From.Z),
		util.MinInt(c.To.Z, co.To.Z),
		util.MaxInt(c.To.Z, co.To.Z),
	}

	var cCuboids []Cuboid
	var sharedCuboid *Cuboid
	var coCuboids []Cuboid
	for xi := 0; xi < len(xvals)-1; xi++ {
		for yi := 0; yi < len(yvals)-1; yi++ {
			for zi := 0; zi < len(zvals)-1; zi++ {
				cuboid := Cuboid{
					From: Coordinate3D{
						X: xvals[xi],
						Y: yvals[yi],
						Z: zvals[zi],
					},
					To: Coordinate3D{
						X: xvals[xi+1],
						Y: yvals[yi+1],
						Z: zvals[zi+1],
					},
				}

				switch {
				case c.Contains(cuboid) && co.Contains(cuboid):
					sharedCuboid = &cuboid
				case !c.Contains(cuboid) && co.Contains(cuboid):
					coCuboids = append(coCuboids, cuboid)
				case c.Contains(cuboid) && !co.Contains(cuboid):
					cCuboids = append(cCuboids, cuboid)
				}
			}
		}
	}

	return cCuboids, sharedCuboid, coCuboids
}

func (c Cuboid) IsOverlapping(co Cuboid) bool {
	minx := util.MaxInt(c.From.X, co.From.X)
	miny := util.MaxInt(c.From.Y, co.From.Y)
	minz := util.MaxInt(c.From.Z, co.From.Z)

	return c.From.X <= minx && c.To.X >= minx &&
		co.From.X <= minx && co.To.X > minx &&
		c.From.Y <= miny && c.To.Y >= miny &&
		co.From.Y <= miny && co.To.Y > miny &&
		c.From.Z <= minz && c.To.Z >= minz &&
		co.From.Z <= minz && co.To.Z > minz
}

func (c Cuboid) Overlapping(co Cuboid) *Cuboid {
	if !c.IsOverlapping(co) {
		return nil
	}

	return &Cuboid{
		From: Coordinate3D{
			X: util.MaxInt(co.From.X, c.From.X),
			Y: util.MaxInt(co.From.Y, c.From.Y),
			Z: util.MaxInt(co.From.Z, c.From.Z),
		},
		To: Coordinate3D{
			X: util.MinInt(co.To.X, c.To.X),
			Y: util.MinInt(co.To.Y, c.To.Y),
			Z: util.MinInt(co.To.Z, c.To.Z),
		},
	}
}
