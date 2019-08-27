package types

type GuideReplicaType int32

const (
	GuideReplicaTypeDefault GuideReplicaType = iota //引导杀完副本怪 (默认)
	GuideReplicaTypeMoJian                          //魔剑副本
	GuideReplicaTypeCatDog                          //猫狗大战
	GuideReplicaTypeRescure                         //救援小医仙
)

func (t GuideReplicaType) Valid() bool {
	switch t {
	case GuideReplicaTypeDefault,
		GuideReplicaTypeMoJian,
		GuideReplicaTypeCatDog,
		GuideReplicaTypeRescure:
		return true
	}

	return false
}

var (
	guideReplicaTypeStringMap = map[GuideReplicaType]string{
		GuideReplicaTypeDefault: "引导杀完副本怪 (默认)",
		GuideReplicaTypeMoJian:  "魔剑副本",
		GuideReplicaTypeCatDog:  "猫狗大战",
		GuideReplicaTypeRescure: "救援小医仙",
	}
)

func (t GuideReplicaType) String() string {
	return guideReplicaTypeStringMap[t]
}

type CatDogKillType int32

const (
	CatDogKillTypeDefault CatDogKillType = iota //默认
	CatDogKillTypeCat                           //猫
	CatDogKillTypeDog                           //狗
)
