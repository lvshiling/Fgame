package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_YUGAO_TYPE), (*uipb.CSChuangShiYuGao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_YUGAO_TYPE), (*uipb.SCChuangShiYuGao)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_YUGAO_BAOMING_TYPE), (*uipb.CSBaoMingChuangShi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_YUGAO_BAOMING_TYPE), (*uipb.SCBaoMingChuangShi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_BAO_MING_TYPE), (*uipb.CSChuangShiShenWangBaoMing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_SHEN_WANG_BAO_MING_TYPE), (*uipb.SCChuangShiShenWangBaoMing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_BAO_MING_LIST_TYPE), (*uipb.CSChuangShiShenWangBaoMingList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_SHEN_WANG_BAO_MING_LIST_TYPE), (*uipb.SCChuangShiShengWangBaoMingList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_TOU_PIAO_TYPE), (*uipb.CSChuangShiShengWangTouPiao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_SHEN_WANG_TOU_PIAO_TYPE), (*uipb.SCChuangShiShengWangTouPiao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_SHEN_WANG_TOU_PIAO_LIST_TYPE), (*uipb.CSChuangShiShenWangTouPiaoList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_SHEN_WANG_TOU_PIAO_LIST_TYPE), (*uipb.SCChuangShiShengWangTouPiaoList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_INFO_TYPE), (*uipb.CSChuangShiInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_INFO_TYPE), (*uipb.SCChuangShiInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_REN_MING), (*uipb.CSChuangShiCityRenMing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_CITY_REN_MING), (*uipb.SCChuangShiCityRenMing)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CAMP_PAY_SCHEDULE_TYPE), (*uipb.CSChuangShiCampPaySchedule)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_CAMP_PAY_SCHEDULE_TYPE), (*uipb.SCChuangShiCampPaySchedule)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_PAY_SCHEDULE_TYPE), (*uipb.CSChuangShiCityPaySchedule)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_CITY_PAY_SCHEDULE_TYPE), (*uipb.SCChuangShiCityPaySchedule)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_MY_PAY_RECEIVE_TYPE), (*uipb.CSChuangShiMyPayReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_MY_PAY_RECEIVE_TYPE), (*uipb.SCChuangShiMyPayReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CAMP_PAY_RECEIVE_TYPE), (*uipb.CSChuangShiCampPayReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_CAMP_PAY_RECEIVE_TYPE), (*uipb.SCChuangShiCampPayReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_JIANSHE_TYPE), (*uipb.CSChuangShiCityJianShe)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_CITY_JIANSHE_TYPE), (*uipb.SCChuangShiCityJianShe)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_POSITION_ADVANCE_TYPE), (*uipb.CSChuangShiPositionAdvance)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_POSITION_ADVANCE_TYPE), (*uipb.SCChuangShiPositionAdvance)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_GUANZHI_REW_TYPE), (*uipb.CSChuangShiGuanZhiRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_GUANZHI_REW_TYPE), (*uipb.SCChuangShiGuanZhiRew)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_ENTER_CITY_TYPE), (*uipb.CSChuangShiEnterCity)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_ENTER_CITY_TYPE), (*uipb.SCChuangShiEnterCity)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_GONGCHENG_TARGET_TYPE), (*uipb.CSChuangShiGongChengTarget)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_GONGCHENG_TARGET_TYPE), (*uipb.SCChuangShiGongChengTarget)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_JOIN_CAMP_TYPE), (*uipb.CSChuangShiJoinCamp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_JOIN_CAMP_TYPE), (*uipb.SCChuangShiJoinCamp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_TIANQI_SET_TYPE), (*uipb.CSChuangShiCityTianQiSet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_TIANQI_SET_TYPE), (*uipb.SCChuangShiCityTianQiSet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_SHEN_WANG_BROADCAST_TYPE), (*uipb.SCChuangShiShenWangBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHUANGSHI_PLAYER_INFO_NOTICE_TYPE), (*uipb.SCChuangShiPlayerInfoNotice)(nil))
}
