/*此类自动生成,请勿修改*/
package template

/*婚礼档次配置*/
type MarryBanquetTemplateVO struct {

	//id
	Id int `json:"id"`

	//版本类型
	HoutaiType int32 `json:"houtai_type"`

	//婚礼类型
	Type int32 `json:"type"`

	//婚礼子类型
	SubType int32 `json:"sub_type"`

	//婚宴名字
	Name string `json:"name"`

	//花费银两
	UseSilver int32 `json:"use_silver"`

	//消耗的绑元
	UseBinggold int32 `json:"use_binggold"`

	//花费元宝
	UseGold int32 `json:"use_gold"`

	//撒喜糖操作时间间隔(毫秒)
	SugarEach int32 `json:"sugar_each"`

	//手动丢撒喜糖的CD(毫秒)
	DropTime int32 `json:"drop_time"`

	//喜糖掉落id
	DropId string `json:"drop_id"`

	//生物表id婚宴
	BanquetBiology int32 `json:"banquet_biology"`

	//刷酒桌的地图id
	MapId int32 `json:"map_id"`

	//酒桌的坐标x1
	PosX1 float64 `json:"pos_x1"`

	//酒桌的坐标y1
	PosY1 float64 `json:"pos_y1"`

	//酒桌的坐标z1
	PosZ1 float64 `json:"pos_z1"`

	//酒桌的坐标x2
	PosX2 float64 `json:"pos_x2"`

	//酒桌的坐标y2
	PosY2 float64 `json:"pos_y2"`

	//酒桌的坐标z2
	PosZ2 float64 `json:"pos_z2"`

	//酒桌的坐标x3
	PosX3 float64 `json:"pos_x3"`

	//酒桌的坐标y3
	PosY3 float64 `json:"pos_y3"`

	//酒桌的坐标z3
	PosZ3 float64 `json:"pos_z3"`

	//是否有游街
	IsYouJie int32 `json:"is_youjie"`

	//酒桌buff
	AddBuffId int32 `json:"add_buff_id"`

	//资源名称
	ZiYuan string `json:"ziyuan"`

	//折扣值(前端显示用)
	Discount int32 `json:"discount"`

	//前端显示的折扣前银两
	FrontSilver int32 `json:"front_silver"`

	//前端显示的折扣前绑元
	FrontBinggold int32 `json:"front_binggold"`

	//前端显示的折扣前元宝
	FrontGold int32 `json:"front_gold"`

	//婚礼结束后奖励物品ID，配置同服务端实际扣除的消耗，取同类型的相加获得
	EndRewId string `json:"end_rew_id"`

	//婚礼结束后奖励物品数量，配置同服务端实际扣除的消耗，取同类型的相加获得
	EndRewCount string `json:"end_rew_count"`
}
