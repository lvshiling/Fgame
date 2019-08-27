package template

import (
	"encoding/json"
	coretypes "fgame/fgame/core/types"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"path/filepath"
)

type mapPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type MapMaskMap struct {
	StartPos *mapPosition `json:"startPos"`
	Mask     [][]int      `json:"mask"`
}

func (mmm *MapMaskMap) Fixed() *MapMaskMap {
	nmmm := &MapMaskMap{
		StartPos: mmm.StartPos,
	}
	maskListOfList := make([][]int, 0, len(mmm.Mask))
	for xIndex, row := range mmm.Mask {
		maskList := make([]int, 0, len(row))
		for zIndex, mask := range row {
			if mask == 0 {
				maskList = append(maskList, mask)
			} else {
				width := len(mmm.Mask)
				tempXIndex := xIndex - 1
				//左边
				tempZIndex := zIndex
				if tempXIndex < 0 {
					maskList = append(maskList, 0)
					continue
				}
				tempRow := mmm.Mask[tempXIndex]
				if tempRow[tempZIndex] == 0 {
					maskList = append(maskList, 0)
					continue
				}
				//右边
				tempXIndex = xIndex + 1
				if tempXIndex >= width {
					maskList = append(maskList, 0)
					continue
				}
				tempRow = mmm.Mask[tempXIndex]
				if tempRow[tempZIndex] == 0 {
					maskList = append(maskList, 0)
					continue
				}
				//上边
				tempZIndex = zIndex - 1
				if tempZIndex < 0 {
					maskList = append(maskList, 0)
					continue
				}
				if row[tempZIndex] == 0 {
					maskList = append(maskList, 0)
					continue
				}
				height := len(row)
				//下边
				tempZIndex = zIndex + 1
				if tempZIndex >= height {
					maskList = append(maskList, 0)
					continue
				}
				if row[tempZIndex] == 0 {
					maskList = append(maskList, 0)
					continue
				}
				maskList = append(maskList, mask)
			}
		}
		maskListOfList = append(maskListOfList, maskList)
	}
	nmmm.Mask = maskListOfList
	return nmmm
}

type mapHeightMap [][]float64

type Map struct {
	maskMap   *MapMaskMap
	heightMap mapHeightMap
	width     int
	length    int

	standList []int32
}

func (m *Map) GetMaskMap() *MapMaskMap {
	return m.maskMap
}

func (m *Map) GetWidth() int {
	return m.width
}

func (m *Map) GetLength() int {
	return m.length
}

func (m *Map) IsValid(x float64, z float64) bool {
	_, _, valid := m.getIndex(x, z)
	if !valid {
		return false
	}
	return true
}

func (m *Map) GetPosition(x, z int32) coretypes.Position {

	posX := float64(x*accuracy) + m.maskMap.StartPos.X
	posZ := m.maskMap.StartPos.Z - float64(z*accuracy)
	return coretypes.Position{
		X: posX,
		Z: posZ,
	}
}

func (m *Map) RandomPosition() coretypes.Position {
	for i := 0; i < 10; i++ {
		standIndex := rand.Intn(len(m.standList))
		index := m.standList[standIndex]
		x := index % int32(m.length)
		z := index / int32(m.length)
		posX := float64(x*accuracy) + m.maskMap.StartPos.X
		posZ := m.maskMap.StartPos.Z - float64(z*accuracy)
		if !m.IsMask(posX, posZ) {
			continue
		}
		return coretypes.Position{
			X: posX,
			Y: m.GetHeight(posX, posZ),
			Z: posZ,
		}
	}
	panic("position error")
}

const (
	accuracy = 1
)

func (m *Map) Accuracy() int {
	return accuracy
}

func (m *Map) getIndex(x float64, z float64) (xIndex int, zIndex int, valid bool) {
	xIndex = int(math.Floor((x - m.maskMap.StartPos.X) / accuracy))
	zIndex = int(math.Floor((m.maskMap.StartPos.Z - z) / accuracy))
	if !m.isIndexValid(xIndex, zIndex) {
		return
	}
	// if zIndex < 0 || xIndex < 0 {
	// 	return
	// }
	// if zIndex >= m.width {
	// 	return
	// }
	// if xIndex >= m.length {
	// 	return
	// }
	valid = true
	return
}

func (m *Map) isIndexValid(xIndex, zIndex int) bool {
	if zIndex < 0 || xIndex < 0 {
		return false
	}
	if zIndex >= m.width {
		return false
	}
	if xIndex >= m.length {
		return false
	}
	return true
}

func (m *Map) GetHeight(x float64, z float64) float64 {
	xIndex, zIndex, valid := m.getIndex(x, z)
	if !valid {
		return 0
	}
	return m.heightMap[zIndex][xIndex]
}

func (m *Map) GetMapMask() [][]int {
	return m.maskMap.Mask
}

func (m *Map) GetStartXZ() (x int, z int) {
	return int(m.maskMap.StartPos.X), int(m.maskMap.StartPos.Z)
}

