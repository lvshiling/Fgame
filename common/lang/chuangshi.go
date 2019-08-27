package lang

const (
	ChuangShiOutOfTime LangCode = ChuangShiBase + iota
	ChuangShiAlreadyBaoMing
	ChuangShiHadShenWangSignUp
	ChuangShiNotShenWang
	ChuangShiNotChengZhu
	ChuangShiNotMengZhu
	ChuangShiNotCamp
	ChuangShiAlreadyCamp
	ChuangShiNotMyRew
	ChuangShiJianSheFullLevel
	ChuangShiJianSheProgressing
	ChuangShiJianSheSigning
	ChuangShiJianSheVoting
	ChuangShiPositionLevelAlreadyTop
	ChuangShiWeiWangNotEnough
	ChuangShiAlreadyReceive
	ChuangShiTemplateNotExist
	ChuangShiNotEnterCityTime
	ChuangShiJiFenNoEnough
	ChuangShiJoinCampFailed
	ChuangShiJianSheSkillSetFailed
	ChuangShiCityNotExist
	ChuangShiCityNotFuShu
	ChuangShiHasChengZhu
	ChuangShiBuChangMailTitle
	ChuangShiBuChangMailContent
	ChuangShiLogGongChengTarget
	ChuangShiLogGongChengBeTarget
	ChuangShiLogPayScheduleCamp
	ChuangShiLogPayScheduleCity
	ChuangShiLogRenMingChengZhu
	ChuangShiLogRenMingShenWang
	ChuangShiLogJianSheCity
	ChuangShiLogJianSheTianQiSet
	ChuangShiCityNotNearCity
	ChuangShiCityTargetFullNum
	ChuangShiCityTargetHadChooes
	ChuangShiLogGongChengWin
	ChuangShiLogGongChengFailed
	ChuangShiJianSheSkillSetSuccess
	ChuangShiCityTargetNotTime
	ChuangShiCampPayScheduleNotResouces
	ChuangShiLogRenMingMengZhuChanged
	ChuangShiShenWangAlreadyVote
	ChuangShiCampChangedNotAllow
	ChuangShiSceneFinishMailTitle
	ChuangShiSceneFinishWinMailContent
	ChuangShiSceneFinishFailedMailContent
)

var (
	chuangShiLangMap = map[LangCode]string{
		ChuangShiOutOfTime:                  "不在创世之战报名时间",
		ChuangShiAlreadyBaoMing:             "玩家已经报名创世之战",
		ChuangShiHadShenWangSignUp:          "已经报名神王竞选",
		ChuangShiShenWangAlreadyVote:        "已经投票",
		ChuangShiNotShenWang:                "不是神王",
		ChuangShiNotChengZhu:                "不是城主",
		ChuangShiNotMengZhu:                 "不是盟主",
		ChuangShiNotCamp:                    "不是阵营成员",
		ChuangShiAlreadyCamp:                "已经是阵营成员",
		ChuangShiNotMyRew:                   "没有可领取的个人奖励",
		ChuangShiHasChengZhu:                "已经是城主了",
		ChuangShiJianSheFullLevel:           "城防建设已经满级",
		ChuangShiJianSheProgressing:         "城防建设中",
		ChuangShiJianSheSigning:             "神王竞选报名中",
		ChuangShiJianSheVoting:              "神王竞选投票中",
		ChuangShiPositionLevelAlreadyTop:    "已是最高官职",
		ChuangShiWeiWangNotEnough:           "升职所需威望不足",
		ChuangShiAlreadyReceive:             "您已领取过该时装",
		ChuangShiTemplateNotExist:           "模板不存在",
		ChuangShiNotEnterCityTime:           "当前不是城池进入时间",
		ChuangShiJiFenNoEnough:              "创世之战积分不足",
		ChuangShiJoinCampFailed:             "加入创世阵营失败",
		ChuangShiJianSheSkillSetFailed:      "创世建设技能设置失败",
		ChuangShiJianSheSkillSetSuccess:     "天气技能设置成功",
		ChuangShiCityNotExist:               "城池不存在",
		ChuangShiCityNotFuShu:               "不是附属城",
		ChuangShiCityNotNearCity:            "不是相邻城池",
		ChuangShiCityTargetFullNum:          "攻城目标数量已满",
		ChuangShiCityTargetHadChooes:        "已经选择该阵营为攻城目标",
		ChuangShiCityTargetNotTime:          "不是设置攻城目标时间",
		ChuangShiCampPayScheduleNotResouces: "没有可分配的阵营工资",
		ChuangShiCampChangedNotAllow:        "攻城前不能更改阵营",

		ChuangShiBuChangMailTitle:             "创世之战更新补偿",
		ChuangShiBuChangMailContent:           "创世之战更新补偿",
		ChuangShiLogGongChengTarget:           "我方阵营神王%s将%s的%s设为了攻城目标，请各位仙友届时准时参与，共谋大业",
		ChuangShiLogGongChengBeTarget:         "%s的神王%s将我方的%s设为了攻城目标，请各位仙友届时准时参与，共卫家园",
		ChuangShiLogPayScheduleCamp:           "我方阵营神王%s对阵营工资进行了分配，各位玩家可前往查看   【前往查看】",
		ChuangShiLogPayScheduleCity:           "我方阵营城主%s对城池工资进行了分配，各位玩家可前往查看   【前往查看】",
		ChuangShiLogRenMingChengZhu:           "我方阵营神王%s任命%s为%s的城主，任命后城主享有分配城池工资的特权",
		ChuangShiLogRenMingShenWang:           "我方阵营%s成功当选为新一任的神王，享有分配阵营工资、进攻城池、任命城主等特权，望各位仙友协力扶持，共谋大业",
		ChuangShiLogJianSheCity:               "我方阵营%s成功将%s等级提升至%s",
		ChuangShiLogJianSheTianQiSet:          "我方%s将%s天气效果设置为%s",
		ChuangShiLogGongChengWin:              "我方阵营总仙友众志成城，成功攻破%s的%s，一统仙界的日子指日可待",
		ChuangShiLogGongChengFailed:           "我方%s被%s攻破，家园沦陷，此等大辱岂可坐视不理，望众仙友下次城战共同参与，夺回家园",
		ChuangShiLogRenMingMengZhuChanged:     "%s的城主%s由于卸任了仙盟盟主职位，因此该城池城主自动移交给该仙盟新一任盟主%s",
		ChuangShiSceneFinishMailTitle:         "创世之战",
		ChuangShiSceneFinishWinMailContent:    "恭喜您所处阵营在【创世之战】活动中成功攻下/守护%s，本阵营所有玩家获得如下奖励，敬请查收",
		ChuangShiSceneFinishFailedMailContent: "很遗憾，您所处阵营在【创世之战】活动中进攻/守护%s失败，本阵营所有玩家获得以下鼓励，期待您下一次城战的参与",
	}
)

func init() {
	mergeLang(chuangShiLangMap)
}
