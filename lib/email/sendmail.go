package email

import (
	"fmt"
	"net/smtp"
)

//SendEmail 寄信功能
func SendEmail() {
	from := "sssh@sssh.tp.edu.tw"
	to := make([]string, 4)
	to[0] = "106061@concords.com.tw"
	to[1] = "106062@concords.com.tw"
	to[2] = "104215@concords.com.tw"
	to[3] = "107094@concords.com.tw"

	var toStrlist string
	for i := 0; i < len(to); i++ {
		toStrlist = toStrlist + to[i] + ","
		fmt.Println(toStrlist)
	}
	msg := []byte("To: " + toStrlist + "\r\nSubject: 海洋邀請信\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n海葵國王邀你打球~~")

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
