package types

// 后台规则系统
type CostLevelRuleType int32

const (
	CostLevelRuleTypeMount          CostLevelRuleType = iota + 1 //坐骑
	CostLevelRuleTypeWing                                        //战翼
	CostLevelRuleTypeAnqi                                        //暗器
	CostLevelRuleTypeBodyShield                                  //护盾
	CostLevelRuleTypeFeather                                     //仙羽
	CostLevelRuleTypeShield                                      //盾刺
	CostLevelRuleTypeLingyu                                      //领域
	CostLevelRuleTypeShenfa                                      //身法
	CostLevelRuleTypeGamble                                      //赌石
	CostLevelRuleTypeChess                                       //真龙棋局10
	CostLevelRuleTypeMarryRing                                   //婚戒
	CostLevelRuleTypeWeaponUpstar                                //兵魂升星
	CostLevelRuleTypeMassacre                                    //戮仙刃
	CostLevelRuleTypeFaBao                                       //法宝
	CostLevelRuleTypeXianTi                                      //仙体
	CostLevelRuleTypeDianXing                                    //点星
	CostLevelRuleTypeShiHunFan                                   //噬魂幡
	CostLevelRuleTypeTianMoTi                                    //天魔体
	CostLevelRuleTypeLingTongWeapon                              //灵兵
	CostLevelRuleTypeLingTongMount                               //灵骑
	CostLevelRuleTypeLingTongWing                                //灵翼
	CostLevelRuleTypeLingTongShenFa                              //灵身
	CostLevelRuleTypeLingTongLingYu                              //灵域
	CostLevelRuleTypeLingTongXianTi                              //灵体
	CostLevelRuleTypeLingTongFaBao                               //灵宝
	CostLevelRuleTypeEquipBaoKu                                  //装备宝库
	CostLevelRuleTypeShenQi                                      //神器
	CostLevelRuleTypeMaterialBaoKu                               //材料宝库
	CostLevelRuleTypeRingBaoKu                                   //特戒宝库

)

// vip等级
type VipLevelType int32

const (
	VipLevelTypeZero VipLevelType = iota
	VipLevelTypeOne
	VipLevelTypeTwo
	VipLevelTypeThree
	VipLevelTypeFour
	VipLevelTypeFive
	VipLevelTypeSix
	VipLevelTypeSeven
	VipLevelTypeEight
	VipLevelTypeNine
	VipLevelTypeTen
	VipLevelTypeEleven
	VipLevelTypeTwelve
	VipLevelTypeThirteen
	VipLevelTypeFourteen
	VipLevelTypeFifteen
)
