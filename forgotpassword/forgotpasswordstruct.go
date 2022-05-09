package forgotpassword

type PasswordReset struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
}
type OtpVerify struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	Otp         string `json:"otp"`
}
type NewPassword struct {
	UniqueId    string `json:"uniqueid"`
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
}
