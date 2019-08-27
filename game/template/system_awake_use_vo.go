/*此类自动生成,请勿修改*/
package template

/*附加使用觉醒丹配置*/
type SystemAwakeUseTemplateVO struct {

	//id
	Id int `json:"id"`

	//系统类型
	SysType int32 `json:"type"`

	//需要的最低系统阶别
	NeedNumber int32 `json:"need_number"`

	//觉醒使用物品id
	UseItem int32 `json:"item_id"`
}
