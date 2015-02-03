package minconf

// MinConf is the main configuration object.
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

// MustArray must convert path value to an interface{} slice or panic.
func (mc *MinConf) MustArray(path string) []interface{} {
	val, err := mc.Array(path)
	mustNotError(err)
	return val
}

// MustBool must convert a path value to bool or panic.
func (mc *MinConf) MustBool(path string) bool {
	val, err := mc.Bool(path)
	mustNotError(err)
	return val
}

// MustFloat must convert a path value to float64 or panic.
func (mc *MinConf) MustFloat(path string) float64 {
	val, err := mc.Float(path)
	mustNotError(err)
	return val
}

// MustInt must convert a path value to int or panic.
func (mc *MinConf) MustInt(path string) int {
	val, err := mc.Int(path)
	mustNotError(err)
	return val
}

// MustMap must convert a path value to map or panic.
func (mc *MinConf) MustMap(path string) map[string]interface{} {
	val, err := mc.Map(path)
	mustNotError(err)
	return val
}

// MustString must convert a path value to string or panic.
func (mc *MinConf) MustString(path string) string {
	val, err := mc.String(path)
	mustNotError(err)
	return val
}

// SafeArray should convert path value to []interface{} or return value.
func (mc *MinConf) SafeArray(path string, value []interface{}) []interface{} {
	if val, err := mc.Array(path); err == nil {
		return val
	}
	return value
}

// SafeBool should convert path value to bool or return value.
func (mc *MinConf) SafeBool(path string, value bool) bool {
	if val, err := mc.Bool(path); err == nil {
		return val
	}
	return value
}

// SafeFloat should convert path value to float64 or return value.
func (mc *MinConf) SafeFloat(path string, value float64) float64 {
	if val, err := mc.Float(path); err == nil {
		return val
	}
	return value
}

// SafeInt should convert path value to int or return value.
func (mc *MinConf) SafeInt(path string, value int) int {
	if val, err := mc.Int(path); err == nil {
		return val
	}
	return value
}

// SafeMap should convert path value to map[string]interface{} or return value.
func (mc *MinConf) SafeMap(path string, value map[string]interface{}) map[string]interface{} {
	if val, err := mc.Map(path); err == nil {
		return val
	}
	return value
}

// SafeString should convert path value to string or return value.
func (mc *MinConf) SafeString(path string, value string) string {
	if val, err := mc.String(path); err == nil {
		return val
	}
	return value
}
