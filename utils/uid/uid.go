package uid

import (
	"regexp"
	"strconv"

	"github.com/rs/xid"
)

func GenerateUID() int {
	guid := xid.New()
	// 取出其中的数字部分
	// 利用正则表达式来取出其中的数字部分
	re := regexp.MustCompile(`\d+`)
	uid, _ := strconv.Atoi(re.FindString(guid.String()))
	return uid
}
