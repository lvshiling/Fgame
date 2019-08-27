package utils

import (
	"fgame/fgame/core/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
)

func Lerp(sourcePos types.Position, targetPos types.Position, t float64) types.Position {

	// if t <= 0 {
	// 	return sourcePos
	// }
	// if t >= 1 {
	// 	return targetPos
	// }

	return types.Position{
		X: t*targetPos.X + (1-t)*sourcePos.X,
		Z: t*targetPos.Z + (1-t)*sourcePos.Z,
	}
}

func Distance(posA types.Position, posB types.Position) float64 {
	deltaX := posA.X - posB.X
	deltaZ := posA.Z - posB.Z
	return math.Sqrt(deltaX*deltaX + deltaZ*deltaZ)
}
func DistanceSquare(posA types.Position, posB types.Position) float64 {
	deltaX := posA.X - posB.X
	deltaZ := posA.Z - posB.Z
	return deltaX*deltaX + deltaZ*deltaZ
}

//获取角度
func GetAngle(sourcePos types.Position, targetPos types.Position) float64 {
	deltaZ := targetPos.Z - sourcePos.Z
	deltaX := targetPos.X - sourcePos.X

	//unity3d计算使用 tan = deltaX/deltaZ
	radian := math.Atan2(deltaX, deltaZ)
	angle := mathutils.RadianToAngle(radian)
	return angle
	// if deltaX == 0 && deltaZ == 0 {
	// 	return 0
	// }
	// deltaAngle := float64(0)
	// if deltaX > 0 {
	// 	if deltaZ == 0 {
	// 		return 90
	// 	}
	// 	if deltaZ < 0 {
	// 		deltaAngle = 180
	// 	}
	// } else if deltaX == 0 {
	// 	if deltaZ < 0 {
	// 		return 180
	// 	}
	// } else {
	// 	if deltaZ == 0 {
	// 		return -90
	// 	}

	// 	if deltaZ < 0 {
	// 		deltaAngle = -180
	// 	}
	// }

	// angle := mathutils.RadianToAngle(math.Atan(deltaX/deltaZ)) + deltaAngle
	// return angle
}

//获取角度
func GetRadian(sourcePos types.Position, targetPos types.Position) float64 {
	deltaZ := targetPos.Z - sourcePos.Z
	deltaX := targetPos.X - sourcePos.X

	radian := math.Atan2(deltaX, deltaZ)
	return radian
}

func PointInPolygon(pos types.Position, polygonPointList []types.Position) bool {
	if len(polygonPointList) < 3 {
		return false
	}
	nCross := 0
	lenOfPoly := len(polygonPointList)
	for i, point := range polygonPointList {
		nextIndex := i + 1
		if nextIndex >= lenOfPoly {
			nextIndex = 0
		}
		point2 := polygonPointList[nextIndex]
		if point.Z == point2.Z {
			continue
		}
		if pos.Z < math.Min(point.Z, point2.Z) {
			continue
		}
		if pos.Z >= math.Max(point.Z, point2.Z) {
			continue
		}
		x := (pos.Z-point.Z)*(point2.X-point.X)/(point2.Z-point.Z) + point.X
		if x > pos.X {
			nCross += 1
		}
	}
	return nCross&1 == 1
}

// func PolygonIntersectPolygon(polygon1 []types.Position, polygon2 []types.Position) bool {
// 	for _, point := range polygon1 {
// 		if PointInPolygon(point, polygon2) {
// 			return true
// 		}
// 	}

// 	for _, point := range polygon2 {
// 		if PointInPolygon(point, polygon1) {
// 			return true
// 		}
// 	}

// 	return false
// }

func checkCross(p1, p2, p3, p4 types.Position) bool {
	return checkTwoCross(p1, p2, p3, p4) <= 0 && (checkTwoCross(p3, p4, p2, p1) <= 0)
}

func crossMul(v1, v2 types.Position) float64 {
	return v1.X*v2.Z - v1.Z*v2.X
}

