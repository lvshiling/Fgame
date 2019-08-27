/*此类自动生成,请勿修改*/
package template

/*属性配置*/
type TradeConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//玩家充值小于额度可以回购
	NeedChongzhi int32 `json:"need_chongzhi"`

	//上架时间超过一定时间
	ShangjiaTime int32 `json:"shangjia_time"`

	//下架时间
	XiajiaTime int32 `json:"xiajia_time"`

	//每天回购总量
	GmMoneyMax int32 `json:"gm_money_max"`

	//每个玩家回购的最大值
	PlayerMoneyMax int32 `json:"player_money_max"`

	//个人上架数量
	PersonalCountMax int32 `json:"personal_count_max"`

	//所有数量
	AllCountMax int32 `json:"all_count_max"`

	//回购最大值
	HuigouPriceMax int32 `json:"huigou_price_max"`
}
