/*
Package config default config
*/
package config

import (
	"io/ioutil"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

// Config 配置
type Config struct {
	AppInfo  appInfo  `json:"AppInfo"`
	Log      logConf  `json:"Log"`
	DB       db       `json:"DB"`
	Redis    redis    `json:"Redis"`
	Security security `json:"Security"`
}

type appInfo struct {
	Env  string `json:"Env"` // example: local, dev, prod
	Addr string `json:"Addr"`
}

type logConf struct {
	LogBasePath string `json:"LogBasePath"`
	LogFileName string `json:"LogFileName"`
}

type security struct {
	Secret string `json:"Secret"`
}

// Conf 配置
var Conf *Config

var filePrefix = "/app/config/"

func init() {
	log.Println("begin init all configs")
	initConf()
	log.Println("over init all configs")
}

func initConf() {
	log.Println("begin init default config")

	Conf = &Config{}
	fileName := "default.json"

	if v, ok := os.LookupEnv("CONFIG_PATH_PREFIX"); ok {
		filePrefix = v
	}
	// read default config
	data, err := ioutil.ReadFile(filePrefix + fileName)
	if err != nil {
		log.Println("config-initConf: read default.json error")
		log.Panic(err)
		return
	}
	err = jsoniter.Unmarshal(data, Conf)
	if err != nil {
		log.Println("config-initConf: unmarshal default.json error")
		log.Panic(err)
		return
	}

	// read env and config path
	if v, ok := os.LookupEnv("ENV"); ok {
		fileName = v + ".json"
	}
	if fileName != "default.json" {

		// read env config
		data, err = ioutil.ReadFile(filePrefix + fileName)
		if err != nil {
			log.Println("config-initConf: read [env].json error")
			log.Panic(err)
			return
		}
		err = jsoniter.Unmarshal(data, Conf)
		if err != nil {
			log.Println("config-initConf: unmarshal [env].json error")
			log.Panic(err)
			return
		}
	}

	if v, ok := os.LookupEnv("MONGO_INITDB_ROOT_USERNAME"); ok {
		Conf.DB.User = v
	}
	if v, ok := os.LookupEnv("MONGO_INITDB_ROOT_PASSWORD"); ok {
		Conf.DB.PW = v
	}
	if v, ok := os.LookupEnv("MONGO_INITDB_DATABASE"); ok {
		Conf.DB.DBName = v
	}
	if v, ok := os.LookupEnv("RedisPass"); ok {
		Conf.Redis.PW = v
	}

	log.Println("over init default config")
}
