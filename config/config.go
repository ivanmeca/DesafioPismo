package config

import "github.com/spf13/viper"

const EnvPrefix = "app"

const (
	DBHost     = "database_host"
	DBPort     = "database_port"
	DBUser     = "database_user"
	DBName     = "database_name"
	DBPassword = "database_password"
	SSLMode    = "ssl_mode"

	QueueLogin    = "queue_login"
	QueuePassword = "queue_password"
	QueueHost     = "queue_host"
	QueuePort     = "queue_port"
	Queue         = "queue"
)

type Config struct {
	db    DBConfig
	queue QueueConfig
}

type QueueConfig struct {
	Login    string
	Password string
	Host     string
	Port     int
	Queue    string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Sslmode  string
	Password string
}

func Load() (*Config, error) {
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	viper.SetConfigName("Config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	cfg := new(Config)

	cfg.db = DBConfig{
		Host:     viper.GetString(DBHost),
		Port:     viper.GetString(DBPort),
		User:     viper.GetString(DBUser),
		Name:     viper.GetString(DBName),
		Sslmode:  viper.GetString(SSLMode),
		Password: viper.GetString(DBPassword),
	}

	cfg.queue = QueueConfig{
		Login:    viper.GetString(QueueLogin),
		Password: viper.GetString(QueuePassword),
		Host:     viper.GetString(QueueHost),
		Port:     viper.GetInt(QueuePort),
		Queue:    viper.GetString(Queue),
	}

	return cfg, nil
}

func (c *Config) GetDB() DBConfig {
	return c.db
}

func (c *Config) GetQueue() QueueConfig {
	return c.queue
}
