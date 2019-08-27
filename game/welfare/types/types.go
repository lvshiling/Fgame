package types

import (
	lang "fgame/fgame/common/lang"
)

//开服活动类型
type OpenActivityType int32

const (
	OpenActivityTypeInvest         OpenActivityType = iota //0投资计划
	OpenActivityTypeFeedback                               //1返利
	OpenActivityTypeRank                                   //2排行
	OpenActivityTypeWelfare                                //3福利礼包
	OpenActivityTypeCycleCharge                            //4每日充值
	OpenActivityTypeAdvanced                               //5升阶返还(多天一场活动，每天的进阶类型固定)
	OpenActivityTypeDiscount                               //6限时礼包
	OpenActivityTypeMergeDrew                              //7抽奖
	OpenActivityTypeLongFeng                               //8龙凤呈祥
	OpenActivityTypeHuHu                                   //9虎虎生风(特殊类型)
	OpenActivityTypeRewards                                //10领奖（通用类型：每充值/消费多少领奖）
	OpenActivityTypeAdvancedRew                            //11进阶奖励（一天一场，策划可拼接进阶类型）
	OpenActivityTypeGroup                                  //12组合活动
	OpenActivityTypeMade                                   //13炼制
	OpenActivityTypeDevelop                                //14培养
	OpenActivityTypeSystemActivate                         //15系统
	OpenActivityTypeAlliance                               //16城战助威
	OpenActivityTypeBoatRace                               //17赛龙舟
	OpenActivityTypeShopDiscount                           //18商店特权
	OpenActivityTypeXiuxianBook                            //19修仙典籍
	OpenActivityTypeTongTianTa                             //20通天塔
)

func (t OpenActivityType) Valid() bool {
	switch t {
	case OpenActivityTypeInvest,
		OpenActivityTypeFeedback,
		OpenActivityTypeRank,
		OpenActivityTypeWelfare,
		OpenActivityTypeCycleCharge,
		OpenActivityTypeAdvanced,
		OpenActivityTypeDiscount,
		OpenActivityTypeMergeDrew,
		OpenActivityTypeLongFeng,
		OpenActivityTypeHuHu,
		OpenActivityTypeRewards,
		OpenActivityTypeAdvancedRew,
		OpenActivityTypeGroup,
		OpenActivityTypeMade,
		OpenActivityTypeDevelop,
		OpenActivityTypeSystemActivate,
		OpenActivityTypeAlliance,
		OpenActivityTypeBoatRace,
		OpenActivityTypeShopDiscount,
		OpenActivityTypeXiuxianBook,
		OpenActivityTypeTongTianTa:
		return true
	default:
		return false
	}
}

func (t OpenActivityType) SubType() int32 {
	return int32(t)
}

//投资计划子类型
type OpenActivityInvestSubType int32

const (
	OpenActivityInvestSubTypeLevel        OpenActivityInvestSubType = iota //投资计划
	OpenActivityInvestSubTypeServenDay                                     //七日投资
	OpenActivityInvestSubTypeNewLevel                                      //新等级投资计划
	OpenActivityInvestSubTypeNewServenDay                                  //新七日投资
)

func (t OpenActivityInvestSubType) Valid() bool {
	switch t {
	case OpenActivityInvestSubTypeLevel,
		OpenActivityInvestSubTypeServenDay,
		OpenActivityInvestSubTypeNewLevel,
		OpenActivityInvestSubTypeNewServenDay:
		return true
	default:
		return false
	}
}

