package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	ylpplayer "fgame/fgame/game/yinglingpu/player"
)

func BuildYlpInfo(tjList []*ylpplayer.YingLingPuObject, spList []*ylpplayer.YingLingPuSuiPianObject) *uipb.SCYingLingPuQuery {
	rst := &uipb.SCYingLingPuQuery{}
	rst.SuiPianList = make([]*uipb.YingLingPuSpInfo, 0)
	rst.TuJianList = make([]*uipb.YingLingPuInfo, 0)
	for _, value := range tjList {
		item := &uipb.YingLingPuInfo{}
		item.Level = &value.Level
		item.TuJianId = &value.TuJianId
		tujianType := int32(value.TuJianType)
		item.TuJianType = &tujianType
		rst.TuJianList = append(rst.TuJianList, item)
	}
	for _, value := range spList {
		item := &uipb.YingLingPuSpInfo{}
		item.SuiPianId = &value.SuiPianId
		item.TuJianId = &value.TuJianId
		tujianType := int32(value.TuJianType)
		item.TuJianType = &tujianType
		rst.SuiPianList = append(rst.SuiPianList, item)
	}
	return rst
}

func BuildYlpXQ(tujianId int32, tujianType int32, suiPianId int32) *uipb.SCYingLingPuSpXiangQian {
	rst := &uipb.SCYingLingPuSpXiangQian{}
	rst.TuJianId = &tujianId
	rst.TuJianType = &tujianType
	rst.SuiPianId = &suiPianId
	return rst
}

func BuildYlpLevel(tujianId int32, tujianType int32, level int32) *uipb.SCYingLingPuUpLevel {
	rst := &uipb.SCYingLingPuUpLevel{}
	rst.TuJianId = &tujianId
	rst.TuJianType = &tujianType
	rst.Level = &level
	return rst
}
