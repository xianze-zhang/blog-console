package model

type Model struct {
	ID          uint32 `gorm:"primary_key" json:"id"`
	CreateUser  string `json:"create_user"`
	CreateTime  uint32 `json:"create_time"`
	UpdateUser  string `json:"update_user"`
	UpdateTime  uint32 `json:"update_time"`
	DeletedTime uint32 `json:"delete_time"`
	IsDel       uint8  `json:"is_del"`
}
