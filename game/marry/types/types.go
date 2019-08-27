package types

import (
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"

	"github.com/golang/protobuf/proto"
)

const (
	//婚礼豪气值top
	HeroisTopLen = int(3)
	//婚礼豪气值列表
	HeroisListLen = int(10)
	//婚戒归还上限时间
	RingTime = int64(90 * common.SECOND)
	//预定婚期归还上限时间
	PreWedTime = int64(40 * common.SECOND)
	//订婚公告间隔时间(单位ms)
	WeddingInterval = int64(5 * common.SECOND)
)

//版本类型
type MarryHoutaiType int32

const (
	MarryHoutaiTypeCommon MarryHoutaiType = 1 + iota
	MarryHoutaiTypeCheep
	MarryHoutaiTypeExp
)

func (mrt MarryHoutaiType) Valid() bool {
	switch mrt {
	case MarryHoutaiTypeCommon,
		MarryHoutaiTypeCheep,
		MarryHoutaiTypeExp:
		return true
	}
	return false
}

var marryHouTaiMap = map[MarryHoutaiType]string{
	MarryHoutaiTypeCommon: "现在版",
	MarryHoutaiTypeCheep:  "廉价版",
	MarryHoutaiTypeExp:    "贵价版",
}

func (mrt MarryHoutaiType) String() string {
	return marryHouTaiMap[mrt]
}

//婚姻阶段状态
type MarryStatusType int32

const (
	//未婚
	MarryStatusTypeUnmarried MarryStatusType = 1 + iota
	//求婚成功
	MarryStatusTypeProposal
	//订婚
	MarryStatusTypeEngagement
	//举办过婚礼
	MarryStatusTypeMarried
	//离婚
	MarryStatusTypeDivorce
)

//婚戒类型
type MarryRingType int32

const (
	//青铜戒
	MarryRingTypeBronze MarryRingType = iota
	//紫金戒
	MarryRingTypePurple
	//龙凤戒
	MarryRingTypeLongFeng
	//免费青铜戒
	MarryRingTypeFreeBronze
	//免费紫金戒
	MarryRingTypeFreePurple
	//免费龙凤戒
	MarryRingTypeFreeLongFeng
	//贵价版青铜戒
	MarryRingTypeLuxuryBronze
	//贵价版紫金戒
	MarryRingTypeLuxuryPurple
	//贵价版龙凤戒
	MarryRingTypeLuxuryLongFeng
)

func (mrt MarryRingType) Valid() bool {
	switch mrt {
	case MarryRingTypeBronze,
		MarryRingTypePurple,
		MarryRingTypeLongFeng,
		MarryRingTypeFreeBronze,
		MarryRingTypeFreePurple,
		MarryRingTypeFreeLongFeng,
		MarryRingTypeLuxuryBronze,
		MarryRingTypeLuxuryPurple,
		MarryRingTypeLuxuryLongFeng:
		return true
	}
	return false
}

func (mrt MarryRingType) BetterThan(mrt2 MarryRingType) bool {
	//都无效
	if mrt < 0 && mrt2 < 0 {
		return false
	}

	if mrt >= 0 && mrt2 >= 0 {
		if mrt >= MarryRingTypeLuxuryBronze && mrt2 >= MarryRingTypeLuxuryBronze {
			return mrt > mrt2
		}
		if mrt <= MarryRingTypeLongFeng && mrt2 <= MarryRingTypeLongFeng {
			return mrt > mrt2
		}
		if mrt <= MarryRingTypeLuxuryBronze && mrt2 <= MarryRingTypeLuxuryBronze && mrt >= MarryRingTypeLongFeng && mrt2 >= MarryRingTypeLongFeng {
			return mrt > mrt2
		}
		return mrt < mrt2
	}
	//一个无效一个有效
	return mrt > mrt2
}

var marryRingHouTaiMap = map[MarryRingType]MarryHoutaiType{
	MarryRingTypeBronze:         MarryHoutaiTypeCommon,
	MarryRingTypePurple:         MarryHoutaiTypeCommon,
	MarryRingTypeLongFeng:       MarryHoutaiTypeCommon,
	MarryRingTypeFreeBronze:     MarryHoutaiTypeCheep,
	MarryRingTypeFreePurple:     MarryHoutaiTypeCheep,
	MarryRingTypeFreeLongFeng:   MarryHoutaiTypeCheep,
	MarryRingTypeLuxuryBronze:   MarryHoutaiTypeExp,
	MarryRingTypeLuxuryPurple:   MarryHoutaiTypeExp,
	MarryRingTypeLuxuryLongFeng: MarryHoutaiTypeExp,
}

