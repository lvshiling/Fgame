/*此类自动生成,请勿修改*/
package template

/*消费等级模板配置*/
type CostLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//消费等级
	Level int32 `json:"level"`

	//下一级
	NextId int32 `json:"next_id"`

	//需要历史消费的元宝
	NeedValue int32 `json:"need_value"`

	//坐骑最小次数系数
	MountMinCount int32 `json:"mount_min_count"`

	//坐骑最大次数系数
	MountMaxCount int32 `json:"mount_max_count"`

	//战翼最小次数系数
	WingMinCount int32 `json:"wing_min_count"`

	//战翼最大次数系数
	WingMaxCount int32 `json:"wing_max_count"`

	//暗器最小次数系数
	AnqiMinCount int32 `json:"anqi_min_count"`

	//暗器最大次数系数
	AnqiMaxCount int32 `json:"anqi_max_count"`

	//护体盾最小次数系数
	BodyShieldMinCount int32 `json:"body_shield_min_count"`

	//护体盾最大次数系数
	BodyShieldMaxCount int32 `json:"body_shield_max_count"`

	//护体仙羽最小次数系数
	FeatherMinCount int32 `json:"feather_min_count"`

	//护体仙羽最大次数系数
	FeatherMaxCount int32 `json:"feather_max_count"`

	//盾刺最小次数系数
	ShieldMinCount int32 `json:"shield_min_count"`

	//盾刺最大次数系数
	ShieldMaxCount int32 `json:"shield_max_count"`

	//领域最小次数系数
	FieldMinCount int32 `json:"field_min_count"`

	//领域最大次数系数
	FieldMaxCount int32 `json:"field_max_count"`

	//身法最小次数系数
	ShenfaMinCount int32 `json:"shenfa_min_count"`

	//身法最大次数系数
	ShenfaMaxCount int32 `json:"shenfa_max_count"`

	//婚戒最小次数系数
	MarryRingMinCount int32 `json:"marry_ring_min_count"`

	//婚戒最大次数系数
	MarryRingMaxCount int32 `json:"marry_ring_max_count"`

	//赌石固定掉落包系数
	DushiCount1Count int32 `json:"dushi_count1_count"`

	//赌石固定掉落包系数
	DushiCount2Count int32 `json:"dushi_count2_count"`

	//赌石固定掉落包系数
	DushiCount3Count int32 `json:"dushi_count3_count"`

	//棋局固定掉落包系数
	QijvCount1Count int32 `json:"qijv_count1_count"`

	//棋局固定掉落包系数
	QijvCount2Count int32 `json:"qijv_count2_count"`

	//棋局固定掉落包系数
	QijvCount3Count int32 `json:"qijv_count3_count"`

	//棋局固定掉落包系数
	QijvCount4Count int32 `json:"qijv_count4_count"`

	//棋局固定掉落包系数
	QijvCount5Count int32 `json:"qijv_count5_count"`

	//棋局固定掉落包系数
	QijvCount6Count int32 `json:"qijv_count6_count"`

	//棋局固定掉落包系数
	QijvCount7Count int32 `json:"qijv_count7_count"`

	//棋局固定掉落包系数
	QijvCount8Count int32 `json:"qijv_count8_count"`

	//兵魂升星的最小次数系数
	WeaponUpstarMinCount int32 `json:"weapon_upstar_min_count"`

	//兵魂升星的最大次数系数
	WeaponUpstarMaxCount int32 `json:"weapon_upstar_max_count"`

	//戮仙刃最小次数系数
	LuxianrenMinCount int32 `json:"luxianren_min_count"`

	//戮仙刃最大次数系数
	LuxianrenMaxCount int32 `json:"luxianren_max_count"`

	//法宝的最小次数系数
	FabaoMinCount int32 `json:"fabao_min_count"`

	//法宝的最大次数系数
	FabaoMaxCount int32 `json:"fabao_max_count"`

	//仙体的最小次数系数
	XianTiMinCount int32 `json:"xianti_min_count"`

	//仙体的最大次数系数
	XianTiMaxCount int32 `json:"xianti_max_count"`

	//点星的最小次数系数
	DianXingMinCount int32 `json:"dianxing_min_count"`

	//点星的最大次数系数
	DianXingMaxCount int32 `json:"dianxing_max_count"`

	//噬魂幡的最小次数系数
	ShiHunFanMinCount int32 `json:"shihunfan_min_count"`

	//噬魂幡的最大次数系数
	ShiHunFanMaxCount int32 `json:"shihunfan_max_count"`

	//点星的最小次数系数
	TianmotiMinCount int32 `json:"tianmoti_min_count"`

	//点星的最大次数系数
	TianmotiMaxCount int32 `json:"tianmoti_max_count"`

	//灵童武器最小次数系数灵兵
	LingTongWeaponMinCount int32 `json:"lingtong_weapon_min_count"`

	//灵童武器最大次数系数灵兵
	LingTongWeaponMaxCount int32 `json:"lingtong_weapon_max_count"`

	//灵童坐骑最小次数系数灵骑
	LingTongMountMinCount int32 `json:"lingtong_mount_min_count"`

	//灵童坐骑最大次数系数灵骑
	LingTongMountMaxCount int32 `json:"lingtong_mount_max_count"`

	//灵童战翼最小次数系数灵翼
	LingTongWingMinCount int32 `json:"lingtong_wing_min_count"`

	//灵童战翼最大次数系数灵翼
	LingTongWingMaxCount int32 `json:"lingtong_wing_max_count"`

	//灵童身法最小次数系数灵身
	LingTongShenFaMinCount int32 `json:"lingtong_shenfa_min_count"`

	//灵童身法最大次数系数灵身
	LingTongShenFaMaxCount int32 `json:"lingtong_shenfa_max_count"`

	//灵童领域最小次数系数灵域
	LingTongLingYuMinCount int32 `json:"lingtong_field_min_count"`

	//灵童领域最大次数系数灵域
	LingTongLingYuMaxCount int32 `json:"lingtong_field_max_count"`

	//灵童法宝最小次数系数灵宝
	LingTongFaBaoMinCount int32 `json:"lingtong_fabao_min_count"`

	//灵童法宝最大次数系数灵宝
	LingTongFaBaoMaxCount int32 `json:"lingtong_fabao_max_count"`

	//灵童仙体最小次数系数灵体
	LingTongXianTiMinCount int32 `json:"lingtong_xianti_min_count"`

	//灵童仙体最大次数系数灵体
	LingTongXianTiMaxCount int32 `json:"lingtong_xianti_max_count"`

	//装备宝库固定掉落包系数
	EquipBaoKuCount1Count string `json:"equip_baoku_count1_count"`

	//神器的最小次数系数
	ShenQiMinCount int32 `json:"shenqi_min_count"`

	//神器的最大次数系数
	ShenQiMaxCount int32 `json:"shenqi_max_count"`

	//材料宝库固定掉落包系数
	MaterialBaoKuCount1Count string `json:"cailiao_baoku_count1_count"`

	//宝库固定掉落包系数
	BaoKuBagCount1Count string `json:"baoku_bag_count1_count"`
}
