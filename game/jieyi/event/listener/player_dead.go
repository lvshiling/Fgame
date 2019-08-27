package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}

// 玩家被杀，声威值掉落，降级
func playerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	attackId := data.(int64)
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取声威值掉落,没有场景对象")
		return
	}

	spl := s.GetPlayer(attackId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取声威值掉落,没有找到击杀者")
		return
	}

	if !s.MapTemplate().CanShengWeiDrop() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取声威值掉落,本地图不能掉")
		return
	}

	itemId, dropNum := jieyilogic.ShengWeiZhiDrop(pl, attackId, spl.GetName(), s.MapId(), pl.GetPos())
	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dropNum":  dropNum,
			}).Warn("massacre:处理获取声威值掉落,掉落数量错误")
		return
	}

	jieyilogic.JieYiPropertyChange(pl)

	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	minStack := int(constantTemp.DropMinStack)
	maxStack := int(constantTemp.DropMaxStack)
	protextTime := constantTemp.DropProtectedTime
	existTime := constantTemp.DropFailTime

	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), attackId, itemId, dropNum, stack, int32(protextTime), int32(existTime))

	return
}
