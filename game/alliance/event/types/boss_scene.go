package types

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

type AllianceBossSceneEventType string

const (
	//盟主召唤仙盟boss成功
	EvnetTypeAllianceBossSummonSucess AllianceBossSceneEventType = "AllianceBossSummonSucess"
	//玩家进入仙盟boss场景
	EventTypePlayerEnterAllianceBossScene = "PlayerEnterAllianceBossScene"
	//伤害排名变化
	EventTypeAllianceBossRankChanged = "AllianceBossRankChanged"
	//仙盟boss场景结束
	EventTypeAllianceBossSceneFinish = "AllianceBossSceneFinish"
	//允许仙盟成员进入仙盟boss
	EventTypeAllowPlayerEnterAllianceBoss = "AllowPlayerEnterAllianceBoss"
)

//仙盟召唤boss
type AllianceBossSummonSucessEventData struct {
	pl player.Player
	sc scene.Scene
}

func CreateAllianceBossSummonSucessEventData(pl player.Player, sc scene.Scene) *AllianceBossSummonSucessEventData {
	d := &AllianceBossSummonSucessEventData{
		pl: pl,
		sc: sc,
	}
	return d
}

func (d *AllianceBossSummonSucessEventData) GetPlayer() player.Player {
	return d.pl
}

func (d *AllianceBossSummonSucessEventData) GetScene() scene.Scene {
	return d.sc
}
