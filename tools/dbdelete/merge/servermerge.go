package merge

import (
	fdb "fgame/fgame/core/db"
	"fgame/fgame/pkg/timeutils"
	"time"

	"github.com/jinzhu/gorm"
)

type IServerMergeOneDB interface {
	MergeChangeServerData(p_fromServerId int, p_toServerId int) error
}

type serverMergeOneDB struct {
	db fdb.DBService
}

func (m *serverMergeOneDB) MergeChangeServerData(p_fromServerId int, p_toServerId int) (err error) {

	serveridList := []string{
		"t_emperor",
		"t_emperor_records",
		"t_marry",
		"t_marry_divorce_consent",
		"t_marry_ring",
		"t_onearena",
		"t_player",
		"t_wedding",
		"t_wedding_card",
		"t_alliance",
		"t_chess_log",
		"t_order",
		"t_privilege_charge",
		"t_marry_pre_wed",
		"t_friend",
	}

	trans := m.db.DB().Begin()
	defer func() {
		if err != nil {
			trans.Rollback()
		}
		return
	}()

	// //仙盟名字修改
	// exdb := trans.Exec("UPDATE t_alliance SET NAME=CONCAT(?,name) WHERE serverId = ?", fmt.Sprintf("S%d.", p_fromServerId), p_fromServerId)
	// if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
	// 	err = exdb.Error
	// 	return
	// }
	// //角色修改
	// exdb = trans.Exec("UPDATE t_player SET NAME=CONCAT(?,name) WHERE serverId = ?", fmt.Sprintf("S%d.", p_fromServerId), p_fromServerId)
	// if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
	// 	err = exdb.Error
	// 	return
	// }
	var exdb *gorm.DB
	//后操作服务器id
	for _, value := range serveridList {
		exdb = trans.Exec("UPDATE "+value+" SET serverId=? where serverId=?", p_toServerId, p_fromServerId)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			err = exdb.Error
			return
		}
	}

	mergeTime := timeutils.TimeToMillisecond(time.Now())
	//修改合服时间和标志
	exdb = trans.Exec("UPDATE t_merge SET merge=1,mergeTime=? WHERE serverId = ?", mergeTime, p_toServerId)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		err = exdb.Error
		return
	}

	trans.Commit()

	return nil
}

// func rollBack(p_db *gorm.DB) {
// 	if p_db.Error != nil && p_db.Error != gorm.ErrRecordNotFound {
// 		p_db.Rollback()
// 	}
// }

func NewServerMergeOneDB(p_db fdb.DBService) IServerMergeOneDB {
	rst := &serverMergeOneDB{
		db: p_db,
	}
	return rst
}
