package config

// Config is 全ての接続情報
type Config struct {
	Postgres *PostgresConfig
}

// PostgresConfig is Postgresへの接続情報
type PostgresConfig struct {
	Host string `env:"POSTGRES_HOST"`
	Port string `env:"POSTGRES_PORT"`
	Pass string `env:"POSTGRES_PASSWORD"`
	User string `env:"POSTGRES_USER"`
	Db   string `env:"POSTGRES_DB"`
}
