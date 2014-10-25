package minconf

import (
	"os"
	"strings"
)

// from https://coderwall.com/p/kjuyqw

var _envmap map[string]string

func envMap(replaceWithDot string) map[string]string {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}

	if _envmap == nil {
		_envmap = getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")

			// allow dot representation `server__port` -> `server.port`
			key = strings.Replace(splits[0], replaceWithDot, ".", -1)
			val = strings.Join(splits[1:], "=")
			return
		})
	}
	return _envmap
}
