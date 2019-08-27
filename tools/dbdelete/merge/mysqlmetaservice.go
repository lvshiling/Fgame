package merge

import (
	"fgame/fgame/tools/dbdelete/model"
	"fmt"
)

type IMySqlMetaService interface {
	GetAllTableName(p_fpid int64, p_fsid int) ([]*model.TableInfo, error)
	GetTableColumn(p_fpid int64, p_fsid int, p_tableName string) ([]*model.TableColumnInfo, error)
}

type mysqlMetaService struct {
	cm IDbConfigManage
}

func (m *mysqlMetaService) GetAllTableName(p_fpid int64, p_fsid int) ([]*model.TableInfo, error) {
	db, err := m.cm.GetDbService(p_fpid, p_fsid)
	if err != nil {
		return nil, err
	}
	config := m.cm.GetDbConfigInfo(p_fpid, p_fsid)
	if config == nil {
		return nil, fmt.Errorf("nil of db")
	}
	sql := `select TABLE_NAME 
	from information_schema.tables 
	where table_schema=?`
	// sql = fmt.Sprintf(sql, config.DBName)
	result := make([]*model.TableInfo, 0)
	errdb := db.DB().Raw(sql, config.DBName).Scan(&result)
	if errdb != nil && errdb.Error != nil {
		return nil, errdb.Error
	}

	return result, nil
}

func (m *mysqlMetaService) GetTableColumn(p_fpid int64, p_fsid int, p_tableName string) ([]*model.TableColumnInfo, error) {
	db, err := m.cm.GetDbService(p_fpid, p_fsid)
	if err != nil {
		return nil, err
	}
	config := m.cm.GetDbConfigInfo(p_fpid, p_fsid)
	if config == nil {
		return nil, fmt.Errorf("nil of db")
	}

	result := make([]*model.TableColumnInfo, 0)
	sql := `select TABLE_NAME,COLUMN_NAME,COLUMN_TYPE
	from information_schema.columns
	where table_schema=? and TABLE_NAME=?`

	dberr := db.DB().Raw(sql, config.DBName, p_tableName).Scan(&result)
	if dberr != nil && dberr.Error != nil {
		return nil, dberr.Error
	}
	return result, nil
}

func NewMySqlMetaService(p_cm IDbConfigManage) IMySqlMetaService {
	rst := &mysqlMetaService{
		cm: p_cm,
	}
	return rst
}
