package quest_guaji

import (
	"fgame/fgame/game/quest/guaji/guaji"
	questtypes "fgame/fgame/game/quest/types"
)

//进入天劫塔X次
func init() {

	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeRealmLevel, guaji.QuestGuaJiFunc(enterRealm))
}
