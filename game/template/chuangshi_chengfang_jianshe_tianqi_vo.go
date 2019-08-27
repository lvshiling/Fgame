/*此类自动生成,请勿修改*/
package template

/*创世城池建设天气模板配置*/
type ChuangShiChengFangJianSheTianQiTemplateVO struct {

	//id
	Id int `json:"id"`

	//守城初始天气buffid
	ShouChengTianqiBuffId int32 `json:"shoucheng_tianqi_buff_id"`

	//攻城初始天气buffid
	GongChengTianqiBuffId int32 `json:"gongcheng_tianqi_buff_id"`

	//攻城使用万法后的天气buffid
	GongChengWanfaBuffId int32 `json:"gongcheng_wanfa_buff_id"`

	//天气激活物品id
	TianqiItemId int32 `json:"tianqi_item_id"`

	//天气激活物品数量
	TianqiItemCount int32 `json:"tianqi_item_count"`

	//天气激活有效时间
	ValidTime int64 `json:"jihuo_chixu_time"`

	//天气名字
	Name string `json:"name"`
}
