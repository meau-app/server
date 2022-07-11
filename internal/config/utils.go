package config

import "os"

// Gets an evironment variable value for e, if the environment value does not
// exist it returns the default set value d
func GetEnvOrDefault(e, d string) string {
	v := os.Getenv(e)

	if len(v) == 0 {
		return d
	}

	return v
}
