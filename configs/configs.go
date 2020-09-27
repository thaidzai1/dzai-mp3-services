package configs

// Config ...
type Config struct {
	Dcrawl ServiceConfig `yaml:"dcrawl"`
}

// ServiceConfig ...
type ServiceConfig struct {
	Address string `yaml:"address"`
}

// DefaultConfig returns default configuration
func DefaultConfig() Config {
	return Config{
		Dcrawl: ServiceConfig{
			Address: ":4001",
		},
	}
}
