package scene

import (
	"fgame/fgame/game/common/common"
	marrytypes "fgame/fgame/game/marry/types"
)

const (
	//婚礼结束后10s清场
	EndDelayTime = int64(10 * common.SECOND)
)

//结婚场景状态
type MarrySceneStatusType int32

const (
	//结婚场景初始化
	MarrySceneStatusTypeInit MarrySceneStatusType = 1 + iota
	//巡游
	MarrySceneStatusCruise
	//婚宴
	MarrySceneStatusBanquet
)

type MarryHeroism struct {
	PlayerId int64
	Name     string
	Heroism  int64
}

//豪气值列表排序
type MarryHeroismList []*MarryHeroism

func (mhl MarryHeroismList) Len() int {
	return len(mhl)
}

func (mhl MarryHeroismList) Less(i, j int) bool {
	return mhl[i].Heroism < mhl[j].Heroism
}

func (mhl MarryHeroismList) Swap(i, j int) {
	mhl[i], mhl[j] = mhl[j], mhl[i]
}

func CreateMarryHeroism(playerId int64, name string, heroism int64) *MarryHeroism {
	d := &MarryHeroism{
		PlayerId: playerId,
		Name:     name,
		Heroism:  heroism,
	}
	return d
}

type MarryData struct {
	Period      int32
	Status      MarrySceneStatusType
	PlayerId    int64
	PlayerName  string
	PlayerRole  int32
	PlayerSex   int32
	SpouseId    int64
	SpouseName  string
	SpouseRole  int32
	SpouseSex   int32
	Grade       marrytypes.MarryBanquetSubTypeWed
	HunCheGrade marrytypes.MarryBanquetSubTypeHunChe
	SugarGrade  marrytypes.MarryBanquetSubTypeSugar
	HeroismList []*MarryHeroism
}

func CreateMarryData(period int32,
	status MarrySceneStatusType,
) *MarryData {
	d := &MarryData{
		Period:      period,
		Status:      status,
		HeroismList: make([]*MarryHeroism, 0, marrytypes.HeroisListLen),
	}
	return d
}

type MarrySceneStatusData struct {
	Id         int64
	Period     int32
	Grade      marrytypes.MarryBanquetSubTypeWed //酒席档次
	Status     MarrySceneStatusType
	PlayerId   int64
	SpouseId   int64
	PlayerName string
	Role       int32
	Sex        int32
	SpouseName string
	SpouseRole int32
	SpouseSex  int32
}

func CreateMarrySceneStatusData(id int64, period int32, grade marrytypes.MarryBanquetSubTypeWed, status MarrySceneStatusType, playerId int64, spouseId int64, playerName string, spouseName string, role int32, sex int32, spouseRole int32, spouseSex int32) *MarrySceneStatusData {
	d := &MarrySceneStatusData{
		Id:         id,
		Period:     period,
		Grade:      grade,
		Status:     status,
		PlayerId:   playerId,
		SpouseId:   spouseId,
		PlayerName: playerName,
		SpouseName: spouseName,
		Role:       role,
		Sex:        sex,
		SpouseRole: spouseRole,
		SpouseSex:  spouseSex,
	}
	return d
}

type MarrySceneGift struct {
	PlayerId int64
	SpouseId int64
	ItemMap  map[int32]int32
	Silver   int64
}

func CreateMarrySceneGift(playerId int64, spouseId int64, roses int32, silver int64) *MarrySceneGift {
	d := &MarrySceneGift{
		PlayerId: playerId,
		SpouseId: spouseId,
		ItemMap:  make(map[int32]int32),
		Silver:   silver,
	}
	return d
}
