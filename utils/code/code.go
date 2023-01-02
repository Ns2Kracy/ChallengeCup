package code

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomPhoneCode() string {
	randomCode := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成6位随机数
	code := fmt.Sprintf("%06v", randomCode.Int31n(1000000))
	return code
}