func checkTwoCross(p1, p2, p3, p4 types.Position) float64 {
	v1 := types.Position{
		X: p1.X - p3.X,
		Z: p1.Z - p3.Z,
	}
	v2 := types.Position{
		X: p2.X - p3.X,
		Y: p2.Z - p3.Z,
	}
	v3 := types.Position{
		X: p4.X - p3.X,
		Y: p4.Z - p3.Z,
	}

	v := crossMul(v1, v3) * crossMul(v2, v3)
	return v
}

var (
	circleSegNum = 12
)

func PolygonIntersectRound(polygon []types.Position, center types.Position, radius float64) bool {
	seg := make([]types.Position, 0, 12)

	for i := 0; i < circleSegNum; i++ {
		angle := float64(i) * math.Pi / float64(circleSegNum) * 2.0
		x := radius * math.Sin(angle)
		z := radius * math.Cos(angle)
		newPos := types.Position{
			center.X + x,
			center.Y,
			center.Z + z,
		}
		seg = append(seg, newPos)
	}
	return PolygonIntersectPolygon(polygon, seg)
}

func PolygonIntersectFan(polygon []types.Position, center types.Position, radius float64, faceAngle float64, fanAngle float64) bool {
	seg := make([]types.Position, 0, 12)

	startAngle := faceAngle - fanAngle/2.0
	endAngle := faceAngle + fanAngle/2.0
	elapseAngle := 360.0 / float64(circleSegNum)
	for angle := startAngle; angle <= endAngle; angle += elapseAngle {
		radian := mathutils.AngleToRadian(angle)
		x := radius * math.Sin(radian)
		z := radius * math.Cos(radian)
		newPos := types.Position{
			center.X + x,
			center.Y,
			center.Z + z,
		}
		seg = append(seg, newPos)
	}
	seg = append(seg, center)
	return PolygonIntersectPolygon(polygon, seg)
}

// func GetRectangle(center types.Position, length float64, width float64, angle float64) []types.Position {
// 	points := make([]types.Position, 0, 4)
// 	radian := mathutils.AngleToRadian(angle)

// 	point1 := types.Position{
// 		X: width/2*math.Sin(radian) - length/2*math.Cos(radian) + center.X,
// 		Y: center.Y,
// 		Z: length/2*math.Cos(radian) + width/2*math.Sin(radian) + center.Z,
// 	}
// 	point2 := types.Position{
// 		X: -width/2*math.Sin(radian) - length/2*math.Cos(radian) + center.X,
// 		Y: center.Y,
// 		Z: length/2*math.Cos(radian) - width/2*math.Sin(radian) + center.Z,
// 	}
// 	point3 := types.Position{
// 		X: -width/2*math.Sin(radian) + length/2*math.Cos(radian) + center.X,
// 		Y: center.Y,
// 		Z: -length/2*math.Cos(radian) - width/2*math.Sin(radian) + center.Z,
// 	}
// 	point4 := types.Position{
// 		X: width/2*math.Sin(radian) + length/2*math.Cos(radian) + center.X,
// 		Y: center.Y,
// 		Z: -length/2*math.Cos(radian) + width/2*math.Sin(radian) + center.Z,
// 	}
// 	points = append(points, point1, point2, point3, point4)
// 	return points
// }

func GetRectangle(center types.Position, length float64, width float64, angle float64) []types.Position {
	points := make([]types.Position, 0, 4)
	radian := mathutils.AngleToRadian(angle)

	point1 := types.Position{
		X: length/2*math.Cos(radian) + width/2*math.Sin(radian) + center.X,
		Y: center.Y,
		Z: -length/2*math.Sin(radian) + width/2*math.Cos(radian) + center.Z,
	}
	point2 := types.Position{
		X: -length/2*math.Cos(radian) + width/2*math.Sin(radian) + center.X,
		Y: center.Y,
		Z: length/2*math.Sin(radian) + width/2*math.Cos(radian) + center.Z,
	}
	point3 := types.Position{
		X: -length/2*math.Cos(radian) - width/2*math.Sin(radian) + center.X,
		Y: center.Y,
		Z: length/2*math.Sin(radian) - width/2*math.Cos(radian) + center.Z,
	}
	point4 := types.Position{
		X: length/2*math.Cos(radian) - width/2*math.Sin(radian) + center.X,
		Y: center.Y,
		Z: -length/2*math.Sin(radian) - width/2*math.Cos(radian) + center.Z,
	}
	points = append(points, point1, point2, point3, point4)
	return points
}