func (t OpenActivityInvestSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityInvestSubType(subType int32) OpenActivitySubType {
	return OpenActivityInvestSubType(subType)
}

//返利子类型
type OpenActivityFeedbackSubType int32

const (
	OpenActivityFeedbackSubTypeCharge               OpenActivityFeedbackSubType = iota //充值返利
	OpenActivityFeedbackSubTypeCost                                                    //消费返利
	OpenActivityFeedbackSubTypeSingleChagre                                            //单笔充值
	OpenActivityFeedbackSubTypeGoldBowl                                                //聚宝盆
	OpenActivityFeedbackSubTypeCycleCharge                                             //连续充值
	OpenActivityFeedbackSubTypeReset                                                   //5.首充返利(没有使用)
	OpenActivityFeedbackSubTypeGoldLaBa                                                //元宝拉霸
	OpenActivityFeedbackSubTypeGoldPig                                                 //养金猪（升级类型的聚宝盆）
	OpenActivityFeedbackSubTypeChargeDevelop                                           //充值养鸡
	OpenActivityFeedbackSubTypeHouseInvest                                             //房产投资
	OpenActivityFeedbackSubTypeChargeReturn                                            //10首充回馈
	OpenActivityFeedbackSubTypeChargeDouble                                            //首充双倍
	OpenActivityFeedbackSubTypeChargeReturnLevel                                       //充值返还（按额度取比例）
	OpenActivityFeedbackSubTypeChargeReturnMultiple                                    //充值返还(每累充)
	OpenActivityFeedbackSubTypeSingleChagreMaxRew                                      //单笔充值（最近档次）
	OpenActivityFeedbackSubTypeHouseExtended                                           //房产活动 ---15

)

func (t OpenActivityFeedbackSubType) Valid() bool {
	switch t {
	case OpenActivityFeedbackSubTypeCharge,
		OpenActivityFeedbackSubTypeCost,
		OpenActivityFeedbackSubTypeSingleChagre,
		OpenActivityFeedbackSubTypeGoldBowl,
		OpenActivityFeedbackSubTypeCycleCharge,
		OpenActivityFeedbackSubTypeReset,
		OpenActivityFeedbackSubTypeGoldLaBa,
		OpenActivityFeedbackSubTypeGoldPig,
		OpenActivityFeedbackSubTypeChargeDevelop,
		OpenActivityFeedbackSubTypeHouseInvest,
		OpenActivityFeedbackSubTypeChargeReturn,
		OpenActivityFeedbackSubTypeChargeDouble,
		OpenActivityFeedbackSubTypeChargeReturnLevel,
		OpenActivityFeedbackSubTypeChargeReturnMultiple,
		OpenActivityFeedbackSubTypeSingleChagreMaxRew,
		OpenActivityFeedbackSubTypeHouseExtended:
		return true
	default:
		return false
	}
}

func (t OpenActivityFeedbackSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityFeedbackSubType(subType int32) OpenActivitySubType {
	return OpenActivityFeedbackSubType(subType)
}

//排行子类型
type OpenActivityRankSubType int32

const (
	OpenActivityRankSubTypeCharge           OpenActivityRankSubType = iota // 0充值排行
	OpenActivityRankSubTypeCost                                            // 1消费排行
	OpenActivityRankSubTypeMount                                           // 2坐骑排行
	OpenActivityRankSubTypeWing                                            // 3战翼排行
	OpenActivityRankSubTypeBodyshield                                      // 4护盾排行
	OpenActivityRankSubTypeLingyu                                          // 5领域排行09-0
	OpenActivityRankSubTypeShenfa                                          // 6身法排行
	OpenActivityRankSubTypeFeather                                         // 7护体仙羽排行
	OpenActivityRankSubTypeShield                                          // 8盾刺排行
	OpenActivityRankSubTypeCharm                                           // 9魅力排行
	OpenActivityRankSubTypeAnqi                                            // 10暗器排行
	OpenActivityRankSubTypeNumber                                          // 11次数排行
	OpenActivityRankSubTypeFaBao                                           // 12法宝排行
	OpenActivityRankSubTypeXianTi                                          // 13仙体排行
	OpenActivityRankSubTypeShiHunFan                                       // 14噬魂幡排行
	OpenActivityRankSubTypeTianMoTi                                        // 15天魔体排行
	OpenActivityRankSubTypeLingBing                                        // 16灵兵排行
	OpenActivityRankSubTypeLingQi                                          // 17灵骑排行
	OpenActivityRankSubTypeLingYi                                          // 18灵翼排行
	OpenActivityRankSubTypeLingBao                                         // 19灵宝排行
	OpenActivityRankSubTypeLingTi                                          // 20灵体排行
	OpenActivityRankSubTypeLingYu                                          // 21灵域排行
	OpenActivityRankSubTypeLingShen                                        // 22灵身排行
	OpenActivityRankSubTypeLevel                                           // 23等级排行
	OpenActivityRankSubTypeMarryDevelop                                    // 24表白排行
	OpenActivityRankSubTypeNumberDay                                       // 25次数排行（每日类型:一场活动对多场排行榜）
	OpenActivityRankSubTypeGoldEquipForce                                  // 26元神金装排行（战力)
	OpenActivityRankSubTypeLingTongForce                                   // 27灵童排行（战力)
	OpenActivityRankSubTypeDianXingForce                                   // 28点星排行（战力)
	OpenActivityRankSubTypeShenQiForce                                     // 29神器排行（战力)
	OpenActivityRankSubTypeMingGeForce                                     // 30命格排行（战力)
	OpenActivityRankSubTypeShengHenForce                                   // 31圣痕排行（战力)
	OpenActivityRankSubTypeZhenFaForce                                     // 32阵法排行（战力)
	OpenActivityRankSubTypeTuLongEquipForce                                // 33屠龙装排行（战力)
	OpenActivityRankSubTypeBabyForce                                       // 34宝宝排行（战力)
	OpenActivityRankSubTypeZhuanSheng                                      // 35转数排行
)

const (
	MinOpenActivityRankSubType = OpenActivityRankSubTypeCharge
	MaxOpenActivityRankSubType = OpenActivityRankSubTypeZhuanSheng
)

func (t OpenActivityRankSubType) Valid() bool {
	switch t {
	case
		OpenActivityRankSubTypeCharge,
		OpenActivityRankSubTypeCost,
		OpenActivityRankSubTypeMount,
		OpenActivityRankSubTypeWing,
		OpenActivityRankSubTypeBodyshield,
		OpenActivityRankSubTypeLingyu,
		OpenActivityRankSubTypeShenfa,
		OpenActivityRankSubTypeFeather,
		OpenActivityRankSubTypeShield,
		OpenActivityRankSubTypeCharm,
		OpenActivityRankSubTypeAnqi,
		OpenActivityRankSubTypeNumber,
		OpenActivityRankSubTypeFaBao,
		OpenActivityRankSubTypeXianTi,
		OpenActivityRankSubTypeShiHunFan,
		OpenActivityRankSubTypeTianMoTi,
		OpenActivityRankSubTypeLingBing,
		OpenActivityRankSubTypeLingQi,
		OpenActivityRankSubTypeLingYi,
		OpenActivityRankSubTypeLingBao,
		OpenActivityRankSubTypeLingTi,
		OpenActivityRankSubTypeLingYu,
		OpenActivityRankSubTypeLingShen,
		OpenActivityRankSubTypeLevel,
		OpenActivityRankSubTypeMarryDevelop,
		OpenActivityRankSubTypeNumberDay,
		OpenActivityRankSubTypeGoldEquipForce,
		OpenActivityRankSubTypeLingTongForce,
		OpenActivityRankSubTypeDianXingForce,
		OpenActivityRankSubTypeShenQiForce,
		OpenActivityRankSubTypeMingGeForce,
		OpenActivityRankSubTypeShengHenForce,
		OpenActivityRankSubTypeZhenFaForce,
		OpenActivityRankSubTypeTuLongEquipForce,
		OpenActivityRankSubTypeBabyForce,
		OpenActivityRankSubTypeZhuanSheng:
		return true
	default:
		return false
	}
}

func (t OpenActivityRankSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityRankSubType(subType int32) OpenActivitySubType {
	return OpenActivityRankSubType(subType)
}

//修仙典籍
type OpenActivityXiuxianBookSubType int32

const (
	OpenActivityXiuxianBookSubTypeEquipStrength OpenActivityXiuxianBookSubType = iota
	OpenActivityXiuxianBookSubTypeEquipOpenLight
	OpenActivityXiuxianBookSubTypeEquipUpStar
	OpenActivityXiuxianBookSubTypeLingTong
	OpenActivityXiuxianBookSubTypeDianXing
	OpenActivityXiuxianBookSubTypeShenQi
	OpenActivityXiuxianBookSubTypeSkillXinFa
	OpenActivityXiuxianBookSubTypeSkillDiHun
)

const (
	MinOpenActivityXiuxianBookSubType = OpenActivityXiuxianBookSubTypeEquipStrength
	MaxOpenActivityXiuxianBookSubType = OpenActivityXiuxianBookSubTypeSkillDiHun
)

func (t OpenActivityXiuxianBookSubType) Valid() bool {
	switch t {
	case OpenActivityXiuxianBookSubTypeEquipStrength,
		OpenActivityXiuxianBookSubTypeEquipOpenLight,
		OpenActivityXiuxianBookSubTypeEquipUpStar,
		OpenActivityXiuxianBookSubTypeLingTong,
		OpenActivityXiuxianBookSubTypeDianXing,
		OpenActivityXiuxianBookSubTypeShenQi,
		OpenActivityXiuxianBookSubTypeSkillXinFa,
		OpenActivityXiuxianBookSubTypeSkillDiHun:
		return true
	default:
		return false
	}
}

func (t OpenActivityXiuxianBookSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityXiuxianBookSubType(subType int32) OpenActivitySubType {
	return OpenActivityXiuxianBookSubType(subType)
}

//福利子类型
type OpenActivityWelfareSubType int32

const (
	OpenActivityWelfareSubTypeLogin      OpenActivityWelfareSubType = iota // 登录奖励
	OpenActivityWelfareSubTypeUpLevel                                      // 升级奖励
	OpenActivityWelfareSubTypeOnline                                       // 在线奖励
	OpenActivityWelfareSubTypeRealm                                        // 天劫塔冲刺
	OpenActivityWelfareSubTypeZhaunSheng                                   // 转生冲刺
)

func (t OpenActivityWelfareSubType) Valid() bool {
	switch t {
	case OpenActivityWelfareSubTypeLogin,
		OpenActivityWelfareSubTypeUpLevel,
		OpenActivityWelfareSubTypeOnline,
		OpenActivityWelfareSubTypeRealm,
		OpenActivityWelfareSubTypeZhaunSheng:
		return true
	default:
		return false
	}
}

func (t OpenActivityWelfareSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityWelfareSubType(subType int32) OpenActivitySubType {
	return OpenActivityWelfareSubType(subType)
}

//助威子类型
type OpenActivityAllianceSubType int32

const (
	OpenActivityAllianceSubTypeAlliance  OpenActivityAllianceSubType = iota //城战助威
	OpenActivityAllianceSubTypeWuLian                                       //1v1
	OpenActivityAllianceSubTypeNewWuLian                                    //1v1 新比武助力
)

func (t OpenActivityAllianceSubType) Valid() bool {
	switch t {
	case OpenActivityAllianceSubTypeAlliance,
		OpenActivityAllianceSubTypeWuLian,
		OpenActivityAllianceSubTypeNewWuLian:
		return true
	default:
		return false
	}
}

func (t OpenActivityAllianceSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityAllianceSubType(subType int32) OpenActivitySubType {
	return OpenActivityAllianceSubType(subType)
}

//升阶返利子类型
type OpenActivityAdvancedSubType int32

const (
	OpenActivityAdvancedSubTypeFeedback       OpenActivityAdvancedSubType = iota //升阶返利
	OpenActivityAdvancedSubTypeAdvancedDouble                                    // 升阶暴击日
	OpenActivityAdvancedSubTypeBlessFeedback                                     // 升阶祝福放送
)

func (t OpenActivityAdvancedSubType) Valid() bool {
	switch t {
	case OpenActivityAdvancedSubTypeFeedback,
		OpenActivityAdvancedSubTypeAdvancedDouble,
		OpenActivityAdvancedSubTypeBlessFeedback:
		return true
	default:
		return false
	}
}

func (t OpenActivityAdvancedSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityAdvancedSubType(subType int32) OpenActivitySubType {
	return OpenActivityAdvancedSubType(subType)
}

//升阶奖励子类型
type OpenActivityAdvancedRewSubType int32

const (
	OpenActivityAdvancedRewSubTypeRew          OpenActivityAdvancedRewSubType = iota //升阶奖励
	OpenActivityAdvancedRewSubTypePower                                              //升阶战力奖励(永久活动随功能开启)
	OpenActivityAdvancedRewSubTypeRewExtended                                        //升阶奖励(永久活动随功能开启)
	OpenActivityAdvancedRewSubTypeBlessCrit                                          //升阶暴击
	OpenActivityAdvancedRewSubTypeExpendReturn                                       //升阶消耗奖励
	OpenActivityAdvancedRewSubTypeTimesReturn                                        //5升阶次数奖励
	OpenActivityAdvancedRewSubTypeRewMax                                             //升阶奖励（只能领取活动初始进阶等级最近档次及更高等级的奖励）
)

func (t OpenActivityAdvancedRewSubType) Valid() bool {
	switch t {
	case OpenActivityAdvancedRewSubTypeRew,
		OpenActivityAdvancedRewSubTypePower,
		OpenActivityAdvancedRewSubTypeRewExtended,
		OpenActivityAdvancedRewSubTypeBlessCrit,
		OpenActivityAdvancedRewSubTypeExpendReturn,
		OpenActivityAdvancedRewSubTypeTimesReturn,
		OpenActivityAdvancedRewSubTypeRewMax:
		return true
	default:
		return false
	}
}

func (t OpenActivityAdvancedRewSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityAdvancedRewSubType(subType int32) OpenActivitySubType {
	return OpenActivityAdvancedRewSubType(subType)
}

//每日充值子类型
type OpenActivityCycleChargeSubType int32

const (
	OpenActivityCycleChargeSubTypeCharge                     OpenActivityCycleChargeSubType = iota //每日充值
	OpenActivityCycleChargeSubTypeSingleCharge                                                     //每日单笔充值
	OpenActivityCycleChargeSubTypeSingleChargeMaxRew                                               //每日单笔充值(只领取最高档次)
	OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple                                       //每日单笔充值(只领取最高档次,多次)
	OpenActivityCycleChargeSubTypeSingleChargeAllRew                                               //每日单笔（能领取所有满足充值档次的奖励，奖励次数叠加，但仅能领取一个档次，线上版本使用）

)

func (t OpenActivityCycleChargeSubType) Valid() bool {
	switch t {
	case OpenActivityCycleChargeSubTypeCharge,
		OpenActivityCycleChargeSubTypeSingleCharge,
		OpenActivityCycleChargeSubTypeSingleChargeMaxRew,
		OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple,
		OpenActivityCycleChargeSubTypeSingleChargeAllRew:
		return true
	default:
		return false
	}
}

func (t OpenActivityCycleChargeSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityCycleChargeSubType(subType int32) OpenActivitySubType {
	return OpenActivityCycleChargeSubType(subType)
}

//炼制子类型
type OpenActivityMadeSubType int32

const (
	OpenActivityMadeSubTypeResource OpenActivityMadeSubType = iota //炼制经验
)

func (t OpenActivityMadeSubType) Valid() bool {
	switch t {
	case OpenActivityMadeSubTypeResource:
		return true
	default:
		return false
	}
}

func (t OpenActivityMadeSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityMadeSubType(subType int32) OpenActivitySubType {
	return OpenActivityMadeSubType(subType)
}

//抽奖活动类型
type OpenActivityDrewSubType int32

const (
	OpenActivityDrewSubTypeTray       OpenActivityDrewSubType = iota //幸运转盘
	OpenActivityDrewSubTypeBombOre                                   //炸矿
	OpenActivityDrewSubTypeChargeDrew                                //充值抽奖
	OpenActivityDrewSubTypeCrazyBox                                  //疯狂宝箱
	OpenActivityDrewSubTypeSmashEgg                                  //砸金蛋
	OpenActivityDrewSubTypeCostDrew                                  //消费抽奖-----5
	OpenActivityDrewSubTypeBaoKuCrit                                 //装备宝库暴击日
	OpenActivityDrewSubTypeRewPools                                  //抽奖池
	OpenActivityDrewSubTypeSmelt                                     //冶炼
)

func (t OpenActivityDrewSubType) Valid() bool {
	switch t {
	case OpenActivityDrewSubTypeTray,
		OpenActivityDrewSubTypeBombOre,
		OpenActivityDrewSubTypeChargeDrew,
		OpenActivityDrewSubTypeCrazyBox,
		OpenActivityDrewSubTypeSmashEgg,
		OpenActivityDrewSubTypeCostDrew,
		OpenActivityDrewSubTypeBaoKuCrit,
		OpenActivityDrewSubTypeRewPools,
		OpenActivityDrewSubTypeSmelt:
		return true
	default:
		return false
	}
}

func (t OpenActivityDrewSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityDrewSubType(subType int32) OpenActivitySubType {
	return OpenActivityDrewSubType(subType)
}

//限时礼包
type OpenActivityDiscountSubType int32

const (
	OpenActivityDiscountSubTypeCommon     OpenActivityDiscountSubType = iota //限时礼包
	OpenActivityDiscountSubTypeZhuanSheng                                    //转生限时礼包
	OpenActivityDiscountSubTypeKanJia                                        //砍价礼包
	OpenActivityDiscountSubTypeTaoCan                                        //超值套餐
	OpenActivityDiscountSubTypeBeach                                         //沙滩挖宝
	OpenActivityDiscountSubTypeYunYin                                        //云隐仙人
)

func (t OpenActivityDiscountSubType) Valid() bool {
	switch t {
	case OpenActivityDiscountSubTypeCommon,
		OpenActivityDiscountSubTypeZhuanSheng,
		OpenActivityDiscountSubTypeKanJia,
		OpenActivityDiscountSubTypeTaoCan,
		OpenActivityDiscountSubTypeBeach,
		OpenActivityDiscountSubTypeYunYin:
		return true
	default:
		return false
	}
}

func (t OpenActivityDiscountSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityDiscountSubType(subType int32) OpenActivitySubType {
	return OpenActivityDiscountSubType(subType)
}

//返利奖励
type OpenActivityRewardsSubType int32

const (
	OpenActivityRewardsSubTypeCharge      OpenActivityRewardsSubType = iota //每充值多少领奖
	OpenActivityRewardsSubTypeCost                                          //每消费多少领奖
	OpenActivityRewardsSubTypeChargeLimit                                   //每充值多少领奖(全服次数限制)
)

func (t OpenActivityRewardsSubType) Valid() bool {
	switch t {
	case OpenActivityRewardsSubTypeCharge,
		OpenActivityRewardsSubTypeCost,
		OpenActivityRewardsSubTypeChargeLimit:
		return true
	default:
		return false
	}
}

func (t OpenActivityRewardsSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityRewardsSubType(subType int32) OpenActivitySubType {
	return OpenActivityRewardsSubType(subType)
}

//特殊类型
type OpenActivitySpecialSubType int32

const (
	OpenActivitySpecialSubTypeDrop      OpenActivitySpecialSubType = iota //boss掉落
	OpenActivitySpecialSubTypeCollect                                     //刷新采集物
	OpenActivitySpecialSubTypeFirstDrop                                   //boss首杀掉落
	OpenActivitySpecialSubTypeQiYu                                        //奇遇副本 ---3
	OpenActivitySpecialSubTypeGoal                                        //任务目标
)

func (t OpenActivitySpecialSubType) Valid() bool {
	switch t {
	case OpenActivitySpecialSubTypeDrop,
		OpenActivitySpecialSubTypeCollect,
		OpenActivitySpecialSubTypeFirstDrop,
		OpenActivitySpecialSubTypeQiYu,
		OpenActivitySpecialSubTypeGoal:
		return true
	default:
		return false
	}
}

func (t OpenActivitySpecialSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivitySpecialSubType(subType int32) OpenActivitySubType {
	return OpenActivitySpecialSubType(subType)
}

//系统激活子类型
type OpenActivitySystemActivateSubType int32

const (
	OpenActivitySystemActivateSubTypeLingYu OpenActivitySystemActivateSubType = iota //领域激活
)

func (t OpenActivitySystemActivateSubType) Valid() bool {
	switch t {
	case OpenActivitySystemActivateSubTypeLingYu:
		return true
	default:
		return false
	}
}

func (t OpenActivitySystemActivateSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivitySystemActivateSubType(subType int32) OpenActivitySubType {
	return OpenActivitySystemActivateSubType(subType)
}

//组合活动子类型
type OpenActivityGroupSubType int32

const (
	OpenActivityGroupSubTypeTimesRew     OpenActivityGroupSubType = iota //次数奖励
	OpenActivityGroupSubTypeCollectPoker                                 //收集扑克牌
)

func (t OpenActivityGroupSubType) Valid() bool {
	switch t {
	case OpenActivityGroupSubTypeTimesRew,
		OpenActivityGroupSubTypeCollectPoker:
		return true
	default:
		return false
	}
}

func (t OpenActivityGroupSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityGroupSubType(subType int32) OpenActivitySubType {
	return OpenActivityGroupSubType(subType)
}

//默认子类型
type OpenActivityDefaultSubType int32

const (
	OpenActivityDefaultSubTypeDefault OpenActivityDefaultSubType = iota // 会员
)

func (t OpenActivityDefaultSubType) Valid() bool {
	return true
}
func (t OpenActivityDefaultSubType) SubType() int32 {
	return int32(t)
}

func CreateOpenActivityDefaultSubType(subType int32) OpenActivitySubType {
	return OpenActivityDefaultSubType(subType)
}

//子类型接口
type OpenActivitySubType interface {
	Valid() bool
	SubType() int32
}

type OpenActivitySubTypeFactory interface {
	CreateOpenActivitySubType(subType int32) OpenActivitySubType
}

type OpenActivitySubTypeFactoryFunc func(subType int32) OpenActivitySubType

func (oaff OpenActivitySubTypeFactoryFunc) CreateOpenActivitySubType(subType int32) OpenActivitySubType {
	return oaff(subType)
}

var (
	openActivitySubTypeFactoryMap = map[OpenActivityType]OpenActivitySubTypeFactory{}
)

func CreateOpenActivitySubType(typ OpenActivityType, subType int32) OpenActivitySubType {
	factory, ok := openActivitySubTypeFactoryMap[typ]
	if !ok {
		return nil
	}

	return factory.CreateOpenActivitySubType(subType)
}

func init() {
	openActivitySubTypeFactoryMap[OpenActivityTypeInvest] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityInvestSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeFeedback] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityFeedbackSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeRank] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityRankSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeWelfare] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityWelfareSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeCycleCharge] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityCycleChargeSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeAdvanced] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityAdvancedSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeDiscount] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDiscountSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeMergeDrew] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDrewSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeLongFeng] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDefaultSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeHuHu] = OpenActivitySubTypeFactoryFunc(CreateOpenActivitySpecialSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeRewards] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityRewardsSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeAdvancedRew] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityAdvancedRewSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeGroup] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityGroupSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeMade] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityMadeSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeDevelop] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDefaultSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeSystemActivate] = OpenActivitySubTypeFactoryFunc(CreateOpenActivitySystemActivateSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeAlliance] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityAllianceSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeBoatRace] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDefaultSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeShopDiscount] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityDefaultSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeXiuxianBook] = OpenActivitySubTypeFactoryFunc(CreateOpenActivityXiuxianBookSubType)
	openActivitySubTypeFactoryMap[OpenActivityTypeTongTianTa] = OpenActivitySubTypeFactoryFunc(CreateTongTianTaSubType)
}

