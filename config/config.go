package config

// Environment variables
var Environment map[string]interface{} = map[string]interface{}{
	// app config
	"app_name":       "skeleton",
	"port":           9000,
	"environment":    "development",
	"tz":             "Asia/Jakarta",
	"max_body_limit": 10, // in MB
	"prefork":        false,
	"aes":            "AES_CIPHER_SECRET_KEY_KUDU_32BIT",
	"salt":           "SALT",

	"basic_auth":       "", // example: user:password
	"swagger_auth":     "swag:swag",
	"limiter_max_hit":  60,
	"limiter_duration": 5,

	// telegram log
	"enable_telegram_log":   false,
	"telegram_bot_endpoint": "https://api.telegram.org/bot",
	"telegram_bot_token":    "",
	"telegram_bot_chatid":   "",

	// redis config
	"redis_host":        "localhost",
	"redis_port":        6379,
	"redis_pass":        "",
	"redis_index":       0,
	"redis_compression": true, // if true, redis value will compressed with GZIP

	// database config
	"enable_migration":      true,       // always set true this value for production usage
	"db_driver":             "postgres", // postgres/mysql/sqlite
	"db_host":               "localhost",
	"db_port":               5432,
	"db_user":               "postgres",
	"db_pass":               "postgres",
	"db_name":               "postgres",
	"db_table_prefix":       "",
	"db_sqlite_path":        "./db.sqlite",
	"db_connection_timeout": 5,    // Only wait for X second for db network connection only
	"db_statement_timeout":  40,   // default statement_timeout set on pool connection level
	"db_max_life_time":      3600, // Seconds max connection life time: 0 = unilimited
	"db_max_idle_time":      300,  // Seconds max connection idle time: 0 = unilimited
	"db_max_idle_conns":     2,    // Max idle connection: 0 = unlimited
	"db_max_open_conns":     3,    // Max open connections: 0 = unlimited
	"db_log_not_found":      false,

	// log config
	"log_level":    "debug",
	"log_max_size": 50,
	"log_path":     "../../logs/app.log",
}
