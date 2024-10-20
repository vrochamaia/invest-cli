package utils

import "os"

const APP_ENV = "APP_ENV"

func IsTestEnv() bool {
	return os.Getenv(APP_ENV) == "test"
}

func SetAppEnv(appEnv string) {
	os.Setenv(APP_ENV, appEnv)
}
