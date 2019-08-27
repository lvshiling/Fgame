/*此类自动生成,请勿修改*/
package template

/*创世城战配置*/
type ChuangShiWarTemplateVO struct {

	//id
	Id int `json:"id"`

	//原地复活的最高次数
	RebornSitu int32 `json:"reborn_situ"`

	//进入后多少毫秒获得奖励
	FristTiem int64 `json:"frist_tiem"`

	//多少毫秒发一次奖励
	RewTiem int64 `json:"rew_tiem"`

	//单次发放经验值
	RewExp int64 `json:"rew_exp"`

	//单次发放的经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int64 `json:"rew_silver"`

	//地图id
	MapId int32 `json:"map_id"`

	//安全区单次发放经验值
	RewExpSafeArea int64 `json:"rew_exp_anquanqu"`

	//安全区单次发放的经验点
	RewExpPointSafeArea int32 `json:"rew_exp_point_anquanqu"`

	//安全区奖励银两
	RewSilverSafeArea int64 `json:"rew_silver_anquanqu"`

	//击杀获得积分
	KillJiFen int32 `json:"kill_jifen"`

	//进入后多少毫秒获得积分
	FristJiFenTime int64 `json:"frist_jifen_time"`

	//多少毫秒发一次积分
	RewJiFenTime int64 `json:"rew_jifen_time"`

	//单次发放积分
	RewJiFen int32 `json:"rew_jifen"`

	//防护罩无敌buff
	ProtectBuffId int32 `json:"wudi_buff_id"`
}
