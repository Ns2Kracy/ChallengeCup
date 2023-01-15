package code

import (
	"crypto/rand"
	"fmt"
)

func RandomCode() string {
	randomCode, err := rand.Prime(rand.Reader, 6)
	if err != nil {
		fmt.Println(err)
	}
	return randomCode.String()[0:6]
}

func PhoneSendCode(phone string, code string) error {
	return nil
}

func MailSendCode(email string, code string) error {
	return nil
}
