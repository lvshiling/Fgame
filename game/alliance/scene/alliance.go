package scene

import (
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancetemplate "fgame/fgame/game/alliance/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

// 初始化防护罩
func (sd *allianceSceneData) initProtectNpc() {
	//
	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	boTemp := warTemp.GetProtectBiologyTemp()
	pos := warTemp.GetProtectPos()
	so := scene.CreateNPC(scenetypes.OwnerTypeAlliance, sd.currentAllianceId, 0, 0, 0, boTemp, pos, 0, 0)
	sd.GetScene().AddSceneObject(so)
	sd.protectNpc = so

	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneInitProtect, sd, so)
}

// 初始化玉玺
func (sd *allianceSceneData) initYuXiNpc() {
	//
	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	boTemp := warTemp.GetYuXiBiologyTemp()
	pos := warTemp.GetYuXiPos()
	so := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, boTemp, pos, 0, 0)
	sd.GetScene().AddSceneObject(so)
	sd.yuxiNpc = so
 
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneInitYuXi, sd, so)
}

//城门死亡
func (sd *allianceSceneData) doorNpcDead(npc scene.NPC) {
	biology := npc.GetBiologyTemplate()
	if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeBuildingMonster {
		return
	}
	sd.currentDoor += 1
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneDoorBroke, sd, nil)
}

//保护罩死亡
func (sd *allianceSceneData) protectNpcDead(npc scene.NPC) {
	biology := npc.GetBiologyTemplate()
	if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeMonster {
		return
	}
	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	if biology.Id != warTemp.GetProtectBiologyTemp().Id {
		return
	}

	//生成玉玺
	sd.initYuXiNpc()
}

// 初始化阵营
func (sd *allianceSceneData) initCamp(p scene.Player) {
	if p.GetAllianceId() == sd.currentAllianceId {
		// PK状态（针对玩家）
		p.SwitchPkState(pktypes.PkStateCamp, AlliancePkCampDefend)
		//设置阵营（针对NPC）
		p.SetFactionType(scenetypes.FactionTypeChengZhanDefendPlayer)
	} else {
		p.SwitchPkState(pktypes.PkStateCamp, AlliancePkCampAttack)
		p.SetFactionType(scenetypes.FactionTypeChengZhanAttackPlayer)
	}
}

//是否占领完成
func (sd *allianceSceneData) ifReliveOccuoyFinish() bool {
	if sd.collectReliveAllianceId == 0 {
		return false
	}

	//已经采集过了
	collectFlagTime := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().OccupyFlagTime //int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceHuangGongOccupyFlagTime))
	now := global.GetGame().GetTimeService().Now()
	elapse := now - sd.collectReliveFlagStartTime
	if elapse > int64(collectFlagTime) {
		return true
	}

	return false
}

//玩家占领复活点旗子了
func (sd *allianceSceneData) onReliveOccupyFinish() {
	collectRelivePlayerId := sd.collectRelivePlayerId
	collectReliveAllianceId := sd.collectReliveAllianceId
	previousAllianceId := sd.currentReliveAllianceId

	sd.collectReliveAllianceId = 0
	sd.collectRelivePlayerId = 0
	sd.currentReliveAllianceId = collectReliveAllianceId
	log.WithFields(
		log.Fields{
			"currentAllianceId":  collectReliveAllianceId,
			"previousAllianceId": previousAllianceId,
		}).Info("alliance:复活点读条完成")

	eventData := CreateAllianceReliveCollectEventData(collectReliveAllianceId, collectRelivePlayerId)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneReliveOccupyFinish, sd, eventData)
}
