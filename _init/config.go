package initial

import (
	"encoding/json"
	"os"
	pkg "server/_pkg"
	cfg "server/_sys/model"
	"sync"
)

var (
	mu_server_config sync.Mutex
)

func LoadServerConfig() *cfg.ServerConfig {
	if O_SERVER_CONFIG == nil {
		mu_server_config.Lock()
		defer mu_server_config.Unlock()
		if O_SERVER_CONFIG == nil {
			if conf, err := pkg.B64Decode(
				os.Getenv("SERVER_CONFIG")); err == nil {
				err = json.Unmarshal(conf, &O_SERVER_CONFIG)
				if err != nil {
					panic("Load Server Config Error: " + err.Error())
				}
			}
		}
	}
	return O_SERVER_CONFIG
}
