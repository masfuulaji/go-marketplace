package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DBHost string `mapstructure:"db_host"`
	DBPort string `mapstructure:"db_port"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBName string `mapstructure:"db_name"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func InitDB() *gorm.DB {
	cnf, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cnf.DBHost, cnf.DBPort, cnf.DBUser, cnf.DBPass, cnf.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	return db
}
