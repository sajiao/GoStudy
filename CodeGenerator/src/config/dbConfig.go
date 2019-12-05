package config

import (
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/mysqlUtil"
	"moqikaka.com/goutil/redisUtil"
)

type DbConfig struct {
	gameModelConfig *mysqlUtil.DBConfig

	gameConfig *mysqlUtil.DBConfig

	redisConfig *redisUtil.RedisConfig
}

var (
	dbConfigObj *DbConfig
)

func (config *DbConfig) GetGameModelConfig() *mysqlUtil.DBConfig {
	return config.gameModelConfig
}

func (config *DbConfig) GetGameConfig() *mysqlUtil.DBConfig {
	return config.gameConfig
}

func (config *DbConfig) GetRedisConfig() *redisUtil.RedisConfig {
	return config.redisConfig
}

func newMysqlConfig(_gameModelConfig *mysqlUtil.DBConfig, _gameConfig *mysqlUtil.DBConfig, _redisConfig *redisUtil.RedisConfig) *DbConfig {
	return &DbConfig{
		gameModelConfig: _gameModelConfig,
		gameConfig:      _gameConfig,
		redisConfig:     _redisConfig,
	}
}

func intDbConfig(configObj *configUtil.XmlConfig) error {
	logUtil.DebugLog("开始加载DbConfig")

}
