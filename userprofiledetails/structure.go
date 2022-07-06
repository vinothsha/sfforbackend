package userprofiledetails

import "github.com/gocql/gocql"

type UserProfile struct {
	ProfileUid   gocql.UUID `json:"profileuid"`
	CountryCode  string     `json:"countrycode"`
	Mobile       string     `json:"mobile"`
	DateOfBirth  string     `json:"dateofbirth"`
	Email        string     `json:"email"`
	FirstName    string     `json:"firstname"`
	LastName     string     `json:"lastname"`
	Gender       string     `json:"gender"`
	Profileimage string     `json:"profileimage"`
	Useruid      string     `json:"useruid"`
}
