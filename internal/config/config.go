package config

import "os"

type Config struct {
	ConfigDirName  string
	ConfigFileName string
}

var Env *Config

func Load() {
	Env = &Config{
		ConfigDirName:  os.Getenv("CONFIG_DIR_NAME"),
		ConfigFileName: os.Getenv("CONFIG_FILE_NAME"),
	}
}