func (mrt MarryRingType) HoutaiType() MarryHoutaiType {
	return marryRingHouTaiMap[mrt]
}

var marryRingMap = map[MarryRingType]string{
	MarryRingTypeBronze:         "青铜戒",
	MarryRingTypePurple:         "紫金戒",
	MarryRingTypeLongFeng:       "龙凤戒",
	MarryRingTypeFreeBronze:     "免费青铜戒",
	MarryRingTypeFreePurple:     "免费紫金戒",
	MarryRingTypeFreeLongFeng:   "免费龙凤戒",
	MarryRingTypeLuxuryBronze:   "贵价青铜戒",
	MarryRingTypeLuxuryPurple:   "贵价紫金戒",
	MarryRingTypeLuxuryLongFeng: "贵价龙凤戒",
}

func (mrt MarryRingType) String() string {
	return marryRingMap[mrt]
}

var marryRingItemSubMap = map[MarryRingType]itemtypes.ItemWedRingSubType{
	MarryRingTypeBronze:         itemtypes.ItemWedRingSubTypeBronze,
	MarryRingTypePurple:         itemtypes.ItemWedRingSubTypePurple,
	MarryRingTypeLongFeng:       itemtypes.ItemWedRingSubTypeLongFeng,
	MarryRingTypeFreeBronze:     itemtypes.ItemWedRingSubTypeBronzeCheap,
	MarryRingTypeFreePurple:     itemtypes.ItemWedRingSubTypePurpleCheap,
	MarryRingTypeFreeLongFeng:   itemtypes.ItemWedRingSubTypeLongFengCheap,
	MarryRingTypeLuxuryBronze:   itemtypes.ItemWedRingSubTypeBronzeLuxury,
	MarryRingTypeLuxuryPurple:   itemtypes.ItemWedRingSubTypePurpleLuxury,
	MarryRingTypeLuxuryLongFeng: itemtypes.ItemWedRingSubTypeLongFengLuxury,
}

func (mrt MarryRingType) ItemWedRingSubType() (ringSubType itemtypes.ItemWedRingSubType) {
	ringSubType, exist := marryRingItemSubMap[mrt]
	if !exist {
		return itemtypes.ItemWedRingSubType(-1)
	}
	return
}

var marryRingItemBanquetRingMap = map[MarryRingType]MarryBanquetSubTypeRing{
	MarryRingTypeBronze:         MarryBanquetSubTypeRingNormal,
	MarryRingTypePurple:         MarryBanquetSubTypeRingMid,
	MarryRingTypeLongFeng:       MarryBanquetSubTypeRingLuxury,
	MarryRingTypeFreeBronze:     MarryBanquetSubTypeRingNormal,
	MarryRingTypeFreePurple:     MarryBanquetSubTypeRingMid,
	MarryRingTypeFreeLongFeng:   MarryBanquetSubTypeRingLuxury,
	MarryRingTypeLuxuryBronze:   MarryBanquetSubTypeRingNormal,
	MarryRingTypeLuxuryPurple:   MarryBanquetSubTypeRingMid,
	MarryRingTypeLuxuryLongFeng: MarryBanquetSubTypeRingLuxury,
}

func (mrt MarryRingType) BanquetSubTypeRing() MarryBanquetSubTypeRing {
	ringSubType, exist := marryRingItemBanquetRingMap[mrt]
	if !exist {
		return MarryBanquetSubTypeRing(-1)
	}
	return ringSubType
}

type MarryBanquetType int32

const (
	//默认
	MarryBanquetTypeDefault MarryBanquetType = iota
	//婚宴
	MarryBanquetTypeWed
	//喜糖
	MarryBanquetTypeSugar
	//婚车
	MarryBanquetTypeHunChe
	//婚戒(仅前端使用)
	MarryBanquetTypeRing
)

func (mbt MarryBanquetType) Valid() bool {
	switch mbt {
	case MarryBanquetTypeDefault,
		MarryBanquetTypeWed,
		MarryBanquetTypeSugar,
		MarryBanquetTypeHunChe,
		MarryBanquetTypeRing:
		break
	default:
		return false
	}
	return true
}

type MarryBanquetSubType interface {
	SubType() int32
	Valid() bool
}

type MarryBanquetSubTypeFactory interface {
	CreateMarryBanquetSubType(subType int32) MarryBanquetSubType
}

type MarryBanquetSubTypeFactoryFunc func(subType int32) MarryBanquetSubType

func (mbsff MarryBanquetSubTypeFactoryFunc) CreateMarryBanquetSubType(subType int32) MarryBanquetSubType {
	return mbsff(subType)
}

