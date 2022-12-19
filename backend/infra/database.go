package infra

import (
	"backend/constant"
	"backend/util"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger = util.NewLogger()

type GormHandler struct {
	DB *gorm.DB
}

func NewGormHandler() *GormHandler {
	conn, err := connectDB()
	if err != nil {
		logger.Fatal(err)
	}
	GormHandler := new(GormHandler)
	GormHandler.DB = conn
	return GormHandler
}

func connectDB() (*gorm.DB, error) {
	dbName := viper.GetString(constant.DBNameEnv)
	dbUser := viper.GetString(constant.DBUserEnv)
	dbHost := viper.GetString(constant.DBHostEnv)

	dsn := fmt.Sprintf("%v@tcp(%s)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", dbUser, dbHost, dbName)

	var db *gorm.DB
	var err error
	for i := 1; i <= constant.DBConnectRetryMaxCount; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(constant.RetryWaitTime * time.Millisecond)
	}

	if err != nil {
		return nil, err
	}

	migrateErr := db.AutoMigrate()

	if migrateErr != nil {
		return nil, migrateErr
	}
	return db, nil
}
