package configs

import (
	"os"
	"os/user"
)

// SysConfig struct
type SysConfig struct{}

// Sys config
func Sys() *SysConfig {
	return &SysConfig{}
}

// Env Sys
func (c *SysConfig) Env(key, def string) string {
	ret := os.Getenv(key)

	if ret == "" {
		return def
	}

	return ret
}

// HomeDir Sys
func (c *SysConfig) HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		return "/"
	}

	return usr.HomeDir
}
