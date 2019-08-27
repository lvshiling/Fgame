package redis_test

import (
	"sync"
	"testing"
	"time"

	. "qipai/redis"
)

const (
	testKey = "test"
)

var (
	rs RedisService
)

var (
	once sync.Once
)

//只会设置1次
func setup(t *testing.T) {

	once.Do(func() {
		rc := &RedisConfig{
			Host:        "127.0.0.1",
			Port:        6379,
			MaxIdle:     200,
			MaxActive:   200,
			Wait:        true,
			IdleTimeout: 200 * int64(time.Second),
		}
		trs, err := NewRedisServiceWithConfig(rc)
		if err != nil {
			t.Errorf("setup error %s", err.Error())
			t.FailNow()
		}
		rs = trs
	})

}

//测试加锁
func TestLock(t *testing.T) {
	setup(t)

	c := make(chan bool)
	for i := 0; i < 1; i++ {
		go HammerMutex(t, 1, c)
	}

	for i := 0; i < 1; i++ {
		<-c
	}
}

func HammerMutex(t *testing.T, loops int, cdone chan bool) {

	conn := rs.Pool().Get()
	if conn.Err() != nil {
		t.Errorf("get redis conn error %s", conn.Err().Error())
		t.FailNow()
	}

	for i := 0; i < loops; i++ {
		_, err := Lock(conn, testKey, 0)
		if err != nil {
			t.Errorf("lock error %s", err.Error())
			t.FailNow()
		}

		Unlock(conn, testKey)
	}
	cdone <- true
}

func BenchmarkMutexUncontended(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		conn := rs.Pool().Get()
		if conn.Err() != nil {
			b.Errorf("get redis conn error %s", conn.Err().Error())
			b.FailNow()
		}

		for pb.Next() {
		 err := Lock(conn, testKey, 0)
			if err != nil {
				b.Errorf("lock error %s", err.Error())
				b.FailNow()
			}

			Unlock(conn, testKey)
		}
	})
}

func benchmarkMutex(b *testing.B, slack, work bool) {

	if slack {
		b.SetParallelism(10)
	}
	b.RunParallel(func(pb *testing.PB) {

		conn := rs.Pool().Get()
		if conn.Err() != nil {
			b.Errorf("get redis conn error %s", conn.Err().Error())
			b.FailNow()
		}
		foo := 0
		for pb.Next() {
			_, err := Lock(conn, testKey, 0)
			if err != nil {
				b.Errorf("lock error %s", err.Error())
				b.FailNow()
			}
			Unlock(conn, testKey)
			if work {
				for i := 0; i < 100; i++ {
					foo *= 2
					foo /= 2
				}
			}
		}
		_ = foo

	})
}

func BenchmarkMutex(b *testing.B) {
	benchmarkMutex(b, false, false)
}

func BenchmarkMutexSlack(b *testing.B) {
	benchmarkMutex(b, true, false)
}

func BenchmarkMutexWork(b *testing.B) {
	benchmarkMutex(b, false, true)
}

func BenchmarkMutexWorkSlack(b *testing.B) {
	benchmarkMutex(b, true, true)
}

func BenchmarkMutexNoSpin(b *testing.B) {
	// This benchmark models a situation where spinning in the mutex should be
	// non-profitable and allows to confirm that spinning does not do harm.
	// To achieve this we create excess of goroutines most of which do local work.
	// These goroutines yield during local work, so that switching from
	// a blocked goroutine to other goroutines is profitable.
	// As a matter of fact, this benchmark still triggers some spinning in the mutex.
	var acc0, acc1 uint64
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		c := make(chan bool)
		var data [4 << 10]uint64
		for i := 0; pb.Next(); i++ {
			if i%4 == 0 {
				m.Lock()
				acc0 -= 100
				acc1 += 100
				m.Unlock()
			} else {
				for i := 0; i < len(data); i += 4 {
					data[i]++
				}
				// Elaborate way to say runtime.Gosched
				// that does not put the goroutine onto global runq.
				go func() {
					c <- true
				}()
				<-c
			}
		}
	})
}

func BenchmarkMutexSpin(b *testing.B) {
	// This benchmark models a situation where spinning in the mutex should be
	// profitable. To achieve this we create a goroutine per-proc.
	// These goroutines access considerable amount of local data so that
	// unnecessary rescheduling is penalized by cache misses.
	var m Mutex
	var acc0, acc1 uint64
	b.RunParallel(func(pb *testing.PB) {
		var data [16 << 10]uint64
		for i := 0; pb.Next(); i++ {
			m.Lock()
			acc0 -= 100
			acc1 += 100
			m.Unlock()
			for i := 0; i < len(data); i += 4 {
				data[i]++
			}
		}
	})
}

func unlock(key string) {

}

func TestUnlock(t *testing.T) {
	setup(t)
}
