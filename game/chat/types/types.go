package types

type ChannelType int32

const (
	//世界频道
	ChannelTypeWorld ChannelType = iota
	//帮派
	ChannelTypeBangPai
	//队伍
	ChannelTypeTeam
	//系统
	ChannelTypeSystem
	//私聊
	ChannelTypePerson
	//答题
	ChannelTypeQuiz
)

var (
	channelTypeMap = map[ChannelType]string{
		ChannelTypeWorld:   "世界",
		ChannelTypeBangPai: "帮派",
		ChannelTypeTeam:    "队伍",
		ChannelTypeSystem:  "系统",
		ChannelTypePerson:  "私聊",
		ChannelTypeQuiz:    "答题",
	}
)

func (ct ChannelType) String() string {
	return channelTypeMap[ct]
}

func (ct ChannelType) Valid() bool {
	switch ct {
	case ChannelTypeWorld,
		ChannelTypeBangPai,
		ChannelTypeTeam,
		ChannelTypeSystem,
		ChannelTypePerson,
		ChannelTypeQuiz:
		return true
	}
	return false
}

type MsgType int32

const (
	//文本
	MsgTypeText MsgType = iota
	//表情
	MsgTypeEmoji
	//声音
	MsgTypeVoice
	//红包
	MsgTypeHongBao
)

var (
	msgTypeMap = map[MsgType]string{
		MsgTypeText:    "文本",
		MsgTypeEmoji:   "表情",
		MsgTypeVoice:   "声音",
		MsgTypeHongBao: "红包",
	}
)

func (mt MsgType) String() string {
	return msgTypeMap[mt]
}

func (mt MsgType) Valid() bool {
	switch mt {
	case MsgTypeText,
		MsgTypeEmoji,
		MsgTypeVoice,
		MsgTypeHongBao:
		return true
	}
	return false
}

//聊天链接类型
type ChatLinkType int32

const (
	//物品链接
	ChatLinkTypeItem ChatLinkType = 1 + iota
	//仙盟链接
	ChatLinkTypeXianMeng
	//跳转至npc位置
	ChatLinkTypeNpc
	//打开界面
	ChatLinkTypeOpenView
	//申请队伍--5
	ChatLinkTypeTeamApply
	//传送世界场景
	ChatLinkToWorldMap
	//跳转仙盟救援
	ChatAllianceRescue
	//点击人名连接
	ChatPlayerName
	//立即前往进场景
	ChatPlayerEnterScene
)

// 聊天常量配置
type ChatConstantType int32

const (
	ChatConstantTypeStopChatStartTime ChatConstantType = 1 //不可聊天的开始时间
	ChatConstantTypeStopChatEndTime                    = 2 //不可聊天的结束时间
)

const (
	ChatConstantTypeMin = ChatConstantTypeStopChatStartTime
	ChatConstantTypeMax = ChatConstantTypeStopChatEndTime
)

var (
	constantTypeMap = map[ChatConstantType]string{
		ChatConstantTypeStopChatStartTime: "不可聊天的开始时间",
		ChatConstantTypeStopChatEndTime:   "不可聊天的结束时间",
	}
)

func (ct ChatConstantType) String() string {
	return constantTypeMap[ct]
}
