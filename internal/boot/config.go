package boot

import (
	"errors"

	"github.com/spf13/viper"
)

func InitConfig() (err error) {
	viper.AddConfigPath("./manifest/config")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
