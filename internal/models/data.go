package models

type Data struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Uuid      int64  `json:"uuid" gorm:"uniqueIndex:idx_uuid"` //uuid
	Name      string `json:"name"`                             //数据集名称
	Type      int    `json:"type"`                             //标注类型
	DataCount int    `json:"data_count"`                       //数据量
	BaseModel
}
