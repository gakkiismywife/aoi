package models

import (
	"AiServer/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func Migrate() {
	_ = utils.DB.AutoMigrate(&Data{})
	_ = utils.DB.AutoMigrate(&DataFiles{})
	_ = utils.DB.AutoMigrate(&DataTag{})
	_ = utils.DB.AutoMigrate(&DataPermission{})
}
