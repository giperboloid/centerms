// Package store provides means for data storage and retrieving.
package store

import (
	"time"

	"github.com/kostiamol/centerms/cfg"

	"github.com/garyburd/redigo/redis"

	"fmt"
)

const (
	partialDevKey       = "device:"
	partialDevCfgKey    = ":cfg"
	partialDevParamsKey = ":params"
)

type (
	// Redis is used to provide a storage based on redis db under the hood.
	Redis struct {
		addr cfg.Addr
		pool *redis.Pool
	}

	// Cfg is used to initialize an instance of Redis.
	Cfg struct {
		Addr             cfg.Addr
		Password         string
		MaxIdlePoolConns uint64
		IdleTimeout      time.Duration
	}
)

// New creates a new instance of Redis store.
func New(c *Cfg) (*Redis, error) {
	r := &Redis{
		addr: c.Addr,
		pool: &redis.Pool{
			MaxIdle:     int(c.MaxIdlePoolConns),
			IdleTimeout: c.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", c.Addr.Host+":"+fmt.Sprint(c.Addr.Port))
				if err != nil {
					return nil, fmt.Errorf("Dial(): %s", err)
				}
				//if _, err := c.Do("AUTH", password); err != nil {
				//	c.Close()
				//	return nil, err
				//}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if _, err := c.Do("PING"); err != nil {
					return fmt.Errorf("PING: %s", err)
				}
				return nil
			},
		},
	}

	if _, err := r.ping(); err != nil {
		return nil, fmt.Errorf("store: PING() failed: %s", err)
	}
	return r, nil
}

// Close releases the resources used by the pool.
func (r *Redis) Close() error {
	return r.pool.Close()
}

func (r *Redis) ping() (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("PING"))
}

func (r *Redis) exists(key string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Bool(redis.Bytes(conn.Do("EXISTS", key)))
}

func (r *Redis) multi() (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(redis.Bytes(conn.Do("MULTI")))
}

func (r *Redis) discard() (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(redis.Bytes(conn.Do("DISCARD")))
}

func (r *Redis) exec() ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(redis.Bytes(conn.Do("EXEC")))
}

func (r *Redis) sadd(key string, members ...interface{}) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(redis.Bytes(conn.Do("SADD", key, members)))
}

func (r *Redis) zadd(key string, args ...interface{}) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(redis.Bytes(conn.Do("ZADD", key, args)))
}

func (r *Redis) hget(key, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(redis.Bytes(conn.Do("HGET", key, field)))
}

func (r *Redis) hmset(key string, fields ...interface{}) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(redis.Bytes(conn.Do("HMSET", key, fields)))
}

func (r *Redis) hmget(key string, fields ...interface{}) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(redis.Bytes(conn.Do("HMGET", key, fields)))
}

func (r *Redis) smembers(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("SMEMBERS", key))
}

func (r *Redis) zrangebyscore(key string, min, max interface{}) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max))
}

func (r *Redis) publish(msg interface{}, channel string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(redis.Bytes(conn.Do("PUBLISH", msg, channel)))
}