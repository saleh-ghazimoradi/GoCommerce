package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var AppConfig *config

type config struct {
	General   General   `mapstructure:"general"`   // general configs
	Databases Databases `mapstructure:"databases"` // databases configs
}

type General struct {
	Listen          string        `mapstructure:"listen"`           // rest listen port
	LogLevel        int8          `mapstructure:"log_level"`        // logger level
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"` // shutdown timeout
}

type Databases struct {
	MongoDB MongoDB `mapstructure:"mongodb"` // postgres configs
}

type MongoDB struct {
	Host         string        `mapstructure:"host"`          // MongoDB host
	Port         string        `mapstructure:"port"`          // MongoDB port
	User         string        `mapstructure:"user"`          // MongoDB user
	Pass         string        `mapstructure:"pass"`          // MongoDB password
	DatabaseName string        `mapstructure:"database_name"` // MongoDB database
	AuthSource   string        `mapstructure:"auth_source"`   // MongoDB auth source (default: "admin")
	MaxPoolSize  uint64        `mapstructure:"max_pool_size"` // MongoDB max pool size
	Timeout      time.Duration `mapstructure:"timeout"`       // Connection timeout
}

func LoadConfig(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name

	if path == "" {
		viper.AddConfigPath("./app/config")          // path to look for the config file in
		viper.AddConfigPath("./config")              // path to look for the config file in
		viper.AddConfigPath("$HOME/.config/chef-go") // call multiple times to add many search paths
		viper.AddConfigPath(".")                     // optionally look for config in the working directory
	} else {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	AppConfig = &config{}
	if err = viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
