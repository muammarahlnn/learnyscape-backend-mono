package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	Addr            string
	Password        string
	DialTimeout     int
	ReadTimeout     int
	WriteTimeout    int
	MinIdleConn     int
	MaxIdleConn     int
	MaxActiveConn   int
	MaxConnLifetime int
}

func NewRedis(opt *RedisOptions) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:            opt.Addr,
		DialTimeout:     time.Duration(opt.DialTimeout) * time.Second,
		ReadTimeout:     time.Duration(opt.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(opt.WriteTimeout) * time.Second,
		MinIdleConns:    opt.MinIdleConn,
		MaxIdleConns:    opt.MaxIdleConn,
		MaxActiveConns:  opt.MaxActiveConn,
		ConnMaxLifetime: time.Duration(opt.MaxConnLifetime) * time.Minute,
	})

	ctx := context.Background()
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("connected to redis:", status)

	return rdb
}
