/*此类自动生成,请勿修改*/
package template

/*四神遗迹常量配置*/
type FourGodTemplateVO struct {

	//id
	Id int `json:"id"`

	//四神遗迹钥匙保存的最大数量
	KeyMax int32 `json:"key_max"`

	//假人刷新时间(毫秒)
	RobotTime int32 `json:"robot_time"`

	//假人属性系数(万分比)
	RobotNum int32 `json:"robot_num"`

	//假人身上钥匙数量区间最小值
	RobotKeyMin int32 `json:"robot_key_min"`

	//假人身上钥匙数量区间最大值
	RobotKeyMax int32 `json:"robot_key_max"`

	//特殊怪刷新时间(毫秒)
	SpecialTime int32 `json:"special_time"`

	//特殊怪持续刷新的时间(毫秒)
	SpecialProbabilityTime int32 `json:"special_probability_time"`

	//特殊怪刷新概率(万分比)
	SpecialProbability int32 `json:"special_probability"`

	//特殊怪的最大数量
	SpecialNum int32 `json:"special_num"`

	//触发黑衣人BUFF的物品id
	ItemId int32 `json:"item_id"`

	//使用蒙面衣获得的BUFFID
	BlackerBuffId int32 `json:"blacker_buff_id"`

	//进入活动后BOSS的刷新时间(毫秒)
	BossTime int32 `json:"boss_time"`

	//boss id
	BossId int32 `json:"boss_id"`

	//特殊怪 id
	SpecialId int32 `json:"special_id"`

	//宝箱 id
	BoxId int32 `json:"box_id"`

	//黑衣人模型id
	BlackerModelId int32 `json:"blacker_model_id"`

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

	//宝箱采集时间(毫秒)
	BoxTime int32 `json:"box_time"`

	//死神BOSS发送公告的血量百分比
	GongGao string `json:"gonggao"`
}