//开服活动时间类型
type OpenTimeType int32

const (
	OpenTimeTypeNotTimeliness       OpenTimeType = iota // 无时效性---0
	OpenTimeTypeOpenActivity                            // 开服活动类型
	OpenTimeTypeSchedule                                // 指定时间段类型
	OpenTimeTypeMerge                                   // 合服活动类型
	OpenTimeTypeXunHuan                                 // 循环活动
	OpenTimeTypeMergeXunHuan                            // 合服循环活动---5
	OpenTimeTypeOpenActivityNoMerge                     //开服活动类型(不合服)
	OpenTimeTypeWeek                                    //每周活动
	OpenTimeTypeMonth                                   //每月活动
)

func (t OpenTimeType) Valid() bool {
	switch t {
	case OpenTimeTypeNotTimeliness,
		OpenTimeTypeOpenActivity,
		OpenTimeTypeSchedule,
		OpenTimeTypeMerge,
		OpenTimeTypeXunHuan,
		OpenTimeTypeMergeXunHuan,
		OpenTimeTypeOpenActivityNoMerge,
		OpenTimeTypeWeek,
		OpenTimeTypeMonth:
		return true
	default:
		return false
	}
}

//进阶返利类型
type AdvancedType int32

const (
	AdvancedTypeMount      AdvancedType = iota // 0坐骑进阶
	AdvancedTypeWing                           // 1战翼进阶
	AdvancedTypeAnqi                           // 2暗器进阶
	AdvancedTypeBodyshield                     // 3护盾进阶
	AdvancedTypeLingyu                         // 4领域进阶
	AdvancedTypeShenfa                         // 5身法进阶
	AdvancedTypeShield                         // 6盾刺进阶
	AdvancedTypeFaBao                          // 7法宝进阶
	AdvancedTypeXianTi                         // 8仙体进阶
	AdvancedTypeFeather                        // 9仙羽进阶
	AdvancedTypeShiHunFan                      // 10噬魂幡进阶
	AdvancedTypeTianMoTi                       // 11天魔体进阶
)

