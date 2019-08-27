/*此类自动生成,请勿修改*/
package template

/*装备宝库积分商城配置*/
type BaoKuJiFenTemplateVO struct {

	//id
	Id int `json:"id"`

	//宝库类型
	Type int32 `json:"type"`

	//最小转数
	ZhuanshuMin int32 `json:"zhuanshu_min"`

	//最大转数
	ZhuanshuMax int32 `json:"zhuanshu_max"`

	//最小等级
	LevelMin int32 `json:"level_min"`

	//最大等级
	LevelMax int32 `json:"level_max"`

	//开天男兑换物品id
	ItemIdKaiTianNan int32 `json:"item_id"`

	//开天女兑换物品id
	ItemIdKaiTianNv int32 `json:"item_id_kaitian_nv"`

	//弈剑男兑换物品id
	ItemIdYiJianNan int32 `json:"item_id_yijian_nan"`

	//弈剑女兑换物品id
	ItemIdYiJianNv int32 `json:"item_id_yijian_nv"`

	//破月男兑换物品id
	ItemIdPoYueNan int32 `json:"item_id_poyue_nan"`

	//破月女兑换物品id
	ItemIdPoYueNv int32 `json:"item_id_poyue_nv"`

	//一次性兑换数量
	BuyCount int32 `json:"buy_count"`

	//消耗积分
	UseJiFen int32 `json:"use_jifen"`

	//每天最多兑换次数
	MaxCount int32 `json:"max_count"`

	//排序
	Pos int32 `json:"pos"`
}
