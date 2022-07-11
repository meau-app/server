package config

var (
	Hostname string
	Port     string
)

func InitEvironment() {
	Hostname = GetEnvOrDefault("MEAU_HOSTNAME", "0.0.0.0")
	Port = GetEnvOrDefault("MEAU_PORT", "8080")
}
