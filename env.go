package minconf

import (
	"os"
	"strings"
)

type envMap struct {
	dotAlias string
	emap     map[string]interface{}
}

// newEnvMap creates a map from process'environment
//
// dotAlias is the key search pattern to replace with a dot or period. A key like "server__key" becomes "server.key"
func newEnvMap(dotAlias string) *envMap {
	return &envMap{dotAlias, nil}
}

func (em *envMap) getenvironment(data []string, getkeyval func(item string) (key, val string)) map[string]interface{} {
	items := map[string]interface{}{}
	for _, item := range data {
		key, val := getkeyval(item)
		items[key] = val
	}
	return items
}

// Config gets the config map from this process' environment.
func (em *envMap) Config() (map[string]interface{}, error) {
	if em.emap == nil {
		em.emap = em.getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")

			// allow dot representation, eg "sever__port" => "server.port"
			key = strings.Replace(splits[0], em.dotAlias, ".", -1)
			val = strings.Join(splits[1:], "=")
			return
		})
	}
	return em.emap, nil
}
