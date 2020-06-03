package config

import (
	"flag"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
)

//运行环境 prod、test、dev
var Environment string

var App app

type app struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Databases databases `json:"databases"`
	Logger    logger    `json:"logger"`
}

type databases struct {
	Driver string     `json:"driver"`
	List   []database `json:"list"`
}

type database struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
}

type logger struct {
	Path      string `json:"path"`
	InfoFile  string `json:"info_file"`
	ErrorFile string `json:"error_file"`
}

func init() {
	flag.StringVar(&Environment, "env", EnvDev, "运行模式")
}

//加载配置
func LoadConfig() {
	var configFileName = "app"
	switch Environment {
	case EnvProd:
		configFileName += "_" + EnvProd
	case EnvTest:
		configFileName += "_" + EnvTest
	default:
		configFileName += "_" + EnvDev
	}
	config := viper.New()
	config.SetConfigName(configFileName)
	config.SetConfigType("json")
	_, filename, _, _ := runtime.Caller(0)
	config.AddConfigPath(path.Dir(filename))
	readErr := config.ReadInConfig()
	if readErr != nil {
		panic("配置文件读取失败，原因：" + readErr.Error())
	}

	unmarshalErr := config.Unmarshal(&App)
	if unmarshalErr != nil {
		panic("配置文件初始化结构体失败，原因：" + unmarshalErr.Error())
	}
}

func Mode() string {
	switch Environment {
	case EnvProd:
		return "release"
	case EnvTest:
		return "test"
	default:
		return "debug"
	}
}
