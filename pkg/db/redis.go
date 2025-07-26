package db

import "github.com/QGchat/qgim_server_merchants_lib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
