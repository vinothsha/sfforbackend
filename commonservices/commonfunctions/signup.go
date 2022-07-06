package commonfunctions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/dongri/phonenumber"
)

func ValidateEmail(email string) bool {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
func ValidateMobile(mobile string) bool {
	number := phonenumber.Parse(mobile, "")
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)
	?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(number)
}

func RandomGenerater() string {
	p, _ := rand.Prime(rand.Reader, 18)
	fmt.Println(p)
	return p.String()

}
func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
