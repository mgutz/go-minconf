package minconf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mgutz/go-nestedjson"
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

// New creates a default instance of Minconf which loads ENV and ARGV.
func New() (*MinConf, error) {
	json := `{
		"$": {
			"envs": {"development": "dev ENV ARGV"},
		},
		"dev": {}
	}`
	return NewFromString(json)
}

// NewFromFile loads configuration from JSON file.
func NewFromFile(jsonFile string) (*MinConf, error) {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	return NewFromString(string(content))
}

// NewFromString loads configuration from JSON string.
func NewFromString(jsonString string) (*MinConf, error) {
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

	envSelector := stringOr(options["envSelector"], "RUN_ENV")
	defaultEnv := stringOr(options["defaultEnv"], "development")
	dotString := stringOr(options["dotString"], "__")

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
					src, err = newArgvMap().Config()
					if err != nil {
						return nil, err
					}

				case "ENV":
					src, err = newEnvMap(dotString).Config()
					if err != nil {
						return nil, err
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

	nj := nestedjson.NewFromMap(base)
	internalMap := nj.Data().(map[string]interface{})

	// some keys have dots such as ENV keys that had "__" replaced with "."
	for key, value := range internalMap {
		if strings.Contains(key, ".") {
			delete(internalMap, key)
			nj.Set(key, value)
		}
	}

	mc := &MinConf{Map: nj, EnvSelector: envSelector, DefaultEnv: defaultEnv, Env: env}
	return mc, nil
}
