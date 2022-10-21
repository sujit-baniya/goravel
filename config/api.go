package config

import "github.com/sujit-baniya/framework/facades"

func init() {
	config := facades.Config
	config.Add("api", map[string]any{
		"prefix":  config.Env("API_PREFIX", "api"),
		"version": config.Env("API_VERSION", "v1"),
	})
}
