package repo

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"richingm/dothings/configs"
	"time"
)

func InitDb(config *configs.Config) *gorm.DB {
	dsn := config.MysqlConfig.Username + ":" + config.MysqlConfig.Password + "@tcp(" + config.MysqlConfig.Path + ")/" + config.MysqlConfig.Dbname
	fmt.Println(dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MysqlConfig.ConnMaxLifetime) * time.Minute)
	return db
}
