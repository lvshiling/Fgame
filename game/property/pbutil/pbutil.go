package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	propertytypes "fgame/fgame/game/property/types"
)

//建立玩家属性数据
func BuildSCPlayerPropertyData(baseChanges map[int32]int64, battleChanges map[int32]int64, force int64) *uipb.SCPlayerProperty {
	spd := &uipb.SCPlayerProperty{}
	for key, val := range baseChanges {
		spd.BasePropertyList = append(spd.BasePropertyList, BuildPropertyData(int32(key), val))
	}
	for key, val := range battleChanges {
		spd.BattlePropertyList = append(spd.BattlePropertyList, BuildPropertyData(int32(key), val))
	}
	if len(spd.BattlePropertyList) != 0 {
		spd.Force = &force
	}
	return spd
}

func BuildProperties(properties map[int32]int64) (ps []*uipb.Property) {
	for key, val := range properties {
		ps = append(ps, BuildPropertyData(key, val))
	}
	return ps
}

func BuildPropertyData(key int32, val int64) *uipb.Property {
	p := &uipb.Property{}
	p.Key = &key
	p.Value = &val
	return p
}

func BuildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	return rewProperty
}