//无子类型
type MarryBanquetCommonSubType int32

const (
	//默认子类型
	MarryBanquetCommonSubTypeDefault MarryBanquetCommonSubType = iota
)

func (tcst MarryBanquetCommonSubType) SubType() int32 {
	return int32(tcst)
}

func (tcst MarryBanquetCommonSubType) Valid() bool {
	switch tcst {
	case
		MarryBanquetCommonSubTypeDefault:
		return true
	}
	return false
}
func CreateMarryBanquetCommonSubType(subType int32) MarryBanquetSubType {
	return MarryBanquetCommonSubType(subType)
}

//婚礼子类型婚宴档次
type MarryBanquetSubTypeWed int32

const (
	//普通婚宴
	MarryBanquetSubTypeWedNormal MarryBanquetSubTypeWed = 1 + iota
	//中档婚宴
	MarryBanquetSubTypeWedMid
	//豪华婚宴
	MarryBanquetSubTypeWedLuxury
)

func (mwt MarryBanquetSubTypeWed) Valid() bool {
	switch mwt {
	case MarryBanquetSubTypeWedNormal,
		MarryBanquetSubTypeWedMid,
		MarryBanquetSubTypeWedLuxury:
		return true
	}
	return false
}

var marryBanquetSubTypeWedMap = map[MarryBanquetSubTypeWed]string{
	MarryBanquetSubTypeWedNormal: "普通婚宴",
	MarryBanquetSubTypeWedMid:    "中档婚宴",
	MarryBanquetSubTypeWedLuxury: "豪华婚宴",
}

func GetMarryWedTypeMap() map[MarryBanquetSubTypeWed]string {
	return marryBanquetSubTypeWedMap
}

func (tstr MarryBanquetSubTypeWed) SubType() int32 {
	return int32(tstr)
}

func CreateMarryBanquetWedSubType(subType int32) MarryBanquetSubType {
	return MarryBanquetSubTypeWed(subType)
}

func (tstr MarryBanquetSubTypeWed) String() string {
	return marryBanquetSubTypeWedMap[tstr]
}

//婚礼子类型喜糖档次
type MarryBanquetSubTypeSugar int32

const (
	//普通喜糖
	MarryBanquetSubTypeSugarNormal MarryBanquetSubTypeSugar = 1 + iota
	//中档喜糖
	MarryBanquetSubTypeSugarMid
	//豪华喜糖
	MarryBanquetSubTypeSugarLuxury
)

func (mst MarryBanquetSubTypeSugar) Valid() bool {
	switch mst {
	case MarryBanquetSubTypeSugarNormal,
		MarryBanquetSubTypeSugarMid,
		MarryBanquetSubTypeSugarLuxury:
		return true
	}
	return false
}

var marryBanquetSubTypeSugarMap = map[MarryBanquetSubTypeSugar]string{
	MarryBanquetSubTypeSugarNormal: "普通喜糖",
	MarryBanquetSubTypeSugarMid:    "中档喜糖",
	MarryBanquetSubTypeSugarLuxury: "豪华喜糖",
}

func GetMarrySugarTypeMap() map[MarryBanquetSubTypeSugar]string {
	return marryBanquetSubTypeSugarMap
}

func (tstr MarryBanquetSubTypeSugar) SubType() int32 {
	return int32(tstr)
}

func CreateMarryBanquetSugarSubType(subType int32) MarryBanquetSubType {
	return MarryBanquetSubTypeSugar(subType)
}

func (tstr MarryBanquetSubTypeSugar) String() string {
	return marryBanquetSubTypeSugarMap[tstr]
}

//婚礼子类型婚车档次
type MarryBanquetSubTypeHunChe int32

const (
	//普通婚车
	MarryBanquetSubTypeHunCheNormal MarryBanquetSubTypeHunChe = 1 + iota
	//中档婚车
	MarryBanquetSubTypeHunCheMid
	//豪华婚车
	MarryBanquetSubTypeHunCheLuxury
)

func (mst MarryBanquetSubTypeHunChe) Valid() bool {
	switch mst {
	case MarryBanquetSubTypeHunCheNormal,
		MarryBanquetSubTypeHunCheMid,
		MarryBanquetSubTypeHunCheLuxury:
		return true
	}
	return false
}

var marryBanquetSubTypeHunCheMap = map[MarryBanquetSubTypeHunChe]string{
	MarryBanquetSubTypeHunCheNormal: "普通婚车",
	MarryBanquetSubTypeHunCheMid:    "中档婚车",
	MarryBanquetSubTypeHunCheLuxury: "豪华婚车",
}

