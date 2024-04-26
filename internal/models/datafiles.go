package models

type DataFiles struct {
	Id     int64  `json:"id" gorm:"primaryKey"`
	Uuid   int64  `json:"uuid"`
	DataId int64  `json:"data_id"`
	Key    string `json:"key"`
	TagId  int64  `json:"tag_id"`
	Label  string `json:"label"`
	BaseModel
}
