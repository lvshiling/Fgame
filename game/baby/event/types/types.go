package types

import (
	commonlog "fgame/fgame/common/log"
	babytypes "fgame/fgame/game/baby/types"
)

// 玩家宝宝（个人事件）
type BabyEventType string

const (
	EventTypeBabyBorn            BabyEventType = "BabyBorn"            //宝宝出生
	EventTypeBabyAdd                           = "BabyAdd"             //宝宝添加
	EventTypeBabyZhuanShi                      = "BabyZhuanShi"        //宝宝转世
	EventTypeBabyTalentChanged                 = "BabyTalentChanged "  //天赋技能变化
	EventTypeBabyUseToy                        = "BabyUseToy"          //宝宝使用玩具
	EventTypeBabyToyUplevel                    = "BabyToyUplevel"      //宝宝玩具升级
	EventTypeBabyLearnUplevel                  = "BabyLearnUplevel"    //宝宝读书升级
	EventTypeBabyBornNotice                    = "BabyBornNotice"      //宝宝出生提醒
	EventTypeBabyPregnantChanged               = "BabyPregnantChanged" //怀孕信息变化
	EventTypeBabyLearnUseItem                  = "BabyLearnUseItem"    //宝宝学习消耗物品
)

//配偶宝宝(全局事件)
type CoupleBabyEventType string

const (
	EventTypeCoupleBabyChanged CoupleBabyEventType = "CoupleBabyChanged" //配偶宝宝信息变化
)

//宝宝日志
type BabyLogEventType string

const (
	EventTypeBabyTalentLog BabyLogEventType = "BabyTalentLog" // 宝宝天赋日志
	EventTypeBabyLearnLog                   = "BabyLearnLog"  // 宝宝读书日志
)

//
type PlayerBabyTalentChangedEventData struct {
	babyId             int64
	oldEffectSkillList map[int32]int32
	newEffectSkillList map[int32]int32
	curTalentList      []*babytypes.TalentInfo
}

func CreatePlayerBabyTalentChangedEventData(babyId int64, oldList, newList map[int32]int32, talentList []*babytypes.TalentInfo) *PlayerBabyTalentChangedEventData {
	d := &PlayerBabyTalentChangedEventData{
		oldEffectSkillList: oldList,
		newEffectSkillList: newList,
		babyId:             babyId,
		curTalentList:      talentList,
	}

	return d
}

func (d *PlayerBabyTalentChangedEventData) GetOldEffectSkillList() map[int32]int32 {
	return d.oldEffectSkillList
}

func (d *PlayerBabyTalentChangedEventData) GetBabyId() int64 {
	return d.babyId
}

func (d *PlayerBabyTalentChangedEventData) GetNewEffectSkillList() map[int32]int32 {
	return d.newEffectSkillList
}

func (d *PlayerBabyTalentChangedEventData) GetTalentList() []*babytypes.TalentInfo {
	return d.curTalentList
}

//
type PlayerBabyToyChangedEventData struct {
	itemId     int32
	suitNumMap map[int32]int32
}

func CreatePlayerBabyToyChangedEventData(suitNumMap map[int32]int32, itemId int32) *PlayerBabyToyChangedEventData {
	d := &PlayerBabyToyChangedEventData{
		itemId:     itemId,
		suitNumMap: suitNumMap,
	}
	return d
}

func (d *PlayerBabyToyChangedEventData) GetSuitNumMap() map[int32]int32 {
	return d.suitNumMap
}

func (d *PlayerBabyToyChangedEventData) GetItemId() int32 {
	return d.itemId
}

//日志
type PlayerBabyTalentLogEventData struct {
	changedTalent []int32
	beforeTalent  []int32
	curTalent     []int32
	reason        commonlog.BabyLogReason
	reasonText    string
}

func CreatePlayerBabyTalentLogEventData(beforeTalent, curTalent, changedTalent []int32, reason commonlog.BabyLogReason, reasonText string) *PlayerBabyTalentLogEventData {
	d := &PlayerBabyTalentLogEventData{
		beforeTalent:  beforeTalent,
		curTalent:     curTalent,
		changedTalent: changedTalent,
		reason:        reason,
		reasonText:    reasonText,
	}
	return d
}

func (d *PlayerBabyTalentLogEventData) GetChangedTalent() []int32 {
	return d.changedTalent
}

func (d *PlayerBabyTalentLogEventData) GetBeforeTalent() []int32 {
	return d.beforeTalent
}

func (d *PlayerBabyTalentLogEventData) GetCurTalent() []int32 {
	return d.curTalent
}

func (d *PlayerBabyTalentLogEventData) GetReason() commonlog.BabyLogReason {
	return d.reason
}

func (d *PlayerBabyTalentLogEventData) GetReasonText() string {
	return d.reasonText
}

//读书日志
type PlayerBabyLevelLogEventData struct {
	beforeLevel int32
	curLevel    int32
	reason      commonlog.BabyLogReason
	reasonText  string
}

func CreatePlayerBabyLevelLogEventData(beforeLevel, curLevel int32, reason commonlog.BabyLogReason, reasonText string) *PlayerBabyLevelLogEventData {
	d := &PlayerBabyLevelLogEventData{
		beforeLevel: beforeLevel,
		curLevel:    curLevel,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerBabyLevelLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerBabyLevelLogEventData) GetCurLevel() int32 {
	return d.curLevel
}

func (d *PlayerBabyLevelLogEventData) GetReason() commonlog.BabyLogReason {
	return d.reason
}

func (d *PlayerBabyLevelLogEventData) GetReasonText() string {
	return d.reasonText
}

//出生提醒
type PlayerBabyBornMsgNoticeEventData struct {
	bornTime   int64
	noticeTime int64
}

func CreatePlayerBabyBornMsgNoticeEventData(bornTime, noticeTime int64) *PlayerBabyBornMsgNoticeEventData {
	d := &PlayerBabyBornMsgNoticeEventData{
		bornTime:   bornTime,
		noticeTime: noticeTime,
	}
	return d
}

func (d *PlayerBabyBornMsgNoticeEventData) GetBornTime() int64 {
	return d.bornTime
}

func (d *PlayerBabyBornMsgNoticeEventData) GetNoticeTime() int64 {
	return d.noticeTime
}

//宝宝学习消耗物品
type PlayerBabyLearnUseItemEventData struct {
	itemMap map[int32]int32
}

func CreatePlayerBabyLearnUseItemEventData(itemMap map[int32]int32) *PlayerBabyLearnUseItemEventData {
	d := &PlayerBabyLearnUseItemEventData{
		itemMap: itemMap,
	}
	return d
}

func (data *PlayerBabyLearnUseItemEventData) GetItemMap() map[int32]int32 {
	return data.itemMap
}
