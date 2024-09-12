package database

import (
	"fmt"
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	return newLogger
}

func StartDB(conf config.DBConfig) (*gorm.DB, error) {

	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", conf.Host, conf.Port, conf.User, conf.Name,
		conf.Sslmode, conf.Password)

	conn, err := gorm.Open(postgres.Open(str), &gorm.Config{
		//Logger: GetLogger(),
	})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
		return nil, err
	}

	db := conn
	config, err := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatal("Error: ", err)
		return nil, err
	}
	return db, nil
}

func CloseConn(db *gorm.DB) error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}
