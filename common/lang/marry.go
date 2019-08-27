package lang

const (
	MarryProposalNoFriend LangCode = MarryBase + iota
	MarrySpouseNoOnline
	MarryHasedSpouse
	MarryIntimacyNoEnough
	MarrySpouseIsNoOpen
	MarrySpouseSameSex
	MarryProposalIsNoOnline
	MarryProposalIsMarried
	MarryDealIsMarried
	MarryDivorceNoOnline
	MarryRingFeedNoMarried
	MarryRingFeedLevelNoEnough
	MarryRingReachLimit
	MarryReserveNoMarried
	MarryReserveIsExist
	MarryReserveIsMarried
	MarryReserveIsDivorced
	MarryReservePeriodIsInvaild
	MarryReservePeriodTimeReachLimit
	MarryRingReplaceNoMarried
	MarryRingReplaceNeedSenior
	MarryTreeFeedNoMarried
	MarryTreeReachLimit
	MarryWedGiftNoWedExist
	MarryClickCarFrequent
	MarryTransferSameScene
	MarryDivorceInWeddingTime
	MarryProposalIsExist
	MarryPeerGiveGift
	MarryWedGiftTitle
	MarryWedGiftContent
	MarryDivorceConsent
	MarryDivorceForce
	MarryProposalSucess
	MarryDueWedding
	MarryWedGiveGiftFlower
	MarryWedGiveGiftSilver
	MarryRingGiveBackTitle
	MarryRingGiveBackContent
	MarryHoldWedNoExist
	MarryWedGiveGiftFireworks
	MarryMergeServerTitle
	MarryMergeServerContent
	MarryMergeServerContentNotice
	MarryWedCancleTitle
	MarryWedCancleContent
	MarryStartInMatchToLeaveTitle
	MarryStartInMatchToLeaveContent
	MarryDealIsSexChanged
	MarryCloseServerTitle
	MarryCloseServerContent
	MarryCloseServerContentNotice
	MarryPreWedSpouseNoOnline
	MarryPreWedNoProposal
	MarryWedIsBeScheduled
	MarryPreWedIsOverdue
	MarryPreWedPeerThinking
	MarryPreWedGiveBackTitle
	MarryPreWedGiveBackRefuseContent
	MarryPreWedGiveBackRobbedContent
	MarryWedSuccessTitle
	MarryWedSuccessContent
	MarryDivorceInPreWed
	MarryNotCouple
	MarryDevelopMaxLevel
	MarryDevelopExpNotEnough
	MarryProposalGiveBackTitle
	MarryProposalGiveBackContent
	MarryXinWuNotExists
	MarryXinWuAlreadyExists
	MarryXinWuItemNotEnough
	MarryXinWuNotSpouse
	MarryXinWuNotSpouseNotOnLine
	MarryXinWuDingQingSuoQuFail
	MarryXinWuDingQingSuoQuFailContent
	MarryXinWuDingQingSuoQuSuccess
	MarryXinWuDingQingSuoQuSuccessContent
	MarryJiNianSendShiZhuangTitle
	MarryJiNianSendShiZhuangContent
	MarryJiNianChengHaoTitle
	MarryJiNianChengHaoContent
	MarryQiuHunNotExists
	MarryPreMarryReturn
	MarryPreMarryReturnContent
	MarryXinWuSuoYaoCd
	MarryPreGiftContent
	MarryXinWuDingQingZengSongSuccess
	MarryXinWuDingQingZengSongSuccessContent
	MarryRingTemplateNotExist
)

