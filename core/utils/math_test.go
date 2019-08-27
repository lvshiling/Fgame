package utils_test

import "testing"
import coretypes "fgame/fgame/core/types"
import . "fgame/fgame/core/utils"

type testPointInPolygonStruct struct {
	Point    coretypes.Position
	Polygons []coretypes.Position
	Result   bool
}

var (
	testPointInPolygonSuits = []*testPointInPolygonStruct{
		&testPointInPolygonStruct{
			Point: coretypes.Position{0, 0, 0},
			Polygons: []coretypes.Position{
				coretypes.Position{1, 0, 1},
				coretypes.Position{1, 0, -1},
				coretypes.Position{-1, 0, -1},
				coretypes.Position{-1, 0, 1},
			},
			Result: true,
		},
		&testPointInPolygonStruct{
			Point: coretypes.Position{1, 0, 1},
			Polygons: []coretypes.Position{
				coretypes.Position{1, 0, 1},
				coretypes.Position{1, 0, -1},
				coretypes.Position{-1, 0, -1},
				coretypes.Position{-1, 0, 1},
			},
			Result: true,
		},
		&testPointInPolygonStruct{
			Point: coretypes.Position{2, 0, 1},
			Polygons: []coretypes.Position{
				coretypes.Position{1, 0, 1},
				coretypes.Position{1, 0, -1},
				coretypes.Position{-1, 0, -1},
				coretypes.Position{-1, 0, 1},
			},
			Result: false,
		},
	}
)

func TestPointInPolygon(t *testing.T) {
	for _, s := range testPointInPolygonSuits {
		result := PointInPolygon(s.Point, s.Polygons)
		if result != result {
			t.Errorf("点[%s]在多边形[%s],期望[%b],得到[%b]", s.Point, s.Polygons, s.Result, result)
		}
	}
}
