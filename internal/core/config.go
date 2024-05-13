package core

import "os"

// Db
type Db struct {
	Dsn string
}

// Add new configs here
type Config struct {
	Port    string
	Env     string
	Version string
	Db      Db
	// AppBranch string
	// AppUrl    string
}

// With NewConfig expect ENV variables passed from Makefile
// but ensure default values are always present in case running
// via go tooling.
// However worth noting that this appliation is tightly coupled
// to a database, without which it will fail.
func NewConfig() *Config {
	return &Config{
		Env:     getEnv("BUILD_REF", "Develop"),
		Port:    getEnv("APIPORT", "8585"),
		Version: getEnv("VERSION", "1.0"),
		Db: Db{
			Dsn: getEnv("DB_DSN", "postgres://postgres:pa55word123@localhost:5432/postgres?sslmode=disable"),
		},
		// Ideally these are also settable after the server is up
		// via some api as well.
		// AppBranch: getEnv("APP-BRANCH", "edu720-azure"),
		// AppUrl:    getEnv("APP-URL", "https://github.com/helxplatform/helx-apps"),
	}
}

// Help set default values
func getEnv(key, defaultValue string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}
	return defaultValue
}
