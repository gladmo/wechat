package settings

import (
	"github.com/spf13/viper"

	"fmt"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Get(key string) string {
	// convert interface to string
	return viper.Get(key).(string)
}

func GetBool(key string) bool {
	return viper.Get(key).(bool)
}

func GetInt(key string) int {
	return viper.Get(key).(int)
}
