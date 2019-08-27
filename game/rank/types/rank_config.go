package types

type RankConfig struct {
	GroupId       int32         //活动id
	StartTime     int64         //开始时间
	EndTime       int64         //结束时间
	MinCondition  int32         //入榜条件
	MaxExpireTime int64         //过期时间
	IsCanExpire   bool          //是否允许过期
	RefreshTime   int64         //刷新时间
	ClassType     RankClassType //排行分类
	RankType      RankType      //排行榜类型
}

func NewRankConfig() *RankConfig {
	rc := &RankConfig{}
	return rc
}

func NewLocalDefaultConfig() *RankConfig {
	rc := &RankConfig{}
	rc.MinCondition = 1
	rc.StartTime = 0
	rc.EndTime = 0
	rc.ClassType = RankClassTypeLocal
	rc.MaxExpireTime = rc.ClassType.RankRedisExpireTime()
	rc.RefreshTime = rc.ClassType.RankRefreshTime()
	rc.IsCanExpire = true
	return rc
}

func NewAreaDefaultConfig() *RankConfig {
	rc := &RankConfig{}
	rc.MinCondition = 1
	rc.StartTime = 0
	rc.EndTime = 0
	rc.ClassType = RankClassTypeArea
	rc.MaxExpireTime = rc.ClassType.RankRedisExpireTime()
	rc.RefreshTime = rc.ClassType.RankRefreshTime()
	rc.IsCanExpire = true
	return rc
}

//榜单排名
type RankingInfo struct {
	playerId   int64  //玩家id
	playerName string //玩家姓名
	ranking    int32  //排名
	rankNum    int64  //排行的数据
}

func CreateRankingInfo(playerId int64, playerName string, ranking int32, rankNum int64) *RankingInfo {
	d := &RankingInfo{}
	d.ranking = ranking
	d.playerName = playerName
	d.rankNum = rankNum
	d.playerId = playerId
	return d
}

func (info *RankingInfo) GetPlayerId() int64 {
	return info.playerId
}

func (info *RankingInfo) GetPlayerName() string {
	return info.playerName
}

func (info *RankingInfo) GetRanking() int32 {
	return info.ranking
}

func (info *RankingInfo) GetRankNum() int64 {
	return info.rankNum
}
