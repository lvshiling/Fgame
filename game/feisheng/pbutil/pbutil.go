package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerfeisheng "fgame/fgame/game/feisheng/player"
)

func BuildSCFeiShengInfo(info *playerfeisheng.PlayerFeiShengObject) *uipb.SCFeiShengInfo {
	feiLevel := info.GetFeiLevel()
	gongDe := info.GetGongDeNum()
	addRate := info.GetAddRate()
	leftQn := info.GetLeftPotential()
	ti := info.GetTiZhi()
	li := info.GetLiDao()
	gu := info.GetJinGu()

	scMsg := &uipb.SCFeiShengInfo{}
	scMsg.FeiLevel = &feiLevel
	scMsg.CurGongDe = &gongDe
	scMsg.ExtralRate = &addRate
	scMsg.LeftQn = &leftQn
	scMsg.QnInfo = buildQnInfo(ti, li, gu)

	return scMsg
}

func BuildSCFeiShengEatDan(eatNum, addRate int32) *uipb.SCFeiShengEatDan {
	scMsg := &uipb.SCFeiShengEatDan{}
	scMsg.ExtralRate = &addRate
	scMsg.EatNum = &eatNum
	return scMsg
}

func BuildSCFeiShengSanGong(addGongGe int64) *uipb.SCFeiShengSanGong {
	scMsg := &uipb.SCFeiShengSanGong{}
	scMsg.AddGongDe = &addGongGe
	return scMsg
}

func BuildSCFeiShengDuJie() *uipb.SCFeiShengDuJie {
	scMsg := &uipb.SCFeiShengDuJie{}
	return scMsg
}

func BuildSCFeiShengResteQn(info *playerfeisheng.PlayerFeiShengObject) *uipb.SCFeiShengRestQn {
	scMsg := &uipb.SCFeiShengRestQn{}
	leftQn := info.GetLeftPotential()
	scMsg.LeftQn = &leftQn
	scMsg.QnInfo = buildQnInfo(info.GetTiZhi(), info.GetLiDao(), info.GetJinGu())
	return scMsg
}

func BuildSCFeiShengSaveQn(ti, li, gu int32) *uipb.SCFeiShengSaveQn {
	scMsg := &uipb.SCFeiShengSaveQn{}
	scMsg.QnInfo = buildQnInfo(ti, li, gu)
	return scMsg
}

func BuildSCFeiShengSanGongBroadcast(plName string, addExp int64) *uipb.SCFeiShengSanGongBroadcast {
	scMsg := &uipb.SCFeiShengSanGongBroadcast{}
	scMsg.PlName = &plName
	scMsg.AddExp = &addExp
	return scMsg
}

func BuildSCFeiShengDuJieNotice(isSuccess bool, feiLevel int32) *uipb.SCFeiShengDuJieNotice {
	scMsg := &uipb.SCFeiShengDuJieNotice{}
	scMsg.IsSuccess = &isSuccess
	scMsg.FeiLevel = &feiLevel
	return scMsg
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func buildQnInfo(ti, li, gu int32) *uipb.QnInfo {
	info := &uipb.QnInfo{}
	info.Ti = &ti
	info.Li = &li
	info.Gu = &gu
	return info
}
