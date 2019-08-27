package types

type PrivilegeLevel int32

func (n PrivilegeLevel) HasChannel() bool {
	if n == PrivilegeLevelChannel ||
		n == PrivilegeLevelPlatform ||
		n == PrivilegeLevelKeFu ||
		n == PrivilegeLevelMinitor ||
		n == PrivilegeLevelCommonKeFu ||
		n == PrivilegeLevelGaoJiKeFu ||
		n == PrivilegeLevelNeiGua ||
		n == PrivilegeLevelGs {
		return true
	}
	return false
}

func (n PrivilegeLevel) HasPlatform() bool {
	if n == PrivilegeLevelPlatform ||
		n == PrivilegeLevelKeFu ||
		n == PrivilegeLevelMinitor ||
		n == PrivilegeLevelCommonKeFu ||
		n == PrivilegeLevelGaoJiKeFu ||
		n == PrivilegeLevelNeiGua ||
		n == PrivilegeLevelGs {
		return true
	}
	return false
}

func (n PrivilegeLevel) HasCanShouYeOrder() bool {
	if n == PrivilegeLevelMinitor ||
		n == PrivilegeLevelKeFu ||
		n == PrivilegeLevelCommonKeFu ||
		n == PrivilegeLevelGaoJiKeFu ||
		n == PrivilegeLevelNeiGua ||
		n == PrivilegeLevelGs {
		return false
	}
	return true
}

const (
	PrivilegeLevelSuperChannel        PrivilegeLevel = 1
	PrivilegeLevelChannel             PrivilegeLevel = 2
	PrivilegeLevelPlatform            PrivilegeLevel = 3
	PrivilegeLevelKeFu                PrivilegeLevel = 4
	PrivilegeLevelMinitor             PrivilegeLevel = 5
	PrivilegeLevelSuperChannelService PrivilegeLevel = 6
	PrivilegeLevelCommonKeFu          PrivilegeLevel = 7
	PrivilegeLevelGaoJiKeFu           PrivilegeLevel = 8
	PrivilegeLevelNeiGua              PrivilegeLevel = 9
	PrivilegeLevelGs                  PrivilegeLevel = 10
	//以防后面要加角色
	PrivilegeLevelAdmin PrivilegeLevel = 999
)

var privilegeLevelStringMap = map[PrivilegeLevel]string{
	PrivilegeLevelSuperChannel:        "总渠道管理员",
	PrivilegeLevelChannel:             "渠道管理员",
	PrivilegeLevelPlatform:            "平台管理员",
	PrivilegeLevelKeFu:                "平台客服",
	PrivilegeLevelMinitor:             "监控",
	PrivilegeLevelAdmin:               "超级管理员",
	PrivilegeLevelSuperChannelService: "总渠道客服",
	PrivilegeLevelCommonKeFu:          "普通客服",
	PrivilegeLevelGaoJiKeFu:           "高级客服",
	PrivilegeLevelNeiGua:              "内挂管理",
	PrivilegeLevelGs:                  "Gs管理",
}

var privilegeLevelCodeMap = map[PrivilegeLevel]string{
	PrivilegeLevelSuperChannel:        "super_channel",
	PrivilegeLevelChannel:             "channel",
	PrivilegeLevelPlatform:            "platform",
	PrivilegeLevelKeFu:                "service",
	PrivilegeLevelMinitor:             "minitor",
	PrivilegeLevelAdmin:               "super_admin",
	PrivilegeLevelSuperChannelService: "super_channel_service",
	PrivilegeLevelCommonKeFu:          "common_service",
	PrivilegeLevelGaoJiKeFu:           "gaoji_service",
	PrivilegeLevelNeiGua:              "neigua",
	PrivilegeLevelGs:                  "gs",
}

var privilegeMap = map[PrivilegeLevel]int64{}

func (pl PrivilegeLevel) Valid() bool {
	switch pl {
	case PrivilegeLevelSuperChannel:
	case PrivilegeLevelChannel:
	case PrivilegeLevelPlatform:
	case PrivilegeLevelKeFu:
	case PrivilegeLevelMinitor:
	case PrivilegeLevelAdmin:
	case PrivilegeLevelSuperChannelService:
	case PrivilegeLevelCommonKeFu:
	case PrivilegeLevelGaoJiKeFu:
	case PrivilegeLevelNeiGua:
	case PrivilegeLevelGs:
	default:
		return false
	}
	return true
}