func GetRectangleByButtom(buttom types.Position, length float64, width float64, angle float64) []types.Position {
	points := make([]types.Position, 0, 4)
	radian := mathutils.AngleToRadian(angle)

	point1 := types.Position{
		X: length/2*math.Cos(radian) + buttom.X,
		Y: buttom.Y,
		Z: -length/2*math.Sin(radian) + buttom.Z,
	}
	point2 := types.Position{
		X: -length/2*math.Cos(radian) + buttom.X,
		Y: buttom.Y,
		Z: length/2*math.Sin(radian) + buttom.Z,
	}
	point3 := types.Position{
		X: -length/2*math.Cos(radian) + width*math.Sin(radian) + buttom.X,
		Y: buttom.Y,
		Z: length/2*math.Sin(radian) + width*math.Cos(radian) + buttom.Z,
	}
	point4 := types.Position{
		X: length/2*math.Cos(radian) + width*math.Sin(radian) + buttom.X,
		Y: buttom.Y,
		Z: -length/2*math.Sin(radian) + width*math.Cos(radian) + buttom.Z,
	}
	points = append(points, point1, point2, point3, point4)
	return points
}

//点乘
func Dot(v1, v2 types.Position) float64 {
	return v1.X*v2.X + v1.Z*v2.Z
}

//点乘
func Normalize(v types.Position) types.Position {
	mag := math.Sqrt(v.X*v.X + v.Z*v.Z)
	return types.Position{v.X / mag, 0, v.Z / mag}
}

//垂直
func Perpendicular(v types.Position) types.Position {
	return types.Position{v.Z, 0, -v.X}
}

func Segment(v1, v2 types.Position) types.Position {
	return types.Position{v2.X - v1.X, 0, v2.Z - v1.Z}
}

func Project(polygon []types.Position, axis types.Position) (min float64, max float64) {
	if len(polygon) < 3 {
		panic(fmt.Errorf("多边形顶点数需要不小于3"))
	}

	for i, p := range polygon {
		val := Dot(p, axis)
		if i == 0 {
			min = val
			max = val
			continue
		}
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return
}

func PolygonIntersectPolygon(polygon1 []types.Position, polygon2 []types.Position) bool {
	edges1 := getEdges(polygon1)
	for _, edge := range edges1 {
		axis := Perpendicular(edge)
		min2, max2 := Project(polygon2, axis)
		min1, max1 := Project(polygon1, axis)
		if max1 < min2 || max2 < min1 {
			return false
		}
	}
	edges2 := getEdges(polygon2)
	for _, edge := range edges2 {
		axis := Perpendicular(edge)
		min2, max2 := Project(polygon2, axis)
		min1, max1 := Project(polygon1, axis)
		if max1 < min2 || max2 < min1 {
			return false
		}
	}
	return true
}

func getEdges(polygon []types.Position) []types.Position {

	if len(polygon) < 3 {
		panic(fmt.Errorf("多边形顶点数需要不小于3"))
	}

	edges := make([]types.Position, 0, len(polygon)-1)
	lenOfPoints := len(polygon)
	for i := 0; i < lenOfPoints-1; i++ {
		p2 := polygon[i+1]
		p1 := polygon[i]
		edges = append(edges, types.Position{p2.X - p1.X, 0, p2.Z - p1.Z})
	}
	return edges
}
