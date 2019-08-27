/*此类自动生成,请勿修改*/
package template

/*结婚常量配置*/
type MarryTemplateVO struct {

	//id
	Id int `json:"id"`

	//结婚亲密度最低要求
	QinmiduHoutaiType string `json:"qinmidu_houtai_type"`

	//结婚亲密度最低要求
	MarryQinmidu string `json:"marry_qinmidu"`

	//离婚亲密度降低万分比
	DivorceQinmidu int32 `json:"divorce_qinmidu"`

	//婚戒培养最大等级差
	RingLevelGap int32 `json:"ring_level_gap"`

	//第一场婚宴开始的时间
	MarryFirstTime string `json:"marry_first_time"`

	//一天内举办的婚礼场次
	MarryAmount int32 `json:"marry_amount"`

	//婚宴场景id
	MarryMapId int32 `json:"marry_map_id"`

	//婚车生物表id
	CarId int32 `json:"car_id"`

	//婚车旅游地图id
	CarMapId int32 `json:"car_map_id"`

	//喜糖掉落半径
	Radius int32 `json:"radius"`

	//每场婚礼的时间(毫秒)
	MarryTime int32 `json:"marry_time"`

	//每场婚礼的清场时间(毫秒)
	QingChangTime int32 `json:"qingchang_time"`

	//结婚纪念满了之后赠送物品
	MarryItem string `json:"marry_item"`

	//结婚纪念满了之后赠送物品数量
	MarryItemCount string `json:"marry_item_count"`

	//结婚纪念满了之后赠送时装Id
	FashionId int32 `json:"fashion_id"`

	//定情信物索要CD时间(毫秒)
	XinwuCd int64 `json:"xinwu_cd"`
}
