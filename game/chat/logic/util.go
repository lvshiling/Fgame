package logic

import (
	coreutils "fgame/fgame/core/utils"
	chattypes "fgame/fgame/game/chat/types"
	"fmt"
)

// 系统模块颜色的公告关键字str格式
func FormatModuleNameNoticeStr(str string) string {
	return coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("【%s】", str))
}

// 邮件的关键字str格式
func FormatMailKeyWordNoticeStr(str string) string {
	return coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("【%s】", str))
}
