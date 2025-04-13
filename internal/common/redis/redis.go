package redis

import (
	"context"
	"sync"

	"github.com/neel4os/warg/internal/common/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	instance *RedisCon
	once     sync.Once
)

type RedisCon struct {
	client redis.UniversalClient
}

func newDataConn(cfg *config.Config) *RedisCon {
	log.Debug().Caller().Msg("Creating new redis connection")
	redisConfig := cfg.RedisConfig
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{redisConfig.Host + ":" + redisConfig.Port},
		DB:       redisConfig.Db,
		Password: redisConfig.Password,
	})
	return &RedisCon{
		client: rdb,
	}
}

func GetRedisCon(cfg *config.Config) *RedisCon {
	once.Do(func() {
		instance = newDataConn(cfg)
	})
	return instance
}

func (d *RedisCon) Ping() {
	_, err := d.client.Ping(context.Background()).Result()
	if err != nil {
		log.Error().Caller().Err(err).Msg("Error pinging redis")
		panic(err)
	}
}

func (d *RedisCon) Close() error {
	// close the redis connection
	err := d.client.Close()
	if err != nil {
		log.Error().Caller().Err(err).Msg("Error closing redis connection")
		return err
	}
	return nil
}

func (d *RedisCon) GetClient() redis.UniversalClient {
	return d.client
}
