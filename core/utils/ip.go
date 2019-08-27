package utils

import (
	"strings"
)

func SplitIp(wholeIp string) (ip string, port string) {
	strList := strings.Split(wholeIp, ":")
	if len(strList) != 2 {
		return
	}
	ip = strList[0]
	port = strList[1]
	return
}
