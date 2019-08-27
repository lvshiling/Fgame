package for_notice_handler

import (
	"fgame/fgame/game/foe/foe"
	"fgame/fgame/game/foe/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	foe.RegisterDefaultHandler(foe.FoeNoticeHandlerFunc(defaultFoeHandler))
}

func defaultFoeHandler(pl, foePl scene.Player, sceneType scenetypes.SceneType) (err error) {

	//仇人信息推送
	noticeScMsg := pbutil.BuildSCFoeNotice(foePl.GetId(), foePl.GetName(), foePl.GetRole(), foePl.GetSex(), int32(sceneType))
	pl.SendMsg(noticeScMsg)
	return
}