var privilegePriorityMap = map[PrivilegeLevel]int32{
	PrivilegeLevelSuperChannelService: 6,
	PrivilegeLevelMinitor:             5,
	PrivilegeLevelSuperChannel:        4,
	PrivilegeLevelChannel:             3,
	PrivilegeLevelPlatform:            2,
	PrivilegeLevelCommonKeFu:          1,
	PrivilegeLevelKeFu:                1,
	PrivilegeLevelGaoJiKeFu:           1,
	PrivilegeLevelNeiGua:              1,
	PrivilegeLevelGs:                  1,
	PrivilegeLevelAdmin:               999,
}

func (pl PrivilegeLevel) Weight() int32 {
	return privilegePriorityMap[pl]
}

func (pl PrivilegeLevel) Priority(pl2 PrivilegeLevel) bool {
	return pl.Weight() > pl2.Weight()
}

func (pl PrivilegeLevel) String() string {
	return privilegeLevelStringMap[pl]
}

func (pl PrivilegeLevel) Code() string {
	return privilegeLevelCodeMap[pl]
}

func (pl PrivilegeLevel) Privilege() int64 {
	return privilegeMap[pl]
}

func (pl PrivilegeLevel) ChildPrivilege() []PrivilegeLevel {
	rst := make([]PrivilegeLevel, 0)
	if pl == PrivilegeLevelAdmin {
		rst = append(rst, PrivilegeLevelSuperChannel)
		rst = append(rst, PrivilegeLevelChannel)
		rst = append(rst, PrivilegeLevelPlatform)
		rst = append(rst, PrivilegeLevelKeFu)
		rst = append(rst, PrivilegeLevelMinitor)
		rst = append(rst, PrivilegeLevelAdmin)
		rst = append(rst, PrivilegeLevelSuperChannelService)
		rst = append(rst, PrivilegeLevelCommonKeFu)
		rst = append(rst, PrivilegeLevelGaoJiKeFu)
		rst = append(rst, PrivilegeLevelNeiGua)
		rst = append(rst, PrivilegeLevelGs)
	}
	if pl == PrivilegeLevelSuperChannel {
		rst = append(rst, PrivilegeLevelChannel)
		rst = append(rst, PrivilegeLevelPlatform)
		rst = append(rst, PrivilegeLevelKeFu)
		rst = append(rst, PrivilegeLevelMinitor)
		rst = append(rst, PrivilegeLevelSuperChannelService)
		rst = append(rst, PrivilegeLevelCommonKeFu)
		rst = append(rst, PrivilegeLevelGaoJiKeFu)
		rst = append(rst, PrivilegeLevelNeiGua)
		rst = append(rst, PrivilegeLevelGs)
	}
	if pl == PrivilegeLevelChannel {
		rst = append(rst, PrivilegeLevelPlatform)
		rst = append(rst, PrivilegeLevelKeFu)
		rst = append(rst, PrivilegeLevelMinitor)
		rst = append(rst, PrivilegeLevelCommonKeFu)
		rst = append(rst, PrivilegeLevelGaoJiKeFu)
		rst = append(rst, PrivilegeLevelNeiGua)
		rst = append(rst, PrivilegeLevelGs)
	}
	if pl == PrivilegeLevelPlatform {
		rst = append(rst, PrivilegeLevelKeFu)
		rst = append(rst, PrivilegeLevelMinitor)
		rst = append(rst, PrivilegeLevelCommonKeFu)
		rst = append(rst, PrivilegeLevelGaoJiKeFu)
		rst = append(rst, PrivilegeLevelNeiGua)
		rst = append(rst, PrivilegeLevelGs)
	}
	return rst
}

type PrivilegeType int64

const (
	PrivilegeTypeUserManage PrivilegeType = 1 << iota
	PrivilegeTypeChannelManage
	PrivilegeTypePlatformManage
	PrivilegeTypePlayerSearch
	PrivilegeTypeChatMinitor
	PrivilegeTypeChatSet
	PrivilegeTypeCenterPlatformManage
	PrivilegeTypeCenterServerManage
	PrivilegeTypeGameLog
	PrivilegeTypeAlliance
	PrivilegeTypeMailApply
	PrivilegeTypeMailApprove
	PrivilegeTypeServerSupportPool
	PrivilegeTypeServerSupportPlayer
	PrivilegeTypeServerCenterOrderList
	PrivilegeTypeOnLineReport
	PrivilegeTypeNotice
	PrivilegeTypeRedeem
	PrivilegeTypeCenterUserQuery
	PrivilegeTypeCenterUser
	PrivilegeTypeGoldChange
	PrivilegeTypeCenternotice
	PrivilegeTypeCenterServerSimpleList
	PrivilegeTypePlatformManageMarry
	PrivilegeTypeServerMetaSet
	PrivilegeTypeCenterServerTradeItem
	PrivilegeTypeRetention
	PrivilegeTypeServerDaily
	PrivilegeTypeRecycle
	PrivilegeTypeJiaoYiZhanQu
	PrivilegeTypeFeedBackFee
	PrivilegeTypeDoubleCharge
)

