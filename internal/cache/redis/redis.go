package cache

import (
	"context"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/hahaclassic/elysium/config"
	"github.com/hahaclassic/elysium/internal/model"
	"github.com/hahaclassic/elysium/pkg/errwrap"
	"github.com/redis/go-redis/v9"
)

type SettingsCache struct {
	conf   *config.RedisConfig
	client *redis.Client
}

func New(ctx context.Context, conf *config.RedisConfig) (*SettingsCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(conf.Host, conf.Port), // адрес вашего Redis-сервера
		Password: conf.Password,                          // пароль, если он установлен
		DB:       conf.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &SettingsCache{
		conf:   conf,
		client: client,
	}, nil
}

func (s *SettingsCache) SetSettings(ctx context.Context, userID int64, settings *model.UserCacheSettings) error {
	strUserID := strconv.FormatInt(userID, 10)

	err := s.client.Set(ctx, strUserID, settings, s.getExpiration()).Err()

	return err
}

func (s *SettingsCache) GetSettings(ctx context.Context, userID int64) (*model.UserCacheSettings, error) {
	strUserID := strconv.FormatInt(userID, 10)

	cmd := s.client.Get(ctx, strUserID)

	if err := cmd.Err(); err != nil {
		return "", err
	}

	return s.client.Result()
}

func (s *SettingsCache) getExpiration() time.Duration {
	return s.conf.Expiration + time.Duration(rand.Int63n(int64(s.conf.Jitter.Seconds())))
}

func (s *SettingsCache) Close() error {
	return errwrap.WrapIfErr(ErrCloseMemcached, rt.client.Close())
}
