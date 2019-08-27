package codec

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func IsLoginMsg(msgType MessageType) bool {
	if MessageType(uipb.MessageType_CS_LOGIN_TYPE) == msgType {
		return true
	}
	if MessageType(uipb.MessageType_CS_TEST_LOGIN_TYPE) == msgType {
		return true
	}
	return false
}

func IsCreateRoleMsg(msgType MessageType) bool {
	if MessageType(uipb.MessageType_CS_SELECT_JOB_TYPE) == msgType {
		return true
	}
	return false
}

func IsCrossLoginMsg(msgType MessageType) bool {
	if MessageType(crosspb.MessageType_SI_LOGIN_TYPE) == msgType {
		return true
	}

	return false
}

const (
	minUIType = 201
	maxUIType = 30000
)

func IsUIMsg(msgType MessageType) bool {
	if msgType >= minUIType && msgType <= maxUIType {
		return true
	}

	return false
}

const (
	minSceneType = 100
	maxSceneType = 200
)

func IsSceneMsg(msgType MessageType) bool {
	if msgType >= minSceneType && msgType <= maxSceneType {
		return true
	}
	return false
}

const (
	minCrossType = 30001
)

func IsCrossMsg(msgType MessageType) bool {
	if msgType >= minCrossType {
		return true
	}

	return false
}
