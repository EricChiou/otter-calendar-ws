package config

import "github.com/EricChiou/config"

type Config struct {
	ServerName      string `key:"SERVER_NAME"`
	ServerPort      string `key:"SERVER_PORT"`
	DBUser          string `key:"DB_USER"`
	DBPwd           string `key:"DB_PWD"`
	DBName          string `key:"DB_NAME"`
	SSLCertFilePath string `key:"SSL_CERT_FILE_PATH"`
	SSLKeyFilePath  string `key:"SSL_KEY_FILE_PATH"`
	JWTKey          string `key:"JWT_KEY"`
	JWTExpire       int    `key:"JWT_EXPIRE"`
	ENV             string `key:"ENV"`
}

var cfg Config

func Load(filePath string) {
	err := config.Load(filePath, &cfg)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	return cfg
}
