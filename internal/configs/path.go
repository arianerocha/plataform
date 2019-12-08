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

// Views path
func (c *PathConfig) Views() string {
	return path.Join(c.App(), "views")
}

// Static path
func (c *PathConfig) Static() string {
	return path.Join(c.App(), "static")
}
