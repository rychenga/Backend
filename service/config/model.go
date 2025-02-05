package config

type GinConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	GinConfig GinConfig `mapstructure:"gin_config"`
}
