package model

type Employee struct {
	BaseModel
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	UserID  string
	User    User `gorm:"foreignKey:UserID"`
}