const (
	AdvancedTypeLingBing AdvancedType = 100 + iota // 灵兵进阶
	AdvancedTypeLingQi                             // 灵骑进阶
	AdvancedTypeLingYi                             // 灵翼进阶
	AdvancedTypeLingBao                            // 灵宝进阶
	AdvancedTypeLingTi                             // 灵体进阶
	AdvancedTypeLingYu                             // 105灵域进阶
	AdvancedTypeLingShen                           // 灵身进阶
)

func (t AdvancedType) Valid() bool {
	switch t {
	case
		AdvancedTypeMount,
		AdvancedTypeWing,
		AdvancedTypeAnqi,
		AdvancedTypeBodyshield,
		AdvancedTypeLingyu,
		AdvancedTypeShenfa,
		AdvancedTypeShield,
		AdvancedTypeFaBao,
		AdvancedTypeXianTi,
		AdvancedTypeFeather,
		AdvancedTypeShiHunFan,
		AdvancedTypeTianMoTi,
		AdvancedTypeLingBing,
		AdvancedTypeLingQi,
		AdvancedTypeLingYi,
		AdvancedTypeLingBao,
		AdvancedTypeLingTi,
		AdvancedTypeLingYu,
		AdvancedTypeLingShen:
		return true
	default:
		return false
	}
}

