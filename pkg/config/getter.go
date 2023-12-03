package config

import "ginblog/config"

func Get() config.Config {
	return configuration
}
