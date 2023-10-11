package domain

type Users struct {
	ID           uint   `json:"id" gorm:"unique;not null"`
	Name         string `json:"name"`
	Email        string `json:"email" validate:"email"`
	Password     string `json:"password" validate:"min=8,max==20"`
	Phone        string `json:"phone"`
	Blocked      bool   `json:"blocked" gorm:"default:false"`
	IsAdmin      bool   `json:"is_admin" gorm:"default:false"`
	ReferralCode string `json:"referral_code"`
}
type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UsersID   uint   `json:"users_id"`
	Users     Users  `json:"-" gorm:"foreignKey:UsersID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"housename" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Default   bool   `json:"default" gorm:"default:false"`
}
