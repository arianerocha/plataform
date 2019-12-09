package configs

import "fmt"

// WebConfig struct
type WebConfig struct{}

// Web config
func Web() *WebConfig {
	return &WebConfig{}
}

// Port Web
func (c *WebConfig) Port() string {
	return Sys().Env("PORT", "3030")
}

// RunnableInterface for the Web server
func (c *WebConfig) RunnableInterface() string {
	return fmt.Sprintf("0.0.0.0:%s", c.Port())
}
