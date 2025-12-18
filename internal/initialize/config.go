package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server port:", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal configuration %w \n", err))
	}
}