func GetMarryHunCheTypeMap() map[MarryBanquetSubTypeHunChe]string {
	return marryBanquetSubTypeHunCheMap
}

func (tstr MarryBanquetSubTypeHunChe) SubType() int32 {
	return int32(tstr)
}

func CreateMarryBanquetHunCheSubType(subType int32) MarryBanquetSubType {
	return MarryBanquetSubTypeHunChe(subType)
}

func (tstr MarryBanquetSubTypeHunChe) String() string {
	return marryBanquetSubTypeHunCheMap[tstr]
}

//婚礼子类型婚戒档次
type MarryBanquetSubTypeRing int32

const (
	//普通婚戒
	MarryBanquetSubTypeRingNormal MarryBanquetSubTypeRing = 1 + iota
	//中档婚戒
	MarryBanquetSubTypeRingMid
	//豪华婚戒
	MarryBanquetSubTypeRingLuxury
)

func (mst MarryBanquetSubTypeRing) Valid() bool {
	switch mst {
	case MarryBanquetSubTypeRingNormal,
		MarryBanquetSubTypeRingMid,
		MarryBanquetSubTypeRingLuxury:
		return true
	}
	return false
}

var marryBanquetSubTypeRingMap = map[MarryBanquetSubTypeRing]string{
	MarryBanquetSubTypeRingNormal: "普通婚戒",
	MarryBanquetSubTypeRingMid:    "中档婚戒",
	MarryBanquetSubTypeRingLuxury: "豪华婚戒",
}

func GetMarryRingTypeMap() map[MarryBanquetSubTypeRing]string {
	return marryBanquetSubTypeRingMap
}

func (tstr MarryBanquetSubTypeRing) SubType() int32 {
	return int32(tstr)
}

func CreateMarryBanquetRingSubType(subType int32) MarryBanquetSubType {
	return MarryBanquetSubTypeRing(subType)
}

func (tstr MarryBanquetSubTypeRing) String() string {
	return marryBanquetSubTypeRingMap[tstr]
}

//贺礼类型
type MarryGiftType int32

const (
	//物品
	MarryGiftTypeItem MarryGiftType = 1 + iota
	//银两
	MarryGiftTypeSilver
	//烟花
	MarryGiftTypeFireworks
)

func (mgt MarryGiftType) Valid() bool {
	switch mgt {
	case MarryGiftTypeItem,
		MarryGiftTypeSilver,
		MarryGiftTypeFireworks:
		return true
	}
	return false
}

//离婚类型
type MarryDivorceType int32

const (
	//强制离婚
	MarryDivorceTypeForce MarryDivorceType = 1 + iota
	//协议离婚
	MarryDivorceTypeConsent
)

func (mdt MarryDivorceType) Valid() bool {
	switch mdt {
	case MarryDivorceTypeConsent,
		MarryDivorceTypeForce:
		return true
	}
	return false
}

//婚期状态
type MarryWedStatusType int32

const (
	//未开始
	MarryWedStatusTypeNoStart MarryWedStatusType = 1 + iota
	//取消
	MarryWedStatusTypeCancle
	//进行中
	MarryWedStatusTypeOngoing
	//举办过
	MarryWedStatusTypeHeld
)

//婚戒状态
type MarryRingStatusType int32

const (
	//进行中
	MarryRingStatusTypeOngoing MarryRingStatusType = 1 + iota
	//失败
	MarryRingStatusTypeFail
	//成功
	MarryRingStatusTypeSucess
)

//婚期预定状态
type MarryPreWedStatusType int32

const (
	//进行中
	MarryPreWedStatusTypeOngoing MarryPreWedStatusType = 1 + iota
	//失败
	MarryPreWedStatusTypeFail
	//成功
	MarryPreWedStatusTypeSucess
)

//决策类型
type MarryResultType int32

const (
	//同意
	MarryResultTypeOk MarryResultType = 1 + iota
	//拒绝
	MarryResultTypeNo
)

func (t MarryResultType) Valid() bool {
	switch t {
	case MarryResultTypeOk,
		MarryResultTypeNo:
		return true
	}
	return false
}

//玩家婚宴状态(婚宴礼服穿卸使用)
type MarryWedStatusSelfType int32

const (
	//不在婚礼
	MarryWedStatusSelfTypeNo MarryWedStatusSelfType = iota
	//巡游
	MarryWedStatusSelfTypeCruise
	//酒席
	MarryWedStatusSelfTypeBanquet
)

