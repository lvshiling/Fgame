package idutil

import (
	"fmt"
	"sync"
	"time"
)

const (
	timeBitLen     = 41
	maxTime        = 1<<timeBitLen - 1
	sequenceBitLen = 7
	maxSequence    = 1<<sequenceBitLen - 1
	workIdBitLen   = 63 - timeBitLen - sequenceBitLen
	maxWorkId      = 1<<workIdBitLen - 1
	cEpoch         = 1483200001000 //2017/1/1 00:00:01

	timeShift   = sequenceBitLen
	workIdShift = sequenceBitLen + timeBitLen
)

type Snowflake interface {
	NextId() int64
}

type snowflake struct {
	m                    sync.Mutex
	workerId             int64
	sequence             int64 //当前序列
	currentStartSequence int64 //当前起始序列
	lastTimeStamp        int64
}

func (f *snowflake) NextId() int64 {
	f.m.Lock()
	defer f.m.Unlock()
Again:
	ts := time.Now().UnixNano() / int64(time.Millisecond)
	if ts == f.lastTimeStamp {
		f.sequence = (f.sequence + 1) & maxSequence
		if f.sequence == f.currentStartSequence {
			//睡眠一毫秒
			time.Sleep(time.Millisecond)
			goto Again
		}
	} else {
		f.currentStartSequence = f.sequence
		f.lastTimeStamp = ts
	}

	id := (ts-cEpoch)<<timeShift | f.workerId<<workIdShift | f.sequence
	return id
}

func NewSnowFlake(workerId int64) (Snowflake, error) {
	f := &snowflake{}
	if workerId < 0 {
		return nil, fmt.Errorf("snowflake:workId[%d] 不能小于0", workerId)
	}
	if workerId > maxWorkId {
		return nil, fmt.Errorf("snowflake:workId[%d] 不能超过 [%d]", workerId, maxWorkId)
	}
	f.workerId = workerId
	return f, nil
}

func Parse(id int64) (t time.Time, workerId int64, seq int64) {
	seq = id & maxSequence
	workerId = (id >> workIdShift) & maxWorkId
	elapse := (id >> timeShift) & maxTime

	ts := elapse + cEpoch
	t = time.Unix(ts/1000, (ts%1000)*int64(time.Millisecond))
	return
}

var (
	once sync.Once
	s    Snowflake
)

func SetupWorker(id int64) error {
	var err error
	once.Do(func() {
		ts, terr := NewSnowFlake(id)
		if terr != nil {
			err = terr
			return
		}
		s = ts
	})
	return err
}
