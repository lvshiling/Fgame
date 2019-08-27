package log

type LevelLogReason int32

const (
	LevelLogReasonGM LevelLogReason = iota + 1
	LevelLogReasonQuestReward
	LevelLogReasonMonsterKilled
	LevelLogReasonMonsterKilledExtra
	LevelLogReasonSoulRuinsRewChapter
	LevelLogReasonSoulRuinsSweep
	LevelLogReasonSoulRuinsFinishAll
	LevelLogReasonSoulRuins
	LevelLogReasonXianfuSaodangRewards
	LevelLogReasonRealmToKillRew
	LevelLogReasonEmperorWorship
	LevelLogReasonMoonloveRankRew
	LevelLogReasonMoonloveTickRew
	LevelLogReasonAllianceSceneDoorReward
	LevelLogReasonSecretCardStarRew
	LevelLogReasonSecretCardFinishAllStarRew
	LevelLogReasonFoundResource
	LevelLogReasonTuMoFinishAll
	LevelLogReasonFourGodOpenBox
	LevelLogReasonXianfuFinishAllRew
	LevelLogReasonBuff
	LevelLogReasonEmailAttachment
	LevelLogReasonGemGambleDrop
	LevelLogReasonSecretCardFinish
	LevelLogReasonSecretCardFinishAll
	LevelLogReasonSoulRuinsSweepDrop
	LevelLogReasonSoulRuinsFinishAllDrop
	LevelLogReasonBoxGet
	LevelLogReasonChessGet
	LevelLogReasonFirstCharge
	LevelLogReasonOpenActivityRew
	LevelLogReasonGiftCode
	LevelLogReasonArenaExpTree
	LevelLogReasonEatUplevelDan
	LevelLogReasonTuLongKillBossRew
	LevelLogReasonCollectRew
	LevelLogReasonTuMoImmediateFinish
	LevelLogReasonEmperorReward
	LevelLogReasonMaterialRew
	LevelLogReasonLivenessStarRew
	LevelLogReasonFriendNoticeRew
	LevelLogReasonFriendAddRew
	LevelLogReasonUnrealBossDrop
	LevelLogReasonDailyQuestReward
	LevelLogReasonDailyQuestFinishAll
	LevelLogReasonBaGuaToKillRew
	LevelLogReasonOutlandBossDrop
	LevelLogReasonSongBuTingRew
	LevelLogReasonInventoryResourceUse
	LevelLogReasonTeamCopyReward
	LevelLogReasonQuizAnswer
	LevelLogReasonActivityTickRew
	LevelLogReasonOpenActivityMadeRes
	LevelLogReasonFeiShengSanGong
	LevelLogReasonFeiShengSanGongGive
	LevelLogReasonHongBaoSnatch
	LevelLogReasonEquipBaoKuGet
	LevelLogReasonEquipBaoKuLuckyBoxGet
	LevelLogReasonMarryPreGift
	LevelLogReasonQuestQiYuRew
	LevelLogReasonXianTaoCommit
	LevelLogReasonGuideRew
	LevelLogReasonShenYuRoundRew
	LevelLogReasonMajorSaoDangGet
	LevelLogReasonMaterialSaodangGet
	LevelLogReasonArenapvpRew
	LevelLogReasonEatSuperUplevelDan
	LevelLogReasonRewPoolsDrew
	LevelLogReasonRingBaoKuGet
	LevelLogReasonRingBaoKuLuckyBoxGet
)

