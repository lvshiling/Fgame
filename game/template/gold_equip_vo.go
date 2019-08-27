/*此类自动生成,请勿修改*/
package template

/*元神金装配置*/
type GoldEquipTemplateVO struct {

	//id
	Id int `json:"id"`

	//套装id
	SuitGroup int32 `json:"suit_group"`

	//强化id
	GoldequipStrenId int32 `json:"goldequip_stren_id"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//生命加成万分比
	HpPercent int32 `json:"hp_percent"`

	//攻击加成万分比
	AttPercent int32 `json:"att_percent"`

	//防御加成万分比
	DefPercent int32 `json:"def_percent"`

	//新版强化起始id
	GoldeuipUpstarId int32 `json:"goldeuip_upstar_id"`

	//开光起始id
	GoldeuipOpenlightId int32 `json:"goldeuip_openlight_id"`

	//附加属性id
	FujiaId int32 `json:"fujia_id"`

	//激活第1条属性所需强化等级
	NeedStrengthen1 int32 `json:"need_strengthen_1"`

	//激活第2条属性所需强化等级
	NeedStrengthen2 int32 `json:"need_strengthen_2"`

	//激活第3条属性所需强化等级
	NeedStrengthen3 int32 `json:"need_strengthen_3"`

	//激活第4条属性所需强化等级
	NeedStrengthen4 int32 `json:"need_strengthen_4"`

	//激活第5条属性所需强化等级
	NeedStrengthen5 int32 `json:"need_strengthen_5"`

	//激活第6条属性所需强化等级
	NeedStrengthen6 int32 `json:"need_strengthen_6"`

	//吞噬经验
	TunshiExp int32 `json:"tushi_exp"`

	//掉落无双神器相关物品
	TunshiDrop int `json:"tunshi_drop"`

	//神铸关联序列ID
	ShenzhuequipId int `json:"shenzhuequip_id"`

	//神铸装备等级
	ShenzhuequipLevel int32 `json:"shenzhuequip_level"`

	//神铸锻魂等级最高级
	DuanhunLevelMax int32 `json:"duanhun_level_max"`
}
