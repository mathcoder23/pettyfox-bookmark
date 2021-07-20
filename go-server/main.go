package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	bookmarkController "pettyfox.top/bookmark/modules/bookmark/controller"
	"pettyfox.top/bookmark/modules/redis"
	"pettyfox.top/bookmark/modules/sonicCli"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 这里写允许的服务器地址，* 号标识任意
		AllowCredentials: true,
	})
	redis.InitRedis()
	sonicCli.InitSonicCli()
	bookmarkApi := app.Party("/bookmark", crs).AllowMethods(iris.MethodOptions)
	{
		bookmarkApi.Use(iris.Compression)
		bookmarkApi.Get("/list", bookmarkController.List)
		bookmarkApi.Get("/search", bookmarkController.Search)
		bookmarkApi.Post("/remove", bookmarkController.Remove)
		bookmarkApi.Post("/save", bookmarkController.Save)
		bookmarkApi.Post("/add", bookmarkController.Save)
	}
	app.Listen(":8080")
}
func after(ctx iris.Context) {
	ctx.WriteString("aa")
}
