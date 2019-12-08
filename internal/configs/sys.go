package configs

import (
	"os/user"
)

// SysConfig struct
type SysConfig struct{}

// Sys config
func Sys() *SysConfig {
	return &SysConfig{}
}

// HomeDir Sys
func (c *SysConfig) HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		return "/"
	}

	return usr.HomeDir
}
