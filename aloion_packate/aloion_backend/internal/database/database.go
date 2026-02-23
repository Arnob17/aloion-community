package database

import (
	"fmt"

	"github.com/arnob17/aloion_backend/internal/config"
	"github.com/arnob17/aloion_backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, // Disable FK constraints during migration
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Payment{},
		&models.Subscription{},
		&models.Enrollment{},
		&models.CourseMaterial{},
		&models.Assignment{},
		&models.Submission{},
		&models.Exam{},
		&models.ExamResult{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
