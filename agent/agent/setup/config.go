package setup

import (
	"bytes"
	"encoding/json"
	gameserver "fgame/fgame/game/server"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mohae/deepcopy"
)

func SetupConfig(tpl string, options *gameserver.GameServerOptions) (content string, err error) {
	tplFile, err := filepath.Abs(tpl)
	if err != nil {
		return
	}
	t, err := template.ParseFiles(tplFile)
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(nil)
	err = t.Execute(buffer, options)
	if err != nil {
		return
	}
	content = buffer.String()
	return
}

type ConfigService interface {
	GetGameServerOptions(serverId int32) (options *gameserver.GameServerOptions, err error)
	SetupGameServer(dir string, serverId int32) (err error)
}

type configService struct {
	baseOptions *gameserver.GameServerOptions
}

const (
	gameServerNumPerDB = 8
	gamePrefix         = "game"
	gameConfig         = "config"
	gameLog            = "logs"
	gameData           = "data"
)

func (s *configService) SetupGameServer(dir string, serverId int32) (err error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return
	}
	f, err := os.Stat(absDir)
	if err != nil {
		return
	}
	if !f.IsDir() {
		return fmt.Errorf("%s 不是目录", dir)
	}
	gameDirName := fmt.Sprintf("%s_%d", gamePrefix, serverId)

	absGameDirName := filepath.Join(absDir, gameDirName)
	_, err = os.Stat(absGameDirName)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		err = nil
		//创建游戏目录
		err = os.Mkdir(absGameDirName, os.ModeDir)
		if err != nil {
			return
		}
	}

	//创建config,log,data
	configDir := filepath.Join(absGameDirName, gameConfig)
	if err = s.createDir(configDir); err != nil {
		return
	}
	logDir := filepath.Join(absGameDirName, gameLog)
	if err = s.createDir(logDir); err != nil {
		return
	}
	dataDir := filepath.Join(absGameDirName, gameData)
	if err = s.createDir(dataDir); err != nil {
		return
	}
	options, err := s.GetGameServerOptions(serverId)
	if err != nil {
		return
	}
	configFileName := fmt.Sprintf("%s_%d.json", gamePrefix, serverId)
	absConfigFileName := filepath.Join(configDir, configFileName)
	configContent, err := json.MarshalIndent(options, "", "    ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(absConfigFileName, configContent, os.ModePerm)
	if err != nil {
		return
	}
	return
}

func (s *configService) createDir(dir string) (err error) {
	_, err = os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		err = nil
		//创建游戏目录
		err = os.Mkdir(dir, os.ModeDir)
		if err != nil {
			return
		}
	}
	return

}

func (s *configService) GetGameServerOptions(serverId int32) (options *gameserver.GameServerOptions, err error) {

	copyOptions := deepcopy.Copy(s.baseOptions)
	options = copyOptions.(*gameserver.GameServerOptions)

	dbIndex := (serverId-1)/gameServerNumPerDB + 1
	options.Game.Db.DBName = fmt.Sprintf("%s_%d", options.Game.Db.DBName, dbIndex)
	options.Game.Remote.Port = options.Game.Remote.Port + serverId - 1
	options.Game.Register.Port = options.Game.Register.Port + serverId - 1
	options.Server.Id = serverId
	options.Server.Port = options.Server.Port + serverId - 1
	options.Server.PprofPort = options.Server.PprofPort + serverId - 1
	options.Server.StatsPort = options.Server.StatsPort + serverId - 1

	return
}

func NewConfigService(baseOptions *gameserver.GameServerOptions) ConfigService {
	s := &configService{}
	s.baseOptions = baseOptions
	return s
}
