package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"

	"fgame/fgame/game/player"
	"fgame/fgame/game/player/logic"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/register/register"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SELECT_JOB_TYPE), dispatch.HandlerFunc(handleCreate))
}

//处理创建角色消息
func handleCreate(s session.Session, msg interface{}) (err error) {
	log.Debug("player:处理创建角色消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(player.Player)
	csSelectJob := msg.(*uipb.CSSelectJob)
	job := csSelectJob.GetJob()
	sex := csSelectJob.GetSex()
	name := csSelectJob.GetName()

	roleType := playertypes.RoleType(job)
	if !roleType.Valid() {
		log.WithFields(
			log.Fields{
				"userId": pl.GetUserId(),
				"job":    job,
				"sex":    sex,
				"name":   name,
			}).Warn("player:处理创建角色消息,角色类型无效")
		playerlogic.SendSystemMessage(pl, lang.JobInvalid)
		return
	}
	sexType := playertypes.SexType(sex)
	if !sexType.Valid() {
		log.WithFields(
			log.Fields{
				"userId": pl.GetUserId(),
				"job":    job,
				"sex":    sex,
				"name":   name,
			}).Warn("player:处理创建角色消息,性别类型无效")
		playerlogic.SendSystemMessage(pl, lang.SexInvalid)
		return
	}
	if len(name) <= 0 {
		log.WithFields(
			log.Fields{
				"userId": pl.GetUserId(),
				"job":    job,
				"sex":    sex,
				"name":   name,
			}).Warn("player:处理创建角色消息,名字无效")
		playerlogic.SendSystemMessage(pl, lang.NameInvalid)
		return
	}
	//判断是否是开放注册
	openRegister := register.GetRegisterService().IsOpenRegister()
	if !openRegister {
		log.WithFields(
			log.Fields{
				"userId": pl.GetUserId(),
				"job":    job,
				"sex":    sex,
				"name":   name,
			}).Warn("player:处理创建角色消息,关闭注册")
		playerlogic.SendExceptionMessage(pl, exception.ExceptionCodeRegisterClose)
		pl.Close(nil)
		return
	}
	err = logic.CreatePlayer(pl, roleType, sexType, name)
	if err != nil {
		log.WithFields(
			log.Fields{
				"userId": pl.GetUserId(),
				"error":  err,
			}).Error("player:处理创建角色消息,创建失败")
		return
	}
	log.WithFields(
		log.Fields{
			"userId": pl.GetUserId(),
			"job":    job,
			"sex":    sex,
			"name":   name,
		}).Debug("player:处理创建角色消息,创建成功")
	return
}
