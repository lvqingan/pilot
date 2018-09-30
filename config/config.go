package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// 当前包内的全局变量，用来保存配置信息
var config = make(map[string]map[string]interface{})

// 读取指定目录下的 json 配置文件，需要先执行 Load，才可以执行 Get
func Load(configPath string) error {
	configFilesInfo, err := ioutil.ReadDir(configPath)

	if err == nil {
		for _, configFileInfo := range configFilesInfo {
			configFileName := configFileInfo.Name()
			configFile, _ := os.Open(configPath + "/" + configFileName)
			configJson, _ := ioutil.ReadAll(configFile)
			configMap := make(map[string]interface{})

			if err := json.Unmarshal(configJson, &configMap); err != nil {
				return err
			}

			config[strings.TrimSuffix(configFileName, ".json")] = configMap
		}

		return nil
	} else {
		return err
	}
}

// 按指定的 key 获取配置内容，. 号分隔
// 譬如 app.json 中包含 "key1": "val" 则按 app.key1 来取值；包含 "key1": {"key2": "val"}，则按 app.key1.key2 取值
// 没找到返回 nil
func Get(key string) interface{} {
	if len(config) == 0 {
		panic("config.Load() not executed")
	}

	if strings.Contains(key, ".") {
		keyParts := strings.Split(key, ".")
		section := keyParts[0]
		keys := keyParts[1:]

		// 支持无限层级查找
		configMap := config[section]
		for _, key := range keys {
			value := configMap[key]

			if reflect.ValueOf(value).Kind() == reflect.Map {
				configMap = value.(map[string]interface{})
			} else {
				return value
			}
		}

		return nil
	} else {
		panic("key must be separated by dot")
	}
}

// 指定返回类型为 string 的 Get 方法
// 没找到返回 ""
func GetString(key string) string {
	value := Get(key)

	if value == nil {
		return ""
	} else {
		return value.(string)
	}
}
