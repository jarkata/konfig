package konfig

import (
	"log/slog"
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
	slog.Info("Init", "FileName", path)
	cache, err := ReadConfig(path)
	if err != nil {
		slog.Error("load config failed:", err)
		return
	}
	slog.Info("Init", "Cache:", cache)
	Config.SetCache(cache)
}

func getFileName() string {
	var path = os.Getenv("CONFIG_PATH")
	slog.Info("ConfigPath", "Path", path)
	if len(path) <= 0 {
		dir, _ := os.Getwd()
		return dir + "/application.properties"
	}
	return path
}
