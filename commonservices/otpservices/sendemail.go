package otpservices

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOtpToEmail(email string, num string) {
	from := mail.NewEmail("StoryFlics", "prasanth.v@outstager.com")
	subject := "OTP from StoryFlics"
	to := mail.NewEmail("Example User", email)
	plainTextContent := "OTP" + num
	htmlContent := "OTP " + num
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} //else {
	// 	fmt.Println(response.StatusCode)
	// 	fmt.Println(response.Body)
	// 	fmt.Println(response.Headers)
	// }
}
