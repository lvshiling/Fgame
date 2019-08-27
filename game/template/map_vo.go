/*此类自动生成,请勿修改*/
package template

/*地图配置*/
type MapTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//资源
	Resource string `json:"resource"`

	//类型
	Type int32 `json:"type"`

	//脚本类型
	ScriptType int32 `json:"script_type"`

	//存在时间
	LastTime int32 `json:"last_time"`

	//副本失败时间
	PointsTime int32 `json:"points_time"`

	//出生地点
	BirthPosX float64 `json:"birthPosX"`

	//出生地点
	BirthPosY float64 `json:"birthPosY"`

	//出生地点
	BirthPosZ float64 `json:"birthPosZ"`

	//回城复活id
	RebornId int32 `json:"reborn_id"`

	//复活地点x
	RebornX float64 `json:"reborn_x"`

	//复活地点y
	RebornY float64 `json:"reborn_y"`

	//复活地点z
	RebornZ float64 `json:"reborn_z"`

	//朝向
	Orientation int32 `json:"orientation"`

	//等级
	Level int32 `json:"level"`

	//限制pk模式
	LimitPkMode int32 `json:"limit_pkmode"`

	//pk模式
	PkMode int32 `json:"pk_mode"`

	//限制重生
	LimitReborn int32 `json:"limit_reborn"`

	//限制打坐
	LimitSit int32 `json:"limit_sit"`

	//限制使用技能类型
	LimitSpellType int32 `json:"limit_spell_type"`

	//限制带出的物品
	LimitItemId int32 `json:"limit_item_id"`

	//限制带出的状态
	LimitStatusId int32 `json:"limit_status_id"`

	//进入地图的默认阵营
	SetFaction int32 `json:"set_faction"`

	//进入地图的需要的银两
	NeedSilver int32 `json:"need_silver"`

	//进入地图的需要的元宝
	NeedGold int32 `json:"need_gold"`

	//需要物品id
	NeedItemId string `json:"need_item_id"`

	//需要物品数量
	NeedItemNum string `json:"need_item_num"`

	//需求等级
	ReqLev int32 `json:"req_lev"`

	//限制使用物品
	LimitUseitem int32 `json:"limit_useitem"`

	//玩家限制数量
	PlayerLimitCount int32 `json:"player_limit_count"`

	//需要前置任务id
	QuestId int32 `json:"quest_id"`

	//最低战斗力限制
	SuggestPower int64 `json:"suggest_power"`

	//需要功能开启id
	ReqModuleId int32 `json:"req_module_id"`

	//安全区
	SafeArea string `json:"safe_area"`

	//动态地图上限
	OpenLimit int32 `json:"open_limit"`

	//限制骑马
	LimitRideHorse int32 `json:"limit_ride_horse"`

	//奖励预览
	RewardPreview string `json:"reward_previwe"`

	//展示boss信息
	ShowBossInfo int32 `json:"show_boss_info"`

	//天气id
	WeatherId string `json:"weather_id"`

	//切换地图显示的图片
	MapNameImage string `json:"map_name_image"`

	//限制挂机
	LimitHook int32 `json:"limit_hook"`

	//小地图偏移
	OffsetX int32 `json:"offsetX"`

	//小地图偏移
	OffsetY int32 `json:"offsetY"`

	//缩放比例
	Scale float64 `json:"scale"`

	//旋转角度
	Rotate int32 `json:"rotate"`

	//旋转角度
	RotationX int32 `json:"rotation_x"`

	//旋转角度
	RotationY int32 `json:"rotation_y"`

	//距离
	Distance int32 `json:"distance"`

	//音效
	Music string `json:"music"`

	//限制弹窗
	LimitWindow int32 `json:"limit_window"`

	//复活类型
	ResurrectionType int32 `json:"resurrection_type"`

	//复活时间
	ResurrectionTime int32 `json:"resurrection_time"`

	//限制区
	XianzhiquyuId string `json:"xianzhiquyu_id"`

	//安全区
	AnquanquId string `json:"anquanqu_id"`

	//该地图是否掉落杀气
	IsShaqiDrop int32 `json:"is_shaqi_drop"`

	//该地图是否掉落声威值
	IsShengWeiDrop int32 `json:"is_shengwei_drop"`

	//是否支持飞鞋传送
	IsChuansong int32 `json:"is_chuansong"`

	//飞鞋消耗
	FeixieCount int32 `json:"feixie_count"`

	//是否支持仙盟求救
	IsFeixie int32 `json:"is_feixie"`

	//是否隐藏称号
	IsTitle int32 `json:"is_title"`

	//pk保护等级
	ProtectLevel int32 `json:"protect_level"`

	//是否pk状态切换场景
	IsPkScene int32 `json:"is_pk_scene"`

	//是否场景过场保护
	IsSceneProtect int32 `json:"is_scene_protect"`

	//是否掉落杀戮之心
	IsShaLuDrop int32 `json:"is_shalu_drop"`

	//血池
	IsXueChi int32 `json:"is_xuechi"`

	//限制传送类型
	SceneXianZhi int64 `json:"scene_xianzhi"`
}
