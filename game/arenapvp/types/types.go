package types

import (
	"fgame/fgame/common/lang"
)

//
type ArenapvpType int32

const (
	ArenapvpTypeInit     ArenapvpType = iota //初始
	ArenapvpTypeElection                     //海选
	ArenapvpTypeTop32                        //32进16
	ArenapvpTypeTop16                        //16进8
	ArenapvpTypeTop8                         //8进4
	ArenapvpTypeTop4                         //4进2
	ArenapvpTypeFinals                       //决赛
	ArenapvpTypeChampion                     //冠军

)

func (t ArenapvpType) Valid() bool {
	switch t {
	case ArenapvpTypeElection,
		ArenapvpTypeTop32,
		ArenapvpTypeTop16,
		ArenapvpTypeTop8,
		ArenapvpTypeTop4,
		ArenapvpTypeFinals:
		return true
	}
	return false
}

var (
	stringMap = map[ArenapvpType]string{
		ArenapvpTypeElection: "海选",
		ArenapvpTypeTop32:    "32进16",
		ArenapvpTypeTop16:    "16进8",
		ArenapvpTypeTop8:     "8进4",
		ArenapvpTypeTop4:     "4进2",
		ArenapvpTypeFinals:   "决赛",
	}
)

func (t ArenapvpType) String() string {
	return stringMap[t]
}

// 转换名次
var (
	toNumberMap = map[ArenapvpType]int32{
		ArenapvpTypeElection: 33,
		ArenapvpTypeTop32:    32,
		ArenapvpTypeTop16:    16,
		ArenapvpTypeTop8:     8,
		ArenapvpTypeTop4:     4,
		ArenapvpTypeFinals:   2,
		ArenapvpTypeChampion: 1,
	}
)

// 邮件用语转换
var (
	rankTypeToLangMap = map[ArenapvpType]lang.LangCode{
		ArenapvpTypeElection: lang.OpenActivityArenapvpRankElection,
		ArenapvpTypeTop32:    lang.OpenActivityArenapvpRankTop32,
		ArenapvpTypeTop16:    lang.OpenActivityArenapvpRankTop16,
		ArenapvpTypeTop8:     lang.OpenActivityArenapvpRankTop8,
		ArenapvpTypeTop4:     lang.OpenActivityArenapvpRankTop4,
		ArenapvpTypeFinals:   lang.OpenActivityArenapvpRankSecond,
		ArenapvpTypeChampion: lang.OpenActivityArenapvpRankFirst,
	}
)

func (t ArenapvpType) GetNumber() int32 {
	return toNumberMap[t]
}

func (t ArenapvpType) GetRankLangCode() lang.LangCode {
	return rankTypeToLangMap[t]
}

// 对战状态
type ArenapvpState int32

const (
	ArenapvpStateInit   ArenapvpState = iota //正常
	ArenapvpStateFailed                      //失败
	ArenapvpStateExit                        //逃跑
)

//竞猜状态
type ArenapvpGuessState int32

const (
	ArenapvpGuessStateInit   ArenapvpGuessState = iota //初始
	ArenapvpGuessStateResult                           //结果
	ArenapvpGuessStateReturn                           //返还
)
