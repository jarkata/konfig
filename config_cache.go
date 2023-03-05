package konfig

import (
	"fmt"
	"os"
)

type GlobalConfig struct {
	cache map[string]string
}

var Config = &GlobalConfig{}

func (config *GlobalConfig) SetCache(cache map[string]string) {
	config.cache = cache
}

func GetCache(key string) string {
	return Config.cache[key]
}

func init() {
	path := getFileName()
	cache, err := ReadConfig(path)
	if err != nil {
		fmt.Println("load config failed:", err)
		return
	}
	fmt.Println("loading config:", cache)
	Config.SetCache(cache)
}

func getFileName() string {
	path := os.Getenv("config_path")
	if len(path) <= 0 {
		dir, _ := os.Getwd()
		return dir + "/konfig.properties"
	}
	return path
}
