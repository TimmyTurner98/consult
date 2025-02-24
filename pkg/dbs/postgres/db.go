package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var config *Config

// InitDBConfig загружает конфигурацию базы данных
func InitDBConfig() error {
	var err error
	config, err = InitConfigs() // Вызываем InitConfigs из config.go
	if err != nil {
		return fmt.Errorf("❌ Ошибка загрузки конфигурации: %w", err)
	}
	return nil
}

// ConnectDB подключается к PostgreSQL через GORM
func ConnectDB() {
	if config == nil {
		if err := InitDBConfig(); err != nil {
			log.Fatalf("%v", err)
		}
	}
	// 🔍 Выводим текущие настройки подключения
	fmt.Printf("🔍 Подключение к PostgreSQL: host=%s port=%s user=%s dbname=%s\n",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.DBname,
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,     // ✅ Теперь порт на месте
		config.Database.Username, // ✅ Теперь username на месте
		config.Database.Password,
		config.Database.DBname,
		config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Ошибка получения *sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("✅ Успешное подключение к PostgreSQL через GORM!")

	DB = db
}

// DBConfig возвращает текущую конфигурацию базы данных
func DBConfig() *Config {
	if config == nil {
		if err := InitDBConfig(); err != nil {
			log.Fatalf("%v", err)
		}
	}
	return config
}
