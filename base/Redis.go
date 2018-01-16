package base

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"time"
)

type RedisConf struct {
	conn redis.Conn
}



func NewRedisClient() (*RedisConf, error) {
	if SCache.Has("redis") {
		return SCache.Get("redis").(*RedisConf), nil
	}
	c, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		return nil, err
	}
	conf := RedisConf{conn:c}
	SCache.Add("redis", &conf, 0*time.Second)
	return &conf, nil
}


func (this *RedisConf) Get(key string) (interface{}, error) {
	res, err:= this.conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	var dat map[string]interface{}
	s,_ := redis.String(res, err)
	json.Unmarshal([]byte(s), &dat)
	return dat, nil

}

func (this *RedisConf) Del(key string) bool {
	_, err := this.conn.Do("DEL", key)
	if err != nil {
		return false
	}
	return true
}

func (this *RedisConf) Set(key string, val interface{}) error {
	data, _ := json.Marshal(val)
	_, err :=this.conn.Do("SET", key, data)
	return err
}


func (this *RedisConf) SetTimeout(key string, val interface{}, t int64) error {
	data, _ := json.Marshal(val)
	_, err :=this.conn.Do("SET", key, data, "EX", t)
	return err
}

func (this *RedisConf) Exist(key string) bool {
	ok , err :=this.conn.Do("EXISTS", key)
	if err != nil {
		return false
	}
	return ok.(int64) != 0
}

func (this *RedisConf) Close()  {
	this.conn.Close()
	SCache.Remove("redis")
}
