package jgmysql

import (
	"fmt"

	"gorm.io/gorm"
)

type Helper struct {
	DB *gorm.DB
}

func NewHelper(db *gorm.DB) *Helper {
	return &Helper{DB: db}
}

// query version of database
func (h *Helper) QueryDbVersion() (string, error) {
	var version string
	err := h.DB.Raw("select version()").Scan(&version).Error
	if err != nil {
		return "", err
	}
	return version, nil
}

// query charset of database
func (h *Helper) QueryDbCharset() (string, error) {
	var records = make([]map[string]interface{}, 0)
	err := h.DB.Raw("show variables like 'character_set_database'").Scan(&records).Error
	if err != nil {
		return "", err
	}
	return records[0]["Value"].(string), nil
}

// query collate of database
func (h *Helper) QueryDbCollate() (string, error) {
	var records = make([]map[string]interface{}, 0)
	err := h.DB.Raw("show variables like 'collation_database'").Scan(&records).Error
	if err != nil {
		return "", err
	}
	return records[0]["Value"].(string), nil
}

// query all tables of database
func (h *Helper) QueryAllTables() ([]string, error) {
	var tables = make([]map[string]interface{}, 0)
	err := h.DB.Raw("show tables").Scan(&tables).Error
	if err != nil {
		return nil, err
	}
	var ret = make([]string, 0)
	for _, table := range tables {
		for _, v := range table {
			ret = append(ret, v.(string))
		}
	}
	return ret, nil
}

// query create table sql
func (h *Helper) QueryCreateTableSql(tableName string) (string, error) {
	var records = make([]map[string]interface{}, 0)
	err := h.DB.Raw(fmt.Sprintf("show create table `%s`", tableName)).Scan(&records).Error
	if err != nil {
		return "", err
	}
	return records[0]["Create Table"].(string), nil
}
