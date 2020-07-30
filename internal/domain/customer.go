package domain

type Customer struct {
	Base
	FirstName   string `json:"first_name" gorm:"column:first_name;type:varchar(255);not null"`
	LastName    string `json:"last_name" gorm:"column:last_name;type:varchar(255);not null"`
	Email       string `json:"email" gorm:"column:email;type:varchar(255);not null"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(15);not null"`
	Password    string `json:"-" gorm:"column:password;type:char(60);not null"`
}
