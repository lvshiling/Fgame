package boss_handler

import (
	"fgame/fgame/cross/arenaboss/arenaboss"
	"fgame/fgame/cross/shareboss/shareboss"
	"fgame/fgame/game/scene/scene"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

type arenaBossService struct {
	// //跨服世界boss
	// shareBossList []scene.NPC
	// //按战斗力排序
	// sortWorldBossList []scene.NPC
}

func (s *arenaBossService) GetShareBossList() []scene.NPC {
	return arenaboss.GetArenaBossService().GetBossList()

}

func (s *arenaBossService) GetShareBossListGroupByMap(mapId int32) (bossArr []scene.NPC) {

	for _, boss := range arenaboss.GetArenaBossService().GetBossList() {
		if boss.GetScene().MapId() == mapId {
			bossArr = append(bossArr, boss)
		}
	}

	return bossArr
}

func (s *arenaBossService) GetShareBoss(biologyId int32) scene.NPC {
	return s.getBoss(biologyId)
}

func (s *arenaBossService) GetGuaiJiShareBossList(force int64) []scene.NPC {
	return nil
}

func (s *arenaBossService) Start() {

	return
}

func (s *arenaBossService) getBoss(biologyId int32) (n scene.NPC) {
	for _, boss := range arenaboss.GetArenaBossService().GetBossList() {
		bossBiologyId := int32(boss.GetBiologyTemplate().TemplateId())
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

var (
	s = &arenaBossService{}
)

func init() {
	shareboss.RegisterShareBossHandler(worldbosstypes.BossTypeArena, s)
}
