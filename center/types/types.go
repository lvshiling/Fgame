package types

type PlatformType int32

const (
	//默认渠道
	PlatformTypeSelf PlatformType = iota + 1
)

//游戏服务器类型
type GameServerType int32

const (
	//游戏服
	GameServerTypeSingle GameServerType = iota
	//临服 跨服游戏服
	GameServerTypeGroup
	//战场组(自定义服务器)
	GameServerTypeRegion
	//平台组
	GameServerTypePlatform
	//所有服务器
	GameServerTypeAll
	//城战跨服
	GameServerTypeChenZhan
)

func (t GameServerType) Valid() bool {
	switch t {
	case GameServerTypeSingle,
		GameServerTypeGroup,
		GameServerTypeRegion,
		GameServerTypePlatform,
		GameServerTypeAll,
		GameServerTypeChenZhan:
		return true
	}
	return false
}

var (
	gameServerTypeMap = map[GameServerType]string{
		GameServerTypeSingle:   "单服",
		GameServerTypeGroup:    "组服",
		GameServerTypeRegion:   "自定义区服",
		GameServerTypePlatform: "平台服",
		GameServerTypeAll:      "全世界",
		GameServerTypeChenZhan: "全世界城战",
	}
)

func (t GameServerType) String() string {
	return gameServerTypeMap[t]
}

var (
	crossServerTypeList = []GameServerType{GameServerTypeGroup, GameServerTypeRegion, GameServerTypePlatform, GameServerTypeAll}
)

func GetCrossServerTypeList() []GameServerType {
	return crossServerTypeList
}

type GameServerStatus int32

const (
	GameServerStatusMaintain GameServerStatus = iota
	GameServerStatusNormal
)

var (
	gameServerStatusMap = map[GameServerStatus]string{
		GameServerStatusMaintain: "维护中",
		GameServerStatusNormal:   "正常",
	}
)

func (t GameServerStatus) String() string {
	return gameServerStatusMap[t]
}

type ServerTag int32

const (
	ServerTagNone ServerTag = iota
	ServerTagNew
	ServerTagHot
)

var (
	serverTagMap = map[ServerTag]string{
		ServerTagNone: "无",
		ServerTagNew:  "新",
		ServerTagHot:  "热",
	}
)

func (t ServerTag) String() string {
	return serverTagMap[t]
}

type ServerStatus int32

const (
	ServerStatusNormal ServerStatus = iota
	ServerStatusFull
	ServerStatusMaintained
)

var (
	serverStatusMap = map[ServerStatus]string{
		ServerStatusNormal:     "流畅",
		ServerStatusFull:       "爆满",
		ServerStatusMaintained: "维护中",
	}
)

func (t ServerStatus) String() string {
	return serverStatusMap[t]
}
