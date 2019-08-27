package types

import "fgame/fgame/core/storage"

//玩家数据管理器类型
type PlayerDataManagerType int32

//统一处理
const (
	//玩家场景数据类型
	PlayerSceneDataManagerType PlayerDataManagerType = iota
	//玩家属性数据类型
	PlayerPropertyDataManagerType
	//玩家技能数据类型
	PlayerSkillDataManagerType
	//玩家buff数据类型
	PlayerBuffDataManagerType
	//玩家仓库数据类型
	PlayerInventoryDataManagerType
	//玩家食丹数据类型
	PlayerDanDataManagerType
	//玩家坐骑数据类型
	PlayerMountDataManagerType
	//玩家战翼数据类型
	PlayerWingDataManagerType
	//玩家护体盾数据类型
	PlayerBShieldDataManagerType
	//玩家称号数据类型
	PlayerTitleDataManagerType
	//任务数据
	PlayerQuestDataManagerType
	//玩家时装数据类型
	PlayerFashionDataManagerType
	//玩家商店购买道具类型
	PlayerShopDataManagerType
	//玩家兵魂类型
	PlayerWeaponDataManagerType
	//玩家pk数据
	PlayerPkDataManagerType
	//玩家帝魂数据
	PlayerSoulDataManagerType
	//玩家功能开启数据类型
	PlayerFuncOpenDataManagerType
	//玩家境界数据
	PlayerRealmDataManagerType
	//玩家好友数据
	PlayerFriendDataManagerType
	//玩家缓存数据
	PlayerCacheDataManagerType
	//玩家绝学数据
	PlayerJueXueDataManagerType
	//玩家心法数据
	PlayerXinFaDataManagerType
	//玩家宝石数据
	PlayerGemDataManagerType
	//玩家邮件数据
	PlayerEmailDataManagerType
	//玩家帝陵遗迹数据
	PlayerSoulRuinsDataManagerType
	//玩家秘境仙府数据
	PlayerXianfuDtatManagerType
	//玩家抢龙椅数据
	PlayerEmperorDataManagerType
	//玩家月下情缘数据
	PlayerMoonloveDataManagerType
	//玩家活动数据
	PlayerActivityDataManagerType
	//玩家仙盟数据
	PlayerAllianceDataManagerType
	//玩家组队数据
	PlayerTeamDataManagerType
	//玩家天机牌数据
	PlayerSecretCardDataManagerType
	//玩家资源回收数据
	PlayerFoundDataManagerType
	//玩家神龙现世数据
	PlayerDragonDataManagerType
	//玩家镖车数据
	PlayerTransportationType
	//玩家四神遗迹数据
	PlayerFourGodDataManagerType
	//玩家结婚数据
	PlayerMarryDataManagerType
	//玩家身法数据
	PlayerShenfaDataManagerType
	//玩家领域数据
	PlayerLingyuDataManagerType
	//玩家元神金装数据
	PlayerGoldEquipDataManagerType
	//玩家苍龙棋局数据
	PlayerChessDataManagerType
	//玩家灵池争夺
	PlayerOneArenaDataManagerType
	//玩家跨服数据
	PlayerCrossDataManagerType
	//玩家复活数据
	PlayerReliveDataManagerType
	//玩家竞技场数据
	PlayerArenaDataManagerType
	//玩家福利数据
	PlayerWelfareDataManagerType
	//玩家充值数据
	PlayerChargeDataManagerType
	//玩家双休数据
	PlayerMajorDataManagerType
	//玩家暗器数据
	PlayerAnqiDataManagerType
	//玩家血池数据
	PlayerXueChiDataManagerType
	//玩家会员数据
	PlayerHuiYuanDataManagerType
	//玩家VIP数据
	PlayerVipDataManagerType
	//玩家幸运符数据
	PlayerLuckyDataManagerType
	//玩家补偿数据
	PlayerCompensateDataManagerType
	//玩家戮仙刃数据
	PlayerMassacreDataManagerType
	//玩家打宝塔数据
	PlayerTowerDataManagerType
	//玩家天书数据
	PlayerTianShuDataManagerType
	//玩家个人BOSS
	PlayerMyBossDataManagerType
	//玩家系统技能
	PlayerSystemSkillDataManagerType
	//玩家附加系统
	PlayerAdditionSysDataManagerType
	//玩家仇人数据
	PlayerFoeDataManagerType
	//玩家法宝数据
	PlayerFaBaoDataManagerType
	//玩家幻境BOSS数据
	PlayerUnrealBossDataManagerType
	//玩家血盾数据
	PlayerXueDunDataManagerType
	//玩家材料副本数据
	PlayerMaterialDataManagerType
	//玩家仙体数据类型
	PlayerXianTiDataManagerType
	//玩家活跃度数据类型
	PlayerLivenessDataManagerType
	//玩家挂机
	PlayerGuaJiManagerType
	//玩家外域BOSS数据
	PlayerOutlandBossDataManagerType
	//玩家八卦秘境
	PlayerBaGuaDataManagerType
	//玩家元宝送不停
	PlayerSongBuTingDataManagerType
	//玩家组队副本数据
	PlayerTeamCopyDataManagerType
	//玩家金银密窟
	PlayerDenseWatDataManagerType
	//玩家点星系统
	PlayerDianXingDataManagerType
	//玩家衣橱数据类型
	PlayerWardrobeDataManagerType
	//玩家天魔数据
	PlayerTianMoDataManagerType
	//玩家噬魂幡数据
	PlayerShiHunFanDataManagerType
	//玩家灵童养成数据类型
	PlayerLingTongDevDataManagerType
	//玩家灵童数据类型
	PlayerLingTongDataManagerType
	//玩家飞升数据类型
	PlayerFeiShengDataManagerType
	//玩家神魔战场数据
	PlayerShenMoWarDataManagerType
	//玩家抢红包数据
	PlayerHongBaoDataManagerType
	//玩家聊天数据
	PlayerChatDataManagerType
	//玩家至尊称号数据
	PlayerSupremeTitleDataManagerType
	//玩家装备宝库数据
	PlayerEquipBaoKuDataManagerType
	//玩家圣痕数据
	PlayerShengHenDataManagerType
	//玩家屠龙装备数据
	PlayerTuLongEquipDataManagerType
	//玩家命格数据
	PlayerMingGeDataManagerType
	//玩家寻宝数据
	PlayerHuntDataManagerType
	//玩家神器数据
	PlayerShenQiDataManagerType
	//英灵普玩家数据
	PlayerYingLingPuManagerType
	//玩家阵法数据
	PlayerZhenFaDataManagerType
	//玩家宝宝数据
	PlayerBabyDataManagerType
	//玩家交易数据
	PlayerTradeDataManagerType
	//玩家仙桃大会数据
	PlayerXianTaoDataManagerType
	//玩家神域数据
	PlayerShenYuDataManagerType
	//玩家房子数据
	PlayerHouseDataManagerType
	//玩家玉玺之战数据
	PlayerYuXiDataManagerType
	//玩家周卡数据
	PlayerWeekDataManagerType
	//玩家八卦符石数据
	PlayerFuShiDataManagerType
	//玩家物品技能数据
	PlayerItemSkillDataManagerType
	//玩家商城促销数据
	PlayerShopDiscountDataManagerType
	//玩家泣血枪
	PlayerQiXueDataManagerType
	//玩家创世之战数据
	PlayerChuangShiDataManagerType
	//玩家跨服pvp数据
	PlayerArenapvpDataManagerType
	//玩家结义数据
	PlayerJieYiDataManagerType
	//玩家逆付费数据
	PlayerFeedbackFeeDataManagerType
	//玩家无双神器数据
	PlayerWushuangWeaponDataManagerType
	//大力丸数据
	PlayerDaLiWanDataManagerType
	//珍惜boss
	PlayerZhenXiDataManagerType
	//玩家boss数据
	PlayerWorldbossManagerType
	//玩家仙尊特权卡数据
	PlayerXianZunCardManagerType
	//玩家上古之灵数据
	PlayerShangguzhilingDataManagerType
	//玩家特戒数据
	PlayerRingDataManagerType
)