func (zslr LevelLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	levelLogReasonMap = map[LevelLogReason]string{
		LevelLogReasonGM:                         "gm修改",
		LevelLogReasonQuestReward:                "任务完成奖励,任务id:%d",
		LevelLogReasonMonsterKilled:              "杀怪奖励,怪物id:%d",
		LevelLogReasonMonsterKilledExtra:         "杀怪额外奖励,怪物id:%d",
		LevelLogReasonSoulRuinsRewChapter:        "帝陵遗迹章节奖励领取chapter:%d, typ:%d",
		LevelLogReasonSoulRuinsSweep:             "帝陵遗迹扫荡,扫荡关卡chapter:%d typ:%d level:%d 扫荡次数:%d",
		LevelLogReasonSoulRuinsFinishAll:         "帝陵遗迹一键完成chapter:%d typ:%d level:%d 次数:%d",
		LevelLogReasonSoulRuins:                  "帝陵遗迹挑战成功奖励chapter:%d typ:%d level:%d",
		LevelLogReasonXianfuSaodangRewards:       "秘境仙府扫荡奖励，仙府id:%d,类型:%d",
		LevelLogReasonRealmToKillRew:             "天劫塔前往击杀奖励,层数:%d",
		LevelLogReasonEmperorWorship:             "膜拜帝王领取奖励",
		LevelLogReasonMoonloveRankRew:            "月下情缘排行榜奖励，名次:%d,类型:%d",
		LevelLogReasonMoonloveTickRew:            "月下情缘定时奖励",
		LevelLogReasonAllianceSceneDoorReward:    "城战城门奖励,仙盟id:%d,门:%d",
		LevelLogReasonSecretCardStarRew:          "天机牌领取运势箱奖励",
		LevelLogReasonSecretCardFinishAllStarRew: "天机牌一键完成剩余运势箱奖励",
		LevelLogReasonFoundResource:              "资源找回获取,资源id:%d,次数:%d,波数：%d",
		LevelLogReasonTuMoFinishAll:              "屠魔任务一键完成,次数:%d",
		LevelLogReasonFourGodOpenBox:             "四神遗迹开启宝箱奖励,钥匙数:%d",
		LevelLogReasonXianfuFinishAllRew:         "秘境仙府一键完成奖励,仙府id:%d,类型:%d",
		LevelLogReasonBuff:                       "buff添加经验,buffId:%d",
		LevelLogReasonEmailAttachment:            "邮件附件奖励",
		LevelLogReasonGemGambleDrop:              "赌石掉落经验",
		LevelLogReasonSecretCardFinish:           "天机牌任务完成",
		LevelLogReasonSecretCardFinishAll:        "天机牌一键完成奖励",
		LevelLogReasonSoulRuinsSweepDrop:         "帝陵遗迹扫荡掉落,扫荡关卡chapter:%d typ:%d level:%d,扫荡次数:%d",
		LevelLogReasonSoulRuinsFinishAllDrop:     "帝陵遗迹一键完成掉落chapter:%d typ:%d level:%d,次数:%d",
		LevelLogReasonBoxGet:                     "宝箱开启奖励",
		LevelLogReasonChessGet:                   "苍龙棋局奖励，类型：%s",
		LevelLogReasonFirstCharge:                "首冲奖励",
		LevelLogReasonOpenActivityRew:            "开服活动奖励,类型:%d，子类型：%d",
		LevelLogReasonGiftCode:                   "礼包兑换奖励",
		LevelLogReasonArenaExpTree:               "3v3竞技场经验树",
		LevelLogReasonEatUplevelDan:              "等级直升丹奖励",
		LevelLogReasonTuLongKillBossRew:          "跨服屠龙击杀boss奖励",
		LevelLogReasonCollectRew:                 "采集物奖励",
		LevelLogReasonTuMoImmediateFinish:        "屠魔任务直接完成",
		LevelLogReasonEmperorReward:              "帝王奖励",
		LevelLogReasonMaterialRew:                "材料副本奖励，type:%s",
		LevelLogReasonLivenessStarRew:            "活跃度领取宝箱奖励",
		LevelLogReasonFriendNoticeRew:            "好友推送反馈奖励,type:%d",
		LevelLogReasonFriendAddRew:               "添加好友奖励",
		LevelLogReasonUnrealBossDrop:             "幻境boss掉落",
		LevelLogReasonDailyQuestReward:           "完成日环任务奖励,类型：%s",
		LevelLogReasonDailyQuestFinishAll:        "一键完成%s任务奖励",
		LevelLogReasonBaGuaToKillRew:             "八卦秘境前往击杀奖励,层数:%d",
		LevelLogReasonOutlandBossDrop:            "外域boss掉落",
		LevelLogReasonSongBuTingRew:              "元宝送不停奖励",
		LevelLogReasonInventoryResourceUse:       "背包资源使用获得",
		LevelLogReasonTeamCopyReward:             "组队副本奖励",
		LevelLogReasonQuizAnswer:                 "仙尊问答奖励,玩家等级:%d ,答题结果:%d",
		LevelLogReasonActivityTickRew:            "活动定时奖励,活动类型:%d",
		LevelLogReasonOpenActivityMadeRes:        "运营活动炼制经验奖励",
		LevelLogReasonFeiShengSanGong:            "飞升散功",
		LevelLogReasonFeiShengSanGongGive:        "玩家飞升散功分享经验",
		LevelLogReasonHongBaoSnatch:              "抢送红包获得",
		LevelLogReasonEquipBaoKuGet:              "装备宝库奖励，等级：%d，转数：%d",
		LevelLogReasonEquipBaoKuLuckyBoxGet:      "装备宝库幸运宝箱奖励",
		LevelLogReasonMarryPreGift:               "结婚游车被赠送获得",
		LevelLogReasonQuestQiYuRew:               "奇遇任务奖励，qiyuId:%d",
		LevelLogReasonXianTaoCommit:              "提交仙桃奖励",
		LevelLogReasonGuideRew:                   "引导副本奖励，类型：%s",
		LevelLogReasonShenYuRoundRew:             "神域之战奖励，参赛轮：%d",
		LevelLogReasonMajorSaoDangGet:            "夫妻副本扫荡%d次奖励,type:%s,fubenId:%d",
		LevelLogReasonMaterialSaodangGet:         "材料副本扫荡%d次奖励，type:%s",
		LevelLogReasonArenapvpRew:                "比武大会奖励，获胜：%v， 类型：%s",
		LevelLogReasonEatSuperUplevelDan:         "等级超级直升丹奖励",
		LevelLogReasonRewPoolsDrew:               "仙人指路-奖池抽奖",
		LevelLogReasonRingBaoKuGet:               "特戒宝库探索获得",
		LevelLogReasonRingBaoKuLuckyBoxGet:       "特戒宝库幸运宝箱奖励",
	}
)

func (slr LevelLogReason) String() string {
	return levelLogReasonMap[slr]
}
