package jgmysql

import (
	"fmt"
	"testing"

	jgconf "github.com/Jinglever/go-config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// open connection to mysql
func openDb() *gorm.DB {
	cfg := struct {
		Host   string `mapstructure:"host"`
		Port   string `mapstructure:"port"`
		User   string `mapstructure:"user"`
		Pass   string `mapstructure:"pass"`
		DbName string `mapstructure:"dbname"`
	}{}
	jgconf.LoadYamlConfig("./_test_data/conf.yml", &cfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// test helper QueryDbVersion
func TestQueryDbVersion(t *testing.T) {
	db := openDb()
	helper := NewHelper(db)
	version, err := helper.QueryDbVersion()
	if err != nil {
		t.Errorf("query db version failed, err: %v", err)
	} else {
		t.Log(version)
	}
}

// test helper QueryDbCharset
func TestQueryDbCharset(t *testing.T) {
	db := openDb()
	helper := NewHelper(db)
	charset, err := helper.QueryDbCharset()
	if err != nil {
		t.Errorf("query db charset failed, err: %v", err)
	} else {
		t.Log(charset)
	}
}

// test helper QueryDbCollate
func TestQueryDbCollate(t *testing.T) {
	db := openDb()
	helper := NewHelper(db)
	collate, err := helper.QueryDbCollate()
	if err != nil {
		t.Errorf("query db collate failed, err: %v", err)
	} else {
		t.Log(collate)
	}
}

// test helper QueryAllTables
func TestQueryAllTables(t *testing.T) {
	db := openDb()
	helper := NewHelper(db)
	tables, err := helper.QueryAllTables()
	if err != nil {
		t.Errorf("query all tables failed, err: %v", err)
	} else {
		t.Log(tables)
	}
}

// test helper QueryCreateTableSql
func TestQueryCreateTableSql(t *testing.T) {
	db := openDb()
	helper := NewHelper(db)
	sql, err := helper.QueryCreateTableSql("user")
	if err != nil {
		t.Errorf("query create table sql failed, err: %v", err)
	} else {
		t.Log(sql)
	}
}
