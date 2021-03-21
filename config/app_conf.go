package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

const ProjectName = "sdk-go"

var AppConf *AppConfig
var defaultAppConfig = []byte(`
app:
  runMode: "release"
  addr: ":8088"
  apiPrefix: "/sdk/api"
`)

func init() {
	initAppConfig()
}

func initAppConfig() {
	AppConf = newAppConfig()
}

//the AppConfig as application layer configuration
//run mode:debug,release,test, release of the default value
type AppConfig struct {
	v *viper.Viper
	Config
}

type Config struct {
	RunMode   string
	Addr      string
	ApiPrefix string
}

func GetAppConfig() *AppConfig {
	return AppConf
}

func (c *AppConfig) GetString(str string) string {
	return c.v.GetString(str)
}

func newAppConfig() *AppConfig {
	conf := &AppConfig{}
	conf.v = newViper(defaultAppConfig, "yaml", "sdk")
	conf.RunMode = conf.v.GetString("app.runMode")
	conf.Addr = conf.v.GetString("app.addr")
	conf.ApiPrefix = conf.v.GetString("app.apiPrefix")

	return conf
}

func newViper(conf []byte, confType, envPrefix string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(confType)
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	if err := v.ReadConfig(bytes.NewBuffer(conf)); err != nil {
		panic(err)
	}
	return v
}
