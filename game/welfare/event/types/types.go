package types

import (
	commonlog "fgame/fgame/common/log"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
)

//开服活动
type ActivityEventType string

const (
	EventTypeFinishFirstCharge         ActivityEventType = "FinishFirstCharge"                //完成首冲
	EventTypeActivityDataChanged                         = "ActivityDataChanged"              //活动数据变化
	EventTypeTimeReddotNotice                            = "TimeReddotNotice"                 //时间条件红点检查
	EventTypeRankEnd                                     = "RankEnd"                          //排行榜结束
	EventTypeAttendDrew                                  = "AttendDrew"                       //参与幸运抽奖
	EventTypeCheckActivityOpenMail                       = "CheckActivityOpenMail"            //检查活动开启邮件
	EventTypeLaBaAddLog                                  = "LaBaAddLog"                       //拉霸添加日志
	EventTypeLaBaAttendLog                               = "LaBaAttendLog"                    //拉霸参与日志
	EventTypeDrewAddLog                                  = "DrewAddLog"                       //抽奖添加日志
	EventTypeCrazyBoxAddLog                              = "CrazyBoxAddLog"                   //疯狂宝箱添加日志
	EventTypeDiscountBuyTaoCao                           = "DiscountBuyTaoCao"                //购买超值套餐
	EventTypeSystemActivate                              = "SystemActivate"                   //运营活动系统激活
	EventTypeDevelopFamousFeed                           = "DevelopFamousFeed"                //名人普培养
	EventTypeDiscountBuyZhuanShengGift                   = "DiscountBuyZhuanShengGift"        //购买转生礼包
	EventTypeActivityOpenTimeChanged                     = "EventTypeActivityOpenTimeChanged" //开服活动时间变化
	EventTypeRefreshXunHuanActivity                      = "RefreshXunHuanActivity"           //刷新循环活动
	EventTypeShopDiscountActivite                        = "ShopDiscountActivite"             //商城促销特权激活
)

//后台拉霸日志
type PlayerLaBaLogEventData struct {
	attendTimes int32
	costGold    int32
	rewGold     int32
	reason      commonlog.LaBaLogReason
	reasonText  string
}

