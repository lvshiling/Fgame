package found

import (
	foundtypes "fgame/fgame/game/found/types"
	"sync"
)

type FoundService interface {
	//可以找回的次数
	CountFoundBackTimes(resType foundtypes.FoundResourceType, resLevel, joinTimes int32) (foundTimes int32)
}

type foundService struct {
}

//初始化
func (s *foundService) init() (err error) {
	return
}

const (
	defaultTimes = 1
)

// 资源找回只计算基础次数
func (s *foundService) CountFoundBackTimes(resType foundtypes.FoundResourceType, maxTimes, joinTimes int32) (foundTimes int32) {
	playModeType := resType.GetPlayModeType()

	if playModeType == foundtypes.PlayModeTypeDailyTasks {
		//次数
		if joinTimes >= maxTimes {
			return 0
		}

		foundTimes = maxTimes - joinTimes
	}
	if playModeType == foundtypes.PlayModeTypeJoinTimesLimitActivity {
		//次数
		if joinTimes >= maxTimes {
			return 0
		}

		foundTimes = maxTimes - joinTimes
	}
	if playModeType == foundtypes.PlayModeTypeFreeJoinActivity {
		//参与过
		if joinTimes > 0 {
			return 0
		}

		foundTimes = int32(defaultTimes)
	}

	return
}

var (
	once sync.Once
	s    *foundService
)

func Init() (err error) {
	once.Do(func() {
		s = &foundService{}
		err = s.init()
	})
	return
}

func GetFoundService() FoundService {
	return s
}
