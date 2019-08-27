/*此类自动生成,请勿修改*/
package template

/*等级配置*/
type CharacterLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int32 `json:"next_id"`

	//经验
	Experience int64 `json:"experience"`

	//移动速度
	SpeedMove int32 `json:"speed_move"`

	//基础攻击
	BaseAttack int32 `json:"base_attack"`

	//基础防御
	BaseDefense int32 `json:"base_defense"`

	//基础暴击
	BaseCritical int32 `json:"base_critical"`

	//基础坚韧
	BaseTough int32 `json:"base_tough"`

	//基础格挡
	BaseBlock int32 `json:"base_block"`

	//基础破格
	BaseBreak int32 `json:"base_break"`

	//需要物品id3
	NeedItemId3 int32 `json:"need_item_id3"`

	//基础最大血量
	BaseMaxhp int32 `json:"base_maxhp"`

	//基础最大体力
	BaseMaxtp int32 `json:"base_maxtp"`

	//uplve点
	HotrexpModul int32 `json:"hotrexp_modul"`

	//生命丹药使用上线
	MedicinesHpMax int32 `json:"medicines_hp_max"`

	//攻击丹药使用上线
	MedicinesAttackMax int32 `json:"medicines_attack_max"`

	//防御丹药使用上线
	MedicinesDefenceMax int32 `json:"medicines_defence_max"`

	//格挡丹药使用上线
	MedicinesBlockMax int32 `json:"medicines_block_max"`

	//破格丹药使用上线
	MedicinesBreakMax int32 `json:"medicines_break_max"`

	//暴击丹药使用上线
	MedicinesCriticalMax int32 `json:"medicines_critical_max"`

	//坚韧丹药使用上线
	MedicinesToughMax int32 `json:"medicines_tough_max"`

	//hp回复时间
	HpReplyTime int32 `json:"hp_reply_time"`

	//hp回复
	HpReply int32 `json:"hp_reply"`

	//tp回复时间
	TpReplyTime int32 `json:"tp_reply_time"`

	//tp回复
	TpReply int32 `json:"tp_reply"`
}