var (
	advancedTypeMap = map[AdvancedType]string{
		AdvancedTypeMount:      "坐骑",
		AdvancedTypeWing:       "战翼",
		AdvancedTypeAnqi:       "暗器",
		AdvancedTypeBodyshield: "护体盾",
		AdvancedTypeLingyu:     "领域",
		AdvancedTypeShenfa:     "身法",
		AdvancedTypeShield:     "神盾尖刺",
		AdvancedTypeFaBao:      "法宝",
		AdvancedTypeXianTi:     "仙体",
		AdvancedTypeFeather:    "仙羽",
		AdvancedTypeShiHunFan:  "噬魂幡进阶",
		AdvancedTypeTianMoTi:   "天魔体进阶",
		AdvancedTypeLingBing:   "灵兵进阶",
		AdvancedTypeLingQi:     "灵骑进阶",
		AdvancedTypeLingYi:     "灵翼进阶",
		AdvancedTypeLingBao:    "灵宝进阶",
		AdvancedTypeLingTi:     "灵体进阶",
		AdvancedTypeLingYu:     "灵域进阶",
		AdvancedTypeLingShen:   "灵身进阶",
	}
)

func (t AdvancedType) String() string {
	return advancedTypeMap[t]
}

// 抽奖随机类型
type DrewResultType int32

