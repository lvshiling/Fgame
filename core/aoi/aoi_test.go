package aoi_test

import (
	. "fgame/fgame/core/aoi"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

const (
	MIN_X                 = -500
	MAX_X                 = 500
	MIN_Y                 = -500
	MAX_Y                 = 500
	DISTANCE              = 100
	NUM_OBJS              = 4000
	VERIFY_NEIGHBOR_COUNT = true
)

func TestXZListAOIManager(t *testing.T) {
	testAOI(t, "XZListAOI", NewXYListAOIManager(DISTANCE), NUM_OBJS)
}

type TestObj struct {
	aoi            *AOI
	Id             int
	neighbors      map[int64]struct{}
	totalNeighbors int64
	nCalc          int64
}

func (obj *TestObj) OnEnterAOI(otheraoi *AOI) {
	if VERIFY_NEIGHBOR_COUNT {
		// idStr := fmt.Sprintf("%d", obj.Id)
		// if strings.EqualFold(idStr, otheraoi.Id) {
		// 	panic("should not enter self")
		// }
		// _, exist := obj.neighbors[otheraoi.Id]
		// if exist {
		// 	panic("duplicate enter aoi")
		// }
		obj.neighbors[otheraoi.Id] = struct{}{}
		obj.totalNeighbors += int64(len(obj.neighbors))
		obj.nCalc += 1
	}
}

func (obj *TestObj) OnLeaveAOI(otheraoi *AOI) {
	if VERIFY_NEIGHBOR_COUNT {
		// idStr := fmt.Sprintf("%d", obj.Id)
		// if strings.EqualFold(idStr, otheraoi.Id) {
		// 	panic("should not leave self")
		// }
		// _, exist := obj.neighbors[otheraoi.Id]
		// if !exist {
		// 	panic("duplicate leave aoi")
		// }
		delete(obj.neighbors, otheraoi.Id)
		obj.totalNeighbors += int64(len(obj.neighbors))
		obj.nCalc += 1
	}
}

func (obj *TestObj) String() string {
	return fmt.Sprintf("TestObj<%d>", obj.Id)
}

func randCoord(min, max float64) float64 {
	return min + float64(rand.Intn(int(max)-int(min)))
}

func testAOI(t *testing.T, manname string, aoiman AOIManager, numAOI int) {
	objs := []*TestObj{}
	for i := 0; i < numAOI; i++ {
		id := int(i + 1)
		obj := &TestObj{Id: id, neighbors: map[int64]struct{}{}}

		obj.aoi = NewAOI(int64(id), 0, 0, obj)

		objs = append(objs, obj)
		aoiman.Enter(obj.aoi, randCoord(MIN_X, MAX_X), randCoord(MIN_Y, MAX_Y))
	}

	proffd, _ := os.OpenFile(manname+".pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer proffd.Close()

	pprof.StartCPUProfile(proffd)
	for i := 0; i < 1; i++ {
		t0 := time.Now()
		for _, obj := range objs {
			aoiman.Move(obj.aoi, obj.aoi.X+randCoord(-10, 10), obj.aoi.Y+randCoord(-10, 10))
		}
		dt := time.Now().Sub(t0)
		t.Logf("%s tick %d objects takes %s", manname, numAOI, dt)
	}

	for _, obj := range objs {
		aoiman.Leave(obj.aoi)
	}

	pprof.StopCPUProfile()

	if VERIFY_NEIGHBOR_COUNT {
		totalCalc := int64(0)
		for _, obj := range objs {
			totalCalc += obj.nCalc
		}
		println("Average calculate count:", totalCalc/int64(len(objs)))
	}
}
