package initialize

import (
	"fmt"
	"os"

	"github.com/paiz4ever/go-web-framework/global"
	"github.com/paiz4ever/go-web-framework/pkg/constant"
	"github.com/spf13/viper"
)

func InitConfig() error {
	env := os.Getenv(constant.SERVER_ENV)
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		return err
	}
	return nil
}
