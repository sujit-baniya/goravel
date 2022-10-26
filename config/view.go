package config

import "github.com/sujit-baniya/framework/facades"

func init() {
	config := facades.Config
	config.Add("view", map[string]any{
		"template":  config.Env("VIEW_TEMPLATE", "resources/views"),
		"layout":    config.Env("VIEW_LAYOUT", "layout"),
		"extension": config.Env("VIEW_FILE_EXTENSION", ".html"),
	})
}
