/*此类自动生成,请勿修改*/
package template

/*抢龙椅配置*/
type DragonChairTemplateVO struct {

	//id
	Id int `json:"id"`

	//元宝底数
	FirstGold int32 `json:"first_gold"`

	//元宝系数1
	CoefficientGold float64 `json:"coefficient_gold"`

	//元宝固定值
	ValueGold int32 `json:"value_gold"`

	//属性底数id
	FirstAttrId int32 `json:"first_attr_id"`

	//属性系数1
	CoefficientAttr float64 `json:"coefficient_attr"`

	//属性固定值id
	ValueAttrId int32 `json:"value_attr_id"`

	//称号id
	TitleId int32 `json:"title_id"`

	//膜拜给予的银两奖励
	SilverWorship int32 `json:"silver_worship"`

	//膜拜给予的绑元奖励
	BindGoldWorship int32 `json:"bind_gold_worship"`

	//膜拜给予的经验奖励
	ExpWorship int32 `json:"exp_worship"`

	//膜拜给予的经验点奖励
	ExpPointWorship int32 `json:"exp_point_worship"`

	//膜拜给予的物品奖励
	ItemWorship string `json:"item_worship"`

	//膜拜给予的物品奖励数量
	ItemCountWorship string `json:"item_count_worship"`

	//每日最大膜拜次数
	WorshipCount int32 `json:"worship_count"`

	//每次膜拜给金库增加银两数量
	WorshipChestSilver int32 `json:"worship_chest_silver"`

	//每次自动往帝王金库添加的银两数量
	AutoSilver int32 `json:"auto_silver"`

	//向帝王金库自动添加银两的间隔时间(毫秒)
	AutoSilverTime int32 `json:"auto_silver_time"`

	//国库存银的最大上限
	ChestMax int32 `json:"chest_max"`

	//当占领时间不在所有区间内的每分钟产出
	ChanchuSilver0 int32 `json:"chanchu_silver0"`

	//占领时间区间
	ZhanlingTime1 string `json:"zhanling_time1"`

	//第一个占领区间产出的银两
	ChanchuSilver1 int32 `json:"chanchu_silver1"`

	//占领时间区间2
	ZhanlingTime2 string `json:"zhanling_time2"`

	//第二个占领区间产出的银两
	ChanchuSilver2 int32 `json:"chanchu_silver2"`

	//占领时间区间3
	ZhanlingTime3 string `json:"zhanling_time3"`

	//第三个占领区间产出的银两
	ChanchuSilver3 int32 `json:"chanchu_silver3"`

	//占领时间区间4
	ZhanlingTime4 string `json:"zhanling_time4"`

	//第四个占领区间产出的银两
	ChanchuSilver4 int32 `json:"chanchu_silver4"`

	//特殊掉落
	SpecialDrop string `json:"special_drop"`

	//普通掉落
	CommonDrop string `json:"common_drop"`

	//获得时间(毫秒)
	DropTime1 int32 `json:"drop_time1"`

	//获得时间(毫秒)
	DropTime2 int32 `json:"drop_time2"`

	//获得时间(毫秒)
	DropTime3 int32 `json:"drop_time3"`

	//获得时间(毫秒)
	DropTime4 int32 `json:"drop_time4"`

	//返还的元宝的万分比
	GoldPercent int32 `json:"gold_percent"`
}
