package gorm

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go-cloud-native/config"
)

// New ...
func New(conf *config.Conf) (*gorm.DB, error) {
	cfg := &mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", conf.Db.Host, conf.Db.Port),
		DBName:               conf.Db.DbName,
		User:                 conf.Db.Username,
		Passwd:               conf.Db.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return gorm.Open("mysql", cfg.FormatDSN())
}
