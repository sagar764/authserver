package models

type Login struct {
	CustomerId    string `json:"customerid"`
	PasswordOrOTP string `json:"password_or_otp"`
}

type Signup struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
