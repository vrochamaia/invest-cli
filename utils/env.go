package utils

import "os"

func IsTestEnv() bool {
	return os.Getenv("APP_ENV") == "test"
}

func SetAppEnv(appEnv string) {
	os.Setenv("APP_ENV", appEnv)
}
