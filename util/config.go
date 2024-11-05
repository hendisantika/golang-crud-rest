package util

type Config struct {
	DbDriver         string `mapstructure:"DB_DRIVER"`
	DbSource         string `mapstructure:"DB_SOURCE"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
}
