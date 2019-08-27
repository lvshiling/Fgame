package check_enter

import (
	activitytypes "fgame/fgame/game/activity/types"
	consttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/foe/foe"
	"fgame/fgame/game/foe/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	playershenyu "fgame/fgame/game/shenyu/player"
	shenyutemplate "fgame/fgame/game/shenyu/template"
	"fgame/fgame/pkg/mathutils"
)

func init() {
	foe.RegisterFoeNoticeHandler(scenetypes.SceneTypeShenYu, foe.FoeNoticeHandlerFunc(foeHandler))
}

func foeHandler(spl, sFoePl scene.Player, sceneType scenetypes.SceneType) (err error) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}
	foePl, ok := sFoePl.(player.Player)
	if !ok {
		return
	}
	//TODO:xzk 统一做活动中
	flag, _ := pl.IfCanKilledInActivity(activitytypes.ActivityTypeShenYu)
	if !flag {
		return
	}

	shenYuManager := pl.GetPlayerDataManager(types.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
	dropNum := shenYuManager.PlayerDead()
	if dropNum <= 0 {
		return
	}
	pl.KilledInActivity(activitytypes.ActivityTypeShenYu)
	itemId := int32(consttypes.ShenYuKey)
	shenYuConstantTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuConstantTemplate()
	minStack := shenYuConstantTemp.MinStack
	maxStack := shenYuConstantTemp.MaxStack + 1
	protectedTime := shenYuConstantTemp.ProtectedTime
	existTime := shenYuConstantTemp.ExistTime

	stack := int32(mathutils.RandomRange(int(minStack), int(maxStack)))
	scenelogic.CustomItemDrop(pl.GetScene(), pl.GetPosition(), foePl.GetId(), itemId, dropNum, stack, protectedTime, existTime)

	//仇人信息推送
	noticeScMsg := pbutil.BuildSCFoeNoticeShenYu(foePl.GetId(), foePl.GetName(), foePl.GetRole(), foePl.GetSex(), int32(sceneType), dropNum)
	pl.SendMsg(noticeScMsg)
	return
}
