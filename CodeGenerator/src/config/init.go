package config

import (
	"moqikaka.com/Framework/configMgr"
	"moqikaka.com/goutil/configUtil"
)

var (
	configObj     *configUtil.XmlConfig
	configManager = configMgr.NewConfigManager()
)

func init() {
	configManager.RegisterInitFunc()
}
