package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	playerlingtong "fgame/fgame/game/lingtong/player"
)

func BuildSCLingTongGet(lingTongId int32,
	lingTongMap map[int32]*playerlingtong.PlayerLingTongInfoObject,
	lingTongFashionMap map[int32]*playerlingtong.PlayerLingTongFashionObject) *uipb.SCLingTongGet {

	lingTongGet := &uipb.SCLingTongGet{}
	lingTongGet.LingTongId = &lingTongId

	for lingTongId, lingTongObj := range lingTongMap {
		lingTongFashion := lingTongFashionMap[lingTongId]
		fashionId := lingTongFashion.GetFashionId()
		lingTongGet.LingTongList = append(lingTongGet.LingTongList, buildLingTong(fashionId, lingTongObj))
	}
	return lingTongGet
}

func buildLingTong(fashionId int32, obj *playerlingtong.PlayerLingTongInfoObject) *uipb.LingTongInfo {
	lingTongInfo := &uipb.LingTongInfo{}
	lingTongId := obj.GetLingTongId()
	lingTongName := obj.GetLingTongName()
	level := obj.GetLevel()
	pro := obj.GetPro()
	peiYangLevel := obj.GetPeiYangLevel()
	peiYangPro := obj.GetPeiYangPro()
	upstarLevel := obj.GetStarLevel()
	upstarPro := obj.GetStarPro()

	lingTongInfo.LingTongId = &lingTongId
	lingTongInfo.LingTongName = &lingTongName
	lingTongInfo.FashinId = &fashionId
	lingTongInfo.Level = &level
	lingTongInfo.LevelPro = &pro
	lingTongInfo.PeiYangLevel = &peiYangLevel
	lingTongInfo.PeiYangPro = &peiYangPro
	lingTongInfo.StarLevel = &upstarLevel
	lingTongInfo.StarPro = &upstarPro
	return lingTongInfo
}

func BuildSCLingTongPowerNotice(power int64) *uipb.SCLingTongPowerNotice {
	scLingTongPowerNotice := &uipb.SCLingTongPowerNotice{}
	scLingTongPowerNotice.Power = &power
	return scLingTongPowerNotice
}

func BuildSCLingTongActivate(fashionId int32, obj *playerlingtong.PlayerLingTongInfoObject) *uipb.SCLingTongActive {
	lingTongActive := &uipb.SCLingTongActive{}
	lingTongId := obj.GetLingTongId()
	lingTongActive.LingTongId = &lingTongId
	lingTongActive.LingTongInfo = buildLingTong(fashionId, obj)
	return lingTongActive
}

func BuildSCLingTongChuZhan(lingTongId int32) *uipb.SCLingTongChuZhan {
	lingTongChuZhan := &uipb.SCLingTongChuZhan{}
	lingTongChuZhan.LingTongId = &lingTongId
	return lingTongChuZhan
}

func BuildSCLingTongUpgrade(lingTongId int32, level int32, progress int32, result bool) *uipb.SCLingTongUpgrade {
	lingTongUpgrade := &uipb.SCLingTongUpgrade{}
	lingTongUpgrade.LingTongId = &lingTongId
	lingTongUpgrade.Level = &level
	lingTongUpgrade.UpPro = &progress
	lingTongUpgrade.Result = &result
	return lingTongUpgrade
}

func BuildSCLingTongUpstar(lingTongId int32, level int32, progress int32, result bool) *uipb.SCLingTongUpstar {
	lingTongUpstar := &uipb.SCLingTongUpstar{}
	lingTongUpstar.LingTongId = &lingTongId
	lingTongUpstar.Level = &level
	lingTongUpstar.UpPro = &progress
	lingTongUpstar.Result = &result
	return lingTongUpstar
}

func BuildSCLingTongPeiYang(lingTongId int32, peiYangLevel int32, peiYangPro int32) *uipb.SCLingTongPeiYang {
	lingTongPeiYang := &uipb.SCLingTongPeiYang{}
	lingTongPeiYang.LingTongId = &lingTongId
	lingTongPeiYang.PeiYangLevel = &peiYangLevel
	lingTongPeiYang.PeiYangPro = &peiYangPro
	return lingTongPeiYang
}

func BuildSCLingTongRename(lingTongId int32, lingTongName string) *uipb.SCLingTongRename {
	lingTongRename := &uipb.SCLingTongRename{}
	lingTongRename.LingTongId = &lingTongId
	lingTongRename.LingTongName = &lingTongName
	return lingTongRename
}

func BuildLingTongCacheInfo(info *lingtongcommon.LingTongInfo) *uipb.LingTongCacheInfo {
	lingTongCacheInfo := &uipb.LingTongCacheInfo{}
	fashionId := info.FashionId
	lingTongId := info.LingTongId
	level := info.Level
	lingTongCacheInfo.FashionId = &fashionId
	lingTongCacheInfo.LingTongId = &lingTongId
	lingTongCacheInfo.Level = &level

	for _, lingTongDetail := range info.LingTongList {
		lingTongCacheInfo.LingTongList = append(lingTongCacheInfo.LingTongList, buildLingTongInfo(lingTongDetail))
	}

	return lingTongCacheInfo
}

func buildLingTongInfo(info *lingtongcommon.LingTongDetail) *uipb.LingTongInfo {
	lingTongInfo := &uipb.LingTongInfo{}
	lingTongId := info.LingTongId
	lingTongName := info.LingTongName
	fashionId := info.FashionId
	level := info.Level
	pro := info.LevelPro
	peiYangLevel := info.PeiYangLevel
	peiYangPro := info.PeiYangPro
	starLevel := info.StarLevel
	starPro := info.StarPro

	lingTongInfo.LingTongId = &lingTongId
	lingTongInfo.LingTongName = &lingTongName
	lingTongInfo.FashinId = &fashionId
	lingTongInfo.Level = &level
	lingTongInfo.LevelPro = &pro
	lingTongInfo.PeiYangLevel = &peiYangLevel
	lingTongInfo.PeiYangPro = &peiYangPro
	lingTongInfo.StarLevel = &starLevel
	lingTongInfo.StarPro = &starPro
	return lingTongInfo

}
