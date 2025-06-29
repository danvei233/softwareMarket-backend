package utils

import "github.com/go-ini/ini"

type AppConfig struct {
	Database Database
	App      App
	Debug    Debug
	cfg      *ini.File
}
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

func (c *AppConfig) ReloadConfig() error {
	return c.cfg.Reload()

}
func (c *AppConfig) GetDsn() string {
	return "host=" + c.Database.Address + " user=" +
		c.Database.User + " password=" + c.Database.Password + " dbname=" +
		c.Database.DBName + " port=" + c.Database.Port + " sslmode=" + c.Database.SSLMode + " TimeZone=" + c.Database.TimeZone

}
func NewAppConfig(filepath string) (*AppConfig, error) {
	f, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}
	app := AppConfig{cfg: f}
	if err := f.MapTo(&app); err != nil {
		return nil, err
	}
	return &app, nil

}
