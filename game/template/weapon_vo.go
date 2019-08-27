/*此类自动生成,请勿修改*/
package template

/*兵魂配置*/
type WeaponTemplateVO struct {

	//id
	Id int `json:"id"`

	//名称1
	Name1 string `json:"name_1"`

	//名称2
	Name2 string `json:"name_2"`

	//名称3
	Name3 string `json:"name_3"`

	//兵魂排序
	Pos int32 `json:"pos"`

	//兵魂标识
	Tag int32 `json:"tag"`

	//兵魂类型
	Type int32 `json:"type"`

	//图标
	Icon int32 `json:"icon"`

	//兵魂描述
	Disc string `json:"disc"`

	//属性
	AttrId int32 `json:"attr_id"`

	//激活所需物品ID
	NeedItemId string `json:"activation_need_item_id"`

	//激活所需物品数量
	NeedItemCount string `json:"activation_need_item_count"`

	//激活所需物品ID2
	NeedItemId2 string `json:"activation_need_item_id2"`

	//激活所需物品数量2
	NeedItemCount2 string `json:"activation_need_item_count2"`

	//激活所需物品ID3
	NeedItemId3 string `json:"activation_need_item_id3"`

	//激活所需物品数量3
	NeedItemCount3 string `json:"activation_need_item_count3"`

	//是否可以觉醒
	IsAwaken int32 `json:"is_awaken"`

	//觉醒所需星级
	NeedStar int32 `json:"awaken_need_star"`

	//觉醒所需物品id
	AwakenItemId int32 `json:"awaken_need_item_id"`

	//觉醒所需物品数量
	AwakenItemNum int32 `json:"awaken_need_item_num"`

	//觉醒成功率
	AwakenSuccessRate int32 `json:"awaken_success_rate"`

	//觉醒属性
	AwakenAttr int32 `json:"awaken_attr"`

	//觉醒光效
	Effect string `json:"effect"`

	//获取途径描述
	Achieve string `json:"achieve"`

	//食丹等级上限
	EatDan int32 `json:"eat_dan"`

	//兵魂起始ID
	WeaponUpgradeBeginId int32 `json:"weapon_upgrade_begin_id"`

	//兵魂培养起始ID
	WeaponPeiYangBeginId int32 `json:"weapon_peiyang_begin_id"`

	//左手武器开天
	WpL1 string `json:"wp_l_1"`

	//右手武器开天
	WpR1 string `json:"wp_r_1"`

	//背部挂点开天
	WpB1 string `json:"wp_b_1"`

	//背部挂点开天
	WpA1 string `json:"wp_a_1"`

	//左手武器奕剑
	WpL2 string `json:"wp_l_2"`

	//右手武器奕剑
	WpR2 string `json:"wp_r_2"`

	//背部挂点奕剑
	WpB2 string `json:"wp_b_2"`

	//背部挂点奕剑
	WpA2 string `json:"wp_a_2"`

	//左手武器破月
	WpL3 string `json:"wp_l_3"`

	//右手武器破月
	WpR3 string `json:"wp_r_3"`

	//背部挂点破月
	WpB3 string `json:"wp_b_3"`

	//背部挂点破月
	WpA3 string `json:"wp_a_3"`
}
