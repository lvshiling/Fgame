package battle

import (
	"fgame/fgame/game/scene/scene"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
)

type TianShuData struct {
	typ  tianshutypes.TianShuType
	rate int32
}

// 天书
type PlayerTianShuManager struct {
	p           scene.Player
	tianshuList []*TianShuData
}

func (m *PlayerTianShuManager) GetTianShuRate(typ tianshutypes.TianShuType) int32 {
	for _, data := range m.tianshuList {
		if data.typ == typ {
			return data.rate
		}
	}

	return 0
}

func (m *PlayerTianShuManager) AddTianShu(typ tianshutypes.TianShuType, rate int32) {
	d := &TianShuData{
		typ:  typ,
		rate: rate,
	}

	m.tianshuList = append(m.tianshuList, d)
}

func (m *PlayerTianShuManager) UplevelTianShu(typ tianshutypes.TianShuType, newRate int32) {
	for _, data := range m.tianshuList {
		if data.typ != typ {
			continue
		}

		data.rate = newRate
	}
}

func CreatePlayerTianShuManagerWithData(p scene.Player, dataMap map[tianshutypes.TianShuType]*playertianshu.PlayerTianShuObject) *PlayerTianShuManager {
	m := &PlayerTianShuManager{
		p: p,
	}

	for typ, obj := range dataMap {
		level := obj.GetLevel()
		tianshuTemp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, level)
		rate := tianshuTemp.Tequan
		m.AddTianShu(typ, rate)
	}

	return m
}

func CreatePlayerTianShuManager(p scene.Player) *PlayerTianShuManager {
	m := &PlayerTianShuManager{
		p: p,
	}

	return m
}
