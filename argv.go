package minconf

import "github.com/mgutz/minimist"

var argvm *argvMap

type argvMap struct {
	mapp map[string]interface{}
}

// NewEnvMap creates a map from process'environment.
func newArgvMap() *argvMap {
	if argvm == nil {
		argvm = &argvMap{minimist.Parse()}
	}
	return argvm
}

// Config gets the map from this process' environment.
func (am *argvMap) Config() (map[string]interface{}, error) {
	return am.mapp, nil
}
