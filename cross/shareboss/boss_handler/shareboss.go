package shareboss

import (
	"fgame/fgame/cross/shareboss/shareboss"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	sharebosstemplate "fgame/fgame/game/shareboss/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

type shareBossService struct {
	//跨服世界boss
	shareBossList []scene.NPC
	//按战斗力排序
	// sortWorldBossList []scene.NPC
}

func (s *shareBossService) GetShareBossList() []scene.NPC {
	return s.shareBossList
}

func (s *shareBossService) GetShareBossListGroupByMap(mapId int32) []scene.NPC {
	var bossArr []scene.NPC
	for _, boss := range s.shareBossList {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}

	return bossArr
}

func (s *shareBossService) GetShareBoss(biologyId int32) scene.NPC {
	return s.getBoss(biologyId)
}

func (s *shareBossService) GetGuaiJiShareBossList(force int64) []scene.NPC {
	return nil
}

const (
	bossType = worldbosstypes.BossTypeShareBoss
)

func (s *shareBossService) Start() {
	// s.sortWorldBossList = make([]scene.NPC, 0, 8)
	s.shareBossList = make([]scene.NPC, 0, 8)
	mapIdList := sharebosstemplate.GetShareBossTemplateService().GetMapIdList(bossType)
	for _, mapId := range mapIdList {
		sc := scene.GetSceneService().GetBossSceneByMapId(mapId)
		if sc == nil {
			continue
		}
		//TODO:shareboss:修改优化
		bossList := sc.GetNPCS(scenetypes.BiologyScriptTypeCrossWorldBoss)
		for _, boss := range bossList {
			s.shareBossList = append(s.shareBossList, boss)
			// s.sortWorldBossList = append(s.sortWorldBossList, boss)
		}
	}
	// sort.Sort(sortWorldBossList(s.sortWorldBossList))
	return
}

func (s *shareBossService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range s.shareBossList {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

var (
	s = &shareBossService{}
)

func init() {
	shareboss.RegisterShareBossHandler(worldbosstypes.BossTypeShareBoss, s)
}
