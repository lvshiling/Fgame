package battle

import (
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeChuangShiZhiZhanFuShu, scene.RelivePointHandlerFunc(RelivePoint))
}

//复活点复活
func RelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return false
	}

	tpl, ok := pl.(player.Player)
	if !ok {
		return false
	}

	//获取城战数据
	fuShuSd := s.SceneDelegate().(chuangshiscene.FuShuSceneData)
	if fuShuSd == nil {
		return
	}

	reliveFlag := fuShuSd.GetReliveFlag()
	if reliveFlag == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:处理对象复活消息,没有设置复活点")
		return false
	}

	//清空原地复活次数
	tpl.SetChuangShiReliveTimes(0)

	camp := pl.GetCamp()
	currentReliveCampType := fuShuSd.GetCurrentReliveCampType()
	if camp == currentReliveCampType {
		//获取场景的复活点
		rebornPos := reliveFlag.GetPosition()
		pl.Reborn(rebornPos)
		return true
	} else {
		rebornPos := s.MapTemplate().GetRebornPos()
		pl.Reborn(rebornPos)
		return true
	}

	return true
}
