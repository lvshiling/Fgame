package model

type TableColumnInfo struct {
	TableName  string `gorm:"column:TABLE_NAME"`
	ColumnName string `gorm:"column:COLUMN_NAME"`
	ColumnType string `gorm:"column:COLUMN_TYPE"`
}
