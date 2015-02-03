package minconf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mgutz/minimist"
	"github.com/mgutz/str"
)

func mustNotError(e error) {
	if e != nil {
		panic(e)
	}
}

// removeComments removes single-line eol comments
func removeComments(json string) string {
	lines := str.Lines(json)
	// remove single line comments
	lines = str.Map(lines, func(line string) string {
		if str.Match(line, `^\s*//`) {
			return ""
		}
		return line
	})
	return strings.Join(lines, "\n")
}

func stringOr(v interface{}, value string) string {
	if v == nil {
		return value
	}
	result := v.(string)
	if result == "" {
		return value
	}
	return result
}

// LoadFile loads configuration from JSON file.
func LoadFile(jsonFile string) (*MinConf, error) {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	return LoadString(string(content))
}

// LoadString loads configuration from JSON string.
func LoadString(jsonString string) (*MinConf, error) {
	const metaKey = "$"
	jsonString = removeComments(jsonString)

	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("invalid configuration")
	}

	// $ = minconf options
	meta := m[metaKey].(map[string]interface{})
	if meta == nil {
		return nil, errors.New(metaKey + " property is required")
	}

	var options map[string]interface{}
	if meta["options"] == nil {
		options = make(map[string]interface{})
	} else {
		options = meta["options"].(map[string]interface{})
	}

	envSelector := stringOr(options["envSelector"], "GO_ENV")
	defaultEnv := stringOr(options["defaultEnv"], "development")
	replaceWithDot := stringOr(options["replaceWithDot"], "__")

	envs := meta["envs"].(map[string]interface{})
	if envs == nil {
		return nil, errors.New("$.envs property is required")
	}

	// get environment from selector
	env := os.Getenv(envSelector)
	if env == "" {
		env = defaultEnv
	}

	base := make(map[string]interface{})
	for key, mergeSpec := range envs {
		if key == env {
			mergeables := strings.Split(mergeSpec.(string), " ")
			for _, mergeable := range mergeables {
				var src map[string]interface{}

				switch mergeable {
				case "ARGV":
					src = make(map[string]interface{})
					args := minimist.Parse(os.Args[1:], nil, nil, nil)
					for key, val := range args {
						src[key] = val
					}

				case "ENV":
					src = make(map[string]interface{})
					for envKey, envVal := range envMap(replaceWithDot) {
						src[envKey] = envVal
					}

				default:
					src = m[mergeable].(map[string]interface{})
				}

				if src != nil {
					base, err = Merge(base, src)
					if err != nil {
						return nil, err
					}
				}
			}

		}
	}

	njson := NewNestedJson(base)
	internalMap := njson.Data()

	// some keys have dots such as ENV keys that had "__" replaced with "."
	// not sure if this is safe as Data() is being modi?
	for key, value := range internalMap {
		if strings.Contains(key, ".") {
			delete(njson.Data(), key)
			njson.Set(key, value)
		}
	}

	mc := &MinConf{NestedJson: njson, EnvSelector: envSelector, DefaultEnv: defaultEnv, Env: env}
	return mc, nil
}
