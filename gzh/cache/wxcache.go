package cache

import (
	"encoding/json"
	fredis "fgame/fgame/core/redis"
	"time"

	"github.com/garyburd/redigo/redis"
)

type MyRedisCache struct {
	_redis fredis.RedisService
}

//Get 获取一个值
func (r *MyRedisCache) Get(key string) interface{} {
	conn := r._redis.Pool().Get()
	defer conn.Close()

	var data []byte
	var err error
	if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}

	return reply
}

//Set 设置一个值
func (r *MyRedisCache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r._redis.Pool().Get()
	defer conn.Close()

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}

	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)

	return
}

//IsExist 判断key是否存在
func (r *MyRedisCache) IsExist(key string) bool {
	conn := r._redis.Pool().Get()
	defer conn.Close()

	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

//Delete 删除
func (r *MyRedisCache) Delete(key string) error {
	conn := r._redis.Pool().Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}

func NewMyRedisCache(p_redis fredis.RedisService) *MyRedisCache {
	rst := &MyRedisCache{
		_redis: p_redis,
	}
	return rst
}
