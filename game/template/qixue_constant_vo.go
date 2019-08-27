/*此类自动生成,请勿修改*/
package template

/*泣血枪常量配置*/
type QiXueConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//杀戮心掉落概率
	DropRate int32 `json:"resources_drop_rate"`

	//杀戮心掉落万分比下限
	DropPercentMin int32 `json:"resources_drop_percent_min"`

	//杀戮心掉落万分比上限
	DropPercentMax int32 `json:"resources_drop_percent_max"`

	//掉落CD
	DropCd int32 `json:"resources_drop_cd"`

	//杀戮心保护时间
	DropProtectedTime int32 `json:"resources_drop_protected_time"`

	//杀戮心存在时间
	DropFailTime int32 `json:"resources_drop_fail_time"`

	//最小堆数
	DropMinStack int32 `json:"resources_drop_min_stack"`

	//最大堆数
	DropMaxStack int32 `json:"resources_drop_max_stack"`

	//背包掉落杀戮心系统回收系数
	DropSystemReturn int32 `json:"resources_drop_system_return"`
}
