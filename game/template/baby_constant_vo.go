/*此类自动生成,请勿修改*/
package template

/*宝宝常量模板配置*/
type BabyConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//字母河水物品
	RiverItem int32 `json:"river_item"`

	//补品物品id
	BupinItemId int32 `json:"bupin_item_id"`

	//补品消耗数量
	BupinItemCount int32 `json:"bupin_item_count"`

	//最大食用次数
	BupinMax int32 `json:"bupin_max"`

	//补品随机值
	BupinQujian string `json:"bupin_qujian"`

	//出生倒计时
	BornTime int64 `json:"time_down"`

	//加速出生消耗元宝
	GoldZaoChan int32 `json:"gold_zaochan"`

	//基础宝宝数量
	BabyCount int32 `json:"baby_count"`

	//转世后天赋技能返还的道具万分比
	ZsTianFuReturnRate int32 `json:"zs_tianfu_return_rate"`

	//超生消耗元宝
	GoldChaoSheng string `json:"gold_chaosheng"`

	//四书五经物品
	BooksId string `json:"books_id"`

	//宝宝改名卡
	GaiMingKaId int32 `json:"gaimingka_id"`

	//天赋技能数量
	LimitSkiiNum int32 `json:"limit_skii_num"`

	//宝宝卡id
	BaoBaoCard int32 `json:"baobao_card"`

	//洗练消耗物品
	XiLianItemId int32 `json:"xilian_item_id"`

	//洗练消耗数量
	XiLianItemCount int32 `json:"xilian_item_count"`

	//洗练消耗系数
	XiLianCoefficient int32 `json:"xilian_coefficient"`

	//洗练消耗固定值
	XiLianCoefficientFixed int32 `json:"xilian_coefficient_2"`

	//洗练次数
	XilianNum int32 `json:"xilian_num"`

	//成长系数
	GrowthCoefficient int32 `json:"growth_coefficient"`

	//出生提示
	TishiChushengTime string `json:"tishi_chusheng_time"`
}
