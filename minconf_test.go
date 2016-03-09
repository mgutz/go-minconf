package minconf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	assert := assert.New(t)

	config := `{
		"$": {
			"envs": {
				"development": "development"
			}
		},

		"development": {
			"server": {
				"port": 8080
			}
		}
	}`

	mc, _ := NewFromString(config)
	assert.Equal(mc.Env, "development")
	assert.Equal(mc.DefaultEnv, "development")
	assert.Equal(mc.EnvSelector, "RUN_ENV")
	assert.Equal(mc.MustInt("server.port"), 8080)
}

func TestMerging(t *testing.T) {
	assert := assert.New(t)

	config := `{
		"$": {
			"envs": {
				"development": "common development"
			}
		},

		"common": {
			"server": {
				"hostname": "localhost",
				"port": 8080
			}
		},

		"development": {
			"server": {
				"hostname": "dev.lan"
			}
		}
	}`

	mc, _ := NewFromString(config)
	assert.Equal(mc.MustInt("server.port"), 8080)
	assert.Equal(mc.MustString("server.hostname"), "dev.lan")
}

func TestEnvMerging(t *testing.T) {
	assert := assert.New(t)

	os.Setenv("server__port", "9000")

	config := `{
		"$": {
			"options": {
				"dotAlias": "__"
			},

			"envs": {
				"development": "common ENV"
			}
		},

		"common": {
			"server": {
				"hostname": "localhost",
				"port": 8080
			}
		}
	}`

	mc, _ := NewFromString(config)
	assert.Equal(mc.MustInt("server.port"), 9000)
	assert.Equal(mc.MustString("server.hostname"), "localhost")
}

func TestArgvMerging(t *testing.T) {
	assert := assert.New(t)

	oldArgs := os.Args
	os.Args = []string{"cmd", "--server.port", "9000", "--server.hostname=localhost"}
	config := `{
		"$": {
			"envs": {
				"development": "common ARGV"
			}
		},

		"common": {
			"server": {
				"hostname": "localhost",
				"port": 8080
			}
		}
	}`
	mc, _ := NewFromString(config)
	assert.Equal(mc.MustInt("server.port"), 9000)
	assert.Equal(mc.MustString("server.hostname"), "localhost")
	os.Args = oldArgs
}

func TestComments(t *testing.T) {
	config := `{
		// this is a comment
		"$": {
			"envs": {
				"development": "common ENV"
			}
		},

		"common": {
			"url": "http://thisisnot/a/comment"
		}
	}`

	mc, _ := NewFromString(config)
	assert.Equal(t, mc.MustString("url"), "http://thisisnot/a/comment")
}
