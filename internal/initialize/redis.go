package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		global.Logger.Error("Failed to connect to Redis:", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Connected to Redis successfully")
	global.Rdb = rdb
}
