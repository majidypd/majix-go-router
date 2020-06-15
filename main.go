package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/my/repo/majix"
	"github.com/my/repo/path"
)

func main() {
	db, _ := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=1")
	defer db.Close()

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	app := majix.NewApp()
	app.Bind("db", db)
	app.Bind("redis", redis)
	app.Bind("router", majix.NewRouter())
	app.Bind("session", majix.NewSessionManager(&majix.RedisProvider{Driver: redis}))
	app.Bind("session", majix.NewSessionManager(&majix.GormProvider{Driver: db}))
	app.Start(":8091", path.Web)
}
