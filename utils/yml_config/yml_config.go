package yml_config

import (
	"Goshop/core/container"
	"Goshop/global/my_errors"
	"Goshop/global/variable"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

type ymlConfig struct {
	viper *viper.Viper
}

func CreateYamlFactory() *ymlConfig {

	yamlConfig := viper.New()
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名
	yamlConfig.SetConfigName("config")
	//设置配置文件类型
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return &ymlConfig{
		yamlConfig,
	}
}

//监听文件变化
func (c *ymlConfig) ConfigFileChangeListen() {
	c.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				c.clearCache()
				lastChangeTime = time.Now()
			}
		}
	})
	c.viper.WatchConfig()
}

// 判断相关键是否已经缓存
func (c *ymlConfig) keyIsCache(keyName string) bool {
	if _, exists := container.CreateContainersFactory().KeyIsExists(variable.ConfigKeyPrefix + keyName); exists {
		return true
	} else {
		return false
	}
}

// 对键值进行缓存
func (c *ymlConfig) cache(keyName string, value interface{}) bool {
	return container.CreateContainersFactory().Set(variable.ConfigKeyPrefix+keyName, value)
}

// 通过键获取缓存的值
func (c *ymlConfig) getValueFromCache(keyName string) interface{} {
	return container.CreateContainersFactory().Get(variable.ConfigKeyPrefix + keyName)
}

// 清空已经窜换的配置项信息
func (c *ymlConfig) clearCache() {
	container.CreateContainersFactory().FuzzyDelete(variable.ConfigKeyPrefix)
}

// Get 一个原始值
func (c *ymlConfig) Get(keyName string) interface{} {
	return c.viper.Get(keyName)
}

// GetString
func (c *ymlConfig) GetString(keyName string) string {
	return c.viper.GetString(keyName)
}

// GetBool
func (c *ymlConfig) GetBool(keyName string) bool {
	return c.viper.GetBool(keyName)
}

// GetInt
func (c *ymlConfig) GetInt(keyName string) int {
	return c.viper.GetInt(keyName)
}

// GetInt32
func (c *ymlConfig) GetInt32(keyName string) int32 {
	return c.viper.GetInt32(keyName)
}

// GetInt64
func (c *ymlConfig) GetInt64(keyName string) int64 {
	return c.viper.GetInt64(keyName)
}

// float64
func (c *ymlConfig) GetFloat64(keyName string) float64 {
	return c.viper.GetFloat64(keyName)
}

// GetDuration
func (c *ymlConfig) GetDuration(keyName string) time.Duration {
	return c.viper.GetDuration(keyName)
}

// GetStringSlice
func (c *ymlConfig) GetStringSlice(keyName string) []string {
	return c.viper.GetStringSlice(keyName)
}
