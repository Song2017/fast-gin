package sys_model

type PgConn struct {
	DSN          string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`                                  // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}
