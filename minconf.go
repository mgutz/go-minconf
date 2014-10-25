package minconf

type MinConf struct {
	*NestedJson

	// Wd is work directory
	Wd string

	// Environment variable name
	EnvSelector string

	// Default environment to use if selector is unset
	DefaultEnv string

	// Environment in use
	Env string
}

func (mc *MinConf) MustArray(path string) []interface{} {
	val, err := mc.Array(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) MustBool(path string) bool {
	val, err := mc.Bool(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) MustFloat(path string) float64 {
	val, err := mc.Float(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) MustInt(path string) int {
	val, err := mc.Int(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) MustMap(path string) map[string]interface{} {
	val, err := mc.Map(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) MustString(path string) string {
	val, err := mc.String(path)
	mustNotError(err)
	return val
}

func (mc *MinConf) SafeArray(path string, value []interface{}) []interface{} {
	val, err := mc.Array(path)
	if err != nil {
		return value
	}
	return val
}

func (mc *MinConf) SafeBool(path string, value bool) bool {
	val, err := mc.Bool(path)
	if err != nil {
		return value
	}
	return val
}

func (mc *MinConf) SafeFloat(path string, value float64) float64 {
	val, err := mc.Float(path)
	if err != nil {
		return value
	}
	return val
}

func (mc *MinConf) SafeInt(path string, value int) int {
	val, err := mc.Int(path)
	if err != nil {
		return value
	}
	return val
}

func (mc *MinConf) SafeMap(path string, value map[string]interface{}) map[string]interface{} {
	val, err := mc.Map(path)
	if err != nil {
		return value
	}
	return val
}

func (mc *MinConf) SafeString(path string, value string) string {
	val, err := mc.String(path)
	if err != nil {
		return value
	}
	return val
}
