package models

type DataPermission struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	UserId     int64  `json:"user_id"`
	DataId     int64  `json:"data_id"`
	Permission string `json:"permission"`
	BaseModel
}
