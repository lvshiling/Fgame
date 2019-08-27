package common

type SkillObject interface {
	GetSkillId() int32
	GetLevel() int32
	SetLevel(level int32)
	GetTianFuList() []*TianFuInfo
	SetTianFuLevel(tianFuId int32, level int32)
}

type TianFuInfo struct {
	TianFuId int32 `json:"tianFuId"`
	Level    int32 `json:"level"`
}

func newTianFuInfo(tianFuId int32, level int32) *TianFuInfo {
	tianFuInfo := &TianFuInfo{
		TianFuId: tianFuId,
		Level:    level,
	}
	return tianFuInfo
}

type SkillObjectImpl struct {
	SkillId    int32         `json:skillId`
	Level      int32         `json:level`
	TianFuList []*TianFuInfo `json:"tianFuList"`
}

func (so *SkillObjectImpl) GetSkillId() int32 {
	return so.SkillId
}

func (so *SkillObjectImpl) GetLevel() int32 {
	return so.Level
}

func (so *SkillObjectImpl) SetLevel(level int32) {
	so.Level = level
}

func (so *SkillObjectImpl) GetTianFuList() []*TianFuInfo {
	return so.TianFuList
}

func (so *SkillObjectImpl) SetTianFuLevel(tianFuId int32, level int32) {
	if so.TianFuList == nil {
		so.TianFuList = make([]*TianFuInfo, 0, 3)
	}
	tianFuInfo := so.getTianFuInfo(tianFuId)
	if tianFuInfo == nil {
		tianFuInfo = newTianFuInfo(tianFuId, level)
		so.TianFuList = append(so.TianFuList, tianFuInfo)
		return
	}
	tianFuInfo.Level = level
}

func (so *SkillObjectImpl) getTianFuInfo(tianFuId int32) *TianFuInfo {
	for _, tianFuInfo := range so.TianFuList {
		if tianFuInfo.TianFuId == tianFuId {
			return tianFuInfo
		}
	}
	return nil
}

func CreateSkillObject(skillId int32, level int32, tianFuList []*TianFuInfo) SkillObject {
	so := &SkillObjectImpl{
		SkillId:    skillId,
		Level:      level,
		TianFuList: tianFuList,
	}
	return so
}
