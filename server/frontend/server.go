package frontend

import (
	"context"
	"demo1/service/service_config"
	"demo1/service/service_file"
	"demo1/service/service_swa"
	"demo1/utils"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	weaver.Implements[weaver.Main]

	swaConfig  service_config.ModSwaConfig
	signingKey []byte

	frontCache local_cache.Cache
	frontRedis *redis.Client

	serviceConfig weaver.Ref[service_config.T]
	serviceSwa    weaver.Ref[service_swa.T]
	serviceFile   weaver.Ref[service_file.T]

	address weaver.Listener
}

func Serve(ctx context.Context, s *server) error {
	s.Logger(ctx).Debug("start swa server...", "address", s.address)
	Router := InitRouters(ctx, s)
	return http.Serve(s.address, Router)
}

func (s *server) Init(ctx context.Context) (err error) {
	s.swaConfig, err = s.serviceConfig.Get().GetCfgSwaConfig(ctx)
	if err != nil {
		fmt.Printf("系统配置失败:%+v\n", err)
	}
	s.signingKey = []byte(s.swaConfig.Jwt.SigningKey)

	s.InitCache(ctx)
	s.InitRedis(ctx)

	s.Logger(ctx).Debug("swa server init ...")

	return err
}

func (s *server) InitCache(_ context.Context) {
	dr, err := utils.ParseDuration(s.swaConfig.Jwt.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(s.swaConfig.Jwt.BufferTime)
	if err != nil {
		panic(err)
	}
	s.frontCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}

func (s *server) InitRedis(ctx context.Context) {
	if s.swaConfig.System.UseMultipoint || s.swaConfig.System.UseRedis {
		redisCfg := s.swaConfig.Redis
		client := redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			s.Logger(ctx).Error("redis连接失败, err:", zap.Error(err))
			s.frontRedis = nil
			panic(err)
			return
		} else {
			s.Logger(ctx).Info("redis 连接成功:", zap.String("pong", pong))
			s.frontRedis = client
		}
	} else {
		s.frontRedis = nil
	}
}
