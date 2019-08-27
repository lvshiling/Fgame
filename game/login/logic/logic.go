package logic

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/login/model"
	accounttypes "fgame/fgame/login/types"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

//登出前的操作
func Logout(p player.Player) {
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家登出")
	//防止奔溃
	defer func() {
		//退出服务器
		player.GetOnlinePlayerManager().PlayerLeaveServer(p)
		//移除用户
		player.GetOnlinePlayerManager().RemovePlayer(p)
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("player:玩家最终登出,成功")
	}()

	flag := p.LogoutSave()
	if !flag {
		panic("login:玩家登出应该成功")
	}
}

//TODO 修改rpc验证
func Login(token string, serverId int32, originServerId int32) (userId int64, platformUserId string, sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, gm int32, iosVersion string, androidVersion string, err error) {
	ctx := context.Background()
	userId, platformUserId, sdkType, deviceType, gm, iosVersion, androidVersion, err = center.GetCenterService().Login(ctx, token, serverId, originServerId)
	return
}

//TODO 修改rpc验证
func TestLogin(userName string, password string) (userId int64, state accounttypes.RealNameState, err error) {
	if len(userName) == 0 {
		return
	}
	if len(password) == 0 {
		return
	}
	m := &model.User{}
	err = global.GetGame().GetDB().DB().Find(m, "name=?", userName).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		//版署暂时改成18+
		m.RealNameState = int32(accounttypes.RealNameStateUp18)
		m.Name = userName
		m.Password = password
		m.CreateTime = global.GetGame().GetTimeService().Now()
		err = global.GetGame().GetDB().DB().Save(m).Error
		if err != nil {
			return
		}
	} else {
		if !strings.EqualFold(m.Password, password) {
			return
		}
	}
	userId = m.Id
	state = accounttypes.RealNameState(m.RealNameState)
	return
}
