package conf

import (
	"encoding/json"
	"os"
	"sync"

	pkg "server/_pkg"
)

type ProjectConfig struct {
	StoreService StoreService `mapstructure:"store-service" json:"store-service" yaml:"store-service"`
}
type StoreService struct {
	BaseURL string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	Token   string `mapstructure:"token" json:"token" yaml:"token"`
	Enable  bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
}

var (
	O_PROJECT_CONFIG *ProjectConfig

	project_config_once sync.Once
)

func InitProjectConfig() *ProjectConfig {
	project_config_once.Do(
		func() {
			if O_PROJECT_CONFIG == nil {
				if conf, err := pkg.B64Decode(
					os.Getenv("PROJECT_CONFIG")); err == nil {
					err = json.Unmarshal(conf, &O_PROJECT_CONFIG)
					if err != nil {
						panic("Load Project Config Error: " + err.Error())
					}
				}
			}
		},
	)
	return O_PROJECT_CONFIG
}

func (s *ProjectConfig) GetStoreSercice() *StoreService {
	if s == nil {
		s = InitProjectConfig()
	}
	return &s.StoreService
}
