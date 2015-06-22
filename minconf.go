package minconf

import "github.com/mgutz/jo/v1"

// MinConf is the main configuration object.
type MinConf struct {
	*jo.Object

	// Wd is work directory
	Wd string

	// Environment variable name
	EnvSelector string

	// Default environment to use if selector is unset
	DefaultEnv string

	// Environment in use
	Env string
}

// ConfigProvider is the interface which config providers must implement.
type ConfigProvider interface {
	// Config gets the map from a concrement implementation of this interface.
	Config() (map[string]interface{}, error)
}
