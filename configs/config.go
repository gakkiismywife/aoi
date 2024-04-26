package configs

import (
	"github.com/spf13/viper"
)

type Cfg struct {
	Server   *Server   `yaml:"server"`
	Jwt      *Jwt      `yaml:"jwt"`
	Database *Database `yaml:"database"`
	Redis    *Redis    `yaml:"redis"`
	ImageDb  *ImageDb  `yaml:"imageDb"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Jwt struct {
	Secret  string `yaml:"secret"`
	Expire  int    `yaml:"expire"`
	Refresh int    `yaml:"refresh"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Index    int    `yaml:"index"`
}

type ImageDb struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var (
	C              *Cfg
	ServerConfig   *Server
	JwtConfig      *Jwt
	DatabaseConfig *Database
	RedisConfig    *Redis
	ImageDbConfig  *ImageDb
)

func InitCfg() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../etc/")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	_ = v.UnmarshalKey("Server", &ServerConfig)
	_ = v.UnmarshalKey("Jwt", &JwtConfig)
	_ = v.UnmarshalKey("Database", &DatabaseConfig)
	_ = v.UnmarshalKey("Redis", &RedisConfig)
	_ = v.UnmarshalKey("ImageDb", &ImageDbConfig)

	C = &Cfg{
		Server:   ServerConfig,
		Jwt:      JwtConfig,
		Database: DatabaseConfig,
		Redis:    RedisConfig,
		ImageDb:  ImageDbConfig,
	}

	return nil
}
