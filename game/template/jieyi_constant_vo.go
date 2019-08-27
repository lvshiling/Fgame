/*此类自动生成,请勿修改*/
package template

/*结义常量配置*/
type JieYiConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//威名改名卡id
	ChangeNameItemId int32 `json:"weiming_gaiming_item_id"`

	//求援CD
	QiuYuanCD int64 `json:"qiuyuan_cd"`

	//发布结义信息CD
	FaBuCD int64 `json:"fabu_jieyi_cd"`

	//结义信息存在最大的时间
	FaBuJieYiMaxTime int64 `json:"fabu_jieyi_time_max"`

	//存储的最大声威值
	ShengWeiMax int64 `json:"shengwei_max"`

	//邀请结义CD
	YaoQingCD int64 `json:"yaoqing_cd"`

	//邀请有效时间
	YaoQingExistTime int64 `json:"yaoqing_guoqi_time"`

	//解除结义CD
	JieChuCD int64 `json:"jiechu_jieyi_cd"`

	//背包内声威掉落概率
	DropRate int32 `json:"shengwei_drop_rate"`

	//背包内声威掉落最小比例
	DropPercentMin int32 `json:"shengwei_drop_percent_min"`

	//背包内声威掉落最大比例
	DropPercentMax int32 `json:"shengwei_drop_percent_max"`

	//声威掉落冷却时间
	DropCD int64 `json:"shengwei_drop_cd"`

	//声威掉落保护时间
	DropProtectedTime int64 `json:"shengwei_drop_protected_time"`

	//声威掉落存活时间
	DropFailTime int64 `json:"shengwei_drop_fail_time"`

	//背包声威掉落最小堆数
	DropMinStack int32 `json:"shengwei_drop_min_stack"`

	//背包声威掉落最大堆数
	DropMaxStack int32 `json:"shengwei_drop_max_stack"`

	//被击杀声威被系统回收比例
	DropSystemReturn int32 `json:"shengwei_drop_system_return"`

	//结义最大人数
	MaxPeopleNum int32 `json:"jieyi_player_max"`
}
