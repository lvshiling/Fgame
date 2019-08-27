package types

import (
	inventorytypes "fgame/fgame/game/inventory/types"
	playertypes "fgame/fgame/game/player/types"
)

// 天赋数据
type TalentInfo struct {
	SkillId int32           `json:"SkillId"` //技能id
	Status  SkillStatusType `json:"Status"`  //技能状态
	Type    SkillType       `json:"Type"`    //技能类型
}

func NewTalentInfo(skillId int32, status SkillStatusType, skType SkillType) *TalentInfo {
	talent := &TalentInfo{
		SkillId: skillId,
		Status:  status,
		Type:    skType,
	}
	return talent
}

//配偶宝宝数据
type CoupleBabyData struct {
	BabyId     int64         //宝宝id
	Quality    int32         //品质
	LearnLevel int32         //读书等级
	Danbei     int32         //属性单倍
	TalentList []*TalentInfo //天赋列表
}

func NewCoupleBabyData(babyId int64, quality, learnLevel, attrBeiShu int32, skillList []*TalentInfo) *CoupleBabyData {
	data := &CoupleBabyData{
		BabyId:     babyId,
		Quality:    quality,
		LearnLevel: learnLevel,
		Danbei:     attrBeiShu,
		TalentList: skillList,
	}
	return data
}

//宝宝卡物品属性
type BabyPropertyData struct {
	*inventorytypes.ItemPropertyDataBase
	Sex        playertypes.SexType `json:"Sex"`        //宝宝性别
	Quality    int32               `json:"Quality"`    //品质
	TalentList []*TalentInfo       `json:"TalentList"` //天赋列表
	Danbei     int32               `json:"Danbei"`     //属性单倍
}

func (gd *BabyPropertyData) InitBase() {
	if gd.ItemPropertyDataBase == nil {
		gd.ItemPropertyDataBase = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
}

func CreateBabyPropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	d := &BabyPropertyData{}
	d.ItemPropertyDataBase = base
	return d
}

//怀孕信息
type PregnantInfo struct {
	PregnantTime int64 `json:"Sex"`      //怀孕时间(PS:json定义写错成Sex)
	TonicPro     int32 `json:"TonicPro"` //补品值
}
