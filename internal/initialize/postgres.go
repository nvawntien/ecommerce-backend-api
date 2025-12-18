package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func InitPostgres() {
	m := global.Config.PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host,
		m.Port,
		m.Username,
		m.Password,
		m.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		global.Logger.Error("Init postgreSQL error", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Init postgreSQL successfully")

	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

	global.Pdbx = db
}
