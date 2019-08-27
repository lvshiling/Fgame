/*此类自动生成,请勿修改*/
package template

/*仙盟捐献配置*/
type UnionDonateTemplateVO struct {

	//id
	Id int `json:"id"`

	//名称
	Name string `json:"name"`

	//类型
	Type int32 `json:"type"`

	//捐献物品id
	DonateItemId int32 `json:"donate_item_id"`

	//捐献物品数量
	DonateItemCount int32 `json:"donate_item_count"`

	//捐献银两
	DonateSilver int32 `json:"donate_silver"`

	//捐献元宝
	DonateGold int32 `json:"donate_gold"`

	//建设性
	DonateBuild int32 `json:"donate_bulid"`

	//贡献
	DonateContribution int32 `json:"donate_contribution"`

	//次数限制
	DonateLimit int32 `json:"donate_limit"`
}
