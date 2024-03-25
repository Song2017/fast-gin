package sys_model

type ServerConfig struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	PgConn PgConn `mapstructure:"pg-conn" json:"pg-conn" yaml:"pg-conn"`

	ServerName string `mapstructure:"server-name" json:"server-name" yaml:"server-name"`
	LogPath    string `mapstructure:"log-path" json:"log-path" yaml:"log-path"`
	TimeOut    int    `mapstructure:"time-out" json:"time-out" yaml:"time-out"`
}
