/*此类自动生成,请勿修改*/
package template

/*结义道具配置*/
type JieYiDaoJuTemplateVO struct {

	//id
	Id int `json:"id"`

	//道具类型
	Type int32 `json:"type"`

	//所需物品id
	UseItemId int32 `json:"use_item_id"`

	//所需物品数量
	UseItemCount int32 `json:"use_item_count"`

	//激活的时装id 0不激活
	FashionId int32 `json:"fashion_id"`

	//获得物品id
	GetItemId string `json:"get_item_id"`

	//获得物品数量
	GetItemCount string `json:"get_item_count"`
}
