package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_OBJECT_ATTACK_TYPE), dispatch.HandlerFunc(handleAttack))
}

//处理攻击包
func handleAttack(s session.Session, msg interface{}) error {
	log.Debugln("scene:处理对象攻击消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csObjectAttack := msg.(*scenepb.CSObjectAttack)

	pos := csObjectAttack.GetPos()

	angle := csObjectAttack.GetAngle()
	skillId := csObjectAttack.GetSkillId()
	uId := csObjectAttack.GetUid()
	aPos := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}

	scenelogic.HandleObjectAttack(tpl, uId, aPos, float64(angle), skillId)

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"pos":      aPos,
			"angle":    angle,
			"skillId":  skillId,
		}).Debug("scene:处理对象攻击消息,完成")
	return nil
}
