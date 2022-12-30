package infra

import (
	"backend/constant"
	"backend/util"
	"fmt"
	"os"
	"time"

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
	dbHost := os.Getenv(constant.DBHostEnv)
	dbPort := os.Getenv(constant.DBPortEnv)
	dBName := os.Getenv(constant.DBNameEnv)
	dbUser := os.Getenv(constant.DBUserEnv)
	dbPassword := os.Getenv(constant.DBPasswordEnv)

	dbAuth := dbUser
	if len(dbPassword) > 0 {
		dbAuth += ":" + dbPassword
	}
	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", dbAuth, dbHost, dbPort, dBName)

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

	// migrateErr := db.AutoMigrate(&model.Article{}, &model.Tag{})

	// if migrateErr != nil {
	// 	return nil, migrateErr
	// }
	return db, nil
}
