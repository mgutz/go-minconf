package minconf

import (
	"os"
	"regexp"
	//"strings"

	"github.com/mgutz/str"
)

var _argvMap map[string]interface{}

func setArg(key string, val interface{}, arg string) {
	_argvMap[key] = val
}

func nextString(list []string, i int) *string {
	if i+1 < len(list) {
		return &list[i+1]
	}
	return nil
}

func contains(strings []string, needle string) bool {
	for _, s := range strings {
		if s == needle {
			return true
		}
	}
	return false
}

func argvMap(bools, strings []string) map[string]interface{} {
	if _argvMap != nil {
		return _argvMap
	}

	_argvMap = make(map[string]interface{})

	_argvMap["_"] = []string{}

	args := os.Args
	l := len(args)
	peek := func(i int) *string {
		if i < l {
			return &args[i+1]
		}
		return nil
	}

	// --port=8000
	longFormEqualRe := regexp.MustCompile(`^--.+=`)
	longFormEqualValsRe := regexp.MustCompile(`^--([^=])=([\s\S]*)$`)

	// --port 8000
	longFormSpaceRe := regexp.MustCompile(`^--.+`)
	longFormSpaceKeyRe := regexp.MustCompile(`^--(.+)`)
	//longFormSpaceValsRe := regexp.MustCompile(`^--([^=])=([\s\S]*)$`)

	// --no-debug
	negateRe := regexp.MustCompile(`^--no-.+`)
	negateValsRe := regexp.MustCompile(`^--no-(.+)`)

	// -ab == -a -b
	shortFormRe := regexp.MustCompile(`^-[^-]+`)

	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]
		if longFormEqualRe.MatchString(arg) {
			// --long-form=value
			m := longFormEqualValsRe.FindStringSubmatch(arg)
			setArg(m[1], m[2], arg)

		} else if negateRe.MatchString(arg) {
			// --no-flag
			m := negateValsRe.FindStringSubmatch(arg)
			setArg(m[1], false, arg)

		} else if longFormSpaceRe.MatchString(arg) {
			// --long-form value
			key := longFormSpaceKeyRe.FindStringSubmatch(arg)[1]
			pnext := peek(i + 1)

			if pnext == nil {
				setArg(key, true, arg)
			} else if str.CharAt(*pnext, 0) == "-" {
				setArg(key, true, arg)
			} else if !contains(bools, key) {
				setArg(key, *pnext, arg)
				i += 1
			} else if *pnext == "true" || *pnext == "false" {
				setArg(key, *pnext, arg)
				i += 1
			} else {
				setArg(key, true, arg)
			}
		} else if shortFormRe.MatchString(arg) {
			// TODO
			key := ""
			for k := 1; k < len(arg); k++ {
				next := str.CharAt(arg, k+2)
				if next == "-" {
					setArg(key, str.CharAt(arg, k), arg)
				}

			}

		}
	}

	return nil
}
