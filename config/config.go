package config

// Environment variables
var Environment map[string]string = map[string]string{
	// app config
	"APP_NAME":       "skeleton",
	"APP_HOST":       "localhost:9009",
	"PORT":           "9009",
	"ENVIRONMENT":    "development",
	"TZ":             "Asia/Jakarta",
	"MAX_BODY_LIMIT": "10", // in MB
	"AES":            "AES_CIPHER_SECRET_KEY_KUDU_32BIT",
	"SALT":           "SALT",

	"LIMITER_MAX_HIT":  "60",
	"LIMITER_DURATION": "5",

	// telegram log
	"ENABLE_TELEGRAM_LOG":   "false",
	"TELEGRAM_BOT_ENDPOINT": "https://api.telegram.org/bot",
	"TELEGRAM_BOT_TOKEN":    "",
	"TELEGRAM_BOT_CHATID":   "",

	// database config
	"ENABLE_MIGRATION":      "true",   // always set true this value for production usage
	"DB_DRIVER":             "sqlite", // postgres/mysql/sqlite
	"DB_HOST":               "localhost",
	"DB_PORT":               "5432",
	"DB_USER":               "postgres",
	"DB_PASS":               "postgres",
	"DB_NAME":               "postgres",
	"DB_TABLE_PREFIX":       "",
	"DB_SQLITE_PATH":        "./db.sqlite",
	"DB_CONNECTION_TIMEOUT": "5",    // Only wait for X second for db network connection only
	"DB_STATEMENT_TIMEOUT":  "30",   // default statement_timeout set on pool connection level
	"DB_MAX_LIFE_TIME":      "3600", // Seconds max connection life time: 0 = unilimited
	"DB_MAX_IDLE_TIME":      "300",  // Seconds max connection idle time: 0 = unilimited
	"DB_MAX_IDLE_CONNS":     "2",    // Max idle connection: 0 = unlimited
	"DB_MAX_OPEN_CONNS":     "3",    // Max open connections: 0 = unlimited
	"DB_LOG_NOT_FOUND":      "false",

	// log config
	"LOG_LEVEL":             "debug",
	"LOG_MAX_SIZE":          "50",
	"LOG_PATH":              "./logs/app.log",
	"TELEGRAM_BOT_LOG_PATH": "./logs/telegram.log",
}
