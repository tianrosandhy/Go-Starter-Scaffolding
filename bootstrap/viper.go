package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LoadEnvironment autoload environment
func NewViperConfig(defaultValues map[string]interface{}) *viper.Viper {
	viperConfig := viper.New()

	viperLoadSystemEnv(viperConfig, defaultValues)
	viperLoadLocalEnv(viperConfig)
	viperLoadParamEnv(viperConfig, defaultValues)
	viperLoadPrivateEnv(viperConfig, defaultValues)
	viperMergeAllToOS(viperConfig)

	return viperConfig
}

// viperLoadSystemEnv Load System Environment
func viperLoadSystemEnv(viperConfig *viper.Viper, defaultValues map[string]interface{}) {
	viperConfig.AutomaticEnv()
	for k := range defaultValues {
		viper.Set(k, viperConfig.Get(k))
	}
}

// viperLoadLocalEnv Load Local Environment
func viperLoadLocalEnv(viperConfig *viper.Viper) {
	// load local env
	viperConfig.SetConfigType("dotenv")
	viperConfig.SetConfigFile(".env")

	configPath := ".env"
	err := viperConfig.ReadInConfig()
	for i := 0; i < 5; i++ {
		if err == nil {
			break
		}
		configPath = "../" + configPath
		viperConfig.SetConfigFile(configPath)
		err = viperConfig.ReadInConfig()
	}
	if err != nil {
		panic("Cannot read .env file. Make sure the config is exists in app")
	}

	viperConfigKeys := viperConfig.AllKeys()
	for i := range viperConfigKeys {
		viper.Set(viperConfigKeys[i], viperConfig.Get(viperConfigKeys[i]))
	}
}

// viperLoadParamEnv Load Parameter Environment
func viperLoadParamEnv(viperConfig *viper.Viper, defaultValues map[string]interface{}) {
	// load parameter env
	viperConfig.AllowEmptyEnv(false)
	for k := range defaultValues {
		if flagKey := strcase.ToKebab(k); nil == pflag.Lookup(flagKey) {
			pflag.String(flagKey, "", k)
		}
	}

	if os.Getenv("ENVIRONMENT_SIMULATION") != "" {
		for k := range defaultValues {
			pflag.CommandLine.Set(strcase.ToKebab(k), viper.GetString(k))
		}
	}

	pflag.Parse()
	if err := viperConfig.BindPFlags(pflag.CommandLine); nil == err {
		viperConfigKeys := viperConfig.AllKeys()
		for i := range viperConfigKeys {
			if stringValue := viperConfig.GetString(viperConfigKeys[i]); stringValue != "" {
				viper.Set(strcase.ToSnake(viperConfigKeys[i]), stringValue)
			}
		}
	}
}

// viperLoadPrivateEnv Load Private Environment
func viperLoadPrivateEnv(viperConfig *viper.Viper, defaultValues map[string]interface{}) {
	// load default env
	for k, v := range defaultValues {
		if !viperConfig.InConfig(k) {
			viperConfig.SetDefault(k, v)
		}
	}
}

// viperMergeAllToOS Merge all System, Local, Parameter and Private Environment
func viperMergeAllToOS(viperConfig *viper.Viper) {
	keys := viperConfig.AllKeys()
	for i := range keys {
		stringValue := viperConfig.GetString(keys[i])
		if stringValue == "" {
			if value := viperConfig.Get(keys[i]); nil != value {
				stringValue = fmt.Sprintf("%v", value)
			}
		}
		os.Setenv(strings.ToUpper(keys[i]), stringValue)
	}
}
