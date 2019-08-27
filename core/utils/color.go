package utils

import (
	"fmt"
)

//设置颜色
func FormatColor(color string, content string) string {
	return fmt.Sprintf("[color=%s]%s[/color]", color, content)
}
