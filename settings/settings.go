package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify" // 监控
	"github.com/spf13/viper"
)

// 用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件以及路径
	// viper.SetConfigName("config") // 指定配置文件名称（不需要要带后缀）
	// viper.AddConfigPath("./conf")   // 指定查找配置文件的路径（可指定多个路径）
	// viper.SetConfigType("yaml")   // 指定配置文件类型（配合远程配置中心使用的）
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Println("viper.ReadInConfig failed err:", err)
		return
	}
	// 把读取到的配置的信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed")
	}
	viper.WatchConfig()                            // 热加载
	viper.OnConfigChange(func(in fsnotify.Event) { // hockfunc
		fmt.Println("配置文件修改了")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal failed")
		}
	})
	return
}