const (
	DrewResultTypeDrop        DrewResultType = iota //掉落包
	DrewResultTypeRatioFirst                        //系数1
	DrewResultTypeRatioSecond                       //系数2
)

//返还类型
type ChargeReturnType int32

const (
	ChargeReturnTypeBindGold ChargeReturnType = iota //绑元
	ChargeReturnTypeGold                             //元宝
)

func (t ChargeReturnType) Valid() bool {
	switch t {
	case ChargeReturnTypeBindGold,
		ChargeReturnTypeGold:
		return true
	default:
		return false
	}
}

//礼包商城类型
type ZhuanShengGiftValue1Type int32

const (
	ZhuanShengGiftValue1TypeDefault ZhuanShengGiftValue1Type = iota //默认
	ZhuanShengGiftValue1TypePoint                                   //积分商城
)

func (t ZhuanShengGiftValue1Type) Valid() bool {
	switch t {
	case ZhuanShengGiftValue1TypeDefault,
		ZhuanShengGiftValue1TypePoint:
		return true
	default:
		return false
	}
}

//邮件内容类型
type EmailContentType int32

const (
	EmailContentTypeDefault            EmailContentType = iota //默认
	EmailContentTypeMarryDevelop                               //表白排行榜
	EmailContentTypeMarryDevelopSpouse                         //表白排行榜配偶
	EmailContentTypeCharm                                      //魅力排行榜
	EmailContentTypeCharmSpouse                                //魅力排行榜配偶
	EmailContentTypeBoatRaceForce                              //赛龙舟排行榜
)

