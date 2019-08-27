package player

// import (
// 	propertycommon "fgame/fgame/game/property/common"
// 	"fgame/fgame/game/scene/scene"
// 	"fmt"
// 	"math"

// 	propertytypes "fgame/fgame/game/property/types"
// )

// //npc属性数据管理器
// type NPCPropertyDataManager struct {
// 	n scene.NPC
// 	//战斗属性组
// 	battlePropertyGroup *propertycommon.BattlePropertyGroup
// 	//当前血量
// 	currentHP int64
// 	//当前体力
// 	currentTP int64
// 	//当前护盾值
// 	currentHuDun int64
// }

// //是否属性变化过
// func (ppdm *NPCPropertyDataManager) isChanged() bool {
// 	return ppdm.battlePropertyGroup.IsChanged()
// }

// //重置改变标记位
// func (ppdm *NPCPropertyDataManager) resetChanged() {
// 	ppdm.battlePropertyGroup.ResetChanged()
// }

// //获取战斗属性
// func (ppdm *NPCPropertyDataManager) GetBattleProperty(battlePropertyType propertytypes.BattlePropertyType) int64 {
// 	return ppdm.battlePropertyGroup.Get(battlePropertyType)
// }

// func (ppdm *NPCPropertyDataManager) NPC() scene.NPC {
// 	return ppdm.n
// }

// func (ppdm *NPCPropertyDataManager) GetHP() int64 {
// 	return ppdm.currentHP
// }
// func (ppdm *NPCPropertyDataManager) GetTP() int64 {
// 	return ppdm.currentTP
// }

// func (ppdm *NPCPropertyDataManager) GetHuDun() int64 {
// 	return ppdm.currentHuDun
// }

// func (ppdm *NPCPropertyDataManager) CostHP(hp int64) bool {
// 	if hp <= 0 {
// 		return false
// 	}
// 	//计算护盾
// 	if ppdm.currentHuDun > 0 {
// 		if hp <= ppdm.currentHuDun {
// 			ppdm.currentHuDun -= hp
// 			return false
// 		}
// 		hp -= ppdm.currentHuDun
// 	}

// 	ppdm.currentHP -= hp
// 	if ppdm.currentHP <= 0 {
// 		ppdm.currentHP = 0
// 		return true
// 	}
// 	return false
// }

// func (ppdm *NPCPropertyDataManager) AddHP(hp int64) int64 {
// 	if hp <= 0 {
// 		return 0
// 	}

// 	oldHp := ppdm.currentHP
// 	currentHp := oldHp + hp
// 	hpMax := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxHP)
// 	if currentHp > hpMax {
// 		currentHp = hpMax
// 	}
// 	ppdm.currentHP = currentHp

// 	return ppdm.currentHP - oldHp
// }

// func (ppdm *NPCPropertyDataManager) Reborn() {
// 	hpMax := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxHP)
// 	ppdm.currentHP = hpMax
// 	tpMax := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxTP)
// 	ppdm.currentTP = tpMax
// }

// //更新战斗属性
// func (ppdm *NPCPropertyDataManager) UpdateBattleProperty(mask uint32) {

// 	for _, effType := range npcPropertyEffectorList {
// 		if effType.Mask()&mask != 0 {
// 			pef := getNPCPropertyEffector(effType)
// 			if pef == nil {
// 				panic(fmt.Errorf("can not find npc property effector %s", effType.String()))
// 			}
// 			//获取相对应属性
// 			p := ppdm.battlePropertyGroup.GetPropertySegment(effType)
// 			p.Clear()
// 			//属性作用
// 			pef(ppdm.n, p)
// 		}
// 	}

// 	hpMaxOld := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxHP)
// 	tpMaxOld := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxTP)
// 	ppdm.battlePropertyGroup.UpdateProperty()

// 	//更新基础属性
// 	hpChanged := ppdm.battlePropertyGroup.IsTypeChanged(propertytypes.BattlePropertyTypeMaxHP)
// 	if hpChanged {
// 		hpMaxNow := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxHP)
// 		hpNow := hpMaxNow
// 		if hpMaxOld != 0 {
// 			hpNow = int64(math.Ceil(float64(hpMaxNow) * (float64(ppdm.GetHP()) / float64(hpMaxOld))))
// 		}
// 		ppdm.currentHP = hpNow
// 	}
// 	tpChanged := ppdm.battlePropertyGroup.IsTypeChanged(propertytypes.BattlePropertyTypeMaxTP)
// 	if tpChanged {
// 		tpMaxNow := ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxTP)
// 		tpNow := tpMaxNow
// 		if tpMaxOld != 0 {
// 			tpNow = int64(math.Ceil(float64(tpMaxNow) * (float64(ppdm.GetTP()) / float64(tpMaxOld))))
// 		}
// 		ppdm.currentTP = tpNow
// 	}

// }

// //心跳
// func (m *NPCPropertyDataManager) Heartbeat() {

// }

// func CreateNPCPropertyDataManager(n scene.NPC) *NPCPropertyDataManager {
// 	ppdm := &NPCPropertyDataManager{}
// 	ppdm.n = n
// 	ppdm.battlePropertyGroup = propertycommon.NewBattlePropertyGroup()
// 	return ppdm
// }
