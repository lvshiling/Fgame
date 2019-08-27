package idutil_test

import (
	. "fgame/fgame/pkg/idutil"
	"fmt"
	"runtime"
	"testing"
	"time"
)

var sf Snowflake
var workerId = int64(0)

func init() {
	tsf, err := NewSnowFlake(workerId)
	if err != nil {
		panic(fmt.Errorf("idutil: init test error[%s]", err.Error()))
	}
	sf = tsf
}

func nextId(t *testing.T) int64 {
	id := sf.NextId()
	return id
}

func TestSnowflakeOnce(t *testing.T) {
	sleepTime := uint64(50)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	id := nextId(t)
	ti, getWorkId, seq := Parse(id)

	if seq != 0 {
		t.Errorf("unexpected sequence: %d", seq)
	}

	if getWorkId != workerId {
		t.Errorf("unexpected machine id: %d", workerId)
	}
	t.Logf("time:%s\n", ti.String())
}

func currentTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TestSnowflakeForSec(t *testing.T) {
	var numID int32
	var lastID int64
	var maxSequence int64

	initial := currentTime()
	current := initial
	for current-initial < 10000 {
		id := nextId(t)
		_, getWorkerId, seq := Parse(id)
		numID++

		if id <= lastID {
			t.Fatal("duplicated id")
		}
		lastID = id

		current = currentTime()

		if maxSequence < seq {
			maxSequence = seq
		}

		if getWorkerId != workerId {
			t.Errorf("unexpected worker id: %d", getWorkerId)
		}
	}

	// if maxSequence != 1<<BitLenSequence-1 {
	// 	t.Errorf("unexpected max sequence: %d", maxSequence)
	// }
	t.Logf("max sequence:%d\n", maxSequence)
	t.Logf("number of id:%d\n", numID)
}

func TestSnowflakeInParallel(t *testing.T) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("number of cpu:", numCPU)

	consumer := make(chan int64)

	const numID = 10000
	generate := func() {
		for i := 0; i < numID; i++ {
			consumer <- nextId(t)
		}
	}

	const numGenerator = 10
	for i := 0; i < numGenerator; i++ {
		go generate()
	}

	idMap := make(map[int64]struct{})
	for i := 0; i < numID*numGenerator; i++ {
		id := <-consumer
		_, exist := idMap[id]
		if exist {
			t.Fatal("duplicated id")
		} else {
			idMap[id] = struct{}{}
		}
	}
	fmt.Println("number of id:", len(idMap))
}

// func pseudoSleep(period time.Duration) {
// 	sf.startTime -= int64(period) / sonyflakeTimeUnit
// }

// func TestNextIDError(t *testing.T) {
// 	year := time.Duration(365*24) * time.Hour
// 	pseudoSleep(time.Duration(174) * year)
// 	nextID(t)

// 	pseudoSleep(time.Duration(1) * year)
// 	_, err := sf.NextID()
// 	if err == nil {
// 		t.Errorf("time is not over")
// 	}
// }
