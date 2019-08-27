package types

//好友系统对方操作日志类型
type FriendLogType int32

const (
	//拉黑
	FriendLogTypeBlack FriendLogType = 1 + iota
	//移出黑名单
	FriendLogTypeRemoveBlack
)

// 好友推送类型
type FriendNoticeType int32

const (
	FriendNoticeTypeUplevel       FriendNoticeType = iota //升级
	FriendNoticeTypeZhuanSheng                            //转生
	FriendNoticeTypeKillBoss                              //击杀boss掉落
	FriendNoticeTypeMountAdvanced                         //坐骑升阶
	FriendNoticeTypeWingAdvanced                          //战翼升阶
)

func (t FriendNoticeType) Valid() bool {
	switch t {
	case FriendNoticeTypeUplevel,
		FriendNoticeTypeZhuanSheng,
		FriendNoticeTypeKillBoss,
		FriendNoticeTypeMountAdvanced,
		FriendNoticeTypeWingAdvanced:
		return true
	}

	return false
}

// 好友推送类型
type FriendFeedbackType int32

const (
	FriendFeedbackTypeFlower FriendFeedbackType = iota //鲜花
	FriendFeedbackTypeEgg                              //鸡蛋
)

func (t FriendFeedbackType) Valid() bool {
	switch t {
	case FriendFeedbackTypeFlower,
		FriendFeedbackTypeEgg:
		return true
	}

	return false
}

// 表白记录数据类型
type MarryDevelopLogType int32

const (
	MarryDevelopLogTypeAll  MarryDevelopLogType = iota //全部
	MarryDevelopLogTypeSend                            //我的表白
	MarryDevelopLogTypeRecv                            //对我的表白
)

func (t MarryDevelopLogType) Valid() bool {
	switch t {
	case MarryDevelopLogTypeAll,
		MarryDevelopLogTypeSend,
		MarryDevelopLogTypeRecv:
		return true
	}

	return false
}
