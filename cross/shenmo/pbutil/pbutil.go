package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

const (
	//生物存在
	bioLive = int32(1)
	//生物死亡
	bioDead = int32(2)
)

func BuildSCShenMoLineUpChanged(beforeNum int32) *uipb.SCShenMoLineUp {
	scShenMoLineUp := &uipb.SCShenMoLineUp{}
	scShenMoLineUp.BeforeNum = &beforeNum
	return scShenMoLineUp
}

func BuildSCShenMoSceneInfo(gongXunNum int32, killNum int32, jiFenNum int32, daQiList []scene.NPC) *uipb.SCShenMoSceneInfo {
	scShenMoGet := &uipb.SCShenMoSceneInfo{}
	scShenMoGet.JiFenNum = &jiFenNum
	scShenMoGet.GongXun = &gongXunNum
	scShenMoGet.KillNum = &killNum
	scShenMoGet.DaQiList = scenepbutil.BuildGeneralCollectInfoListByList(daQiList)
	return scShenMoGet
}

func BuildSCShenMoSceneEnd() *uipb.SCShenMoSceneEnd {
	scShenMoSceneEnd := &uipb.SCShenMoSceneEnd{}
	return scShenMoSceneEnd
}

func BuildSCShenMoGongXunNumChanged(gongXunNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.GongXun = &gongXunNum
	return scShenMoSceneDataChanged
}

func BuildSCShenMoKillNumChanged(killNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.KillNum = &killNum
	return scShenMoSceneDataChanged
}

func BuildSCShenMoJiFenNumChanged(jiFenNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.JiFenNum = &jiFenNum
	return scShenMoSceneDataChanged
}

func BuildSCShenMoBioBroadcast(npc scene.NPC) *uipb.SCShenMoBioBroadcast {
	shenMoBioBroadcast := &uipb.SCShenMoBioBroadcast{}
	shenMoBioBroadcast.Bio = scenepbutil.BuildGeneralCollectInfo(npc)
	return shenMoBioBroadcast
}
