package monitor

import (
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type IUserServerManager interface {
	SetUserServer(p_userid int64, p_serverList []int32)
	ClearUser(p_userid int64)
	GetUserServerList(p_userid int64) []int32
	GetServerUserList(p_serverId int32) []int64

	Log()
}

type userServerManager struct {
	rwm           sync.RWMutex
	userServerMap map[int64][]int32
	serverUserMap map[int32]map[int64]int64
}

func (m *userServerManager) SetUserServer(p_userid int64, p_serverList []int32) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	if preUserServer, ok := m.userServerMap[p_userid]; ok {
		for _, preServerId := range preUserServer {
			if _, serOK := m.serverUserMap[preServerId]; serOK {
				delete(m.serverUserMap[preServerId], p_userid)
			}
		}
	}
	m.userServerMap[p_userid] = p_serverList
	for _, value := range p_serverList {
		if _, ok := m.serverUserMap[value]; !ok {
			m.serverUserMap[value] = make(map[int64]int64)
		}
		m.serverUserMap[value][p_userid] = p_userid
	}
}

func (m *userServerManager) ClearUser(p_userid int64) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	if preUserServer, ok := m.userServerMap[p_userid]; ok {
		for _, preServerId := range preUserServer {
			if _, serOK := m.serverUserMap[preServerId]; serOK {
				delete(m.serverUserMap[preServerId], p_userid)
			}
		}
	}
	delete(m.userServerMap, p_userid)
}

func (m *userServerManager) GetUserServerList(p_userid int64) []int32 {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	return m.userServerMap[p_userid]
}

func (m *userServerManager) GetServerUserList(p_serverId int32) []int64 {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	rst := make([]int64, 0)
	for _, value := range m.serverUserMap[p_serverId] {
		rst = append(rst, value)
	}
	return rst
}

func (m *userServerManager) Log() {
	if len(m.userServerMap) == 0 {
		log.Debug("玩家状态为空")
	}
	for key, value := range m.userServerMap {
		log.WithFields(log.Fields{
			"玩家ID":   key,
			"玩家服务列表": joinArray(value),
		}).Debug("玩家状态")
	}

	if len(m.serverUserMap) == 0 {
		log.Debug("服务器状态为空")
	}
	for key, value := range m.serverUserMap {
		log.WithFields(log.Fields{
			"服务器ID":    key,
			"服务器中的玩家：": joinMap(value),
		}).Debug("服务器状态")
	}
}

func joinArray(p_data []int32) string {
	rst := ""
	for _, value := range p_data {
		rst += fmt.Sprintf("%d,", value)
	}
	return rst
}
func joinMap(p_data map[int64]int64) string {
	rst := ""
	for _, value := range p_data {
		rst += fmt.Sprintf("%d,", value)
	}
	return rst
}

func NewUserServerManage() IUserServerManager {
	rst := &userServerManager{}
	rst.userServerMap = make(map[int64][]int32)
	rst.serverUserMap = make(map[int32]map[int64]int64)
	return rst
}