//关联lang.LangCode
var emailContentTypeMap = map[EmailContentType]map[bool]lang.LangCode{
	EmailContentTypeDefault: map[bool]lang.LangCode{
		true:  lang.EmailOpenActivityRankContent,
		false: lang.OpenActivityRankContentNotOnLevel,
	},
	EmailContentTypeMarryDevelop: map[bool]lang.LangCode{
		true:  lang.OpenActivityRankContentMarryDevelop,
		false: lang.OpenActivityRankContentMarryDevelopNotOnLevel,
	},
	EmailContentTypeMarryDevelopSpouse: map[bool]lang.LangCode{
		true:  lang.OpenActivityRankContentMarryDevelopSpouse,
		false: lang.OpenActivityRankContentMarryDevelopSpouseNotOnLevel,
	},
	EmailContentTypeCharm: map[bool]lang.LangCode{
		true:  lang.OpenActivityRankContentCharm,
		false: lang.OpenActivityRankContentCharmNotOnLevel,
	},
	EmailContentTypeCharmSpouse: map[bool]lang.LangCode{
		true:  lang.OpenActivityRankContentCharmSpouse,
		false: lang.OpenActivityRankContentCharmSpouseNotOnLevel,
	},
	EmailContentTypeBoatRaceForce: map[bool]lang.LangCode{
		true:  lang.OpenActivityRankContentBoatRaceForce,
		false: lang.OpenActivityRankContentBoatRaceForceNotOnLevel,
	},
}

