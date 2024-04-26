package models

type DataTag struct {
	Id     int64  `json:"id" gorm:"primaryKey"`
	Uuid   int64  `json:"uuid" gorm:"uniqueIndex:idx_uuid"`
	Name   string `json:"name"`
	DataId int64  `json:"data_id"`
	BaseModel
}
