package main

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

//html,plain
func SendMail(title, user, pswd, smtpserver, port, from, to, subject, body, format string) error {
	bs64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	header["From"] = title + "<" + from + ">"
	header["To"] = to
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", bs64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/" + format + "; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	data := ""
	for k, v := range header {
		data += k + ": " + v + "\r\n"
	}
	data += "\r\n" + bs64.EncodeToString([]byte(body))

	err := smtp.SendMail(smtpserver+":"+port, smtp.PlainAuth("", user, pswd, smtpserver), from, []string{to}, []byte(data))
	return err
}

func main() {

	title := "Monitor"
	from := "yanghu@dyhqc.com"
	to := "yanghu@huataidun.com"
	subject := "网站监控"
	body := "123"
	smtpserver := "smtp.mxhichina.com"
	pswd := "yanghu1989##"
	err := SendMail(title, from, pswd, smtpserver, "25", from, to, subject, body, "plain")
	fmt.Println(err)
}
