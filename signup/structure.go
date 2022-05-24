package signup

import "github.com/gocql/gocql"

type Otp struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
}

type CreateAccount struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	OtpNumber   string `json:"otp"`
}

type Passwd struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
}
type Result struct {
	Status      bool       `json:"content"`
	Message     string     `json:"message"`
	UserId      gocql.UUID `josn:"userid"`
	Usermail    string     `json:"usrmail"`
	UserMobile  string     `json:"usermobile"`
	CountryCode string     `json:"countrycode"`
}
