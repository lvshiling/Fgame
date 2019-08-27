package logic

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//复活
func OnReborn(bo scene.BattleObject, pos coretypes.Position) {
	bo.GetScene().OnReborn(bo, pos)

	switch boObj := bo.(type) {
	case scene.Player:
		{
			onPlayerReborn(boObj, pos)
		}
		break
	}
}

//推送复活消息
func onPlayerReborn(p scene.Player, pos coretypes.Position) {
	scPlayerRelive := pbutil.BuildSCPlayerRelive(p)
	BroadcastNeighborIncludeSelf(p, scPlayerRelive)
	if p.GetLingTong() != nil && !p.IsLingTongHidden() {
		if CheckIfLingTongAndPlayerSameScene(p.GetLingTong()) {
			FixPosition(p.GetLingTong(), pos)
		}
	}

	scUIPlayerRelive := pbutil.BuildSCUIPlayerRelive(p)
	p.SendMsg(scUIPlayerRelive)
	return
}

//复活
func RebornBack(pl scene.Player) bool {
	s := pl.GetScene()
	if s == nil {
		return false
	}
	mapTemplate := s.MapTemplate()
	rebornPos := s.MapTemplate().GetRebornPos()
	os := pl.GetScene()
	s = scene.GetSceneService().GetWorldSceneByMapId(mapTemplate.RebornId)
	if s == nil {
		s = scene.GetSceneService().GetActivitySceneByMapId(mapTemplate.RebornId)
	}
	if s == nil {
		s = scene.GetSceneService().GetBossSceneByMapId(mapTemplate.RebornId)
	}

	if s == nil {
		s = scene.GetSceneService().GetTowerSceneByMapId(mapTemplate.RebornId)
	}

	if s == nil {
		s = scene.GetSceneService().GetSceneByMapId(mapTemplate.RebornId)
	}

	if s == nil {
		return false
	}

	if os != s {
		PlayerEnterScene(pl, s, mapTemplate.GetRebornPos())
	} else {
		pl.Reborn(rebornPos)
	}
	return true
}

func EnterEntryRelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return false
	}
	mapSceneType := s.MapTemplate().GetMapType()
	switch mapSceneType {
	case scenetypes.SceneTypeCrossTuLong,
		scenetypes.SceneTypeCrossTeamCopy:
		{
			reliveEntryPointHandler := scene.GetReliveEntryPointHandler(mapSceneType)
			return reliveEntryPointHandler.ReliveEntryPoint(pl)
		}
	default:
		{
			pl.Reborn(s.MapTemplate().GetBornPos())
			return true
		}
	}
}

func EnterRelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return false
	}
	mapSceneType := s.MapTemplate().GetMapType()
	switch mapSceneType {
	case scenetypes.SceneTypeChengZhan,
		scenetypes.SceneTypeHuangGong,
		scenetypes.SceneTypeCrossLianYu,
		scenetypes.SceneTypeCrossGodSiege,
		scenetypes.SceneTypeCrossDenseWat,
		scenetypes.SceneTypeCrossShenMo:
		relivePointHandler := scene.GetRelivePointHandler(mapSceneType)
		return relivePointHandler.RelivePoint(pl)
	default:
		return RebornBack(pl)
	}
	return false
}

//自动复活
func AutoReborn(pl scene.Player) bool {

	s := pl.GetScene()
	if s == nil {
		return false
	}
	mapTemplate := s.MapTemplate()
	autoReliveHandler := scene.GetAutoReliveHandler(s.MapTemplate().GetMapType())
	if autoReliveHandler != nil {
		flag := autoReliveHandler.AutoRelive(pl)
		if !flag {
			return false
		}
	}

	reliveType := mapTemplate.GetReliveType()
	switch reliveType {
	case scenetypes.ReliveTypeImmediate:
		pl.Reborn(pl.GetPosition())
		return true

	case scenetypes.ReliveTypeBack:
		return RebornBack(pl)

	case scenetypes.ReliveTypeEnterPoint:
		// pl.Reborn(mapTemplate.GetBornPos())
		// return true
		return EnterEntryRelivePoint(pl)
	case scenetypes.ReliveTypeRelivePoint:
		//特殊处理
		return EnterRelivePoint(pl)
	}
	return false
}
