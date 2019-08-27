package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	qixuelogic "fgame/fgame/game/qixue/logic"
	qixuetemplate "fgame/fgame/game/qixue/template"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

//玩家死亡杀气掉落，泣血枪降等级
func playerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	attackId := data.(int64)
	s := pl.GetScene()
	if s == nil {
		return
	}

	spl := s.GetPlayer(attackId)
	if spl == nil {
		return
	}
	if !s.MapTemplate().IfCanShaLuDrop() {
		return
	}

	itemId, dropNum := qixuelogic.QiXueProcessDrop(pl, attackId, spl.GetName())
	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dropNum":  dropNum,
			}).Warn("qixue:处理获取泣血枪掉落,掉落数量错误1")
		return
	}

	minStack := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropMinStack)
	maxStack := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropMaxStack) + 1
	protectedTime := qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropProtectedTime
	existTime := qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropFailTime
	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), attackId, itemId, int32(dropNum), stack, protectedTime, existTime)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}
