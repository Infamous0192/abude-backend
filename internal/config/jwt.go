package config

type JwtConfig struct {
	Issuer string `mapstructure:"JWT_ISSUER" yaml:"issuer" env:"JWT_ISSUER" env-default:"Tarkiz Paz Banua"`
	Secret string `mapstructure:"JWT_SECRET" yaml:"secret" env:"JWT_SECRET" env-default:"thisisasecret"`
	Expire int    `mapstructure:"JWT_EXPIRE" yaml:"expire" env:"JWT_EXPIRE" env-default:"3600"`
}
