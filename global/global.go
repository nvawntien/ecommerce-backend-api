package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"

	"github.com/jmoiron/sqlx"
)

var (
	Config setting.Config
	Logger *logger.Logger
	Pdbx   *sqlx.DB
)
