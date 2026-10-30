// ได้ค่าสั่ง Connect Database
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	username = "root"
	password = ""
	dbname   = "sau_db"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("DB connect successful.....")

	//ทำ Migration Database กับตัว Model ที่สร้างไว้
	if err := db.AutoMigrate(&Student{}); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Migration successful.....")

	return db, err
}