func CreatePlayerLaBaLogEventData(attendTimes, costGold, rewGold int32, reason commonlog.LaBaLogReason, reasonText string) *PlayerLaBaLogEventData {
	d := &PlayerLaBaLogEventData{
		attendTimes: attendTimes,
		costGold:    costGold,
		rewGold:     rewGold,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerLaBaLogEventData) GetAttendTimes() int32 {
	return d.attendTimes
}

func (d *PlayerLaBaLogEventData) GetCostGold() int32 {
	return d.costGold
}

func (d *PlayerLaBaLogEventData) GetRewGold() int32 {
	return d.rewGold
}

func (d *PlayerLaBaLogEventData) GetReason() commonlog.LaBaLogReason {
	return d.reason
}

func (d *PlayerLaBaLogEventData) GetReasonText() string {
	return d.reasonText
}

// 拉霸日志
type LaBaAddLogEventData struct {
	playerName string
	costGold   int32
	rewGold    int32
}

func CreateLaBaAddLogEventData(plName string, costGold, rewGold int32) *LaBaAddLogEventData {
	d := &LaBaAddLogEventData{
		playerName: plName,
		costGold:   costGold,
		rewGold:    rewGold,
	}
	return d
}

func (d *LaBaAddLogEventData) GetPlayerName() string {
	return d.playerName
}

func (d *LaBaAddLogEventData) GetCostGold() int32 {
	return d.costGold
}

func (d *LaBaAddLogEventData) GetRewGold() int32 {
	return d.rewGold
}

// 抽奖数据
type PlayerAttendDrewEventData struct {
	groupId   int32
	attendNum int32
}

func CreatePlayerAttendDrewEventData(groupId, attendNum int32) *PlayerAttendDrewEventData {
	d := &PlayerAttendDrewEventData{
		groupId:   groupId,
		attendNum: attendNum,
	}
	return d
}

func (d *PlayerAttendDrewEventData) GetAttendNum() int32 {
	return d.attendNum
}

func (d *PlayerAttendDrewEventData) GetGroupId() int32 {
	return d.groupId
}

// 抽奖日志
type DrewAddLogEventData struct {
	playerName string
	itemId     int32
	itemNum    int32
}

func CreateDrewAddLogEventData(plName string, itemId, itemNum int32) *DrewAddLogEventData {
	d := &DrewAddLogEventData{
		playerName: plName,
		itemId:     itemId,
		itemNum:    itemNum,
	}
	return d
}

func (d *DrewAddLogEventData) GetPlayerName() string {
	return d.playerName
}

func (d *DrewAddLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *DrewAddLogEventData) GetItemNum() int32 {
	return d.itemNum
}

// 疯狂宝箱日志
type CrazyBoxAddLogEventData struct {
	playerName string
	itemId     int32
	itemNum    int32
}

func CreateCrazyBoxAddLogEventData(plName string, itemId, itemNum int32) *CrazyBoxAddLogEventData {
	d := &CrazyBoxAddLogEventData{
		playerName: plName,
		itemId:     itemId,
		itemNum:    itemNum,
	}
	return d
}

func (d *CrazyBoxAddLogEventData) GetPlayerName() string {
	return d.playerName
}

func (d *CrazyBoxAddLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *CrazyBoxAddLogEventData) GetItemNum() int32 {
	return d.itemNum
}

// 名人喂养数据
type PlayerFamousFeedEventData struct {
	groupId           int32
	totalFavorbaleNum int32
	dayFavorbaleNum   int32
}

func CreatePlayerFamousFeedEventData(groupId, addFavorbaleNum, dayFavorbaleNum int32) *PlayerFamousFeedEventData {
	d := &PlayerFamousFeedEventData{
		groupId:           groupId,
		totalFavorbaleNum: addFavorbaleNum,
		dayFavorbaleNum:   dayFavorbaleNum,
	}
	return d
}

func (d *PlayerFamousFeedEventData) GetTotalFavorbaleNum() int32 {
	return d.totalFavorbaleNum
}

func (d *PlayerFamousFeedEventData) GetDayFavorbaleNum() int32 {
	return d.dayFavorbaleNum
}

func (d *PlayerFamousFeedEventData) GetGroupId() int32 {
	return d.groupId
}

// 城战助威
type PlayerAllianceCheerEventData struct {
	groupId  int32
	giftType int32
	costGold int32
}

func CreatePlayerAllianceCheerEventData(groupId, giftType, costGold int32) *PlayerAllianceCheerEventData {
	d := &PlayerAllianceCheerEventData{
		groupId:  groupId,
		giftType: giftType,
		costGold: costGold,
	}
	return d
}

func (d *PlayerAllianceCheerEventData) GetCostGold() int32 {
	return d.costGold
}

func (d *PlayerAllianceCheerEventData) GetGroupId() int32 {
	return d.groupId
}

func (d *PlayerAllianceCheerEventData) GetGiftType() int32 {
	return d.giftType
}

// 商城促销特权激活
type PlayerShopDiscountEventData struct {
	groupId   int32
	typ       shopdiscounttypes.ShopDiscountType
	startTime int64
	endTime   int64
}

func CreatePlayerShopDiscountEventData(groupId int32, typ shopdiscounttypes.ShopDiscountType, startTime int64, endTime int64) *PlayerShopDiscountEventData {
	d := &PlayerShopDiscountEventData{
		groupId:   groupId,
		typ:       typ,
		startTime: startTime,
		endTime:   endTime,
	}
	return d
}

func (d *PlayerShopDiscountEventData) GetGroupId() int32 {
	return d.groupId
}

func (d *PlayerShopDiscountEventData) GetShopDiscountType() shopdiscounttypes.ShopDiscountType {
	return d.typ
}

func (d *PlayerShopDiscountEventData) GetStartTime() int64 {
	return d.startTime
}

func (d *PlayerShopDiscountEventData) GetEndTime() int64 {
	return d.endTime
}
