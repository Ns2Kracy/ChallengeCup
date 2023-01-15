package verify

import (
	"regexp"
)

func VerifyPhone(phone string) bool {
	if phone == "" || len(phone) != 11 {
		return false
	}
	regular := `^((13[0-9])|(14[5,7])|(15([0-3]|[5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
