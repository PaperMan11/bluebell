package mysql

import (
	"bluebell/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 驱动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns) // 设置最大连接数
	db.SetMaxIdleConns(cfg.MaxIdleConns) // 设置最大存活数
	return
}

func Close() {
	db.Close()
}
