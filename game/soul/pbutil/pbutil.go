package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playersoul "fgame/fgame/game/soul/player"
	soultypes "fgame/fgame/game/soul/types"
)

func BuildSCSoulGet(embedList []int32, soulInfoList map[soultypes.SoulType]*playersoul.PlayerSoulObject) *uipb.SCSoulGet {
	soulGet := &uipb.SCSoulGet{}
	for _, embedId := range embedList {
		soulGet.EmbedIdList = append(soulGet.EmbedIdList, int32(embedId))
	}
	for _, soul := range soulInfoList {
		soulGet.SoulList = append(soulGet.SoulList, buildSoul(soul))
	}
	return soulGet
}

func BuildSCSoulActive(soulObj *playersoul.PlayerSoulObject) *uipb.SCSoulActive {
	soulActive := &uipb.SCSoulActive{}
	soulActive.SoulInfo = buildSoul(soulObj)
	return soulActive
}

func BuildSCSoulEmbed(soulId int32) *uipb.SCSoulEmbed {
	soulEmbed := &uipb.SCSoulEmbed{}
	soulEmbed.SoulId = &soulId
	return soulEmbed
}

func BuildSCSoulFeed(soulTag int32, level int32, exp int32) *uipb.SCSoulFeed {
	soulFeed := &uipb.SCSoulFeed{}
	soulFeed.SoulTag = &soulTag
	soulFeed.Level = &level
	soulFeed.Experience = &exp
	return soulFeed
}

func BuildSCSoulAwaken(soulTag int32, isAwaken bool) *uipb.SCSoulAwaken {
	soulAwaken := &uipb.SCSoulAwaken{}
	soulAwaken.SoulTag = &soulTag
	soulAwaken.IsAwaken = &isAwaken
	return soulAwaken
}

func BuildSCSoulUpgrade(soulTag int32, level int32) *uipb.SCSoulUpgrade {
	soulUpgrade := &uipb.SCSoulUpgrade{}
	soulUpgrade.SoulTag = &soulTag
	soulUpgrade.Level = &level
	return soulUpgrade
}

func BuildSCSoulStrengthen(soulTag int32, level int32, process int32) *uipb.SCSoulStrengthen {
	scSoulStrengthen := &uipb.SCSoulStrengthen{}
	scSoulStrengthen.SoulTag = &soulTag
	scSoulStrengthen.Level = &level
	scSoulStrengthen.Process = &process
	return scSoulStrengthen
}

func buildSoul(soul *playersoul.PlayerSoulObject) *uipb.SoulInfo {
	soulInfo := &uipb.SoulInfo{}
	soulTag := soul.SoulTag
	level := soul.Level
	experience := soul.Experience
	awakenOrder := soul.AwakenOrder
	strengthenLevel := soul.StrengthenLevel
	strengthenExp := soul.StrengthenPro

	isAwaken := false
	if soul.IsAwaken == 1 {
		isAwaken = true
	}

	soulInfo.SoulTag = &soulTag
	soulInfo.Level = &level
	soulInfo.Experience = &experience
	soulInfo.AwakenOrder = &awakenOrder
	soulInfo.IsAwaken = &isAwaken
	soulInfo.StrengthenLevel = &strengthenLevel
	soulInfo.StrengthenExp = &strengthenExp
	return soulInfo
}

func BuildAllSoulInfo(info *soultypes.AllSoulInfo) *uipb.AllSoulInfo {
	allSoulInfo := &uipb.AllSoulInfo{}

	for _, typ := range info.EmbedIdList {
		allSoulInfo.EmbedIdList = append(allSoulInfo.EmbedIdList, int32(typ))
	}
	for _, tempInfo := range info.SoulList {
		allSoulInfo.SoulList = append(allSoulInfo.SoulList, BuildSoulInfo(tempInfo))
	}
	return allSoulInfo
}

func BuildSoulInfo(info *soultypes.SoulInfo) *uipb.SoulInfo {
	soulInfo := &uipb.SoulInfo{}
	soulTag := int32(info.SoulTag)
	level := info.Level
	experience := info.Experience
	awakenOrder := info.AwakenOrder
	isAwaken := false
	if info.IsAwaken != 0 {
		isAwaken = true
	}

	strengthenLevel := info.StrengthenLevl
	strengthenExp := info.StrengthenPro
	soulInfo.SoulTag = &soulTag
	soulInfo.Level = &level
	soulInfo.Experience = &experience
	soulInfo.AwakenOrder = &awakenOrder
	soulInfo.IsAwaken = &isAwaken
	soulInfo.StrengthenLevel = &strengthenLevel
	soulInfo.StrengthenExp = &strengthenExp
	return soulInfo
}
