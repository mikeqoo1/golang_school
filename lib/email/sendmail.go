package email

import (
	"fmt"
	"net/smtp"
)

//SendEmail 寄信功能
func SendEmail() {
	from := "sssh@sssh.tp.edu.tw"
	to := make([]string, 10)
	to[0] = "105209@concords.com.tw"
	to[1] = "105233@concords.com.tw"
	to[2] = "106061@concords.com.tw"
	to[3] = "106062@concords.com.tw"
	to[4] = "107094@concords.com.tw"
	to[5] = "109007@concords.com.tw"
	to[6] = "109038@concords.com.tw"
	to[7] = "105105@concords.com.tw"
	to[8] = "105211@concords.com.tw"
	to[9] = "105110@concords.com.tw"

	var toStrlist string
	for i := 0; i < len(to); i++ {
		toStrlist = toStrlist + to[i] + ","
		fmt.Println(toStrlist)
	}
	msg := []byte("To: " + toStrlist + "\r\nSubject: 打球復建趣\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n此為系統發信, 明天松山見 ^_^")

	host := "exgmail.concords.com.tw" // Address from socketLabs

	port := "25" // Port from socketlabs

	// user := "" // smtp Username from socketlabs

	// pswd := "" // smtp Password from socketlabs

	hostWithPort := host + ":" + port

	//auth := smtp.CRAMMD5Auth(user, pswd)

	err := smtp.SendMail(hostWithPort, nil, from, to, msg)

	if err != nil {
		fmt.Println("SEND ENAIL ~~~~", err)
	}
	fmt.Println("done")
}
