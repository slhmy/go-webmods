# Config

Read configuration from config files or environment variables.

Config namings are recommended to be in snake_case and separated by dots.

You can set configs in `conifg` file and override them in `override` file
(multiple file extensions are supported, such as `.yaml`, `.toml`, `.json`, etc.).

In some cases, using environment variables is more convenient.
You can override the final config by making every character uppercase and replacing `.` with `_`
(e.g. `server.port` becomes `SERVER_PORT`),
environment variables have the highest priority.

## Example

```go
os.Setenv("SERVER_PORT", "443")
port = gwm_app.Config().GetInt("server.port")
// port = 443
```

## Customize

Customize the config file path by setting the `GWM_RELATIVE_CONFIG_PATH` environment variable.
Set the override config file name by setting the `GWM_OVERRIDE_CONFIG_NAME` environment variable.

The following environment variables should be set before running the application.
