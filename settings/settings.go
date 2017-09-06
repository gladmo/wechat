package settings

import (
	"github.com/spf13/viper"

	"fmt"
)

func Get(key string) string {
	viper.AddConfigPath("conf")
	viper.SetConfigName("databases")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// convert interface to string
	return viper.Get(key).(string)
}
