package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ConfigPath = flag.String("cfg", getConfigPath(), "Path to configuration file")
)

func getConfigPath() string {
	configFileName := "config.json"
	// if isdelve.Enabled {
	// 	return path.Join(getWorkingDirectoryConfigPath(), configFileName)
	// } else {

	// 	return path.Join(getDefaultConfisgPath(), configFileName)
	// }

	return filepath.Join(GetWorkingDirectory(), configFileName)
}

func GetExecutablePath() string {

	if exepath, err := os.Executable(); err != nil {
		panic(err)
	} else {
		return filepath.Dir(exepath)
	}

}

func GetWorkingDirectory() string {
	if dir, err := os.Getwd(); err != nil {

		panic(err)
	} else {
		return dir
	}
}

func LoadConfig(filepath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(filepath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

type ServerConfig struct {
	Port int
}

var DefalutServerConfig = ServerConfig{
	Port: 8000,
}

type LogConfig struct {
	Service string     `json:"service" yaml:"service"`
	Zap     zap.Config `json:"zap" yaml:"zap"`
	Sinks   []string   `json:"sinks" yaml:"sinks"`
}

type BaseConfig struct {
	Server           ServerConfig
	ConnectionString string
	Log              LogConfig `json:"log" yaml:"log"`
	Secret           string
}

var Values BaseConfig

type IAppConfig interface {
	GetBaseConfig() BaseConfig
}

/*
Initialize Flag so all the flag parameters should be populated before calling this method
*
*/
func InitializeFlag() {
	flag.Parse()
}

/*
Should be called first to initialize the config.
Value is populated from config file provided through flag and overridden by env
*/
func Initialize(appConfig IAppConfig) {

	fmt.Println("Initializing config")
	InitializeFlag()
	initializeDefaultValue()
	if v, err := LoadConfig(*ConfigPath); err != nil {
		panic(err)
	} else {
		if appConfig == nil {
			v.Unmarshal(&Values)
		} else {
			v.Unmarshal(&appConfig)
			Values = appConfig.GetBaseConfig()
			//fmt.Printf("baseconfig %v\n", appConfig.GetBaseConfig())
		}

	}

	fmt.Println("Config Initialized Success fully")
}

func initializeDefaultValue() {
	Values = BaseConfig{
		Server: DefalutServerConfig,
		Log: LogConfig{
			Zap: DefaultZapConfig,
		},
	}
}

var DefaultZapConfig = zap.Config{
	Encoding:         "json",
	OutputPaths:      []string{"stdout", "log.log"},
	ErrorOutputPaths: []string{"stderr", "log.log"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
	},
}
