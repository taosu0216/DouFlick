package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var globalConfig = new(GlobalConfig)

type GlobalConfig struct {
	*SvrConfig    `mapstructure:"svr_config"`
	*MysqlConfig  `mapstructure:"mysql"`
	*LogConfig    `mapstructure:"log"`
	*RedisConfig  `mapstructure:"redis"`
	*ConsulConfig `mapstructure:"consul"`
	*MinioConfig  `mapstructure:"minio"`
}

type MinioConfig struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            string `mapstructure:"port" json:"port"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key"`
	VideoBuckets    string `mapstructure:"video_buckets" json:"video_buckets"`
	PicBuckets      string `mapstructure:"pic_buckets" json:"pic_buckets"`
	VideoPath       string `mapstructure:"video_path" json:"video_path"`
	PicPath         string `mapstructure:"pic_path" json:"pic_path"`
}
type SvrConfig struct {
	Name string `mapstructure:"name"` // 服务name
	Host string `mapstructure:"host"` // 服务host
	Port int    `mapstructure:"port"`
	// UserSvrName     string `mapstructure:"user_svr_name"` // 用户服务name
	RelationSvrName string `mapstructure:"relation_svr_name"`
	FavoriteSvrName string `mapstructure:"favorite_svr_name"`
	// 雪花算法 机器ID
	MachineID int `mapstructure:"machine_id"`
}
type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Database    string `mapstructure:"database"`
	UserName    string `mapstructure:"username"`
	PassWord    string `mapstructure:"password"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleTime int    `mapstructure:"max_idle_time"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	LogPath    string `mapstructure:"log_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type RedisConfig struct {
	DB       int    `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"pool_size"`
	Host     string `mapstructure:"host"`
	PassWord string `mapstructure:"password"`
}
type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host" yaml:"host"`
	Port int      `mapstructure:"port" json:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" json:"tags" yaml:"tags"`
}

func Init() (err error) {
	configFile := GetRootDir() + "/config/config.yaml"
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config error: ", err)
		return fmt.Errorf("read config error: %v", err)
	}

	//反序列化
	if err = viper.Unmarshal(globalConfig); err != nil {
		fmt.Println("unmarshal config error: ", err)
		return fmt.Errorf("unmarshal config error: %v", err)
	}

	//热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
		if err = viper.Unmarshal(globalConfig); err != nil {
			fmt.Println("unmarshal config error: ", err)
		}
	})
	return nil
}
func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}
