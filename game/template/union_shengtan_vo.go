/*此类自动生成,请勿修改*/
package template

/*仙盟圣坛模板配置*/
type UnionShengTanTemplateVO struct {

	//id
	Id int `json:"id"`

	//进入多少毫秒后获得奖励
	FirstTime int64 `json:"frist_tiem"`

	//多少毫秒发放一次奖励
	RewTime int64 `json:"rew_tiem"`

	//百年仙酿使用上限
	ExpAddItemLimit int32 `json:"exp_add_item_limit"`

	//刷新时间
	XiaoguaiTime int32 `json:"xiaoguai_time"`

	//圣坛id
	ShengtanId int32 `json:"shengtan_id"`

	//公告比例
	GonggaoHpPercent int32 `json:"gonggao_hp_percent"`

	//刷怪开始时间
	ShuaguaiBeginTime int32 `json:"shuaguai_begin_time"`
}
