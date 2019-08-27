package aoi

import (
	"container/list"
	"fgame/fgame/core/types"
	"fmt"
	"sync"
)

type xzAOI struct {
	AOI          AOI
	XElement     *list.Element
	ZElement     *list.Element
	neighbors    map[*xzAOI]struct{}
	markVal      int
	markOuterVal int
}

func (aoi *xzAOI) Reset() {
	aoi.AOI = nil
	aoi.XElement = nil
	aoi.ZElement = nil
	aoi.neighbors = make(map[*xzAOI]struct{})
	aoi.markVal = 0
	aoi.markOuterVal = 0
}
func (aoi *xzAOI) Mark() {
	aoi.markVal += 1
}

func (aoi *xzAOI) MarkOut() {
	aoi.markOuterVal += 1
}

func (aoi *xzAOI) IsNeighbor() bool {
	return aoi.markVal == 2
}

func (aoi *xzAOI) IsStillNeighbor() bool {
	return aoi.markOuterVal == 2
}

func (aoi *xzAOI) Clear() {
	aoi.markVal = 0
	aoi.markOuterVal = 0
}

func (aoi *xzAOI) RemoveNeighbor(neighbor *xzAOI, complete bool) {
	_, exist := aoi.neighbors[neighbor]
	if !exist {
		panic("repeat remove same neighbor")
	}
	delete(aoi.neighbors, neighbor)
	aoi.AOI.OnLeaveAOI(neighbor.AOI, complete)
}

func (aoi *xzAOI) AddNeighbor(neighbor *xzAOI) {
	_, exist := aoi.neighbors[neighbor]
	if exist {
		panic("repeat add same neighbor")
	}
	aoi.neighbors[neighbor] = struct{}{}
	aoi.AOI.OnEnterAOI(neighbor.AOI)
}

type XYListAOIManager struct {
	xzAOIMap      map[int64]*xzAOI
	xList         *list.List
	zList         *list.List
	enterDistance float64
	exitDistance  float64
}

//TODO 修改为缓冲距离
func NewXZListAOIManager(enterDistance float64, exitDistance float64) *XYListAOIManager {
	m := &XYListAOIManager{}
	m.xzAOIMap = make(map[int64]*xzAOI)
	m.xList = list.New()
	m.zList = list.New()
	m.enterDistance = enterDistance
	m.exitDistance = exitDistance
	return m
}

func (m *XYListAOIManager) getAOIData(id int64) (aoi *xzAOI) {
	aoi, ok := m.xzAOIMap[id]
	if !ok {
		return nil
	}
	return
}

func (m *XYListAOIManager) addAOIData(aoi *xzAOI) {
	m.xzAOIMap[aoi.AOI.GetId()] = aoi
}

func (m *XYListAOIManager) removeAOIData(aoi *xzAOI) {
	delete(m.xzAOIMap, aoi.AOI.GetId())
}

var xzAOIPool = &sync.Pool{
	New: func() interface{} {
		return &xzAOI{
			neighbors: make(map[*xzAOI]struct{}),
		}
	},
}

func (m *XYListAOIManager) Enter(aoi AOI, pos types.Position) {

	xyaoi := m.getAOIData(aoi.GetId())
	if xyaoi != nil {
		return
	}

	// aoi.X = x
	// aoi.Y = y
	aoi.SetPosition(pos)
	xzAOI := xzAOIPool.Get().(*xzAOI)
	xzAOI.AOI = aoi
	m.insertX(xzAOI)
	m.insertZ(xzAOI)
	m.addAOIData(xzAOI)
	m.adjust(xzAOI)
	return
}

func (m *XYListAOIManager) debugXList() {
	for e := m.xList.Front(); e != nil; e = e.Next() {
		xyaoi := e.Value.(*xzAOI)
		fmt.Println(xyaoi.AOI)
	}
}
func (m *XYListAOIManager) debugYList() {
	for e := m.zList.Front(); e != nil; e = e.Next() {
		xyaoi := e.Value.(*xzAOI)
		fmt.Println(xyaoi.AOI)
	}
}

func (m *XYListAOIManager) insertX(aoi *xzAOI) {
	x := aoi.AOI.GetPosition().X
	var te *list.Element
	for e := m.xList.Front(); e != nil; e = e.Next() {
		taoi, _ := e.Value.(*xzAOI)
		if x < taoi.AOI.GetPosition().X {
			te = e
			break
		}
	}
	var nel *list.Element
	if te == nil {
		nel = m.xList.PushBack(aoi)
	} else {
		nel = m.xList.InsertBefore(aoi, te)
	}
	//方便删除
	aoi.XElement = nel
}

