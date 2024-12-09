package infra

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		log.Println("db creation initiated")
		connection_str := "host=localhost user=postgres password=1234 port=5432 dbname=glad sslmode=disable" // todo: move to .env
		db, err = gorm.Open(postgres.Open(connection_str), &gorm.Config{})
		log.Println("in function", db)
		if err != nil {
			log.Println("there was an error with the database", err)
			return
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("there was an error configuring the db")
			return
		}
		// set up connection pooling
		sqlDB.SetMaxIdleConns(10)           // max number of connections in idle connection pool
		sqlDB.SetMaxOpenConns(100)          // max number of open connections
		sqlDB.SetConnMaxLifetime(time.Hour) // max amount of time a connection may be reused
	})
	log.Println(db, "is the conenciton state")
	return db, err
}
