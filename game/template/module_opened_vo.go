/*此类自动生成,请勿修改*/
package template

/*功能开启配置*/
type ModuleOpenedTemplateVO struct {

	//id
	Id int `json:"id"`

	//功能id
	FuncId int32 `json:"func_id"`

	//名字
	Name string `json:"name"`

	//功能开启id
	OpenedQuestId int32 `json:"opened_quest_id"`

	//开启等级
	OpenedLevel int32 `json:"opened_level"`

	//开启转数
	OpenedZhuanShu int32 `json:"opened_zhuanshu"`

	//开启时间
	OpenedTime int64 `json:"opened_time"`

	//开服时间
	KaifuTime int64 `json:"kaifu_time"`

	//开启后奖励物品id
	RewItem string `json:"rew_item"`

	//开启后奖励物品数量
	RewItemCount string `json:"rew_item_count"`

	//开启后元宝奖励
	RewGold int64 `json:"rew_gold"`

	//开启后绑元奖励
	RewBindgold int64 `json:"rew_bindgold"`

	//开启后银两奖励
	RewYinliang int64 `json:"rew_yinliang"`

	//开启首冲
	OpenedShouchong int32 `json:"opened_shouchong"`

	//开启物品id
	OpenedItemId int32 `json:"opened_item_id"`

	//开启物品数量
	OpenedItemCount int32 `json:"opened_item_count"`

	//前置功能id
	ParentId string `json:"parent_id"`

	//显示功能预告
	IsPreview int32 `json:"is_preview"`

	//描述
	Description string `json:"description"`

	//显示提示框
	Show string `json:"show"`

	//邮件奖励
	MailRewItem string `json:"mail_rew_item"`

	//邮件奖励数量
	MailRewItemCount string `json:"mail_rew_item_count"`

	//邮件内容
	MailDes string `json:"mail_des"`

	//邮件标题
	MailTitle string `json:"mail_title"`

	//关联运营活动时间id
	OpensvId string `json:"opensv_id"`

	//建号几天后开启
	JianHaoDay int32 `json:"jianhao_day"`

	//是否自动开启
	JianHaoIsAuto int32 `json:"jianhao_is_auto"`

	//功能开启所需系统类型
	NeedSysType int32 `json:"need_sys_type"`

	//功能开启所需阶别
	NeedSysNum int32 `json:"need_sys_num"`
}
