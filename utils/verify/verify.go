package verify

import (
	"regexp"
)

func VerifyPhone(phone string) bool {
	if phone == "" || len(phone) != 11 {
		return false
	}
	regular := `/^(?:(?:\+|00)86)?1[3-9]\d{9}$/`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

func VerifyEmail(email string) bool {
	if email == "" {
		return false
	}
	regular := `/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+$/`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(email)
}
