package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	centertypes "fgame/fgame/game/center/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	propertytypes "fgame/fgame/game/property/types"
)

func BuildSCHuiYuanInfo(isHuiyuanInterim, isReceiveInterim, isReceivePlus, isHuiyuanPlus, isBuyTodayInterim, isBuyTodayPlus bool, expireTime int64, houtaiType centertypes.ZhiZunType) *uipb.SCHuiYuanInfo {
	houtai := int32(houtaiType)
	scHuiYuanInfo := &uipb.SCHuiYuanInfo{}
	scHuiYuanInfo.HuiyuanInfo = append(scHuiYuanInfo.HuiyuanInfo, buildHuiYuanInfo(isHuiyuanInterim, isReceiveInterim, expireTime, huiyuantypes.HuiYuanTypeInterim, isBuyTodayInterim))
	scHuiYuanInfo.HuiyuanInfo = append(scHuiYuanInfo.HuiyuanInfo, buildHuiYuanInfo(isHuiyuanPlus, isReceivePlus, 0, huiyuantypes.HuiYuanTypePlus, isBuyTodayPlus))
	scHuiYuanInfo.HoutaiType = &houtai

	return scHuiYuanInfo
}

func BuildSCBuyHuiYuan(expireTiem int64, houtaiType centertypes.ZhiZunType) *uipb.SCBuyHuiYuan {
	houtai := int32(houtaiType)
	scBuyHuiYuan := &uipb.SCBuyHuiYuan{}
	scBuyHuiYuan.ExpireTime = &expireTiem
	scBuyHuiYuan.HoutaiType = &houtai
	return scBuyHuiYuan
}

func BuildSCHuiYuanReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32, houtaiType centertypes.ZhiZunType) *uipb.SCHuiYuanReceiveRew {
	houtai := int32(houtaiType)
	scHuiYuanReceiveRew := &uipb.SCHuiYuanReceiveRew{}
	for itemId, num := range itemMap {
		scHuiYuanReceiveRew.DropInfo = append(scHuiYuanReceiveRew.DropInfo, buildDropInfo(itemId, num))
	}
	scHuiYuanReceiveRew.RewInfo = buildRewProperty(rd)
	scHuiYuanReceiveRew.HoutaiType = &houtai

	return scHuiYuanReceiveRew
}

func buildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
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

func buildDropInfo(itemId, num int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num

	return dropInfo
}
func buildHuiYuanInfo(isHuiYuan, isReceive bool, expireTime int64, huiyuanType huiyuantypes.HuiYuanType, isFirstRew bool) *uipb.HuiYuanInfo {
	huiYuanInfo := &uipb.HuiYuanInfo{}
	huiYuanInfo.IsHuiYuan = &isHuiYuan
	huiYuanInfo.IsReceive = &isReceive
	huiYuanInfo.ExpireTime = &expireTime
	typ := int32(huiyuanType)
	huiYuanInfo.HuiyuanType = &typ
	huiYuanInfo.IsFirstRew = &isFirstRew

	return huiYuanInfo
}