func (m *XYListAOIManager) insertZ(aoi *xzAOI) {

	z := aoi.AOI.GetPosition().Z
	var te *list.Element
	for e := m.zList.Front(); e != nil; e = e.Next() {
		taoi, ok := e.Value.(*xzAOI)
		if !ok {
			panic("never reach here")
		}
		if z < taoi.AOI.GetPosition().Z {
			te = e
			break
		}
	}
	var nel *list.Element
	if te == nil {
		nel = m.zList.PushBack(aoi)
	} else {
		nel = m.zList.InsertBefore(aoi, te)
	}
	aoi.ZElement = nel
}

func (m *XYListAOIManager) Move(aoi AOI, pos types.Position) {
	xyaoi := m.getAOIData(aoi.GetId())
	if xyaoi == nil {
		//aoi no exist
		return
	}
	x := pos.X
	z := pos.Z
	oldX := aoi.GetPosition().X
	// oldY := aoi.GetY()
	oldZ := aoi.GetPosition().Z
	xyaoi.AOI.SetPosition(pos)

	dirty := false
	if oldX != x {
		m.moveX(xyaoi, oldX)
		dirty = true
	}
	if oldZ != z {
		m.moveZ(xyaoi, oldZ)
		dirty = true
	}
	if dirty {
		m.adjust(xyaoi)
	}
}

func (m *XYListAOIManager) markX(aoi *xzAOI) {

	for prevX := aoi.XElement.Prev(); prevX != nil; prevX = prevX.Prev() {
		prevXAOI := prevX.Value.(*xzAOI)
		distance := aoi.AOI.GetPosition().X - prevXAOI.AOI.GetPosition().X
		if distance > m.exitDistance {
			break
		}
		prevXAOI.MarkOut()
		if distance > m.enterDistance {
			continue
		}
		prevXAOI.Mark()
	}

	for nextX := aoi.XElement.Next(); nextX != nil; nextX = nextX.Next() {
		nextXAOI := nextX.Value.(*xzAOI)
		distance := nextXAOI.AOI.GetPosition().X - aoi.AOI.GetPosition().X
		if distance > m.exitDistance {
			break
		}
		nextXAOI.MarkOut()
		if distance > m.enterDistance {
			continue
		}
		nextXAOI.Mark()
	}

}

func (m *XYListAOIManager) markZ(aoi *xzAOI) {
	for prevZ := aoi.ZElement.Prev(); prevZ != nil; prevZ = prevZ.Prev() {
		prevZAOI := prevZ.Value.(*xzAOI)
		distance := aoi.AOI.GetPosition().Z - prevZAOI.AOI.GetPosition().Z
		if distance > m.exitDistance {
			break
		}
		prevZAOI.MarkOut()
		if distance > m.enterDistance {
			continue
		}
		prevZAOI.Mark()
	}
	for nextZ := aoi.ZElement.Next(); nextZ != nil; nextZ = nextZ.Next() {
		nextZAOI := nextZ.Value.(*xzAOI)
		distance := nextZAOI.AOI.GetPosition().Z - aoi.AOI.GetPosition().Z
		if distance > m.exitDistance {
			break
		}
		nextZAOI.MarkOut()
		if distance > m.enterDistance {
			continue
		}
		nextZAOI.Mark()
	}
}

func (m *XYListAOIManager) newMarkedNeighborAndClear(aoi *xzAOI) {

	for nextX := aoi.XElement.Next(); nextX != nil; nextX = nextX.Next() {
		nextXAOI := nextX.Value.(*xzAOI)
		distance := nextXAOI.AOI.GetPosition().X - aoi.AOI.GetPosition().X
		if distance > m.exitDistance {
			break
		}
		if nextXAOI.IsNeighbor() {
			aoi.AddNeighbor(nextXAOI)
			nextXAOI.AddNeighbor(aoi)
		}
		nextXAOI.Clear()
	}
	for prevX := aoi.XElement.Prev(); prevX != nil; prevX = prevX.Prev() {
		prevXAOI := prevX.Value.(*xzAOI)
		distance := aoi.AOI.GetPosition().X - prevXAOI.AOI.GetPosition().X
		if distance > m.exitDistance {
			break
		}
		if prevXAOI.IsNeighbor() {
			aoi.AddNeighbor(prevXAOI)
			prevXAOI.AddNeighbor(aoi)
		}
		prevXAOI.Clear()
	}

	for prevZ := aoi.ZElement.Prev(); prevZ != nil; prevZ = prevZ.Prev() {
		prevZAOI := prevZ.Value.(*xzAOI)
		distance := aoi.AOI.GetPosition().Z - prevZAOI.AOI.GetPosition().Z
		if distance > m.exitDistance {
			break
		}
		prevZAOI.Clear()
	}
	for nextZ := aoi.ZElement.Next(); nextZ != nil; nextZ = nextZ.Next() {
		nextZAOI := nextZ.Value.(*xzAOI)
		distance := nextZAOI.AOI.GetPosition().Z - aoi.AOI.GetPosition().Z
		if distance > m.exitDistance {
			break
		}
		nextZAOI.Clear()
	}
}

