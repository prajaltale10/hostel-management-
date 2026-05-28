package db
import (
	"fmt"
	"hostel-saas/config"
	"hostel-saas/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func InitPostgres(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		return err
	}
	DB = db
	logger.Log.Info("Connected to PostgreSQL database successfully")
	return nil
}
