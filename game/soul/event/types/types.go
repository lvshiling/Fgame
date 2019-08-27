package types

import (
	soultypes "fgame/fgame/game/soul/types"
)

type SoulEventType string

const (
	//帝魂激活
	EventTypeSoulActive SoulEventType = "soulActive"
	//帝魂镶嵌
	EventTypeSoulEmbed SoulEventType = "soulEmbed"
	//帝魂升级
	EventTypeSoulUpgrade SoulEventType = "soulUpgrade"
	//帝魂觉醒
	EventTypeSoulAwaken SoulEventType = "soulAwaken"
	//帝魂强化
	EventTypeSoulStrengthen SoulEventType = "soulStrengthen"
)

type SoulEmbedEventData struct {
	oldSoulTag    soultypes.SoulType
	newOldSoulTag soultypes.SoulType
}

func (s *SoulEmbedEventData) GetOldTag() soultypes.SoulType {
	return s.oldSoulTag
}

func (s *SoulEmbedEventData) GetNewTag() soultypes.SoulType {
	return s.newOldSoulTag
}

func CreateSoulEmbedEventData(oldSoulTag soultypes.SoulType, newOldSoulTag soultypes.SoulType) *SoulEmbedEventData {
	d := &SoulEmbedEventData{
		oldSoulTag:    oldSoulTag,
		newOldSoulTag: newOldSoulTag,
	}
	return d
}

type SoulUpgradeEventData struct {
	soulTag  soultypes.SoulType
	oldOrder int32
	newOrder int32
}

func (s *SoulUpgradeEventData) GetSoulTag() soultypes.SoulType {
	return s.soulTag
}

func (s *SoulUpgradeEventData) GetOldOrder() int32 {
	return s.oldOrder
}

func (s *SoulUpgradeEventData) GetNewOrder() int32 {
	return s.newOrder
}

func CreateSoulUpgradeEventData(soulTag soultypes.SoulType, oldOrder int32, newOrder int32) *SoulUpgradeEventData {
	d := &SoulUpgradeEventData{
		soulTag:  soulTag,
		oldOrder: oldOrder,
		newOrder: newOrder,
	}
	return d
}
