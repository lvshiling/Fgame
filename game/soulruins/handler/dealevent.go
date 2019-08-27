package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	"fgame/fgame/game/soulruins/pbutil"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_DEALEVENT_TYPE), dispatch.HandlerFunc(handleSoulRuinsDealEvent))
}

//处理帝陵遗迹事件处理消息
func handleSoulRuinsDealEvent(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹事件处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulRuinsDealEvent := msg.(*uipb.CSSoulRuinsDealEvent)
	eventType := csSoulRuinsDealEvent.GetEventType()
	acceptFlag := csSoulRuinsDealEvent.GetAccept()

	err = soulRuinsDealEvent(tpl, soulruinstypes.SoulRuinsEventType(eventType), acceptFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"eventType":  eventType,
				"acceptFlag": acceptFlag,
				"error":      err,
			}).Error("soulruins:处理获取帝陵遗迹事件处理消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"eventType":  eventType,
			"acceptFlag": acceptFlag,
		}).Debug("soulruins:处理获取帝陵遗迹事件处理消息完成")
	return nil

}

//获取帝陵遗迹事件处理的逻辑
func soulRuinsDealEvent(pl player.Player, eventType soulruinstypes.SoulRuinsEventType, acceptFlag bool) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFuBenSoulRuins {
		return
	}
	//判断状态
	sd := s.SceneDelegate()
	sceneData := sd.(*soulruinslogic.SoulRuinsSceneData)

	//判断阶段
	stageType := sceneData.GetStageType()
	if stageType != soulruinstypes.SoulRuinsStageTypeEvent {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"eventType":  eventType,
				"acceptFlag": acceptFlag,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断事件
	curEventType := sceneData.GetEventType()
	if curEventType != eventType {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"eventType":  eventType,
				"acceptFlag": acceptFlag,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	switch eventType {
	case soulruinstypes.SoulRuinsEventTypeNot: //未触发出事件
		{
			return
		}
	case soulruinstypes.SoulRuinsEventTypeRobber: //马贼事件
		{
			if !acceptFlag {
				break
			}
			//上交银两
			isReturn := soulruinslogic.GiveSilverToRobber(pl, sceneData)
			if isReturn {
				return
			}
			break
		}
	case soulruinstypes.SoulRuinsEventTypeSoul: //帝魂降临事件
		{
			if !acceptFlag {
				s.Finish(true)
				return
			}
			break
		}
	case soulruinstypes.SoulRuinsEventTypeBoss: //Boss事件
		{
			buffId := sceneData.GetSoulRuinsTemplate().SpecialBossBuff
			scenelogic.AddBuff(pl, buffId, pl.GetId(), common.MAX_RATE)
			break
		}
	}
	flag := sceneData.AcceptEvent(acceptFlag)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsDealEvent AcceptEvent should be ok"))
	}
	scSoulRuinsDealEvent := pbutil.BuildSCSoulRuinsDealEvent(int32(eventType), acceptFlag)
	pl.SendMsg(scSoulRuinsDealEvent)
	return
}