func (t EmailContentType) ConvertToRankEmailContentLangCodeType(isOnLevel bool) lang.LangCode {
	typ, ok := emailContentTypeMap[t][isOnLevel]
	if !ok {
		typ = lang.EmailOpenActivityRankContent
	}
	return typ
}

//通天塔子类型
type TongTianTaSubType int32

const (
	TongTianTaSubTypeLingTong TongTianTaSubType = iota
	TongTianTaSubTypeMingGe
	TongTianTaSubTypeTuLong
	TongTianTaSubTypeShengHen
	TongTianTaSubTypeZhenFa
	TongTianTaSubTypeBaby
	TongTianTaSubTypeDianXing
)

func (t TongTianTaSubType) Valid() bool {
	switch t {
	case TongTianTaSubTypeLingTong,
		TongTianTaSubTypeMingGe,
		TongTianTaSubTypeTuLong,
		TongTianTaSubTypeShengHen,
		TongTianTaSubTypeZhenFa,
		TongTianTaSubTypeBaby,
		TongTianTaSubTypeDianXing:
		return true
	default:
		return false
	}
}

func (t TongTianTaSubType) SubType() int32 {
	return int32(t)
}

func CreateTongTianTaSubType(subType int32) OpenActivitySubType {
	return TongTianTaSubType(subType)
}

var TongTianTaSubTypeMap = map[TongTianTaSubType]string{
	TongTianTaSubTypeLingTong: "灵童战力",
	TongTianTaSubTypeMingGe:   "命格战力",
	TongTianTaSubTypeTuLong:   "屠龙战力",
	TongTianTaSubTypeShengHen: "圣痕战力",
	TongTianTaSubTypeZhenFa:   "阵法战力",
	TongTianTaSubTypeBaby:     "宝宝战力",
	TongTianTaSubTypeDianXing: "点星战力",
}

func (t TongTianTaSubType) String() string {
	return TongTianTaSubTypeMap[t]
}

const (
	MinTongTianTaSubType = TongTianTaSubTypeLingTong
	MaxTongTianTaSubType = TongTianTaSubTypeDianXing
)

type ReceiveType int32

const (
	ReceiveTypeOnce ReceiveType = iota
	ReceiveTypeTen              //十连
)

func (t ReceiveType) Valid() bool {
	switch t {
	case ReceiveTypeOnce,
		ReceiveTypeTen:
		return true
	default:
		return false
	}
}

var receiveTypeToInt32Map = map[ReceiveType]int32{
	ReceiveTypeOnce: int32(1),
	ReceiveTypeTen:  int32(10),
}

func (t ReceiveType) ToInt32() int32 {
	return receiveTypeToInt32Map[t]
}
