package otpservices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendOtpToMobile(mob string, num string) {
	accountSid := "ACac40d86f1e4383335d6e208ffe96c130"
	authToken := "0713eb9d2f5077cb7efadbd9e7fcc052"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// rand.Seed(time.Now().Unix())
	msgData := url.Values{}
	msgData.Set("To", mob) //vicky--9629381169 hussain--9094501317
	msgData.Set("From", "+16592013522")
	msgData.Set("Body", num)
	fmt.Println(mob)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