var (
	dataManagerMap = map[PlayerDataManagerType]string{
		PlayerSceneDataManagerType:     "玩家场景数据类型",
		PlayerPropertyDataManagerType:  "玩家属性数据类型",
		PlayerSkillDataManagerType:     "玩家技能数据类型",
		PlayerBuffDataManagerType:      "玩家buff数据类型",
		PlayerInventoryDataManagerType: "玩家仓库数据类型",
		PlayerDanDataManagerType:       "玩家食丹数据类型",
		PlayerMountDataManagerType:     "玩家坐骑数据类型",
		PlayerWingDataManagerType:      "玩家战翼数据类型",
		PlayerBShieldDataManagerType:   "玩家护体盾数据类型",
		PlayerTitleDataManagerType:     "玩家称号数据类型",

		PlayerQuestDataManagerType: "任务数据",

		PlayerFashionDataManagerType: "玩家时装数据类型",

		PlayerShopDataManagerType: "玩家商店购买道具类型",

		PlayerWeaponDataManagerType: "玩家兵魂类型",

		PlayerPkDataManagerType: "玩家pk数据",

		PlayerSoulDataManagerType: "玩家帝魂数据",

		PlayerFuncOpenDataManagerType: "玩家功能开启数据类型",

		PlayerRealmDataManagerType: "玩家境界数据",

		PlayerFriendDataManagerType: "玩家好友数据",

		PlayerCacheDataManagerType: "玩家缓存数据",

		PlayerJueXueDataManagerType: "玩家绝学数据",

		PlayerXinFaDataManagerType: "玩家心法数据",

		PlayerGemDataManagerType: "玩家宝石数据",

		PlayerEmailDataManagerType: "玩家邮件数据",

		PlayerSoulRuinsDataManagerType: "玩家帝陵遗迹数据",

		PlayerXianfuDtatManagerType: "玩家秘境仙府数据",

		PlayerEmperorDataManagerType: "玩家抢龙椅数据",

		PlayerMoonloveDataManagerType: "玩家月下情缘数据",

		PlayerActivityDataManagerType: "玩家活动数据",

		PlayerAllianceDataManagerType: "玩家仙盟数据",

		PlayerTeamDataManagerType:       "玩家组队数据",
		PlayerSecretCardDataManagerType: "玩家天机牌数据",

		PlayerFoundDataManagerType: "玩家资源回收数据",

		PlayerDragonDataManagerType: "玩家神龙现世数据",

		PlayerTransportationType: "玩家镖车数据",

		PlayerFourGodDataManagerType: "玩家四神遗迹数据",

		PlayerMarryDataManagerType: "玩家结婚数据",

		PlayerShenfaDataManagerType: "玩家身法数据",

		PlayerLingyuDataManagerType: "玩家领域数据",

		PlayerGoldEquipDataManagerType: "玩家元神金装数据",

		PlayerChessDataManagerType: "玩家苍龙棋局数据",

		PlayerOneArenaDataManagerType: "玩家灵池争夺",

		PlayerCrossDataManagerType:   "玩家跨服数据",
		PlayerReliveDataManagerType:  "玩家复活数据",
		PlayerArenaDataManagerType:   "玩家竞技场数据",
		PlayerWelfareDataManagerType: "玩家福利数据",
		PlayerChargeDataManagerType:  "玩家充值数据",

		PlayerMajorDataManagerType:          "玩家双休数据",
		PlayerAnqiDataManagerType:           "玩家暗器数据",
		PlayerXueChiDataManagerType:         "玩家血池数据",
		PlayerHuiYuanDataManagerType:        "玩家会员数据",
		PlayerVipDataManagerType:            "玩家VIP数据",
		PlayerLuckyDataManagerType:          "玩家幸运符数据",
		PlayerCompensateDataManagerType:     "玩家补偿数据",
		PlayerMassacreDataManagerType:       "玩家戮仙刃数据",
		PlayerTowerDataManagerType:          "玩家打宝塔数据",
		PlayerTianShuDataManagerType:        "玩家天书数据",
		PlayerMyBossDataManagerType:         "玩家个人数据",
		PlayerSystemSkillDataManagerType:    "玩家系统技能数据",
		PlayerAdditionSysDataManagerType:    "玩家附加系统",
		PlayerFoeDataManagerType:            "玩家仇人数据",
		PlayerFaBaoDataManagerType:          "玩家法宝数据",
		PlayerUnrealBossDataManagerType:     "玩家幻境BOSS数据",
		PlayerXueDunDataManagerType:         "玩家血盾数据",
		PlayerMaterialDataManagerType:       "玩家材料副本数据",
		PlayerXianTiDataManagerType:         "玩家仙体数据类型",
		PlayerLivenessDataManagerType:       "玩家活跃度类型",
		PlayerGuaJiManagerType:              "玩家挂机类型",
		PlayerOutlandBossDataManagerType:    "玩家外域BOSS数据",
		PlayerBaGuaDataManagerType:          "玩家八卦秘境",
		PlayerSongBuTingDataManagerType:     "玩家元宝送不停",
		PlayerTeamCopyDataManagerType:       "玩家组队副本数据",
		PlayerDenseWatDataManagerType:       "玩家金银密窟数据",
		PlayerDianXingDataManagerType:       "玩家点星系统",
		PlayerWardrobeDataManagerType:       "玩家衣橱数据",
		PlayerTianMoDataManagerType:         "玩家天魔数据",
		PlayerShiHunFanDataManagerType:      "玩家噬魂幡数据",
		PlayerLingTongDevDataManagerType:    "玩家灵童养成数据",
		PlayerLingTongDataManagerType:       "玩家灵童数据类型",
		PlayerFeiShengDataManagerType:       "玩家飞升数据类型",
		PlayerShenMoWarDataManagerType:      "玩家神魔战场数据类型",
		PlayerHongBaoDataManagerType:        "玩家抢红包数据",
		PlayerChatDataManagerType:           "玩家聊天数据",
		PlayerSupremeTitleDataManagerType:   "玩家至尊称号数据",
		PlayerEquipBaoKuDataManagerType:     "玩家装备宝库数据",
		PlayerShengHenDataManagerType:       "玩家圣痕数据",
		PlayerTuLongEquipDataManagerType:    "玩家屠龙装备数据",
		PlayerMingGeDataManagerType:         "玩家命格数据",
		PlayerHuntDataManagerType:           "玩家寻宝数据",
		PlayerShenQiDataManagerType:         "玩家神器数据",
		PlayerYingLingPuManagerType:         "英灵普数据",
		PlayerZhenFaDataManagerType:         "玩家阵法数据",
		PlayerBabyDataManagerType:           "玩家宝宝数据",
		PlayerTradeDataManagerType:          "玩家交易数据",
		PlayerXianTaoDataManagerType:        "玩家仙桃大会数据",
		PlayerShenYuDataManagerType:         "玩家神域数据",
		PlayerHouseDataManagerType:          "玩家房子数据",
		PlayerYuXiDataManagerType:           "玩家玉玺之战数据",
		PlayerWeekDataManagerType:           "玩家周卡数据",
		PlayerFuShiDataManagerType:          "玩家八卦符石数据",
		PlayerItemSkillDataManagerType:      "玩家物品技能数据",
		PlayerShopDiscountDataManagerType:   "玩家商城促销数据",
		PlayerQiXueDataManagerType:          "玩家泣血枪数据",
		PlayerChuangShiDataManagerType:      "玩家创世之战数据",
		PlayerArenapvpDataManagerType:       "玩家跨服pvp数据",
		PlayerJieYiDataManagerType:          "玩家结义数据",
		PlayerFeedbackFeeDataManagerType:    "玩家逆付费数据",
		PlayerWushuangWeaponDataManagerType: "玩家无双神器数据",
		PlayerDaLiWanDataManagerType:        "玩家大力丸数据",
		PlayerZhenXiDataManagerType:         "珍稀boss数据",
		PlayerWorldbossManagerType:          "玩家boss数据",
		PlayerXianZunCardManagerType:        "玩家仙尊特权卡数据",
		PlayerShangguzhilingDataManagerType: "玩家上古之灵数据",
		PlayerRingDataManagerType:           "玩家特戒数据",
	}
)

func (t PlayerDataManagerType) String() string {
	return dataManagerMap[t]
}

//玩家数据库数据
type PlayerDataEntity interface {
	storage.Entity
	GetPlayerId() int64
}

//玩家业务数据
type PlayerDataPersistanceObject interface {
	storage.PersistanceObject
	GetPlayerId() int64
}

type PlayerOnlineState int32

const (
	PlayerOnlineStateOffline = iota
	PlayerOnlineStateOnline
)
