package postgres

import (
	"fmt"
	"store-dashboard-service/config"
	"store-dashboard-service/util/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dbConfig := config.GetConfig().PostgresDB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.GetLogger().Error("postgres_connection", "error connect to postgres", err)
		panic(err)
	}

	return db
}
