package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	BaseModel
}
