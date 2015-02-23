## Sample Configuration

```js
{
    "$": {
        // these are the default options and this entire section may be omitted
        "options": {
            // SERVER__PORT => SERVER.PORT
            "dotAlias": "__",
            // selects config environment based on the value of this environment variable
            "envSelector": "RUN_ENV",
            // default environment if selector is unset or empty
            "defaultEnv": "development"
        },

        // named config environments, REQUIRED
        "envs": {
            // common <- ENV (environment variables) <- ARGV (command line argv)
            "development": "common ENV ARGV",
            "test": "common test ENV ARGV"
        }
    },

    "common": {
        "server": {
            "host": "localhost",
            "port": 8080
        }
    },

    "test": {
        "server": {
            // overrides port 8080 in common
            "port": 80
        }
    }
}
```

To use development runtime environment (default)

```sh
yourapp
```

To use test runtime environment

```sh
RUN_ENV=test yourapp
```

To override with command line args

```sh
yourapp --server.port=9000
```

To override using environment variables

```sh
SERVER__PORT=9000 yourapp
```


