/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 17:00:00
 * @LastEditTime: 	2021-04-19 17:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	main
 * @FilePath: 		go-liziwei01-appui/main.go
 */

package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {

}

func Send() error {
	fmt.Println("Start Send!")

	e := &email.Email{
		ReplyTo:     []string{},
		From:        "Ziwei Li <liziwei.work@gmail.com>",
		To:          []string{"118010160@link.cuhk.edu.cn"},
		Bcc:         []string{},
		Cc:          []string{},
		Subject:     "Test Subject",
		Text:        []byte("Text Body is, of course, supported!"),
		HTML:        []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Sender:      "",
		Headers:     map[string][]string{},
		Attachments: []*email.Attachment{},
		ReadReceipt: []string{},
	}
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "liziwei.work@gmail.com", "Liziwei01", "smtp.gmail.com"))

	fmt.Println("Send success!")
	if err != nil {
		return err
	}
	return nil
}
