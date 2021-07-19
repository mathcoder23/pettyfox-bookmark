package main

import (
	"github.com/kataras/iris/v12"
	bookmarkController "pettyfox.top/bookmark/modules/bookmark/controller"
	"pettyfox.top/bookmark/modules/redis"
	"pettyfox.top/bookmark/modules/sonicCli"
)

func main() {
	app := iris.New()
	redis.InitRedis()
	sonicCli.InitSonicCli()
	bookmarkApi := app.Party("/bookmark")
	{
		bookmarkApi.Use(iris.Compression)
		bookmarkApi.Get("/list", bookmarkController.List)
		bookmarkApi.Post("/remove", bookmarkController.Remove)
		bookmarkApi.Post("/save", bookmarkController.Save)
		bookmarkApi.Post("/add", bookmarkController.Save)
	}
	app.Listen(":8080")
}
func after(ctx iris.Context) {
	ctx.WriteString("aa")
}
