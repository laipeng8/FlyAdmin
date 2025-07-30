package config

import (
	"golang.org/x/time/rate"
)

type Config struct {
	Env string `yaml:"env"`
	Db  struct {
		Type         string `yaml:"type"`
		MaxIdleConns int    `yaml:"max-idle-conns"`
		MaxOpenConns int    `yaml:"max-open-conns"`
		Port         string `yaml:"port"`
		Host         string `yaml:"host"`
		TablePrefix  string `yaml:"table_prefix"`
		Database     string `yaml:"database"`
		User         string `yaml:"name"`
		PassWord     string `yaml:"password"`
	}
	MyJwt struct {
		Secret    string `yaml:"secret"`
		ExpiresAt int64  `yaml:"expires_at"`
	}
	App struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		UploadFile string `yaml:"uploadFile"`
		ImgUrl     string `yaml:"imgUrl"`
	}
	Rate struct {
		Limit rate.Limit `yaml:"limit"`
		Burst int        `yaml:"burst"`
	}
	Logger struct {
		Drive  string `yaml:"drive"`
		Path   string `yaml:"path"`
		Size   int    `yaml:"size"`
		MaxAge int    `yaml:"maxAge"`
		StdOut bool   `yaml:"stdOut"`
	}
	WX struct {
		MCHID          string `yaml:"MCHID"`
		PRIVATE_KEY    string `yaml:"PRIVATE_KEY"`
		CERT_SERIAL_NO string `yaml:"CERT_SERIAL_NO"`
		APIV3_KEY      string `yaml:"APIV3_KEY"`
		APPID          string `yaml:"APPID"`
		AppSecret      string `yaml:"AppSecret"`
		Code2Session   string `yaml:"code2Session"`
		Url            string `yaml:"URL"`
		CERT_DIR       string `yaml:"CERT_DIR"`
		PARTNER_MODE   string `yaml:"PARTNER_MODE"`
		PROXY          string `yaml:"PROXY"`
		TIMEOUT        string `yaml:"TIMEOUT"`
	}
	Cron struct {
		OrderStatusUpdate string `yaml:"order_status_update"`
	}
}
