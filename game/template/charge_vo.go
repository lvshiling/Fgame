/*此类自动生成,请勿修改*/
package template

/*充值模板配置*/
type ChargeTemplateVO struct {

	//id
	Id int `json:"id"`

	//平台类型
	Type int32 `json:"type"`

	//服务器排序
	SubType int32 `json:"sub_type"`

	//金额
	Rmb int32 `json:"rmb"`

	//元宝
	Gold int32 `json:"gold"`

	//首次返还元宝
	FanhuanGold int32 `json:"fanhuan_gold"`

	//首次返还绑元
	FanhuanBindGold int32 `json:"fanhuan_bindgold"`
}
