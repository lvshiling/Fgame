package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	coretypes "fgame/fgame/core/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_PET_ATTACK_TYPE), dispatch.HandlerFunc(handlePetAttack))
}

//处理宠物攻击包
func handlePetAttack(s session.Session, msg interface{}) error {
	log.Debugln("scene:处理宠物攻击消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csPetAttack := msg.(*scenepb.CSPetAttack)
	objType := csPetAttack.GetObjectType()
	pos := csPetAttack.GetPos()

	angle := csPetAttack.GetAngle()
	skillId := csPetAttack.GetSkillId()

	aPos := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}

	//用户攻击
	petAttack(tpl, objType, aPos, float64(angle), skillId)

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"pos":      aPos,
			"angle":    angle,
			"skillId":  skillId,
		}).Debug("scene:处理宠物攻击消息,完成")
	return nil
}

//攻击
func petAttack(pl scene.Player, objType int32, pos coretypes.Position, angle float64, skillId int32) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理宠物攻击消息,场景为空")
		return
	}

	// skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	// if skillTemplate == nil {
	// 	log.WithFields(log.Fields{
	// 		"playerId": pl.GetId(),
	// 		"pos":      pos,
	// 		"angle":    angle,
	// 		"skillId":  skillId,
	// 	}).Warn("scene:处理宠物攻击消息,技能不存在")
	// 	return
	// }
	// //被动技能
	// if skillTemplate.IsPassive() {
	// 	log.WithFields(log.Fields{
	// 		"playerId": pl.GetId(),
	// 		"pos":      pos,
	// 		"angle":    angle,
	// 		"skillId":  skillId,
	// 	}).Warnln("scene:处理宠物攻击消息,被动技能")
	// 	return
	// }

	flag := scenelogic.PetAttack(pl, objType, pos, angle, skillId, true)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"skillId":  skillId,
		}).Warnln("scene:处理宠物攻击消息,攻击失败")
		return
	}
}
