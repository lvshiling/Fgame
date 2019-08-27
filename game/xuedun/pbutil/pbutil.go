package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerxuedun "fgame/fgame/game/xuedun/player"
)

func BuildSCXueDunGet(xueDunInfo *playerxuedun.PlayerXueDunObject) *uipb.SCXueDunGet {
	xueDunGet := &uipb.SCXueDunGet{}
	xueDunGet.XunDunInfo = buildXueDun(xueDunInfo)
	return xueDunGet
}

func BuildSCXueDunUpgrade(sucess bool, xueDunInfo *playerxuedun.PlayerXueDunObject) *uipb.SCXueDunUpgrade {
	xueDunUpgrade := &uipb.SCXueDunUpgrade{}
	xueDunUpgrade.Sucess = &sucess
	xueDunUpgrade.XunDunInfo = buildXueDun(xueDunInfo)
	return xueDunUpgrade
}

func BuildSCXueDunPeiYang(xueDunInfo *playerxuedun.PlayerXueDunObject) *uipb.SCXueDunPeiYang {
	xueDunPeiYang := &uipb.SCXueDunPeiYang{}
	xueDunPeiYang.XunDunInfo = buildXueDun(xueDunInfo)
	return xueDunPeiYang
}

func BuildSCXueDunBlood(blood int64) *uipb.SCXueDunBloodChanged {
	xueDunBloodChanged := &uipb.SCXueDunBloodChanged{}
	xueDunBloodChanged.Blood = &blood
	return xueDunBloodChanged
}

func BuildSCXueDunCacheInfo(xueDunInfo *playerxuedun.PlayerXueDunObject) *uipb.XueDunCacheInfo {
	xueDunCacheInfo := &uipb.XueDunCacheInfo{}
	blood := xueDunInfo.GetBlood()
	number := xueDunInfo.GetNumber()
	star := xueDunInfo.GetStar()
	starPro := xueDunInfo.GetStarPro()
	culLevel := xueDunInfo.GetCulLevel()
	culPro := xueDunInfo.GetCulPro()
	xueDunCacheInfo.Blood = &blood
	xueDunCacheInfo.Number = &number
	xueDunCacheInfo.Star = &star
	xueDunCacheInfo.StarPro = &starPro
	xueDunCacheInfo.CulLevel = &culLevel
	xueDunCacheInfo.CulPro = &culPro
	return xueDunCacheInfo
}

func buildXueDun(obj *playerxuedun.PlayerXueDunObject) *uipb.XueDunInfo {
	xueDunInfo := &uipb.XueDunInfo{}
	number := obj.GetNumber()
	blood := obj.GetBlood()
	star := obj.GetStar()
	starNum := obj.GetStarNum()
	starPro := obj.GetStarPro()
	culLevel := obj.GetCulLevel()
	culNum := obj.GetCulNum()
	culPro := obj.GetCulPro()

	xueDunInfo.Blood = &blood
	xueDunInfo.Number = &number
	xueDunInfo.Star = &star
	xueDunInfo.StarNum = &starNum
	xueDunInfo.StarPro = &starPro
	xueDunInfo.CulLevel = &culLevel
	xueDunInfo.CulNum = &culNum
	xueDunInfo.CulPro = &culPro
	return xueDunInfo
}
