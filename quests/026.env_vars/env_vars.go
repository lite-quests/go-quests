package env_vars

// TODO: Implement LoadConfig
// Read README.md for the instructions

type Config struct {
	DBHost    string
	DBPort    int
	DebugMode bool
}

func LoadConfig() (Config, error) {
	return Config{}, nil
}
