package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeMainQuest, guaji.GuaJiEnterCheckHandlerFunc(mainQuestEnterCheck))
}

const (
	mainLevel = 75
)

func mainQuestEnterCheck(pl player.Player) bool {
	//判断死亡次数
	//判断等级
	if pl.GetLevel() >= mainLevel {
		return false
	}
	return true
}