func (m *XYListAOIManager) adjust(aoi *xzAOI) {
	m.markX(aoi)
	m.markZ(aoi)

	for neighbor, _ := range aoi.neighbors {
		//still neighbor
		if neighbor.IsStillNeighbor() {
			neighbor.Clear()
			continue
		}
		aoi.RemoveNeighbor(neighbor, false)
		neighbor.RemoveNeighbor(aoi, false)
	}

	//判断新加入和清除标记
	m.newMarkedNeighborAndClear(aoi)

}

func (m *XYListAOIManager) moveX(aoi *xzAOI, oldX float64) {
	ox := aoi.AOI.GetPosition().X
	if ox == oldX {
		return
	}

	if ox < oldX {
		prev := aoi.XElement.Prev()
		if prev == nil {
			//keep
			return
		}
		//find suitable aoi
		prevData, _ := prev.Value.(*xzAOI)
		for ox < prevData.AOI.GetPosition().X {
			//remove and insert before prev
			if prev.Prev() == nil {
				m.xList.MoveBefore(aoi.XElement, prev)
				return
			}
			prev = prev.Prev()
			prevData, _ = prev.Value.(*xzAOI)
		}
		m.xList.MoveAfter(aoi.XElement, prev)
	} else {
		//moving to next
		next := aoi.XElement.Next()
		if next == nil {
			//keep
			return
		}
		nextData, _ := next.Value.(*xzAOI)
		for ox > nextData.AOI.GetPosition().X {
			if next.Next() == nil {
				m.xList.MoveAfter(aoi.XElement, next)
				return
			}
			next = next.Next()
			nextData, _ = next.Value.(*xzAOI)
		}
		m.xList.MoveBefore(aoi.XElement, next)
	}

}

func (m *XYListAOIManager) moveZ(aoi *xzAOI, oldZ float64) {
	oz := aoi.AOI.GetPosition().Z
	if oz == oldZ {
		return
	}

	if oz < oldZ {
		prev := aoi.ZElement.Prev()
		if prev == nil {
			//keep
			return
		}
		//find suitable aoi
		prevData, _ := prev.Value.(*xzAOI)
		for oz < prevData.AOI.GetPosition().Z {
			//remove and insert before prev
			if prev.Prev() == nil {
				m.zList.MoveBefore(aoi.ZElement, prev)
				return
			}
			prev = prev.Prev()
			prevData, _ = prev.Value.(*xzAOI)
		}
		m.zList.MoveAfter(aoi.ZElement, prev)
	} else {
		//moving to next
		next := aoi.ZElement.Next()
		if next == nil {
			//keep
			return
		}
		nextData, _ := next.Value.(*xzAOI)
		for oz > nextData.AOI.GetPosition().Z {
			if next.Next() == nil {
				m.zList.MoveAfter(aoi.ZElement, next)
				return
			}
			next = next.Next()
			nextData, _ = next.Value.(*xzAOI)
		}
		m.zList.MoveBefore(aoi.ZElement, next)
	}
}

func (m *XYListAOIManager) Leave(aoi AOI) {
	// fmt.Printf("%d 退出aoi前\n", aoi.GetId())
	//get element
	xyaoi := m.getAOIData(aoi.GetId())
	if xyaoi == nil {
		//aoi no exist
		return
	}
	// fmt.Printf("%d 退出aoi后\n", aoi.GetId())
	//移除
	for neighbor, _ := range xyaoi.neighbors {
		neighbor.RemoveNeighbor(xyaoi, true)
	}
	//移除从列表
	m.removeX(xyaoi)
	m.removeZ(xyaoi)
	//移除缓存数据
	m.removeAOIData(xyaoi)
	xyaoi.Reset()
	xzAOIPool.Put(xyaoi)
	return
}

func (m *XYListAOIManager) removeX(aoi *xzAOI) {
	m.xList.Remove(aoi.XElement)
}
func (m *XYListAOIManager) removeZ(aoi *xzAOI) {
	m.zList.Remove(aoi.ZElement)
}
