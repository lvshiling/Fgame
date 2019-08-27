package logic

import (
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/common/pbutil"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//培养使用
func GetStatusAndProgress(curTimesNum int32, curBless int32, timesMin int32, timesMax int32, addMin int32, addMax int32, updateRate int32, blessMax int32) (pro int32, randBless int32, sucess bool) {
	sucess = false
	pro = 0
	randBless = 0

	if timesMax <= 0 {
		panic(fmt.Errorf("最大次数不能小于1"))
	}

	if (curTimesNum + 1) >= timesMax { //成功
		sucess = true
	} else if (curTimesNum + 1) >= timesMin { //成功随机
		if mathutils.RandomHit(common.MAX_RATE, int(updateRate)) {
			sucess = true
		}
	}
	if !sucess { //进入祝福值随机
		randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
		curBless += randBless
		if curBless >= blessMax {
			pro = curBless - blessMax
			sucess = true
		} else {
			pro = randBless
			sucess = false
		}
	}
	return
}

//进阶使用
func AdvancedStatusAndProgress(curTimesNum int32, curBless int32, timesMin int32, timesMax int32, randBless int32, updateRate int32, blessMax int32) (pro int32, sucess bool) {
	sucess = false
	pro = 0

	if timesMax <= 0 {
		panic(fmt.Errorf("最大次数不能小于1"))
	}

	if curTimesNum >= timesMax { //成功
		sucess = true
	} else if curTimesNum >= timesMin { //成功随机
		if mathutils.RandomHit(common.MAX_RATE, int(updateRate)) {
			sucess = true
		}
	}

	//进入祝福值随机
	if !sucess {
		curBless += randBless
		if curBless != 0 && curBless >= blessMax {
			pro = curBless - blessMax
			sucess = true
		} else {
			pro = randBless
			sucess = false
		}
	}
	return
}

//祝福丹道具计算
func AdvancedBlessDan(curBless int32, randBless int32, blessMax int32) (sucess bool, pro int32) {
	curBless += randBless
	if curBless != 0 && curBless >= blessMax {
		pro = curBless - blessMax
		sucess = true
	} else {
		pro = randBless
		sucess = false
	}
	return
}

//TODO 修改为没有错误
func SendSystemMessage(pl scene.Player, code lang.LangCode, args ...string) (err error) {
	content := lang.GetLangService().ReadLang(code)
	scSystemMessage := pbutil.BuildSCSystemMessage(content, args...)
	err = pl.SendMsg(scSystemMessage)
	if err != nil {
		return
	}
	return
}

//TODO 修改为没有错误
func SendSessionSystemMessage(gs gamesession.Session, code lang.LangCode, args ...string) (err error) {
	content := lang.GetLangService().ReadLang(code)
	scSystemMessage := pbutil.BuildSCSystemMessage(content, args...)
	gs.Send(scSystemMessage)
	return
}

//TODO 修改为没有错误
func SendExceptionMessage(pl player.Player, code exception.ExceptionCode) (err error) {
	content := lang.GetLangService().ReadLang(code.LangCode())
	scException := pbutil.BuildSCException(content, code)
	err = pl.SendMsg(scException)
	if err != nil {
		return
	}
	return
}

//TODO 修改为没有错误
func SendSessionExceptionMessage(gs gamesession.Session, code exception.ExceptionCode) (err error) {
	content := lang.GetLangService().ReadLang(code.LangCode())
	scException := pbutil.BuildSCException(content, code)
	gs.Send(scException)
	return
}

// 模拟奖池抽奖路径
func SimulateRewPools(rewPools commontypes.RewPools, location int32, loop int32, curBackTimes int32, maxBackTimesCauseForward int32) (forwardTimes int32, backTimes int32, path []commontypes.PathType, err error) {
	backTimesRecord := curBackTimes
	backTimes = int32(0)
	forwardTimes = int32(0)
	path = []commontypes.PathType{}
	locat := location
	for i := int32(0); i < int32(len(rewPools)); i++ {
		node, ok := rewPools[i]
		if !ok {
			err = fmt.Errorf("no found rewNode,location:[%d]", i)
			return
		}
		if i == int32(0) {
			list := node.GetRateList()
			backRate := list[2]
			if backRate != int64(0) {
				err = fmt.Errorf("level 0 backRate should be 0,location:[%d]", i)
				return
			}
		}
		if i == int32(len(rewPools))-1 {
			list := node.GetRateList()
			forwardRate := list[1]
			if forwardRate != int64(0) {
				err = fmt.Errorf("level max forwardRate should be 0,location:[%d]", i)
				return
			}
		}
	}
	_, ok := rewPools[location]
	if !ok {
		err = fmt.Errorf("no found rewNode,location:[%d]", location)
		return
	}
	if loop <= 0 {
		err = fmt.Errorf("loop should more than 0,loop:[%d]", loop)
		return
	}
	if curBackTimes < 0 {
		err = fmt.Errorf("curBackTimes should more than 0,curBackTimes:[%d]", curBackTimes)
		return
	}
	if maxBackTimesCauseForward <= 0 {
		err = fmt.Errorf("maxBackTimesCauseForward should more than 0,maxBackTimesCauseForward:[%d]", maxBackTimesCauseForward)
		return
	}

	for i := int32(0); i < loop; i++ {
		node, _ := rewPools[locat]
		index := mathutils.RandomWeights(node.GetRateList())
		pathType := commontypes.PathType(index)

		if backTimesRecord%maxBackTimesCauseForward == 0 && backTimesRecord > 0 {
			pathType = commontypes.PathTypeBackTimesEnoughForward
		}

		switch pathType {
		case commontypes.PathTypeForward:
			forwardTimes++
			locat++
		case commontypes.PathTypeBack:
			backTimes++
			locat--
			backTimesRecord++
		case commontypes.PathTypeStill:
			break
		case commontypes.PathTypeBackTimesEnoughForward:
			forwardTimes++
			locat++
			backTimesRecord = 0
		default:
			panic(fmt.Errorf("no found pathType"))
			return
		}

		path = append(path, pathType)
	}
	return
}
