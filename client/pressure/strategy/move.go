package strategy

import (
	"fgame/fgame/client/player/player"
	playerproperty "fgame/fgame/client/property/player"
	"fgame/fgame/client/scene/pbutil"
	playerscene "fgame/fgame/client/scene/player"
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"math"
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"
)

type moveStrategy struct {
	p      *player.Player
	t      *time.Ticker
	elapse time.Duration
}

func (s *moveStrategy) GetPlayer() *player.Player {
	return s.p
}

func (s *moveStrategy) Run() {
	// for {
	// 	select {
	// 	case <-s.t.C:
	// 		s.randomMove()
	// 	}
	// }
}

func (s *moveStrategy) randomMove() {
	elapse := float64(s.elapse) / float64(time.Second)
	//获取当前位置
	sceneManager := s.p.GetManager(player.PlayerDataKeyScene).(*playerscene.PlayerSceneDataManager)
	mapId := sceneManager.GetMapId()
	tempMapTemplate := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil))
	if tempMapTemplate == nil {
		log.WithFields(
			log.Fields{
				"mapId": mapId,
			}).Warn("pressure:找不到地图")
	}
	mapTemplate := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil)).(*gametemplate.MapTemplate)

	curPos := sceneManager.GetPos()
	//随机角度
	angle := rand.Float64() * math.Pi * 2
	propertyManager := s.p.GetManager(player.PlayerDataKeyProperty).(*playerproperty.PlayerPropertyDataManager)
	//获取移动速度
	moveSpeed := float64(propertyManager.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) / float64(1000)
	//计算下一个地点
	//移动的距离
	distance := float64(elapse) * moveSpeed
	destPos := coretypes.Position{
		curPos.X + math.Cos(angle)*distance,
		curPos.Y,
		curPos.Z + math.Sin(angle)*distance,
	}

	destPos.Y = mapTemplate.GetMap().GetHeight(destPos.X, destPos.Z)
	faceAngle := angle / math.Pi * 180

	//保留原地
	if !mapTemplate.GetMap().IsMask(destPos.X, destPos.Z) {
		return
	}

	//移动或攻击
	uId := s.p.GetPlayerId()

	csObjectMove := pbutil.BuildCSObjectMove(uId, destPos, float32(faceAngle), scenetypes.MoveTypeNormal)
	s.p.SendMessage(csObjectMove)
}

func (s *moveStrategy) OnError(code int32) {

}

func (s *moveStrategy) OnItemChanged() {

}

func CreateMoveStrategy(p *player.Player) player.Strategy {
	s := &moveStrategy{
		p: p,
	}
	s.elapse = time.Millisecond * 500
	s.t = time.NewTicker(s.elapse)

	return s
}
