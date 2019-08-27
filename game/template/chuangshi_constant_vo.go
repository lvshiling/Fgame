/*此类自动生成,请勿修改*/
package template

/*创世常量模板配置*/
type ChuangShiConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//阵营更换消耗物品id
	GenghunUseItemId int32 `json:"genghun_use_item_id"`

	//阵营更换消耗物品数量
	GenghunUseItemCount int32 `json:"genghun_use_item_count"`

	//阵营发言vip
	ZhenyingFayanVipLevel int32 `json:"zhenyingfayan_vip_level"`

	//阵营发言玩家等级
	ZhenyingFayanPlayerLevel int32 `json:"zhenyingfayan_player_level"`

	//战区发言vip
	ZhanquFayanVipLevel int32 `json:"zhanqufayan_vip_level"`

	//战区发言玩家等级
	ZhanquFayanPlayerLevel int32 `json:"zhanqufayan_player_level"`

	//战区发言物品id
	ZhanquFayanUseItemId int32 `json:"zhanqufayan_use_item_id"`

	//战区发言物品数量
	ZhanquFayanUseItemCount int32 `json:"zhanqufayan_use_item_count"`

	//弹劾神王离线时间
	TanheNeedTime int64 `json:"tanhe_need_time"`

	//佣兵最多存在数量
	YongbingCountMax int32 `json:"yongbing_count_max"`

	//神王竞选报名日
	BaomingWeekday string `json:"baoming_weekday"`

	//神王竞选报名时间
	BaomingTime string `json:"baoming_time"`

	//神王选举日
	XuanjuWeekday string `json:"xuanju_weekday"`

	//神王选举时间
	XuanjuTime string `json:"xuanju_time"`

	//城池产出间隔时间
	CityRewTime int64 `json:"city_rew_time"`

	//选择攻城目标限制时间
	GongchengTime int64 `json:"gongcheng_time"`

	//更换阵营限制时间
	JinzhiGenghuanTime int64 `json:"jinzhigenghuan_time"`

	//跟随盟主更换阵营限制时间
	MengzhuGenghuanTime int64 `json:"mengzhu_genghuan_time"`

	//个人工资最大时间
	GerenZiyuanDay int64 `json:"geren_ziyuan_day"`
}
