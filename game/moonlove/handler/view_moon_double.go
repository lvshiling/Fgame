package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/core/utils"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/moonlove/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOONLOVE_VIEW_DOUBLE_STATE_TYPE), dispatch.HandlerFunc(handlerMoonlveViewDoubleState))
}

const (
	MAX_DISTANCE = 5
)

//双人赏月状态变更
func handlerMoonlveViewDoubleState(s session.Session, msg interface{}) (err error) {
	log.Debug("moonlove:处理双人赏月状态变更请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMoonViewDoubleState := msg.(*uipb.CSMoonloveViewDoubleState)
	targetPlayerId := csMoonViewDoubleState.GetTargetPlayerId()

	err = moonlveViewDoubleState(tpl, targetPlayerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("moonlove:处理双人赏月状态变更请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("moonlove:处理双人赏月状态变更请求完成")

	return
}

//双人赏月状态变更逻辑
func moonlveViewDoubleState(pl player.Player, targetPlayerId int64) (err error) {
	if pl.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("月下情缘：玩家不在月下情缘场景中")
		playerlogic.SendSystemMessage(pl, lang.MoonloveNotInScene)
		return
	}

	sceneData := pl.GetScene().SceneDelegate().(moonlovelogic.MoonloveSceneData)
	targetPlayer := pl.GetScene().GetPlayer(targetPlayerId)
	if targetPlayer == nil {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"targetPlayerId": targetPlayerId,
			}).Warn("月下情缘：目标玩家不在月下情缘场景中")
		playerlogic.SendSystemMessage(pl, lang.MoonloveNotInScene)
		return
	}

	//是否单人状态
	tarState := sceneData.IsCouple(targetPlayerId)
	if tarState {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"tarState": tarState,
			}).Warn("月下情缘：玩家已不是单人状态")
		moonloveViewDoubleState := pbutil.BuildMoonloveViewDoubleState(targetPlayerId, false)
		pl.SendMsg(moonloveViewDoubleState)
		return
	}

	//比较位置是否小于一码
	targetPosition := targetPlayer.GetPosition()
	myPosition := pl.GetPosition()
	distance := utils.Distance(targetPosition, myPosition)
	if distance > MAX_DISTANCE {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"tarState": tarState,
			}).Warn("月下情缘：与目标距离太远")
		moonloveViewDoubleState := pbutil.BuildMoonloveViewDoubleState(targetPlayerId, false)
		pl.SendMsg(moonloveViewDoubleState)
		return
	}

	//状态变更
	sceneData.CombineCouple(pl.GetId(), targetPlayerId)
	return
}
