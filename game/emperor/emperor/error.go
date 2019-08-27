package emperor

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	errorEmperorRobbedByOther    = gamecommon.CodeError(lang.EmperorRobbedByOther)
	errorEmperorOpenBoxRobbed    = gamecommon.CodeError(lang.EmperorOpenBoxRobbed)
	errorEmperorOpenBoxNoStorage = gamecommon.CodeError(lang.EmperorOpenBoxNoStorage)
)
