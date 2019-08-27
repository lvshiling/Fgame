package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	massacrelogic "fgame/fgame/game/massacre/logic"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

//玩家死亡杀气掉落，戮仙刃降等级
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
			}).Warn("massacre:处理获取戮仙刃掉落,没有场景对象")
		return
	}

	spl := s.GetPlayer(attackId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取戮仙刃掉落,没有找到击杀者")
		return
	}

	if s.MapTemplate().CanShaqiDrop() == false {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取戮仙刃掉落,本地图不能掉")
		return
	}

	itemId, dropNum := massacrelogic.MassacreProcessDrop(pl, attackId, spl.GetName())
	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dropNum":  dropNum,
			}).Warn("massacre:处理获取戮仙刃掉落,掉落数量错误1")
		return
	}

	// manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	// manager.SetMassacreDrop(itemNum) //设置掉落时间，剩下的杀气
	minStack := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqMinStack))
	maxStack := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqMaxStack)) + 1
	protectedTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqProtectedTime)
	existTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqExistTime)
	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), attackId, itemId, int32(dropNum), stack, protectedTime, existTime)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}
