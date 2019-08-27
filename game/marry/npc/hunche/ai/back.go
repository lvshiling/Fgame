package marry

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//使用默认
func init() {
	scene.RegisterAction(scenetypes.BiologyScriptTypeWeddingCar, scene.NPCStateBack, scene.NPCActionHandler(backAction))
}

//返回
func backAction(n scene.NPC) {
	n.Idle()
}
