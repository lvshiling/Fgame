/*此类自动生成,请勿修改*/
package template

/*活跃度配置*/
type HuoYueTemplateVO struct {

	//id
	Id int `json:"id"`

	//关联功能开启
	ModuleOpenedId int32 `json:"module_opened_id"`

	//完成次数上限
	RewardCountLimit int32 `json:"reward_count_limit"`

	//备注
	Remarks string `json:"remarks"`

	//任务前端排列顺序
	pos int32 `json:"pos"`
}
