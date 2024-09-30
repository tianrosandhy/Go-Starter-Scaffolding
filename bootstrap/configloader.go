package bootstrap

import (
	"github.com/tianrosandhy/goconfigloader"
)

// LoadEnvironment autoload environment
func NewConfigLoader(defaultValues map[string]string) *goconfigloader.Config {
	config := goconfigloader.NewConfigLoader(defaultValues)
	return config
}
