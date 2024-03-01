package config

type DatabaseDriver struct {
	Driver   string `yaml:"driver" env:"DB_DRIVER"`
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     int    `yaml:"port" env:"DB_PORT"`
	Username string `yaml:"username" env:"DB_USERNAME"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Name     string `yaml:"name" env:"DB_NAME"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver" env:"DB_DRIVER"`
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     int    `yaml:"port" env:"DB_PORT"`
	Username string `yaml:"username" env:"DB_USERNAME"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Name     string `yaml:"name" env:"DB_NAME"`
	Debug    string `yaml:"debug" env:"DEBUG" env-default:"false"`
}
