package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func CreateConnection(env Env) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		username,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		// Logger: logger.GetGormLogger(),
	})

	if err != nil {
		panic(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host,
			username,
			password,
			dbname,
			port,
		))
		// fmt.Print("\nCould not conenct to the DB\n")

		// logger.Info("Url: ", url)
		// logger.Panic(err)
	}

	fmt.Print("Successful connect ")

	// logger.Info("Database connection established")

	return Database{
		DB: db,
	}
}