var privilegeTypeStringMap = map[PrivilegeType]string{
	PrivilegeTypeUserManage:             "用户管理",
	PrivilegeTypeChannelManage:          "渠道管理",
	PrivilegeTypePlatformManage:         "平台管理",
	PrivilegeTypePlayerSearch:           "玩家查询",
	PrivilegeTypeChatMinitor:            "聊天监控",
	PrivilegeTypeChatSet:                "聊天设置",
	PrivilegeTypeCenterPlatformManage:   "中心平台配置",
	PrivilegeTypeCenterServerManage:     "中心服务配置",
	PrivilegeTypeGameLog:                "日志查询",
	PrivilegeTypeAlliance:               "仙盟查询",
	PrivilegeTypeMailApply:              "邮件提交",
	PrivilegeTypeMailApprove:            "邮件审核",
	PrivilegeTypeServerSupportPool:      "扶持池",
	PrivilegeTypeServerSupportPlayer:    "扶持玩家",
	PrivilegeTypeServerCenterOrderList:  "中心订单查询",
	PrivilegeTypeOnLineReport:           "在线人数报表",
	PrivilegeTypeNotice:                 "公告发送",
	PrivilegeTypeRedeem:                 "兑换码",
	PrivilegeTypeCenterUserQuery:        "中心用户查询",
	PrivilegeTypeCenterUser:             "中心用户信息",
	PrivilegeTypeGoldChange:             "元宝变化报表",
	PrivilegeTypeCenternotice:           "中心公告配置",
	PrivilegeTypeCenterServerSimpleList: "服务器列表",
	PrivilegeTypePlatformManageMarry:    "平台结婚配置",
	PrivilegeTypeServerMetaSet:          "平台参数配置",
	PrivilegeTypeCenterServerTradeItem:  "中心交易平台",
	PrivilegeTypeRetention:              "存留率统计",
	PrivilegeTypeServerDaily:            "分服详细",
	PrivilegeTypeRecycle:                "回收元宝",
	PrivilegeTypeJiaoYiZhanQu:           "交易战区",
	PrivilegeTypeFeedBackFee:            "兑换记录",
	PrivilegeTypeDoubleCharge:           "首冲翻倍",
}

func (pt PrivilegeType) String() string {
	return privilegeTypeStringMap[pt]
}

func init() {
	initPrivilegeMap()
}

func initPrivilegeMap() {

	privilegeMap[PrivilegeLevelKeFu] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeCenterServerSimpleList |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelCommonKeFu] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelGaoJiKeFu] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelPlatform] = int64(
		PrivilegeTypeUserManage |
			PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeMailApprove |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeRedeem |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeCenterServerSimpleList |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelChannel] = int64(
		PrivilegeTypeUserManage |
			PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeMailApprove |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeRedeem |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeCenterServerSimpleList |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelSuperChannel] = int64(
		PrivilegeTypeUserManage |
			PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeAlliance |
			PrivilegeTypeMailApply |
			PrivilegeTypeMailApprove |
			PrivilegeTypeServerSupportPool |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeGameLog |
			PrivilegeTypeNotice |
			PrivilegeTypeRedeem |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeCenterServerSimpleList |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelSuperChannelService] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeChatSet |
			PrivilegeTypeMailApply |
			PrivilegeTypeAlliance |
			PrivilegeTypeNotice |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeGameLog |
			PrivilegeTypeRedeem |
			PrivilegeTypeCenterServerSimpleList |
			PrivilegeTypeCenterServerTradeItem |
			PrivilegeTypeServerDaily |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelMinitor] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeGameLog |
			PrivilegeTypeAlliance |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeChatSet |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)

	privilegeMap[PrivilegeLevelNeiGua] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeGameLog |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeServerSupportPlayer |
			PrivilegeTypeRecycle |
			PrivilegeTypeFeedBackFee)
	privilegeMap[PrivilegeLevelGs] = int64(
		PrivilegeTypePlayerSearch |
			PrivilegeTypeGameLog |
			PrivilegeTypeChatMinitor |
			PrivilegeTypeServerCenterOrderList |
			PrivilegeTypeCenterUserQuery |
			PrivilegeTypeOnLineReport |
			PrivilegeTypeFeedBackFee)
	privilegeMap[PrivilegeLevelAdmin] = ^0

}
