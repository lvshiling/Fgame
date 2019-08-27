package types

import (
	commonlog "fgame/fgame/common/log"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

type MarryEventType string

const (
	//求婚
	EventTypeMarryProposal MarryEventType = "marryProposal"
	//被求婚者决策
	EventTypeMarryProposalDeal MarryEventType = "marryProposalDeal"
	//玩家离婚
	EventTypeMarryDivorce MarryEventType = "marryDivorce"
	//被离婚者决策
	EventTypeMarryDivorceDeal MarryEventType = "marryDivorceDeal"
	//预定婚礼
	EventTypeMarryWed MarryEventType = "marryWedding"
	//婚礼开始
	EventTypeMarryWedStart MarryEventType = "marryWeddingStart"
	//婚礼取消
	EventTypeMarryWedCancle MarryEventType = "marryWeddingCancle"
	//婚礼结束
	EventTypeMarryWedEnd MarryEventType = "marryWeddingEnd"
	//婚戒替换
	EventTypeMarryRingReplace MarryEventType = "marryRingReplace"
	//婚车结束
	EventTypeMarryHunCheEnd MarryEventType = "marryHunCheEnd"
	//喜糖刷新时间
	EventTypeMarrySugarTime MarryEventType = "marrySugarTime"
	//场景婚礼开始
	EventTypeMarrySceneWedBegin MarryEventType = "marrySceneWedBegin"
	//场景婚礼结束
	EventTypeMarrySceneWedEnd MarryEventType = "marrySceneWedEnd"
	//玩家进入结婚场景
	EventTypePlayerEnterMarryScene MarryEventType = "playerEnterMarryScene"
	//豪气值排行变化
	EventTypeMarryHeriosmChange MarryEventType = "marryHeriosmChange"
	//婚戒归还
	EventTypeMarryRingGiveBack MarryEventType = "marryRingGiveBack"
	//婚戒培养升级
	EventTypeMarryRingFeedUpgrade MarryEventType = "marryFeedRingUpgrade"
	//结婚合服
	EventTypeMarryMergeServer MarryEventType = "marryMergeServer"
	//玩家名字改变
	EventTypeMarryPlayerNameChanged MarryEventType = "marryPlayerNameChanged"
	//婚宴场景玩家名字变化
	EventTypeMarryScenePlayerNameChanged MarryEventType = "marryScenePlayerNameChanged"
	//意外关服玩家婚礼补偿
	EventTypeMarryCloseServer MarryEventType = "marryCloseServer"
	//玩家婚期预定
	EventTypeMarryPreWed MarryEventType = "marryPreWed"
	//配偶决策婚期预定
	EventTypeMarryPreWedDeal MarryEventType = "marryPreWedDeal"
	//预定婚期返还
	EventTypeMarryPreWedGiveBack MarryEventType = "marryPreWedGiveBack"
	//有一方不在线婚礼取消归还
	EventTypeMarryStartOffline MarryEventType = "marryStartOffline"
	//配偶表白等级变化
	EventTypeMarryDevelopLevelChanged MarryEventType = "MarryDevelopLevelChanged"
	//定情信物激活
	EventTypeDingQingTokenActivite MarryEventType = "DingQingTokenActivite"
)

const (
	EventTypePlayerMarrySpouseChange        MarryEventType = "PlayerMarrySpouseChange"
	EventTypePlayerMarryWedStatusChange     MarryEventType = "PlayerMarryWedStatusChange"
	EventTypePlayerMarryRingChange          MarryEventType = "PlayerMarryRingChange"
	EventTypePlayerMarryDevelopLevelChanged MarryEventType = "PlayerMarryDevelopLevelChanged" //表白等级变化
	EventTypePlayerMarryDevelopExpAdd       MarryEventType = "PlayerMarryDevelopExpAdd"       //表白经验增加
	EventTypePlayerMarryDevelopLog          MarryEventType = "PlayerMarryDevelopLog"          //表白等级日志
	EventTypePlayerMarryDevelopExpLog       MarryEventType = "PlayerMarryDevelopExp"          //表白经验升级等级日志
	EventTypePlayerMarryBiaoBai             MarryEventType = "PlayerMarryBiaoBai "            //表白

)

type MarryRingReplaceEventData struct {
	ringType marrytypes.MarryRingType
	spouseId int64
}

func (d *MarryRingReplaceEventData) GetRingType() marrytypes.MarryRingType {
	return d.ringType
}

func (d *MarryRingReplaceEventData) GetSpouseId() int64 {
	return d.spouseId
}

func CreateMarryRingReplaceEventData(ringType marrytypes.MarryRingType, spouseId int64) *MarryRingReplaceEventData {
	d := &MarryRingReplaceEventData{
		ringType: ringType,
		spouseId: spouseId,
	}
	return d
}

type MarryWedStartEventData struct {
	owerId      int64
	period      int32
	grade       int32
	hunCheGrade int32
	sugarGrade  int32
	playerId    int64
	playerName  string
	playerRole  playertypes.RoleType
	playerSex   playertypes.SexType
	spouseId    int64
	spouseName  string
	spouseRole  playertypes.RoleType
	spouseSex   playertypes.SexType
	isFirst     bool
}

func (d *MarryWedStartEventData) GetOwerId() int64 {
	return d.owerId
}

func (d *MarryWedStartEventData) GetGrade() int32 {
	return d.grade
}

func (d *MarryWedStartEventData) GetHunCheGrade() int32 {
	return d.hunCheGrade
}

func (d *MarryWedStartEventData) GetSugarGrade() int32 {
	return d.sugarGrade
}

func (d *MarryWedStartEventData) GetPeriod() int32 {
	return d.period
}

func (d *MarryWedStartEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *MarryWedStartEventData) GetPlayerRole() playertypes.RoleType {
	return d.playerRole
}

func (d *MarryWedStartEventData) GetPlayerSex() playertypes.SexType {
	return d.playerSex
}

func (d *MarryWedStartEventData) GetSpouseId() int64 {
	return d.spouseId
}

func (d *MarryWedStartEventData) GetPlayerName() string {
	return d.playerName
}

func (d *MarryWedStartEventData) GetSpouseName() string {
	return d.spouseName
}

func (d *MarryWedStartEventData) GetSpouseRole() playertypes.RoleType {
	return d.spouseRole
}

func (d *MarryWedStartEventData) GetSpouseSex() playertypes.SexType {
	return d.spouseSex
}

func (d *MarryWedStartEventData) GetIsFirst() bool {
	return d.isFirst
}

func CreateMarryWedStartEventData(
	owerId int64,
	period int32,
	grade int32,
	hunCheGrade int32,
	sugarGrade int32,
	pl player.Player,
	spl player.Player,
	firstFlag bool) *MarryWedStartEventData {
	d := &MarryWedStartEventData{
		owerId:      owerId,
		period:      period,
		grade:       grade,
		hunCheGrade: hunCheGrade,
		sugarGrade:  sugarGrade,
		playerId:    pl.GetId(),
		playerName:  pl.GetName(),
		playerRole:  pl.GetRole(),
		playerSex:   pl.GetSex(),
		spouseId:    spl.GetId(),
		spouseName:  spl.GetName(),
		spouseRole:  spl.GetRole(),
		spouseSex:   spl.GetSex(),
		isFirst:     firstFlag,
	}
	return d
}

type MarryDivorceDealEventData struct {
	agree      bool
	spouseId   int64
	spouseName string
	marryList  []*MarryDivorceDealEventDataItem
}

func (d *MarryDivorceDealEventData) GetAgree() bool {
	return d.agree
}

func (d *MarryDivorceDealEventData) GetSpouseId() int64 {
	return d.spouseId
}

func (d *MarryDivorceDealEventData) GetSpouseName() string {
	return d.spouseName
}

func (d *MarryDivorceDealEventData) GetPreMarryList() []*MarryDivorceDealEventDataItem {
	return d.marryList
}

type MarryDivorceDealEventDataItem struct {
	grade          marrytypes.MarryBanquetSubTypeWed
	hunCheGrade    marrytypes.MarryBanquetSubTypeHunChe
	sugarGrade     marrytypes.MarryBanquetSubTypeSugar
	returnPlayerId int64
}

func (d *MarryDivorceDealEventDataItem) GetMarryBanquetSubTypeWed() marrytypes.MarryBanquetSubTypeWed {
	return d.grade
}

func (d *MarryDivorceDealEventDataItem) GetMarryBanquetSubTypeHunChe() marrytypes.MarryBanquetSubTypeHunChe {
	return d.hunCheGrade
}

func (d *MarryDivorceDealEventDataItem) GetMarryBanquetSubTypeSugar() marrytypes.MarryBanquetSubTypeSugar {
	return d.sugarGrade
}

func (d *MarryDivorceDealEventDataItem) GetReturnPlayerId() int64 {
	return d.returnPlayerId
}

func CreateMarryDivorceDealEventData(agree bool, spouseId int64, spouseName string, preMarryList []*MarryDivorceDealEventDataItem) *MarryDivorceDealEventData {
	d := &MarryDivorceDealEventData{
		agree:      agree,
		spouseId:   spouseId,
		spouseName: spouseName,
		marryList:  preMarryList,
	}
	return d
}

func CreateMarryDivorceDealEventDataItem(playerId int64, grade marrytypes.MarryBanquetSubTypeWed, hunCheGrade marrytypes.MarryBanquetSubTypeHunChe, sugarGrade marrytypes.MarryBanquetSubTypeSugar) *MarryDivorceDealEventDataItem {
	d := &MarryDivorceDealEventDataItem{
		grade:          grade,
		hunCheGrade:    hunCheGrade,
		sugarGrade:     sugarGrade,
		returnPlayerId: playerId,
	}
	return d
}

type MarryDivorceEventData struct {
	spouseId    int64
	divorceType marrytypes.MarryDivorceType
	marryList   []*MarryDivorceDealEventDataItem
}

func (d *MarryDivorceEventData) GetSpouseId() int64 {
	return d.spouseId
}

func (d *MarryDivorceEventData) GetDivorceType() marrytypes.MarryDivorceType {
	return d.divorceType
}

func (d *MarryDivorceEventData) GetPreMarryList() []*MarryDivorceDealEventDataItem {
	return d.marryList
}

func CreateMarryDivorceEventData(spouseId int64, divorceType marrytypes.MarryDivorceType, preMarryList []*MarryDivorceDealEventDataItem) *MarryDivorceEventData {
	d := &MarryDivorceEventData{
		spouseId:    spouseId,
		divorceType: divorceType,
		marryList:   preMarryList,
	}
	return d
}

type MarryProposalDealEventData struct {
	agree      bool
	proposalId int64
}

func (d *MarryProposalDealEventData) GetAgree() bool {
	return d.agree
}

func (d *MarryProposalDealEventData) GetProposalId() int64 {
	return d.proposalId
}

func CreateMarryProposalDealEventData(agree bool, proposalId int64) *MarryProposalDealEventData {
	d := &MarryProposalDealEventData{
		agree:      agree,
		proposalId: proposalId,
	}
	return d
}

type MarryProposalEventData struct {
	playerId int64
	spouseId int64
	ringType marrytypes.MarryRingType
}

func (d *MarryProposalEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *MarryProposalEventData) GetSpouseId() int64 {
	return d.spouseId
}

func (d *MarryProposalEventData) GetRingType() marrytypes.MarryRingType {
	return d.ringType
}

func CreateMarryProposalEventData(playerId int64, spouseId int64, ringType marrytypes.MarryRingType) *MarryProposalEventData {
	d := &MarryProposalEventData{
		playerId: playerId,
		spouseId: spouseId,
		ringType: ringType,
	}
	return d
}

type MarryWedEventData struct {
	period     int32
	spl        player.Player
	playerId   int64
	playerName string
	holdTime   string
}

func (d *MarryWedEventData) GetPeriod() int32 {
	return d.period
}

func (d *MarryWedEventData) GetPlayer() player.Player {
	return d.spl
}

func (d *MarryWedEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *MarryWedEventData) GetPlayerName() string {
	return d.playerName
}

func (d *MarryWedEventData) GetHoldTime() string {
	return d.holdTime
}

func CreateMarryWedEventData(period int32, spl player.Player, playerId int64, playerName string, holdTime string) *MarryWedEventData {
	d := &MarryWedEventData{
		period:     period,
		spl:        spl,
		playerId:   playerId,
		playerName: playerName,
		holdTime:   holdTime,
	}
	return d
}

type MarryPreWedEventData struct {
	period     int32
	playerName string
	spouseId   int64
	marryGrade *marrytypes.MarryGrade
}

func (d *MarryPreWedEventData) GetPeriod() int32 {
	return d.period
}

func (d *MarryPreWedEventData) GetPlayerName() string {
	return d.playerName
}

func (d *MarryPreWedEventData) GetSpouseId() int64 {
	return d.spouseId
}

func (d *MarryPreWedEventData) GetMarryGrade() *marrytypes.MarryGrade {
	return d.marryGrade
}

func CreateMarryPreWedEventData(period int32, playerName string, spouseId int64, marryGrade *marrytypes.MarryGrade) *MarryPreWedEventData {
	d := &MarryPreWedEventData{
		period:     period,
		playerName: playerName,
		spouseId:   spouseId,
		marryGrade: marryGrade,
	}
	return d
}

type MarryPreWedGiveBackEventData struct {
	isRefuse    bool
	grade       marrytypes.MarryBanquetSubTypeWed
	hunCheGrade marrytypes.MarryBanquetSubTypeHunChe
	sugarGrade  marrytypes.MarryBanquetSubTypeSugar
}

func (d *MarryPreWedGiveBackEventData) GetIsRefuse() bool {
	return d.isRefuse
}

func (d *MarryPreWedGiveBackEventData) GetGrade() marrytypes.MarryBanquetSubTypeWed {
	return d.grade
}

func (d *MarryPreWedGiveBackEventData) GetHunCheGrade() marrytypes.MarryBanquetSubTypeHunChe {
	return d.hunCheGrade
}

func (d *MarryPreWedGiveBackEventData) GetSugarGrade() marrytypes.MarryBanquetSubTypeSugar {
	return d.sugarGrade
}

func CreateMarryPreWedGiveBackEventData(isRefuse bool, grade marrytypes.MarryBanquetSubTypeWed, hunCheGrade marrytypes.MarryBanquetSubTypeHunChe, sugarGrade marrytypes.MarryBanquetSubTypeSugar) *MarryPreWedGiveBackEventData {
	d := &MarryPreWedGiveBackEventData{
		isRefuse:    isRefuse,
		grade:       grade,
		hunCheGrade: hunCheGrade,
		sugarGrade:  sugarGrade,
	}
	return d
}

type MarryPreWedGiftEventData struct {
	giftType       int32
	exp            int64
	expPoint       int64
	giftPlayerId   int64
	giftPlayerName string
}

func (m *MarryPreWedGiftEventData) GetGiftType() int32 {
	return m.giftType
}

func (m *MarryPreWedGiftEventData) GetExp() int64 {
	return m.exp
}

func (m *MarryPreWedGiftEventData) GetExpPoint() int64 {
	return m.expPoint
}

func (m *MarryPreWedGiftEventData) GetGiftPlayerId() int64 {
	return m.giftPlayerId
}

func (m *MarryPreWedGiftEventData) GetGiftPlayerName() string {
	return m.giftPlayerName
}

func CreateMarryPreWedGiftEventData(giftType int32, exp int64, expPoint int64, giftPlayerId int64, giftPlayerName string) *MarryPreWedGiftEventData {
	rst := &MarryPreWedGiftEventData{
		giftType:       giftType,
		exp:            exp,
		expPoint:       expPoint,
		giftPlayerId:   giftPlayerId,
		giftPlayerName: giftPlayerName,
	}
	return rst
}

type MarryGiveBackRingEventData struct {
	ringType marrytypes.MarryRingType
	peerName string
}

func (m *MarryGiveBackRingEventData) GetRingType() marrytypes.MarryRingType {
	return m.ringType
}
func (m *MarryGiveBackRingEventData) GetPeerName() string {
	return m.peerName
}
func CreateMarryGiveBackRingEventData(ringType marrytypes.MarryRingType, peerName string) *MarryGiveBackRingEventData {
	rst := &MarryGiveBackRingEventData{
		ringType: ringType,
		peerName: peerName,
	}
	return rst
}

// 配偶表白等级变化
type MarryDevelopLevelChangedEventData struct {
	changedPlayerId int64
	developLevel    int32
}

func (m *MarryDevelopLevelChangedEventData) GetChangedPlayerId() int64 {
	return m.changedPlayerId
}
func (m *MarryDevelopLevelChangedEventData) GetDevelopLevel() int32 {
	return m.developLevel
}

func CreateMarryDevelopLevelChangedEventData(changedPlayerId int64, developLevel int32) *MarryDevelopLevelChangedEventData {
	rst := &MarryDevelopLevelChangedEventData{
		changedPlayerId: changedPlayerId,
		developLevel:    developLevel,
	}
	return rst
}

// 玩家表白等级日志
type PlayerDevelopLevelLogEventData struct {
	beforeDevelopLevel int32
	curDevelopLevel    int32
	reason             commonlog.MarryLogReason
	reasonText         string
}

func CreatePlayerDevelopLevelLogEventData(beforeDevelopLevel, curDevelopLevel int32, reason commonlog.MarryLogReason, reasonText string) *PlayerDevelopLevelLogEventData {
	d := &PlayerDevelopLevelLogEventData{
		beforeDevelopLevel: beforeDevelopLevel,
		curDevelopLevel:    curDevelopLevel,
		reason:             reason,
		reasonText:         reasonText,
	}
	return d
}

func (d *PlayerDevelopLevelLogEventData) GetBeforeDevelopLevel() int32 {
	return d.beforeDevelopLevel
}

func (d *PlayerDevelopLevelLogEventData) GetCurDevelopLevel() int32 {
	return d.curDevelopLevel
}

func (d *PlayerDevelopLevelLogEventData) GetReason() commonlog.MarryLogReason {
	return d.reason
}

func (d *PlayerDevelopLevelLogEventData) GetReasonText() string {
	return d.reasonText
}

// 玩家表白经验根据物品增加日志
type PlayerDevelopExpLogEventData struct {
	beforeDevelopExp int32
	curDevelopExp    int32
	reason           commonlog.MarryLogReason
	reasonText       string
}

func CreatePlayerDevelopExpLogEventData(beforeExp int32, curExp int32, reason commonlog.MarryLogReason, reasonText string) *PlayerDevelopExpLogEventData {
	d := &PlayerDevelopExpLogEventData{
		beforeDevelopExp: beforeExp,
		curDevelopExp:    curExp,

		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerDevelopExpLogEventData) GetBeforeDevelopExp() int32 {
	return d.beforeDevelopExp
}

func (d *PlayerDevelopExpLogEventData) GetCurDevelopExp() int32 {
	return d.curDevelopExp
}

func (d *PlayerDevelopExpLogEventData) GetReason() commonlog.MarryLogReason {
	return d.reason
}

func (d *PlayerDevelopExpLogEventData) GetReasonText() string {
	return d.reasonText
}
