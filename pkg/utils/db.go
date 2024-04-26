package utils

import (
	"AiServer/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var DB *gorm.DB

var once sync.Once

func InitDbEngine(c *configs.Database) error {
	var err error
	once.Do(func() {
		s := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
		dsn := fmt.Sprintf(s, c.Host, c.Username, c.Password, c.Database, c.Port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	return err
}
