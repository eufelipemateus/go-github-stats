package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API           APPConfig
	DB            dbConfig
	Github  github
}

type APPConfig struct {
	CollectLimit int
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type github struct {
	Token string
	Tokens []string
}

func init() {
	viper.SetDefault("collect.limit", "1000")
	viper.SetDefault("database.host", "locahost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APPConfig{
		CollectLimit: viper.GetInt("collect.limit"),
	}

	cfg.DB = dbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	cfg.Github =  github {
		Token: viper.GetString("github.token"),
		Tokens: viper.GetStringSlice("github.tokens"),
	}

	return nil
}

func GetDB() dbConfig {
	return cfg.DB
}

func GetCollectLimit() int {
	return cfg.API.CollectLimit
}

func GetGithubToken() string {
	return cfg.Github.Token
}

func GetGithubTokens() []string {
	return cfg.Github.Tokens
}