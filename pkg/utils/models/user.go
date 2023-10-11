package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassWord string `json:"confirmpassword"`
	ReferralCode    string `json:"referral_code"`
}

type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

// user details along with embedded token which can be used by the user to access protectrd routed
type TokenUsers struct {
	Users UserDeatilsResponse
	Token string
}

// user details shown after logging in
type UserDeatilsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserLoign struct {
	Email    string `json:"email" validate:"email"`
	PassWord string `json:"password"`
}

type UserSignInResponse struct {
	Id       uint   `json:"id"`
	UserId   string `json:"userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserDetailsAdmin struct {
	Id         int    `json:"id"`
	Name       string `josn:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	BlockUsers string `json:"blockusers"`
}
type AddAddress struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type ChangePassword struct {
	Oldpassword string `json:"old_password"`
	Password    string `json:"password"`
	Repassword  string `json:"re_password"`
}

type ForgotPassword struct {
	Phone string `json:"phone"`
}

type ForgotPasswordSend struct {
	Phone string `json:"phone"`
}

type ForgotVarify struct {
	Phone       string `json:"phone"`
	Otp         string `json:"otp"`
	NewPassword string `json:"newpassword"`
}

type EditName struct {
	Name string `json:"name"`
}

type EditEmail struct {
	Email string `json:"email"`
}

type Editphones struct {
	Phone string `json:"phone"`
}

type GetCart struct {
	ProductName     string  `json:"product_name"`
	Category_id     int     `json:"category_id"`
	Quantity        int     `json:"quantity"`
	Totel           float64 `json:"totel_price"`
	DiscountedPrice float64 `json:"discounted_price"`
}
type CheckOut struct {
	Addresses      []Address
	Products       []GetCart
	PaymentMethord []PaymentMethod
	TotelPrice     float64
}
type Search struct {
	Key string `json:"searchkey"`
}
