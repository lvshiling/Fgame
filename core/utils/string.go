package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitAsIntArray(s string) (numArr []int32, err error) {
	if len(s) == 0 {
		return
	}
	strArr := strings.Split(s, ",")
	for _, str := range strArr {
		num, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}
		numArr = append(numArr, int32(num))
	}
	return
}

func SplitAsFloatArray(s string) (numArr []float64, err error) {
	if len(s) == 0 {
		return
	}
	strArr := strings.Split(s, ",")
	for _, str := range strArr {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		numArr = append(numArr, num)
	}
	return
}

func FormatHref(args []int64) string {
	paramStr := ""
	for index, param := range args {
		paramStr += fmt.Sprintf("%d", param)
		if index != len(args)-1 {
			paramStr += ","
		}
	}

	return paramStr
}

func FormatLink(linkText string, args []int64) string {
	href := FormatHref(args)
	link := fmt.Sprintf("<a href='%s'>%s</a>", href, linkText)
	return link
}

func FormatNoticeStr(str string) string {
	return fmt.Sprintf("【%s】", str)
}

func FormatStrUnderline(str string) string {
	return fmt.Sprintf("[u]%s[/u]", str)
}

func FormatStrPosiotn(x, z string) string {
	return fmt.Sprintf("(%s,%s)", x, z)
}

func FormatNoticeStrUnderline(str string) string {
	strLine := FormatStrUnderline(str)
	return FormatNoticeStr(strLine)
}

func FormatParamsAsString(params ...interface{}) string {
	args := ""
	for _, param := range params {
		str := fmt.Sprintf("%v,", param)
		args += str
	}
	if len(args) > 0 {
		args = args[:len(args)-1]
	}
	return args
}
