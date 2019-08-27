package lang

const (
	JieYiNotJieYi LangCode = JieYiBase + iota
	JieYiCanNotInvite
	JieYiAlreadyJieYi
	JieYiNotOnline
	JieYiTemplateNotExist
	JieYiNameIllegal
	JieYiNameDirty
	JieYiNameRepetitive
	JieYiMemberAlreadyFull
	JieYiInviteTimeCD
	JieYiObjectNotExist
	JieYiNotIsLaoDa
	JieYiNameAlreadyTopLevel
	JieYiShengWeiZhiNotEnough
	JieYiTokenAlreadyActivite
	JieYiTokenNotActivite
	JieYiPostMessageCD
	JieYiLeaveWordIllegal
	JieYiLeaveWordDirty
	JieYiTokenNotChange
	JieYiNotIsSameJieYi
	JieYiDaoJuNotChange
	JieYiQiuYuanTimeCD
	JieYiMemberNotChuanSong
	JieYiLeaveWordGongGao
	JieYiJieChuJieYiCD
	JieYiDuiFangJieChuJieYiCD
	JieYiChangeNameSuccess
	JieYiDuiFangAlreadyJieYi
	JieYiInivteExprise
	JieYiInivteFail
	JieYiBuChaJiaFail
	JieYiLiuYanFail
	JieYiUseGaiMingKaFail
)

var (
	jieyiLangMap = map[LangCode]string{
		JieYiNotJieYi:             "玩家未结义",
		JieYiCanNotInvite:         "您不是结义老大",
		JieYiAlreadyJieYi:         "您已经结义",
		JieYiNotOnline:            "对方不在线",
		JieYiTemplateNotExist:     "模板不存在",
		JieYiNameIllegal:          "威名字数需要2~6个字",
		JieYiNameDirty:            "结义威名含有脏字",
		JieYiNameRepetitive:       "结义威名重复",
		JieYiMemberAlreadyFull:    "结义人数已满",
		JieYiInviteTimeCD:         "距离上一次邀请未满%s分钟",
		JieYiObjectNotExist:       "结义对象不存在",
		JieYiNotIsLaoDa:           "玩家不是老大",
		JieYiNameAlreadyTopLevel:  "威名已经是最高等级",
		JieYiShengWeiZhiNotEnough: "声威值不足，无法升级",
		JieYiTokenAlreadyActivite: "信物已经激活",
		JieYiTokenNotActivite:     "信物未激活",
		JieYiPostMessageCD:        "距离上一次发布结义消息未满%s分钟",
		JieYiLeaveWordIllegal:     "玩家结义留言不合法",
		JieYiLeaveWordDirty:       "玩家结义留言含有脏字",
		JieYiTokenNotChange:       "低级信物无法替换高级信物",
		JieYiNotIsSameJieYi:       "不在同一结义阵营",
		JieYiDaoJuNotChange:       "低级道具无法替换高级道具",
		JieYiQiuYuanTimeCD:        "距离上一次求援未满%s分钟",
		JieYiMemberNotChuanSong:   "该地图不支持传送",
		JieYiLeaveWordGongGao:     "千金易得，挚友难寻，%s发布了结义公告，望能够寻找有志之士共同结义 %s",
		JieYiJieChuJieYiCD:        "您当前与他人解除结义未满%s小时，无法继续结义",
		JieYiDuiFangJieChuJieYiCD: "对方当前与他人解除结义未满%s小时，无法继续结义",
		JieYiChangeNameSuccess:    "改名成功",
		JieYiDuiFangAlreadyJieYi:  "对方已经结义",
		JieYiInivteExprise:        "邀请已过期",
		JieYiInivteFail:           "正在等待对方回应...",
		JieYiBuChaJiaFail:         "补差价升级失败，购买次数不足",
		JieYiLiuYanFail:           "结义留言失败",
		JieYiUseGaiMingKaFail:     "使用改名卡失败",
	}
)

func init() {
	mergeLang(jieyiLangMap)
}
