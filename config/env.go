package config

import "os"

const (
	envKey   = "APP_ENV"
	EnvProd  = "prod"
	EnvLocal = "local"
)

var env = GetEnv(envKey, EnvLocal)

func Env() string {
	return env
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}

	return def
}
