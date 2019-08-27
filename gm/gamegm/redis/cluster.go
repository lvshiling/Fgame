package redis

import (
	"math/rand"
	"time"

	"github.com/chasex/redis-go-cluster"
)

type RedisClusterOptions struct {
	StartNodes []string `json:"startNodes"`

	ConnTimeout  int64 `json:"connTimeout"`
	ReadTimeout  int64 `json:"readTimeout"`
	WriteTimeout int64 `json:"writeTimeout"`

	KeepAlive int   `json:"keepAlive"`
	AliveTime int64 `json:"aliveTime"`
}

type RedisClusterService interface {
	Cluster() *redis.Cluster
}

type redisClusterService struct {
	cluster *redis.Cluster
}

func (rcs *redisClusterService) Cluster() *redis.Cluster {
	return rcs.cluster
}

func NewRedisClusterService(rco *RedisClusterOptions) (rs RedisClusterService, err error) {
	trs := &redisClusterService{}
	opt := &redis.Options{
		StartNodes:   rco.StartNodes,
		ConnTimeout:  time.Duration(rco.ConnTimeout * int64(time.Millisecond)),
		ReadTimeout:  time.Duration(rco.ReadTimeout * int64(time.Millisecond)),
		WriteTimeout: time.Duration(rco.WriteTimeout * int64(time.Millisecond)),
		KeepAlive:    rco.KeepAlive,
		AliveTime:    time.Duration(rco.AliveTime * int64(time.Millisecond)),
	}
	c, err := redis.NewCluster(opt)
	if err != nil {
		return
	}
	trs.cluster = &c

	rs = trs
	return
}

func ClusterLockDefault(conn *redis.Cluster, key string) (bool, error) {
	return ClusterLock(conn, key, defaultTimeout, defaulRetryTimes)
}

//加锁 毫毛为单位
func ClusterLock(conn *redis.Cluster, key string, timeout int64, retryTimes int) (bool, error) {
	//不允许不设置过期时间，否则将会死锁
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	//拼接lock key
	lockKeyStr := Join(lockKey, key)
	//循环拿锁
	for i := 0; i < retryTimes; i++ {

		now := time.Now().UnixNano() / int64(time.Millisecond)
		timeout = now + timeout + 1
		setnxrst, _ := (*conn).Do("setnx", lockKeyStr, timeout)
		val, err := redis.Int(setnxrst, nil)

		//redis错误
		if err != nil {
			return false, err
		}

		//获得锁了
		if val == 1 {
			return true, nil
		}

		//获取过期时间戳
		getValue, err := redis.Int64((*conn).Do("get", lockKeyStr))
		//redis 错误
		if err != nil {
			//重试
			if err == redis.ErrNil {
				continue
			}
			return false, err
		}

		//过期了
		if getValue < now {

			//原子操作取值设置新值
			getOldValue, err := redis.Int64((*conn).Do("getset", lockKeyStr, now))
			//redis 错误
			if err != nil {
				//重试
				if err == redis.ErrNil {
					continue
				}

				return false, err
			}

			// 拿到锁了
			if getOldValue == getValue {
				return true, nil
			}
			//被人抢先了
		}

		//睡眠
		rand.Seed(time.Now().UnixNano())
		sleepInterval := rand.Intn(sleepMaxInterval-sleepMinInterval) + sleepMinInterval
		time.Sleep(time.Duration(sleepInterval) * time.Millisecond)
	}
	return false, nil
}

//解锁
func ClusterUnlock(conn *redis.Cluster, key string) (bool, error) {
	lockKeyStr := Join(lockKey, key)

	now := time.Now().Unix() / int64(time.Millisecond)
	getValue, err := redis.Int64((*conn).Do("get", lockKeyStr))
	if err != nil {
		return false, err
	}

	//已经不是自己的锁了
	if now > getValue {
		return true, nil
	}

	//删除key
	_, err = redis.Int((*conn).Do("del", lockKeyStr))
	if err != nil {
		return false, err
	}

	return true, nil
}
