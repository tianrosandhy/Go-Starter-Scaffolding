package bootstrap

import (
	"fmt"
	"regexp"
	"skeleton/src/database"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tianrosandhy/goconfigloader"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	Config *goconfigloader.Config
	Logger *logrus.Logger
}

func NewDatabase(cfg *goconfigloader.Config, logger *logrus.Logger) *gorm.DB {
	db := Database{
		Config: cfg,
		Logger: logger,
	}

	db.Migrate()
	return db.Connect()
}

func (db *Database) Connect(mode ...string) *gorm.DB {
	logLevel := logger.Info

	switch db.Config.GetString("ENVIRONMENT") {
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   db.Config.GetString("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	tx, err := gorm.Open(dbDriverDialect(db.Config), &config)
	if nil != err {
		panic(err)
	}

	if len(mode) == 0 || mode[0] != "migration" {
		db.Logger.Infof("Start database connection to %s", db.Config.GetString("DB_DRIVER"))
	}

	if nil != db {
		sqlDB, _ := tx.DB()
		sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNS"))
		sqlDB.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE_CONNS"))
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(viper.GetInt("DB_MAX_IDLE_TIME")))
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(viper.GetInt("DB_MAX_LIFE_TIME")))
	}

	return tx
}

func dbDriverDialect(cfg *goconfigloader.Config) gorm.Dialector {
	tz := cfg.GetString("TZ")
	if len(tz) == 0 {
		tz = "Asia/Jakarta"
	}

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		cfg.GetString("DB_USER"),
		cfg.GetString("DB_PASS"),
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_PORT"),
		cfg.GetString("DB_NAME"),
	) + "&loc=" + strings.ReplaceAll(tz, "/", "%2F")

	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s application_name=%s connect_timeout=%d",
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_PORT"),
		cfg.GetString("DB_USER"),
		cfg.GetString("DB_PASS"),
		cfg.GetString("DB_NAME"),
		tz,
		strings.ToLower(regexp.MustCompile(`[^A-z0-9-]`).ReplaceAllString(cfg.GetString("APP_NAME"), "-")),
		cfg.GetInt("DB_CONNECTION_TIMEOUT"),
	)

	sqliteDSN := "file::memory:?cache=shared"
	sqlitePath := cfg.GetString("DB_SQLITE_PATH")
	if sqlitePath != "" {
		sqliteDSN = fmt.Sprintf("%s?cache=shared", sqlitePath)
	}

	switch cfg.GetString("DB_DRIVER") {
	case "mysql":
		return mysql.Open(mysqlDSN)
	case "postgres":
		return postgres.Open(postgresDSN)
	case "sqlite":
		return sqlite.Open(sqliteDSN)
	}

	// fallback
	return mysql.Open(mysqlDSN)
}

func (db *Database) Migrate() {
	// only run dbMigrate() while "enable_migration" is set to true
	enableMigration := db.Config.GetBool("ENABLE_MIGRATION")
	if !enableMigration {
		db.Logger.Printf("MIGRATION IS DISABLED")
		return
	}

	tx := db.Connect("migration")

	if nil != db && len(database.EntityMigrations) > 0 {
		db.Logger.Printf("AUTOMIGRATE START")
		start := time.Now().UnixNano()
		err := tx.AutoMigrate(database.EntityMigrations...)
		end := time.Now().UnixNano()
		db.Logger.Printf("AUTOMIGRATE FINISH IN : %.3f s", ((float64(end) - float64(start)) / 1000000000))

		if nil != err {
			panic(err)
		}

		seeds := database.DataSeeds()
		if len(seeds) > 0 {
			db.Logger.Printf("AUTO SEEDER START")
			start := time.Now().UnixNano()

			for i := range seeds {
				trx := tx.Begin()
				defer func() {
					if r := recover(); r != nil {
						trx.Rollback()
					}
				}()
				if err := trx.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(seeds[i], 100).Error; nil != err {
					trx.Rollback()
				}
				if err := trx.Commit().Error; nil != err {
					trx.Rollback()
				}
			}

			end := time.Now().UnixNano()
			db.Logger.Printf("AUTO SEEDER FINISH IN : %.3f s", ((float64(end) - float64(start)) / 1000000000))
		}

		tx.Migrator().DropTable("schema_migration")
		sqlDB, _ := tx.DB()
		defer sqlDB.Close()
	}
}
