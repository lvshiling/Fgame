/*此类自动生成,请勿修改*/
package template

/*物品配置*/
type ItemTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//类型
	Type int32 `json:"type"`

	//子类型
	SubType int32 `json:"sub_type"`

	//背包类型
	BagType int32 `json:"bag_type"`

	//任务物品标记
	QuestFlag int32 `json:"quest_flag"`

	//是否能够放入仓库
	Storage int32 `json:"storage"`

	//掉落时效
	LiveTime int32 `json:"live_time"`

	//使用组合字段
	UseFlag int32 `json:"use_flag"`

	//绑定类型
	BindType int32 `json:"bind_type"`

	//冷却类型
	CdGroup int32 `json:"cd_group"`

	//冷却时间
	CdTime int32 `json:"cd_time"`

	//购买银两
	BuySilver int32 `json:"buy_silver"`

	//购买元宝
	BuyGold int32 `json:"buy_gold"`

	//购买元宝
	BuyBindgold int32 `json:"buy_bindgold"`

	//销售比例
	SaleRate int32 `json:"sale_rate"`

	//最大叠加上限
	MaxOverlap int32 `json:"max_overlap"`

	//物品生效时间
	ShengXiaoTime string `json:"shengxiao_time"`

	//物品时效类型
	LimitTimeType int32 `json:"limit_time_type"`

	//物品时效
	LimitTime string `json:"limit_time"`

	//职业需求
	NeedProfession int32 `json:"need_profession"`

	//性别需求
	NeedGender int32 `json:"need_gender"`

	//转生需要
	NeedZhuanShu int32 `json:"need_zhuanshu"`

	//等级需求
	NeedLevel int32 `json:"need_level"`

	//相关参数1
	TypeFlag1 int32 `json:"type_flag1"`

	//相关参数2
	TypeFlag2 int32 `json:"type_flag2"`

	//相关参数3
	TypeFlag3 int32 `json:"type_flag3"`

	//相关参数4
	TypeFlag4 int32 `json:"type_flag4"`

	//相关参数5
	TypeFlag5 int32 `json:"type_flag5"`

	//装备模型
	BoyModel int32 `json:"boy_model"`

	//装备模型
	GirlModel int32 `json:"girl_model"`

	//品质
	Quality int32 `json:"quality"`

	//每日限制次数
	LimitTimeDay int32 `json:"limit_time_day"`

	//总限制次数
	LimitTimeAll int32 `json:"limit_time_all"`

	//物品图标
	IconItem int32 `json:"icon_item"`

	//掉落图标
	IconDrop int32 `json:"icon_drop"`

	//批量使用
	UseBatched int32 `json:"use_batched"`

	//分类
	Classify int32 `json:"classify"`

	//是否提示使用
	IsPopup int32 `json:"is_popup"`

	//标签
	Tag int32 `json:"tag"`

	//将物品放入仓库获得的贡献值
	UnionGet int32 `json:"union_get"`

	//将物品取出仓库消耗的贡献值
	UnionUse int32 `json:"union_use"`

	//关联运营活动排行榜id
	RankGroup string `json:"rank_group"`

	//交易表id
	MarketId int32 `json:"market_id"`

	//最小交易价值
	MarketMinPrice int32 `json:"market_min_price"`

	//回购价格
	HuigouPrice int32 `json:"huigou_price"`

	//拆解获得物品id
	ChaiJieItemId string `json:"chaijie_item_id"`

	//拆解获得数量
	ChaiJieItemCount string `json:"chaijie_item_count"`
}
