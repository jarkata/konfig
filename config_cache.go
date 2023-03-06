package konfig

import (
	"fmt"
	"os"
	"sync"
)

type GlobalConfig struct {
	cache map[string]string
}

var Config = &GlobalConfig{}

func (config *GlobalConfig) SetCache(cache map[string]string) {
	config.cache = cache
}

var once sync.Once

func GetCache(key string) string {
	val := Config.cache[key]
	if len(val) <= 0 {
		once.Do(Init)
		val = Config.cache[key]
	}
	return val
}

func init() {
	Init()
}

func Init() {
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
