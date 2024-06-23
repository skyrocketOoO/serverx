package boot

import (
	"web-server-template/manifest/config"
)

func InitConfig() {
	config.ReadConfig()
}
