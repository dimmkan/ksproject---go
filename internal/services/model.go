package services

type Service struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Client_name string `json:"client_name"`
	Client_segment string `json:"client_segment"`
}

