package config

import (
	"os"

	"github.com/caarlos0/env"
)

// Environment ...
type Environment string

// Conf 設定情報グローバル変数
var Conf *Config

const (
	develop Environment = "develop"
	local   Environment = "local"
	test    Environment = "test"
)

// GetGoEnv is os.Getenvのラッパー関数です。
func GetGoEnv() Environment {
	return Environment(os.Getenv("GO_ENV"))
}

// NewConfig is 接続情報を初期化する
func NewConfig(e Environment) error {
	pc, err := newPostgresConf(e)
	if err != nil {
		return err
	}

	Conf = &Config{
		Postgres: pc,
	}
	return nil
}

func newPostgresConf(e Environment) (*PostgresConfig, error) {
	c := &PostgresConfig{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
