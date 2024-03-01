package config

type ServerConfig struct {
	Name        string `mapstructure:"APP_NAME" yaml:"name" env:"APP_NAME" env-default:"Tarkiz Paz Banua"`
	Version     string `mapstructure:"APP_VERSION" yaml:"version" env:"APP_VERSION" env-default:"dev"`
	Env         string `mapstructure:"APP_ENV" yaml:"env" env:"APP_ENV" env-default:"development"`
	Url         string `mapstructure:"APP_URL" yaml:"url" env:"APP_URL"`
	Host        string `mapstructure:"APP_HOST" yaml:"host" env:"APP_HOST" env-default:"localhost"`
	Port        string `mapstructure:"APP_PORT" yaml:"port" env:"APP_PORT" env-default:"3001"`
	Path        string `mapstructure:"APP_PATH" yaml:"path" env:"APP_PATH"`
	Debug       bool   `mapstructure:"APP_DEBUG" yaml:"debug" env:"APP_DEBUG" env-default:"true"`
	PublicPath  string `mapstructure:"PUBLIC_PATH" yaml:"public_path" env:"PUBLIC_PATH" env-default:"web/public"`
	UploadPath  string `mapstructure:"UPLOAD_PATH" yaml:"upload_path" env:"UPLOAD_PATH" env-default:"web/upload"`
	ExecPath    bool   `mapstructure:"EXEC_PATH" yaml:"exec_path" env:"EXEC_PATH" env-default:"false"`
	UploadLimit int    `mapstructure:"UPLOAD_LIMIT" yaml:"upload_limit" env:"UPLOAD_LIMIT" env-default:"8"`
}
