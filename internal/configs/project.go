package configs

// ProjectConfig struct
type ProjectConfig struct{}

// Project config
func Project() *ProjectConfig {
	return &ProjectConfig{}
}

// Name Project
func (c *ProjectConfig) Name() string {
	return "plataform"
}

// Org Project
func (c *ProjectConfig) Org() string {
	return "krakenlab"
}
