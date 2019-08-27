package reddot

// import (
// 	"fgame/fgame/game/player"
// 	"fgame/fgame/game/reddot/reddot"
// 	scenetemplate "fgame/fgame/game/scene/template"
// 	welfarelogic "fgame/fgame/game/welfare/logic"
// 	welfaretypes "fgame/fgame/game/welfare/types"
// 	"fgame/fgame/game/welfare/welfare"
// )

// func init() {
// 	reddot.Register(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeFirstDrop, reddot.HandlerFunc(handleRedDotBossFirstKill))
// }

// //Boss首杀红点
// func handleRedDotBossFirstKill(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
// 	if !welfarelogic.IsOnActivityTime(groupId) {
// 		return
// 	}

// 	bossNum := int32(0)
// 	allBoss := scenetemplate.GetSceneTemplateService().GetWorldBossTemplate()
// 	for _, boss := range allBoss {
// 		for _, relateGroupId := range boss.GetGroupIdList() {
// 			if relateGroupId != groupId {
// 				continue
// 			}
// 			bossNum += 1
// 		}
// 	}

// 	curRecord := welfare.GetWelfareService().GetBossFirstKillRecord(groupId)
// 	if len(curRecord) >= int(bossNum) {
// 		return
// 	}

// 	isNotice = true
// 	return
// }
