package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//@author: yj
//@function: GenValidateCode
//@description: 生成短信验证码
//@param: int
//@return: string
func GenValidateCode(longer int) string {
	num := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(num)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < longer; i++ {
		fmt.Fprintf(&sb, "%d", num[rand.Intn(r)])
	}
	return sb.String()
}
