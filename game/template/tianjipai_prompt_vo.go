/*此类自动生成,请勿修改*/
package template

/*天机牌提示配置*/
type TianJiPaiPromptTemplateVO struct {

	//id
	Id int `json:"id"`

	//关联到功能开启表
	ModuleOpenedId int32 `json:"module_opened_id"`

	//是否可以一键完成
	IsAuto int32 `json:"is_auto"`

	//奖励预览
	RewardPreview string `json:"reward_preview"`

	//备注
	Remarks string `json:"remarks"`
}