//婚戒类型映射
var ItemWedRingMap = map[itemtypes.ItemWedRingSubType]string{
	itemtypes.ItemWedRingSubTypeBronze:   "青铜对戒",
	itemtypes.ItemWedRingSubTypePurple:   "紫金对戒",
	itemtypes.ItemWedRingSubTypeLongFeng: "龙凤对戒",
}

var (
	marryBanquetSubTypeFactoryMap = make(map[MarryBanquetType]MarryBanquetSubTypeFactory)
)

func CreateMarryBanquetSubType(typ MarryBanquetType, subType int32) MarryBanquetSubType {
	factory, ok := marryBanquetSubTypeFactoryMap[typ]
	if !ok {
		return CreateMarryBanquetCommonSubType(subType)
	}
	return factory.CreateMarryBanquetSubType(subType)
}

func init() {
	marryBanquetSubTypeFactoryMap[MarryBanquetTypeDefault] = MarryBanquetSubTypeFactoryFunc(CreateMarryBanquetCommonSubType)
	marryBanquetSubTypeFactoryMap[MarryBanquetTypeWed] = MarryBanquetSubTypeFactoryFunc(CreateMarryBanquetWedSubType)
	marryBanquetSubTypeFactoryMap[MarryBanquetTypeSugar] = MarryBanquetSubTypeFactoryFunc(CreateMarryBanquetSugarSubType)
	marryBanquetSubTypeFactoryMap[MarryBanquetTypeHunChe] = MarryBanquetSubTypeFactoryFunc(CreateMarryBanquetHunCheSubType)
	marryBanquetSubTypeFactoryMap[MarryBanquetTypeRing] = MarryBanquetSubTypeFactoryFunc(CreateMarryBanquetRingSubType)
}

type MarryGrade struct {
	Grade       MarryBanquetSubTypeWed
	HunCheGrade MarryBanquetSubTypeHunChe
	SugarGrade  MarryBanquetSubTypeSugar
}

func CreateMarryGrade(grade MarryBanquetSubTypeWed, hunCheGrade MarryBanquetSubTypeHunChe, sugardGrade MarryBanquetSubTypeSugar) *MarryGrade {
	d := &MarryGrade{
		Grade:       grade,
		HunCheGrade: hunCheGrade,
		SugarGrade:  sugardGrade,
	}
	return d
}

type MarryMsgRelateType int32

const (
	//喜帖
	MarryMsgRelateTypeWedCard MarryMsgRelateType = 1 + iota
	//婚帖
	MarryMsgRelateTypeMarry
	//婚期取消
	MarryMsgRelateTypeMarryCancle
)

type MarryPairMsgRelated struct {
	PlayerId         int64
	SpouseId         int64
	AllianceId       int64
	SpouseAllianceId int64
	Type             MarryMsgRelateType
	MsgShow          proto.Message
	Msg              proto.Message
}

func CreateMarryPairMsgRelated(playerId int64, spouseId int64, allianceId int64, spouseIdAllianceId int64, typ MarryMsgRelateType, showMsg proto.Message, msg proto.Message) *MarryPairMsgRelated {
	d := &MarryPairMsgRelated{
		PlayerId:         playerId,
		SpouseId:         spouseId,
		AllianceId:       allianceId,
		SpouseAllianceId: spouseIdAllianceId,
		Type:             typ,
		MsgShow:          showMsg,
		Msg:              msg,
	}
	return d
}

type MarryPushWedRelated struct {
	WedId            int64
	PlayerId         int64
	SpouseId         int64
	PlayerAllianceId int64
	SpouseAllianceId int64
	Msg              proto.Message
}

func CreateMarryPushWedRelated(wedId int64, playerId int64, spouseId int64, msg proto.Message, playerAllianceId int64, spouseAllianceId int64) *MarryPushWedRelated {
	d := &MarryPushWedRelated{
		WedId:            wedId,
		PlayerId:         playerId,
		SpouseId:         spouseId,
		Msg:              msg,
		PlayerAllianceId: playerAllianceId,
		SpouseAllianceId: spouseAllianceId,
	}
	return d
}

type MarryPreGiftType int32

const (
	MarryPreGiftTypeFlower MarryPreGiftType = 1 + iota
	MarryPreGiftTypeHeLi
)

var ItemMarryPreGiftTypeMap = map[MarryPreGiftType]string{
	MarryPreGiftTypeFlower: "鲜花",
	MarryPreGiftTypeHeLi:   "贺礼",
}

func (t MarryPreGiftType) Valid() bool {
	switch t {
	case MarryPreGiftTypeFlower,
		MarryPreGiftTypeHeLi:
		return true
	}
	return false
}

func (t MarryPreGiftType) String() string {
	return ItemMarryPreGiftTypeMap[t]
}
