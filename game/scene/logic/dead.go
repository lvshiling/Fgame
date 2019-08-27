package logic

import (
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//死亡
func Dead(defenceObject scene.BattleObject, attackId int64) {
	s := defenceObject.GetScene()
	if s == nil {
		return
	}
	killName := ""
	attackObject := s.GetSceneObject(attackId)
	if attackObject != nil {
		switch attackObj := attackObject.(type) {
		case scene.Player:
			{
				killName = attackObj.GetName()
				switch defendObj := defenceObject.(type) {
				case scene.NPC:
					{

						// monsterKillEvent := sceneeventtypes.CreateMonsterKilledEventData(attackObj, int32(defendObj.GetBiologyTemplate().TemplateId()))
						gameevent.Emit(sceneeventtypes.EventTypeMonsterKilled, attackObj, defendObj)
						break
					}
				case scene.Player:
					{
						// playerKilledEvent := sceneeventtypes.CreatePlayerKilledEventData(attackObj, defendObj)
						gameevent.Emit(sceneeventtypes.EventTypePlayerKilled, attackObj, defenceObject)
						break
					}
				}
				break
			}
		case scene.NPC:
			{
				killName = attackObj.GetBiologyTemplate().Name
			}
		}
	}
	if killName == "" {
		killName = scenetypes.GetKillName(attackId)
	}
	//发送击杀事件
	switch defendObj := defenceObject.(type) {
	case scene.Player:
		//发送击杀
		scPlayerKilled := pbutil.BuildSCPlayerKilled(killName)
		defendObj.SendMsg(scPlayerKilled)
		break
	case scene.NPC:
		{
			for _, so := range defendObj.GetEnemies() {
				if so.GetHate() <= 1 {
					continue
				}
				switch hurtObj := so.BattleObject.(type) {
				case scene.Player:
					// monsterHurtedEvent := sceneeventtypes.CreateMonsterHurtedEventData(hurtObj, int32(defendObj.GetBiologyTemplate().TemplateId()))
					gameevent.Emit(sceneeventtypes.EventTypeMonsterHurted, hurtObj, int32(defendObj.GetBiologyTemplate().TemplateId()))
					break
				}
			}

			ownerId := int64(0)
			if attackObject != nil {
				switch attackObj := attackObject.(type) {
				case scene.Player:
					ownerId = attackObj.GetId()
					//打宝塔特殊处理
					if !attackObj.IsOnDabao() && defenceObject.GetScene().MapTemplate().IsTower() {
						goto AfterDrop
					}
					break
				}
			}

			if defendObj.GetBiologyTemplate().GetDropType() == scenetypes.DropTypeAfterDead &&
				defendObj.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeGeneralCollect &&
				defendObj.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeArenaTreasure &&
				defendObj.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeUnrealBoss &&
				defendObj.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeOutlandBoss {
				switch defendObj.GetBiologyTemplate().GetDropJudgeType() {
				case scenetypes.DropJudgeTypeKiller:
					{
						//掉落归属
						Drop(defendObj, scenetypes.DropOwnerTypePlayer, ownerId, 1)
						break
					}
				case scenetypes.DropJudgeTypeKillerOrTeam:
					{
						teamId := int64(0)
						switch attackObj := attackObject.(type) {
						case scene.Player:
							teamId = attackObj.GetTeamId()
							break
						}

						if teamId != 0 {
							//掉落归属
							Drop(defendObj, scenetypes.DropOwnerTypeTeam, teamId, 1)
						} else {
							Drop(defendObj, scenetypes.DropOwnerTypePlayer, ownerId, 1)
						}
						break
					}
				case scenetypes.DropJudgeTypeOpener,
					scenetypes.DropJudgeTypeOpenerOrTeam:
					//TODO: 需要记录开怪者
					Drop(defendObj, scenetypes.DropOwnerTypePlayer, ownerId, 1)
					break
				case scenetypes.DropJudgeTypeMaxHurt:
					{
						maxDamage := int64(0)
						for attackId, damage := range defenceObject.GetAllDamages() {
							so := s.GetSceneObject(attackId)
							if so == nil {
								continue
							}
							_, ok := so.(scene.Player)
							if !ok {
								continue
							}
							if damage > maxDamage {
								ownerId = attackId
								maxDamage = damage
							}
						}
						Drop(defendObj, scenetypes.DropOwnerTypePlayer, ownerId, 1)
					}
					break
				case scenetypes.DropJudgeTypeMaxHurtOrTeam:
					{
						teamDamageMap := make(map[int64]int64)
						s := defenceObject.GetScene()
						ownerType := scenetypes.DropOwnerTypePlayer
						maxDamage := int64(0)
						//个人数据
						for attackId, damage := range defenceObject.GetAllDamages() {
							so := s.GetSceneObject(attackId)
							if so == nil {
								continue
							}
							pl, ok := so.(scene.Player)
							if !ok {
								continue
							}
							if pl.GetTeamId() != 0 {
								teamDamageMap[pl.GetTeamId()] += damage
							} else {
								if damage > maxDamage {
									ownerId = attackId
									maxDamage = damage
								}
							}
						}

						for teamId, damage := range teamDamageMap {
							if damage > maxDamage {
								ownerType = scenetypes.DropOwnerTypeTeam
								ownerId = teamId
								maxDamage = damage
							}
						}

						//需要记录伤害
						Drop(defendObj, ownerType, ownerId, 1)
					}
					break
				}

			}
			break
		}

	}
AfterDrop:
	defenceObject.GetScene().OnDead(defenceObject)
}
