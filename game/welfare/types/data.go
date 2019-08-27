package types

//排行榜奖励邮件回调数据
type RankEmailData struct {
	Ranking     int32
	RankNum     int64
	GroupId     int32
	EndTime     int64
	ContentType EmailContentType
}

func NewRankEmailData(endTime int64, ranking int32, rankNum int64, groupId int32, contentType EmailContentType) *RankEmailData {
	d := &RankEmailData{
		Ranking:     ranking,
		RankNum:     rankNum,
		GroupId:     groupId,
		EndTime:     endTime,
		ContentType: contentType,
	}
	return d
}

// 全局次数信息
type TimesLimitInfo struct {
	Key   int32
	Times int32
}

// 数据接口
type OpenActivityData interface{}
