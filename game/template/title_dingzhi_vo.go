/*此类自动生成,请勿修改*/
package template

/*称号定制配置*/
type TitleDingZhiTemplateVO struct {

	//id
	Id int `json:"id"`

	//称号名称
	Name string `json:"name"`

	//排序
	Sort int32 `json:"sort"`

	//需要物品id
	NeedItemId int32 `json:"need_item_id"`

	//需要物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//生命上限
	Hp int64 `json:"hp"`

	//攻击值
	Attack int64 `json:"attack"`

	//防御值
	Defence int64 `json:"defence"`

	//buff的id
	BuffId int32 `json:"buff_id"`

	//获取途径
	Access string `json:"access"`
}
