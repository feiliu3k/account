package rediscache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/hatlonely/account/internal/mysqldb"
	"time"
)

type Option struct {
	Address    string
	Timeout    time.Duration
	Retries    int
	PoolSize   int
	Password   string
	DB         int
	Expiration time.Duration
}

type RedisCache struct {
	client *redis.Client
	option *Option
}

//func NewRedisCache(addr string, timeout time.Duration, retries int, poolSize int, password string, db int) (*RedisCache, error) {
func NewRedisCache(option *Option) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         option.Address,
		DialTimeout:  option.Timeout,
		ReadTimeout:  option.Timeout,
		WriteTimeout: option.Timeout,
		MaxRetries:   option.Retries,
		PoolSize:     option.PoolSize,
		Password:     option.Password,
		DB:           option.DB,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
		option: option,
	}, nil
}

func (rc *RedisCache) SetAccount(token string, account *mysqldb.Account) error {
	buf, err := json.Marshal(account)
	if err != nil {
		return err
	}

	return rc.client.Set(token, buf, rc.option.Expiration).Err()
}

func (rc *RedisCache) DelAccount(token string) error {
	return rc.client.Del(token).Err()
}

func (rc *RedisCache) GetAccount(token string) (*mysqldb.Account, error) {
	buf, err := rc.client.Get(token).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	account := &mysqldb.Account{}
	if err := json.Unmarshal([]byte(buf), account); err != nil {
		return nil, err
	}

	return account, nil
}
