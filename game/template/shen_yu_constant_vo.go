/*此类自动生成,请勿修改*/
package template

/*神域常量配置*/
type ShenYuConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//神域钥匙保存的最大数量
	KeyMax int32 `json:"key_max"`

	//神域钥匙最小死亡不掉落数量
	KeyKeepMin int32 `json:"key_min"`

	//神域钥匙死亡掉落万分比
	KeyDropPercent int32 `json:"key_percent"`

	//神域钥匙死亡掉落CD
	KeyDropCD int32 `json:"bekill_cd"`

	//击杀玩家掉落最小堆数
	MinStack int32 `json:"min_stack"`

	//击杀玩家掉落最大堆数
	MaxStack int32 `json:"max_stack"`

	//掉落存活时间
	ExistTime int32 `json:"exist_time"`

	//掉落保护时间
	ProtectedTime int32 `json:"protected_time"`

	//掉落物品失效时间
	FailTime int32 `json:"fail_time"`

	//初始钥匙数量
	InitKeyNum int32 `json:"key_begin"`

	//重置钥匙数量
	ResetKeyNum int32 `json:"key_jiesuan"`
}
