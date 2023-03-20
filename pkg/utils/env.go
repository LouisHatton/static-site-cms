package utils

import "os"

func GetEnv(env string, alt string) string {
	value, exists := os.LookupEnv(env)
	if exists {
		return value
	} else {
		return alt
	}
}
