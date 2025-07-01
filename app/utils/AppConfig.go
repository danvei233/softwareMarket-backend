package utils

import (
	"errors"
	"github.com/go-ini/ini"
)

type AppConfig struct {
	Database Database
	Dir      Dir
	App      App
	Debug    Debug
	cfg      *ini.File
}

var conf AppConfig

type Database struct {
	Address  string
	Port     string // strconv too stupid
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}
type App struct {
	Port string // do not strconv
	Addr string
}
type Debug struct {
	DisableProductionMode bool
}

type Dir struct {
	DataRelativePath   string
	AssertRelativePath string
}

func (c *AppConfig) ReloadConfig() error {
	return c.cfg.Reload()

}
func (c *AppConfig) GetDsn() string {
	return "host=" + c.Database.Address + " user=" +
		c.Database.User + " password=" + c.Database.Password + " dbname=" +
		c.Database.DBName + " port=" + c.Database.Port + " sslmode=" + c.Database.SSLMode + " TimeZone=" + c.Database.TimeZone

}
func GetConfig() (*AppConfig, error) {
	if conf.cfg == nil {
		return nil, errors.New("not init")
	}
	return &conf, nil
}
func InitAppConfig(filepath string) (*AppConfig, error) {
	f, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}
	conf := AppConfig{cfg: f}
	if err := f.MapTo(&conf); err != nil {
		return nil, err
	}
	return &conf, nil

}
