package nav

import (
	"fmt"
	"math"

	astar "github.com/beefsack/go-astar"
)

type KindType int32

// Kind* constants refer to tile kinds for input and output.
const (
	KindTypePlain KindType = iota
	KindTypeBlocker
	KindTypeMoveItem
)

var (
	kindTypeMap = map[KindType]string{
		KindTypePlain:    "地面",
		KindTypeBlocker:  "阻挡物",
		KindTypeMoveItem: "攻击物",
	}
)

func (kt KindType) String() string {
	return kindTypeMap[kt]
}

// KindCosts map tile kinds to movement costs.
var kindCosts = map[KindType]float64{
	KindTypePlain:    1.0,
	KindTypeBlocker:  2.0,
	KindTypeMoveItem: 2.0,
}

func (kt KindType) Cost() float64 {
	return kindCosts[kt]
}

// A Tile is a tile in a grid which implements Pather.
type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	kindType KindType
	// X and Y are the coordinates of the tile.
	x, z int
	// W is a reference to the World that the tile is a part of.
	w *World
}

func (t *Tile) IsMask() bool {
	return t.kindType == KindTypePlain
}
func (t *Tile) String() string {
	return fmt.Sprintf("tile:type:%s,x:%d,z%d", t.kindType.String(), t.x, t.z)
}

func (t *Tile) GetXZ() (x int, z int) {
	return t.x, t.z
}

func NewTile(w *World, kt KindType, x int, z int) *Tile {
	t := &Tile{
		w:        w,
		kindType: kt,
		x:        x,
		z:        z,
	}
	return t
}

// func (t *Tile) Stand() {
// 	if t.kindType == KindTypePlain {
// 		t.kindType = KindTypeMoveItem
// 	}
// }

// func (t *Tile) Move() {
// 	if t.kindType == KindTypeMoveItem {
// 		t.kindType = KindTypePlain
// 	}
// }

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	leftTile := t.w.Tile(t.z, t.x-1)
	if leftTile != nil && leftTile.kindType == KindTypePlain {
		neighbors = append(neighbors, leftTile)
	}
	leftUpTile := t.w.Tile(t.z+1, t.x-1)
	if leftUpTile != nil && leftUpTile.kindType == KindTypePlain {
		neighbors = append(neighbors, leftUpTile)
	}
	leftDownTile := t.w.Tile(t.z-1, t.x-1)
	if leftDownTile != nil && leftDownTile.kindType == KindTypePlain {
		neighbors = append(neighbors, leftDownTile)
	}
	rightTile := t.w.Tile(t.z, t.x+1)
	if rightTile != nil && rightTile.kindType == KindTypePlain {
		neighbors = append(neighbors, rightTile)
	}
	rightUpTile := t.w.Tile(t.z+1, t.x+1)
	if rightUpTile != nil && rightUpTile.kindType == KindTypePlain {
		neighbors = append(neighbors, rightUpTile)
	}
	rightDownTile := t.w.Tile(t.z-1, t.x+1)
	if rightDownTile != nil && rightDownTile.kindType == KindTypePlain {
		neighbors = append(neighbors, rightDownTile)
	}
	upTile := t.w.Tile(t.z+1, t.x)
	if upTile != nil && upTile.kindType == KindTypePlain {
		neighbors = append(neighbors, upTile)
	}
	downTile := t.w.Tile(t.z-1, t.x)
	if downTile != nil && downTile.kindType == KindTypePlain {
		neighbors = append(neighbors, downTile)
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.x - t.x
	absZ := toT.z - t.z
	distance := math.Sqrt(float64(absX*absX + absZ*absZ))
	return toT.kindType.Cost() * distance
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absZ := toT.z - t.z
	if absZ < 0 {
		absZ = -absZ
	}
	return float64(absX + absZ)
}

// World is a two dimensional map of Tiles.
type World struct {
	startX      float64
	startZ      float64
	accuracy    int
	tileOfTiles map[int]map[int]*Tile
}

// Tile gets the tile at the given coordinates in the world.
func (w *World) Tile(z, x int) *Tile {
	ts, ok := w.tileOfTiles[z]
	if !ok {
		return nil
	}

	t, ok := ts[x]
	if !ok {
		return nil
	}
	return t
}

func (w *World) GetPositionForTile(tile *Tile) (x, z float64) {
	xIndex, zIndex := tile.GetXZ()
	return w.startX + float64(xIndex*w.accuracy), w.startZ - float64(zIndex*w.accuracy)
}

func (w *World) TileForPosition(x float64, z float64) *Tile {
	xIndex := int(math.Floor((x - w.startX) / float64(w.accuracy)))
	zIndex := int(math.Floor((w.startZ - z) / float64(w.accuracy)))
	return w.Tile(zIndex, xIndex)
}

func NewWorld(startX int, startZ int, accuracy int, maskMap [][]int) *World {

	w := &World{
		startX:   float64(startX),
		startZ:   float64(startZ),
		accuracy: accuracy,
	}

	w.tileOfTiles = make(map[int]map[int]*Tile)
	for z, xRow := range maskMap {
		tiles := make(map[int]*Tile)
		for x, mask := range xRow {
			kt := KindTypePlain
			if mask == 0 {
				kt = KindTypeBlocker
			}

			t := NewTile(w, kt, x, z)
			tiles[x] = t
		}

		w.tileOfTiles[z] = tiles
	}
	return w
}

func TileFromWorld(w *World, x float64, z float64) *Tile {
	// xInt := int(math.Floor(x))
	// zInt := int(math.Floor(z))

	return w.TileForPosition(x, z)
}
