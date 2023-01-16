package verify

import (
	"regexp"
)

func VerifyPhone(phone string) bool {
	if phone == "" || len(phone) != 11 {
		return false
	}
	regular := `/^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[1589]))\d{8}$/`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

func VerifyEmail(email string) bool {
	if email == "" {
		return false
	}
	regular := `/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(email)
}
