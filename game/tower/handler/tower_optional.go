package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
	towertypes "fgame/fgame/game/tower/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TOWER_DA_BAO_TYPE), dispatch.HandlerFunc(handleTowerDaBao))
}

//处理打宝操作
func handleTowerDaBao(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理处理打宝操作消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTowerDaBao)
	typ := csMsg.GetTyp()

	operateType := towertypes.TowerOperationType(typ)
	if !operateType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tower:处理处理打宝操作消息,类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = operateTower(tpl, operateType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tower:处理处理打宝操作消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tower:处理处理打宝操作消息完成")
	return nil

}

func operateTower(pl player.Player, operateType towertypes.TowerOperationType) (err error) {

	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	switch operateType {
	case towertypes.TowerOperationTypeBegin:
		{
			remainTime := towerManager.GetRemainTime()
			if remainTime <= 0 {
				log.WithFields(
					log.Fields{
						"playerId":    pl.GetId(),
						"operateType": operateType,
					}).Warn("chess:处理获取打宝塔操作请求,没有打宝时间")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}
			towerManager.StartDaBao()

			// 打宝光效
			buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDaBaoBuff))
			scenelogic.AddBuff(pl, buffId, pl.GetId(), common.MAX_RATE)
		}
	case towertypes.TowerOperationTypeEnd:
		{
			towerManager.EndDaBao()

			// 移除光效
			buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDaBaoBuff))
			scenelogic.RemoveBuff(pl, buffId)
		}
	}

	remainTime := towerManager.GetRemainTime()
	scMsg := pbutil.BuildSCTowerDaBao(remainTime)
	pl.SendMsg(scMsg)
	return
}