func (m *Map) IsMask(x float64, z float64) bool {
	xIndex, zIndex, valid := m.getIndex(x, z)
	if !valid {
		return false
	}
	if m.maskMap.Mask[zIndex][xIndex] == 0 {
		return false
	}
	return true
}

func (m *Map) IsWalkable(x float64, z float64) bool {

	xIndex, zIndex, valid := m.getIndex(x, z)
	if !valid {
		return false
	}

	if m.maskMap.Mask[zIndex][xIndex] != 0 {
		return true
	}
	//左边
	leftTileXIndex, leftTileZIndex := xIndex-1, zIndex
	if m.isIndexValid(leftTileXIndex, leftTileZIndex) {
		if m.maskMap.Mask[leftTileZIndex][leftTileXIndex] != 0 {
			return true
		}
	}
	leftUpXIndex, leftUpZIndex := xIndex-1, zIndex+1
	if m.isIndexValid(leftUpXIndex, leftUpZIndex) {
		if m.maskMap.Mask[leftUpZIndex][leftUpXIndex] != 0 {
			return true
		}
	}
	leftDownXIndex, leftDownZIndex := xIndex-1, zIndex-1
	if m.isIndexValid(leftDownXIndex, leftDownZIndex) {
		if m.maskMap.Mask[leftDownZIndex][leftDownXIndex] != 0 {
			return true
		}
	}
	rightXIndex, rightZIndex := xIndex+1, zIndex
	if m.isIndexValid(rightXIndex, rightZIndex) {
		if m.maskMap.Mask[rightZIndex][rightXIndex] != 0 {
			return true
		}
	}
	rightUpXIndex, rightUpZIndex := xIndex+1, zIndex+1
	if m.isIndexValid(rightUpXIndex, rightUpZIndex) {
		if m.maskMap.Mask[rightUpZIndex][rightUpXIndex] != 0 {
			return true
		}
	}
	rightDownXIndex, rightDownZIndex := xIndex+1, zIndex-1
	if m.isIndexValid(rightDownXIndex, rightDownZIndex) {
		if m.maskMap.Mask[rightDownZIndex][rightDownXIndex] != 0 {
			return true
		}
	}
	upXIndex, upZIndex := xIndex, zIndex+1
	if m.isIndexValid(upXIndex, upZIndex) {
		if m.maskMap.Mask[upZIndex][upXIndex] != 0 {
			return true
		}
	}
	downXIndex, downZIndex := xIndex, zIndex-1
	if m.isIndexValid(downXIndex, downZIndex) {
		if m.maskMap.Mask[downZIndex][downXIndex] != 0 {
			return true
		}
	}
	return false

}

func (m *Map) valid() error {
	maskWidth := len(m.maskMap.Mask)
	maskLength := 0
	for i, xMask := range m.maskMap.Mask {
		if maskLength == 0 {
			maskLength = len(xMask)
			continue
		}
		if maskLength != len(xMask) {
			return fmt.Errorf("mask x row %d length expect %d,but get %d", i, maskLength, len(xMask))
		}
	}
	if len(m.heightMap) != maskWidth {
		return fmt.Errorf("height z length expect %d,but get %d", maskWidth, len(m.heightMap))
	}
	for i, xHeight := range m.heightMap {
		if len(xHeight) != maskLength {
			return fmt.Errorf("height x row %d length expect %d,but get %d", i, maskLength, len(xHeight))
		}
	}
	m.width = maskWidth
	m.length = maskLength

	for z, zMask := range m.maskMap.Mask {
		for x, mask := range zMask {
			if mask != 0 {
				m.standList = append(m.standList, int32(z)*int32(maskLength)+int32(x))
			}
		}
	}
	return nil
}

func NewMap(mask *MapMaskMap, height mapHeightMap) (m *Map, err error) {
	m = &Map{
		maskMap:   mask,
		heightMap: height,
	}
	err = m.valid()
	if err != nil {
		return
	}

	return
}

const (
	mapHeight = "_height"
	mapExt    = ".txt"
)

func (ts *templateService) ReadMap(mapFile string) (m *Map, err error) {

	mapMaskFile := filepath.Join(ts.mapDir, mapFile+mapExt)
	bs, err := ioutil.ReadFile(mapMaskFile)
	if err != nil {
		return
	}
	mm := &MapMaskMap{}
	err = json.Unmarshal(bs, mm)
	if err != nil {
		return
	}
	mm = mm.Fixed()
	mapHeightFile := filepath.Join(ts.mapDir, mapFile+mapHeight+mapExt)
	bs, err = ioutil.ReadFile(mapHeightFile)
	if err != nil {
		return
	}

	var mhm mapHeightMap
	err = json.Unmarshal(bs, &mhm)
	if err != nil {
		return
	}
	m, err = NewMap(mm, mhm)
	if err != nil {
		return
	}
	return
}
