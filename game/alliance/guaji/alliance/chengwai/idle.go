package chengwai

import (
	coreutils "fgame/fgame/core/utils"
	alliancelogic "fgame/fgame/game/alliance/logic"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeChengWai, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

const (
	yuxiZhouMaxDistance = 30
)

func (a *idleAction) GuaJi(p scene.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}
	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	yuxiPos := warTemp.GetYuXiPos()
	allianceSubSceneData := s.SceneDelegate().(alliancescene.AllianceSceneData)
	currentDoor := allianceSubSceneData.GetCurrentDoor()
	if currentDoor >= 3 {
		//玉玺周围查找敌人
		if coreutils.Distance(p.GetPos(), yuxiPos) <= yuxiZhouMaxDistance {
			//获取玉玺位置
			e := scenelogic.FindHatestEnemy(p)
			if e != nil {
				p.SetAttackTarget(e.BattleObject)
				p.GuaJiTrace()
				return
			}
			//查找默认目标
			bo := p.GetDefaultAttackTarget()
			if bo != nil {
				p.SetAttackTarget(bo)
				p.GuaJiTrace()
				return
			}
		}
		//判断是否正在移动
		if p.IsMove() {
			return
		}
		pos := s.MapTemplate().GetMap().RandomPosition()
		flag, _ := alliancelogic.ChenZhanCheckMove(p, pos)
		if flag {
			flag = p.SetDestPosition(pos)
			if !flag {
				return
			}
		}
		return
	}

	e := scenelogic.FindHatestEnemy(p)
	//查找敌人

	//TODO 判断安全区
	if e != nil {
		p.SetAttackTarget(e.BattleObject)
		p.GuaJiTrace()
		return
	}
	//查找默认目标
	bo := p.GetDefaultAttackTarget()
	if bo != nil {
		p.SetAttackTarget(bo)
		p.GuaJiTrace()
		return
	}

	//判断是否正在移动
	if p.IsMove() {
		return
	}

	isDefend := allianceSubSceneData.GetCurrentDefendAllianceId() == p.GetAllianceId()
	// switch allianceSubSceneData.GetState() {
	// case alliancetypes.AllianceSceneStateBangPai:
	// 	//皇宫是否关闭
	// 	if alliancelogic.IfHuGongClose(s) {
	// 		//随机走动
	// 		pos := s.MapTemplate().GetMap().RandomPosition()
	// 		flag, _ := alliancelogic.ChenZhanCheckMove(p, pos)
	// 		if flag {
	// 			flag = p.SetDestPosition(pos)
	// 			if !flag {
	// 				return
	// 			}
	// 		}
	// 	} else {
	// 		//进皇宫
	// 		portalMap := scenetemplate.GetSceneTemplateService().GetPortalTemplateMapByMapId(s.MapId())
	// 		for _, portal := range portalMap {
	// 			//同场景
	// 			if portal.MapId == s.MapId() {
	// 				continue
	// 			}

	// 			scenelogic.MoveToPortal(p, portal)
	// 			return
	// 		}
	// 	}

	// 	break
	// case alliancetypes.AllianceSceneStateGongCheng:

	//守方 守城门
	if isDefend {
		pos := s.MapTemplate().GetMap().RandomPosition()
		flag, _ := alliancelogic.ChenZhanCheckMove(p, pos)
		if flag {
			flag = p.SetDestPosition(pos)
			if !flag {
				return
			}
		}
	} else {
		//TODO zrc:修改为统一方法
		//随机行为
		flag := mathutils.RandomOneHit(0.5)
		if flag {
			//攻方 打城门
			buildNpcs := s.GetNPCS(scenetypes.BiologyScriptTypeBuildingMonster)
			//闲逛
			doorReward := alliancetemplate.GetAllianceTemplateService().GetDoorRewardTemplate(currentDoor)
			for _, buildNpc := range buildNpcs {
				if buildNpc.GetBiologyTemplate().Id == int(doorReward.BiologyId) {
					p.SetAttackTarget(buildNpc)
					p.GuaJiTrace()
					return
				}
			}
		} else {
			//随机逛
			pos := s.MapTemplate().GetMap().RandomPosition()
			flag, _ := alliancelogic.ChenZhanCheckMove(p, pos)
			if flag {
				flag = p.SetDestPosition(pos)
				if !flag {
					return
				}
			}
		}
	}

	// }

	return
}

func (a *idleAction) OnExit() {
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
