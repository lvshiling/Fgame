package types

import (
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/pkg/mathutils"
)

//棋局类型
type ChessType int32

const (
	ChessTypeSilver   ChessType = iota //银两棋局
	ChessTypeGold                      //元宝棋局
	ChessTypeBindGold                  //绑元棋局
)

const (
	MinChessType = ChessTypeSilver
	MaxChessType = ChessTypeBindGold
)

func (t ChessType) Valid() bool {
	switch t {
	case ChessTypeSilver,
		ChessTypeGold,
		ChessTypeBindGold:
		return true
	default:
		return false
	}
}

var chessTypeMap = map[ChessType]string{
	ChessTypeSilver:   "银两棋局",
	ChessTypeGold:     "元宝棋局",
	ChessTypeBindGold: "绑元棋局",
}

func RandomChessType() ChessType {
	flag := mathutils.RandomHit(100, 50)
	if flag {
		return ChessTypeSilver
	} else {
		return ChessTypeGold
	}
}
func (t ChessType) String() string {
	return chessTypeMap[t]
}

func (t ChessType) GetChangeConstantType() constanttypes.ConstantType {
	switch t {
	case ChessTypeGold:
		return constanttypes.ConstantTypeChessChangedUseGold
	case ChessTypeBindGold:
		return constanttypes.ConstantTypeChessChangedUseBindGold
	default:
		return -1
	}
}

func (t ChessType) GetAttendLimitConstantType() constanttypes.ConstantType {
	switch t {
	case ChessTypeSilver:
		return constanttypes.ConstantTypeChessSilverAttendTimesMax
	case ChessTypeBindGold:
		return constanttypes.ConstantTypeChessBindGoldAttendTimesMax
	default:
		return -1
	}
}
