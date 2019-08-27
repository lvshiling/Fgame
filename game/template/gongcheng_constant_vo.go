/*此类自动生成,请勿修改*/
package template

/*神兽攻城配置*/
type GongChengConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//麒麟地图
	MapId1 int32 `json:"map_id1"`

	//火凤地图
	MapId2 int32 `json:"map_id2"`

	//毒龙地图
	MapId3 int32 `json:"map_id3"`

	//麒麟id
	BossId1 int32 `json:"boss_id1"`

	//火凤id
	BossId2 int32 `json:"boss_id2"`

	//毒龙id
	BossId3 int32 `json:"boss_id3"`

	//BOSS刷新时间
	BossTime int32 `json:"boss_time"`

	//活动人数上限
	PlayerLimitCount int32 `json:"player_limit_count"`

	//麒麟来袭对应的系统
	SystemName1 string `json:"system_name1"`

	//火凤来袭对应的系统
	SystemName2 string `json:"system_name2"`

	//毒龙来袭对应的系统
	SystemName3 string `json:"system_name3"`

	//显示固定最高级掉宝符ID
	YulanId int32 `json:"yulan_id"`

	//金银秘窟地图
	MapId4 int32 `json:"map_id4"`

	//金银秘窟活动人数上限
	MoneyLimitCount int32 `json:"money_limit_count"`

	//采集次数
	CaiJiCountLimit int32 `json:"caiji_count_limit"`

	//特殊采集生物id
	CaijiBiologyId string `json:"caiji_biology_id"`
}
