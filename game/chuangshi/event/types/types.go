package types

type ChuangShiEventType string

const (
	EventTypeChuangShiShenWangSignUp   ChuangShiEventType = "ChuangShiShenWangSignUp"   //神王竞选报名成功
	EventTypeChuangShiShenWangVote                        = "ChuangShiShenWangVote"     //神王竞选投票成功
	EventTypeChuangShiChengFangJianShe                    = "ChuangShiChengFangJianShe" //城池建设成功
	EventTypeChuangShiCampPaySchedule                     = "ChuangShiCampPaySchedule"  //阵营工资分配
	EventTypeChuangShiCityPaySchedule                     = "ChuangShiCityPaySchedule"  //城池工资分配
	EventTypeChuangShiShenWangChanged                     = "ChuangShiShenWangChanged"  //神王变更
)
