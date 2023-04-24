package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	InitConf()
}

// InitConf 初始化配置文件
func InitConf() {
	viper.SetConfigName("configs/server")
	viper.AddConfigPath(".") // 添加搜索路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
