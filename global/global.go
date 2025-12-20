package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.Logger
	Pdbx   *sqlx.DB
	Rdb    *redis.Client
)
