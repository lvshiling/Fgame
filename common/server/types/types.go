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
)

var (
	gameServerTypeMap = map[GameServerType]string{
		GameServerTypeSingle:   "单服",
		GameServerTypeGroup:    "组服",
		GameServerTypeRegion:   "自定义区服",
		GameServerTypePlatform: "平台服",
		GameServerTypeAll:      "全世界",
	}
)

func (t GameServerType) String() string {
	return gameServerTypeMap[t]
}
