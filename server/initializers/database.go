package initializers

import (
	"context"
	"fmt"
	"igclone/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Custom SQL Logger
type sqlLogger struct {
	logger.Interface
}

// Implement the Trace method for the custom logger
func (s sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)
	if err != nil {
		fmt.Printf("SQL Error: %v | Query: %s | Rows: %d | Time: %s\n", err, sql, rows, elapsed)
	} else {
		fmt.Printf("Executed SQL: %s | Rows: %d | Time: %s\n", sql, rows, elapsed)
	}
}

func ConnectToDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("Connecting to database with DSN:", dsn)

	// Open the database connection with custom logger
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: sqlLogger{logger.Default.LogMode(logger.Info)},
	})

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserProfile{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established!")
}
