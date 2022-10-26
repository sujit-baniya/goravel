package helpers

import (
	"github.com/sujit-baniya/framework/contracts/cache"
	"github.com/sujit-baniya/framework/facades"
	"github.com/sujit-baniya/redisCache"
)

func GetCache() cache.Store {
	defaultStore := facades.Config.GetString("cache.default")
	driver := facades.Config.GetString("cache.stores." + defaultStore + ".driver")
	switch driver {
	case "redis":
		connection := facades.Config.GetString("cache.stores." + facades.Config.GetString("cache.default") + ".connection")
		if connection == "" {
			connection = "default"
		}
		host := facades.Config.GetString("database.redis." + connection + ".host")
		port := facades.Config.GetString("database.redis." + connection + ".port")
		password := facades.Config.GetString("database.redis." + connection + ".password")
		db := facades.Config.GetInt("database.redis." + connection + ".database")
		prefix := facades.Config.GetString("cache.prefix" + ":")
		store, _ := redisCache.New(redisCache.Config{
			Prefix:   prefix,
			Host:     host,
			Port:     port,
			DB:       db,
			Password: password,
		})
		return store
	}
	return nil
}
