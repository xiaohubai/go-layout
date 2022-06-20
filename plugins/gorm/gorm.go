package gorm

import (
	"fmt"

	"github.com/xiaohubai/go-layout/configs/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Gorm 数据库组件
func Init() *gorm.DB {
	m := global.Cfg.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig())
	if err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s \n", err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}

func gormConfig() *gorm.Config {
	cfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,  //禁用表复数形式
			TablePrefix:   "tb_", //表前缀
		},
	}
	switch global.Cfg.Mysql.LogMode {
	case "error", "Error":
		cfg.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		cfg.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		cfg.Logger = Default.LogMode(logger.Info)
	default:
		cfg.Logger = Default.LogMode(logger.Info)
	}
	return cfg
}
