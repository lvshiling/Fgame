package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
)

func BuildSCZhenFaGet(zhenFaMap map[zhenfatypes.ZhenFaType]*playerzhenfa.PlayerZhenFaObject) *uipb.SCZhenFaGet {
	scZhenFaGet := &uipb.SCZhenFaGet{}
	for _, obj := range zhenFaMap {
		scZhenFaGet.ZhenFaList = append(scZhenFaGet.ZhenFaList, buildZhenFaInfo(obj))
	}
	return scZhenFaGet
}

func buildZhenFaInfo(obj *playerzhenfa.PlayerZhenFaObject) *uipb.ZhenFaInfo {
	zhenFaInfo := &uipb.ZhenFaInfo{}
	zhenFaType := int32(obj.GetZhenFaType())
	level := obj.GetLevel()
	zhenFaInfo.ZhenFaType = &zhenFaType
	zhenFaInfo.Level = &level
	return zhenFaInfo
}

func BuildSCZhenQiGet(zhenQiMap map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]*playerzhenfa.PlayerZhenQiObject) *uipb.SCZhenQiGet {
	scZhenQiGet := &uipb.SCZhenQiGet{}
	for zhenFaType, zhenFaZhenQiMap := range zhenQiMap {
		scZhenQiGet.ZhenQiList = append(scZhenQiGet.ZhenQiList, buildZhenQiInfo(int32(zhenFaType), zhenFaZhenQiMap))
	}
	return scZhenQiGet
}

func buildZhenQiInfo(zhenFaType int32, zhenQiMap map[zhenfatypes.ZhenQiType]*playerzhenfa.PlayerZhenQiObject) *uipb.ZhenQiInfo {
	zhenQiInfo := &uipb.ZhenQiInfo{}
	zhenQiInfo.ZhenFaType = &zhenFaType

	for _, obj := range zhenQiMap {
		zhenQiInfo.ZhenQiPosList = append(zhenQiInfo.ZhenQiPosList, buildZhenQiPosInfo(obj))
	}
	return zhenQiInfo
}

func buildZhenQiPosInfo(obj *playerzhenfa.PlayerZhenQiObject) *uipb.ZhenQiPosInfo {
	zhenQiPosInfo := &uipb.ZhenQiPosInfo{}
	zhenQiPos := int32(obj.GetZhenQiPos())
	number := obj.GetNumber()
	zhenQiPosInfo.ZhenQiPos = &zhenQiPos
	zhenQiPosInfo.Number = &number
	return zhenQiPosInfo
}

func BuildSCZhenQiXianHuoGet(zhenQiXianHuoMap map[zhenfatypes.ZhenFaType]*playerzhenfa.PlayerZhenQiXianHuoObject) *uipb.SCZhenQiXianHuoGet {
	scZhenQiXianHuoGet := &uipb.SCZhenQiXianHuoGet{}

	for _, obj := range zhenQiXianHuoMap {
		scZhenQiXianHuoGet.ZhenQiXianHuoList = append(scZhenQiXianHuoGet.ZhenQiXianHuoList, buildZhenQiXianHuo(obj))
	}
	return scZhenQiXianHuoGet
}

func buildZhenQiXianHuo(obj *playerzhenfa.PlayerZhenQiXianHuoObject) *uipb.ZhenQiXianHuoInfo {
	zhenQiXianHuoInfo := &uipb.ZhenQiXianHuoInfo{}
	zhenFaType := int32(obj.GetZhenFaType())
	level := obj.GetLevel()
	luckyStar := obj.GetLuckyStar()

	zhenQiXianHuoInfo.ZhenFaType = &zhenFaType
	zhenQiXianHuoInfo.Level = &level
	zhenQiXianHuoInfo.LuckyStar = &luckyStar
	return zhenQiXianHuoInfo
}

func BuildSCZhenFaActivate(obj *playerzhenfa.PlayerZhenFaObject) *uipb.SCZhenFaActivate {
	scZhenFaActivate := &uipb.SCZhenFaActivate{}
	scZhenFaActivate.ZhenFaInfo = buildZhenFaInfo(obj)
	return scZhenFaActivate
}

func BuildSCZhenFaShengJi(sucess bool, obj *playerzhenfa.PlayerZhenFaObject) *uipb.SCZhenFaShengJi {
	scZhenFaShengJi := &uipb.SCZhenFaShengJi{}
	scZhenFaShengJi.Sucess = &sucess
	scZhenFaShengJi.ZhenFaInfo = buildZhenFaInfo(obj)
	return scZhenFaShengJi
}

func BuildSCZhenQiAdvanced(sucess bool, obj *playerzhenfa.PlayerZhenQiObject) *uipb.SCZhenQiAdvanced {
	scZhenQiAdvanced := &uipb.SCZhenQiAdvanced{}
	zhenFaType := int32(obj.GetZhenFaType())
	scZhenQiAdvanced.ZhenFaType = &zhenFaType
	scZhenQiAdvanced.ZhenQiPosInfo = buildZhenQiPosInfo(obj)
	scZhenQiAdvanced.Sucess = &sucess
	return scZhenQiAdvanced
}

func BuidlSCZhenFaXianHuoShengJi(sucess bool, obj *playerzhenfa.PlayerZhenQiXianHuoObject) *uipb.SCZhenQiXianHuoShengJi {
	scZhenQiXianHuoShengJi := &uipb.SCZhenQiXianHuoShengJi{}
	scZhenQiXianHuoShengJi.Sucess = &sucess
	scZhenQiXianHuoShengJi.ZhenQiXianHuoInfo = buildZhenQiXianHuo(obj)
	return scZhenQiXianHuoShengJi
}
