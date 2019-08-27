package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerqixue "fgame/fgame/game/qixue/player"
	qixuetypes "fgame/fgame/game/qixue/types"
)

func BuildSCQiXueGet(obj *playerqixue.PlayerQiXueObject) *uipb.SCQiXueGet {
	scMsg := &uipb.SCQiXueGet{}
	scMsg.QiXueInfo = buildObjQiXueInfo(obj)
	return scMsg
}

func BuildSCQiXueAdavanced(obj *playerqixue.PlayerQiXueObject) *uipb.SCQiXueAdvanced {
	scMsg := &uipb.SCQiXueAdvanced{}
	scMsg.QiXueInfo = buildObjQiXueInfo(obj)
	return scMsg
}

// func BuildSCQiXueWeaponLose(old_id int32, new_id int32, attacker string) *uipb.SCQiXueWeaponLoseInfo {
// 	qixueWeaponLose := &uipb.SCQiXueWeaponLoseInfo{}
// 	qixueWeaponLose.OldAdvancedId = &old_id
// 	qixueWeaponLose.NewAdvancedId = &new_id
// 	qixueWeaponLose.KillerName = &attacker

// 	return qixueWeaponLose
// }

func BuildSCQiXueShaQiDrop(obj *playerqixue.PlayerQiXueObject, costStar int32, bagDropNum int32, attackName string) *uipb.SCQiXueShaQiDrop {
	scMsg := &uipb.SCQiXueShaQiDrop{}
	scMsg.QiXueInfo = buildObjQiXueInfo(obj)
	scMsg.DropStar = &costStar
	scMsg.DropNum = &bagDropNum
	scMsg.KillerName = &attackName

	return scMsg
}

func BuildSCQiXueShaQiVary(num int64) *uipb.SCQiXueShaQiVary {
	qixueShaQiVary := &uipb.SCQiXueShaQiVary{}
	qixueShaQiVary.ShaLuNum = &num

	return qixueShaQiVary
}

func buildQiXueInfo(info *qixuetypes.QiXueInfo) *uipb.QiXueInfo {

	slNum := info.ShaLuNum
	lev := info.CurrLevel
	star := info.CurrStar

	qixueInfo := &uipb.QiXueInfo{}
	qixueInfo.ShaLuNum = &slNum
	qixueInfo.Leve = &lev
	qixueInfo.Star = &star

	return qixueInfo
}

func buildObjQiXueInfo(obj *playerqixue.PlayerQiXueObject) *uipb.QiXueInfo {
	slNum := obj.GetShaLuNum()
	lev := obj.GetLevel()
	star := obj.GetStar()
	timesNum := obj.GetTimesNum()

	qixueInfo := &uipb.QiXueInfo{}
	qixueInfo.Leve = &lev
	qixueInfo.Star = &star
	qixueInfo.ShaLuNum = &slNum
	qixueInfo.TimesNum = &timesNum
	return qixueInfo
}
