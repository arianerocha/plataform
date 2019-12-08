package configs

import (
	"path"
)

// PathConfig struct
type PathConfig struct{}

// Path config
func Path() *PathConfig {
	return &PathConfig{}
}

// App path
func (c *PathConfig) App() string {
	return path.Join(Sys().HomeDir(), Project().Org(), Project().Name())
}
