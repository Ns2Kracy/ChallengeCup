package code

import (
	"crypto/rand"
	"fmt"
	"time"

	"ChallengeCup/dao"
)

func RandomCode() string {
	randomCode, err := rand.Prime(rand.Reader, 6)
	if err != nil {
		fmt.Println(err)
	}
	dao.RedisClient.Set(randomCode.String()[0:6], "validate_code", 300*time.Second)
	return randomCode.String()[0:6]
}