var (
	marryLangMap = map[LangCode]string{
		MarryProposalNoFriend:                    "%s未添加您为好友,无法向对方发起求婚",
		MarrySpouseNoOnline:                      "玩家不在线,无法求婚",
		MarryHasedSpouse:                         "您或对方已拥有伴侣,无法进行求婚",
		MarryIntimacyNoEnough:                    "%s亲密度以上才可进行求婚",
		MarrySpouseIsNoOpen:                      "对方结婚功能未开启",
		MarrySpouseSameSex:                       "只能向异性求婚",
		MarryProposalIsNoOnline:                  "求婚者已下线",
		MarryProposalIsMarried:                   "求婚者已婚",
		MarryDealIsMarried:                       "您处于已婚状态",
		MarryDivorceNoOnline:                     "您的配偶当前不在线,无法进行协议离婚",
		MarryRingFeedNoMarried:                   "未结婚,无法培养婚戒",
		MarryRingFeedLevelNoEnough:               "夫妻戒指等级差不满足,无法提升",
		MarryRingReachLimit:                      "婚戒等级达到最高价",
		MarryReserveNoMarried:                    "未结婚,无法预定婚礼举办时间",
		MarryReserveIsExist:                      "已经预定过婚礼举办时间",
		MarryReserveIsMarried:                    "已经举办过婚礼了",
		MarryReserveIsDivorced:                   "您目前处于离婚状态",
		MarryReservePeriodIsInvaild:              "预定的举办场次已经过期了",
		MarryReservePeriodTimeReachLimit:         "预定的举办场次距离开始时间小于1分钟",
		MarryRingReplaceNoMarried:                "未婚,无法替换",
		MarryRingReplaceNeedSenior:               "婚戒替换需要更高级的",
		MarryTreeFeedNoMarried:                   "未结婚,无法培养爱情树",
		MarryTreeReachLimit:                      "爱情树等级达到最高阶",
		MarryWedGiftNoWedExist:                   "婚宴场次不存在,无法赠送贺礼",
		MarryClickCarFrequent:                    "点击婚车过于频繁",
		MarryTransferSameScene:                   "您当前已在该场景",
		MarryDivorceInWeddingTime:                "婚礼期间无法离婚",
		MarryProposalIsExist:                     "您已经向%s求婚,同一时间只能求婚一个",
		MarryPeerGiveGift:                        "%s对%s送出鲜花,亲密度上升%d,情比金坚",
		MarryWedGiftTitle:                        "新婚贺礼",
		MarryWedGiftContent:                      "祝福两位新人永结同心,百年好合!",
		MarryDivorceConsent:                      "%s与%s感情破裂,通过协商,和平分手,双方亲密度扣除%d%%",
		MarryDivorceForce:                        "%s与%s感情破裂,一纸休书,休了%s",
		MarryProposalSucess:                      "%s和%s一起许下爱的誓言，从此修仙路上又多了一对神仙眷侣，羡煞旁人!",
		MarryDueWedding:                          "%s与%s喜结连理,普天同庆,将于%s宴请广大侠士",
		MarryWedGiveGiftFlower:                   "%s豪气冲天,赠送了%s,为新郎新娘送上最美好的祝福!",
		MarryWedGiveGiftSilver:                   "%s豪气冲天,赠送了%d银两,为新郎新娘送上最美好的祝福!",
		MarryRingGiveBackTitle:                   "婚戒返还",
		MarryRingGiveBackContent:                 "求婚失败",
		MarryHoldWedNoExist:                      "目前无正在举办的婚礼",
		MarryMergeServerTitle:                    "合服结婚重置",
		MarryMergeServerContent:                  "由于当前服务器进行合服,结婚进行重置,您失去了原先的预定的婚礼,系统补偿您的预定的花费,请进行查收",
		MarryMergeServerContentNotice:            "由于当前服务器进行合服,结婚进行重置,您的婚期已取消",
		MarryWedCancleTitle:                      "婚礼取消",
		MarryWedCancleContent:                    "您的婚礼因为有一方不在线而被取消,预约婚礼的元宝将全数返还给预约者,请两位再择吉日结婚,祝幸福!",
		MarryStartInMatchToLeaveTitle:            "离队通知",
		MarryStartInMatchToLeaveContent:          "您的婚礼开始时,您处于3v3匹配,系统帮您自动离队",
		MarryDealIsSexChanged:                    "您或者对方使用了变性卡,无法结婚",
		MarryCloseServerTitle:                    "预约婚礼取消",
		MarryCloseServerContent:                  "由于当前服务器关服进行维护,您的婚期受到影响,为了防止您的婚礼出问题,特此帮您取消婚礼并且把预约婚礼的消耗以邮件的形式返还给您,请查收！",
		MarryCloseServerContentNotice:            "由于当前服务器关服进行维护,您的婚期受到影响,为了防止您的婚礼出问题,特此帮您取消婚礼并且把预约婚礼的消耗以邮件的形式返还给您的配偶,请通知您的配偶重新预约婚礼！",
		MarryPreWedSpouseNoOnline:                "您的爱人不在线,无法预定婚期",
		MarryPreWedNoProposal:                    "您不是求婚者发起人,无法预定婚礼,赶快向配偶催办婚礼吧",
		MarryWedIsBeScheduled:                    "本场次婚礼已被预定,您可以预定其它场次的",
		MarryPreWedIsOverdue:                     "预定婚期请求已失效",
		MarryPreWedPeerThinking:                  "您的爱人正在思考中,请过一会儿再发送",
		MarryPreWedGiveBackTitle:                 "预定婚礼归还",
		MarryPreWedGiveBackRefuseContent:         "您的爱人婉拒了您的婚礼预约,归还您的预定花费",
		MarryPreWedGiveBackRobbedContent:         "您的预定的婚礼被其它人抢先预定了,归还您的预定花费",
		MarryWedSuccessTitle:                     "婚礼预定成功",
		MarryWedSuccessContent:                   "您已经成功预约%s举行婚礼,请与您的爱人准时在本服进行婚礼",
		MarryDivorceInPreWed:                     "当前请求举办婚宴中,无法离婚",
		MarryNotCouple:                           "不是您的伴侣",
		MarryDevelopMaxLevel:                     "您当前等级已达上限，无法继续升级",
		MarryDevelopExpNotEnough:                 "表白经验不足",
		MarryProposalGiveBackTitle:               "求婚失败",
		MarryProposalGiveBackContent:             "您向%s发起的求婚被拒绝，求婚失败，系统将求婚的元宝以邮件形式返还给您，请查收",
		MarryXinWuNotExists:                      "信物模板不存在",
		MarryXinWuAlreadyExists:                  "已经存在信物",
		MarryXinWuItemNotEnough:                  "信物不足",
		MarryXinWuNotSpouse:                      "还未结婚",
		MarryXinWuNotSpouseNotOnLine:             "您的伴侣不在线",
		MarryXinWuDingQingSuoQuFail:              "索要拒绝通知",
		MarryXinWuDingQingSuoQuSuccess:           "赠送通知",
		MarryXinWuDingQingSuoQuFailContent:       "非常遗憾，由于您的伴侣囊中羞涩，您向您的伴侣发起的【%s】的索要申请已被拒绝。",
		MarryXinWuDingQingSuoQuSuccessContent:    "您的伴侣财大气粗，您向您的伴侣发起的【%s】的索要申请已确定赠送激活，战力提升%d。",
		MarryJiNianSendShiZhuangTitle:            "结婚纪念",
		MarryJiNianSendShiZhuangContent:          "您完美的完成所有结婚纪念任务，赠送您一套限定唯美时装，请查收！",
		MarryJiNianChengHaoTitle:                 "结婚纪念称号",
		MarryJiNianChengHaoContent:               "您和您的伴侣举行了%d次%s婚礼，获得了结婚纪念称号，请查收！",
		MarryQiuHunNotExists:                     "求婚类型不存在",
		MarryPreMarryReturn:                      "预定婚礼归还",
		MarryPreMarryReturnContent:               "您和您的伴侣已经离婚，原先预约的婚宴消耗以邮件形式返还给您，请查收！",
		MarryXinWuSuoYaoCd:                       "索要信物冷却中，请稍后再试",
		MarryPreGiftContent:                      "您的背包空间已满，结婚祝福宝箱无法获得，已通过邮件发到您的邮箱，请查收",
		MarryXinWuDingQingZengSongSuccess:        "赠送通知",
		MarryXinWuDingQingZengSongSuccessContent: "您的伴侣财大气粗，您的伴侣给您赠送了【%s】，信物成功激活，战力提升%d。",
		MarryRingTemplateNotExist:                "婚戒模板不存在",
	}
)

func init() {
	mergeLang(marryLangMap)
}
