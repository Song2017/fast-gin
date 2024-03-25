package initial

import (
	cfg "server/_sys/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func InitGormPg() *gorm.DB {
	p := LoadServerConfig().PgConn
	return GormPgSqlByConfig(p)
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p cfg.PgConn) *gorm.DB {
	if p.DSN == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.DSN, // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(
		pgsqlConfig), GormEx.GormConfig("pg", false)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
